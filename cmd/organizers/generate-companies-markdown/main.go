package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/generated/organizers"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
	"github.com/readytotouch/readytotouch/internal/templates/dev"
)

const (
	outputDir = "./docs/companies"
	siteURL   = "https://readytotouch.com"
)

type language struct {
	organizer domain.Organizer
	title     string // display name in README
	markdown  bool   // generate docs/companies/{alias}.md
}

type readme struct {
	path    string
	headers [2]string // language and companies column titles
	list    string
}

var (
	// The order is the order of rows in the README table
	languages = []language{
		{organizer: domain.OrganizerGolang, title: "Go", markdown: true},
		{organizer: domain.OrganizerRust, title: "Rust", markdown: true},
		{organizer: domain.OrganizerScala, title: "Scala", markdown: true},
		{organizer: domain.OrganizerElixir, title: "Elixir", markdown: true},
		{organizer: domain.OrganizerClojure, title: "Clojure", markdown: true},
		{organizer: domain.OrganizerErlang, title: "Erlang"},
		{organizer: domain.OrganizerZig, title: "Zig"},
		{organizer: domain.OrganizerHaskell, title: "Haskell"},
		{organizer: domain.OrganizerFSharp, title: "F#"},
		{organizer: domain.OrganizerOCaml, title: "OCaml"},
	}

	readmes = []readme{
		{path: "./README.md", headers: [2]string{"Language", "Companies"}, list: "Markdown list"},
		{path: "./README.uk.md", headers: [2]string{"Мова", "Компаній"}, list: "Список Markdown"},
	}

	tableRegexp = regexp.MustCompile(`(?s)<!-- languages-table:start -->\n.*?<!-- languages-table:end -->`)
)

func main() {
	var (
		companies = db.CloneCompanies()
		counts    = make(map[string]int, len(languages))
	)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		panic(err)
	}

	for _, lang := range languages {
		data := buildData(lang.organizer, companies)
		counts[lang.organizer.Alias] = data.Total

		if !lang.markdown {
			continue
		}

		path := filepath.Join(outputDir, lang.organizer.Alias+".md")
		if err := os.WriteFile(path, []byte(dev.Companies(data)), 0644); err != nil {
			panic(err)
		}

		fmt.Printf("%s: %d companies\n", path, data.Total)
	}

	for _, r := range readmes {
		updateReadme(r, counts)
	}
}

func updateReadme(r readme, counts map[string]int) {
	content, err := os.ReadFile(r.path)
	if err != nil {
		panic(err)
	}

	if !tableRegexp.Match(content) {
		panic(fmt.Sprintf("%s: languages-table markers not found", r.path))
	}

	block := "<!-- languages-table:start -->\n" + languagesTable(r, counts) + "<!-- languages-table:end -->"

	updated := tableRegexp.ReplaceAllLiteral(content, []byte(block))
	if err := os.WriteFile(r.path, updated, 0644); err != nil {
		panic(err)
	}
}

func languagesTable(r readme, counts map[string]int) string {
	rows := [][3]string{{r.headers[0], r.headers[1], r.list}}

	for _, lang := range languages {
		var (
			alias    = lang.organizer.Alias
			count    = "—"
			markdown = "—"
		)

		if counts[alias] > 0 {
			count = fmt.Sprint(counts[alias])
		}
		if lang.markdown {
			markdown = fmt.Sprintf("[%s](docs/companies/%s.md)", lang.title, alias)
		}

		rows = append(rows, [3]string{
			fmt.Sprintf("[%s](%s/%s/companies)", lang.title, siteURL, alias),
			count,
			markdown,
		})
	}

	var widths [3]int
	for _, row := range rows {
		for i, cell := range row {
			widths[i] = max(widths[i], utf8.RuneCountInString(cell))
		}
	}

	var (
		sb  strings.Builder
		pad = func(cell string, width int) string {
			return cell + strings.Repeat(" ", width-utf8.RuneCountInString(cell))
		}
	)

	for i, row := range rows {
		sb.WriteString(fmt.Sprintf("| %s | %s | %s |\n", pad(row[0], widths[0]), pad(row[1], widths[1]), pad(row[2], widths[2])))

		if i == 0 {
			sb.WriteString(fmt.Sprintf("|%s|%s|%s|\n", strings.Repeat("-", widths[0]+2), strings.Repeat("-", widths[1]+2), strings.Repeat("-", widths[2]+2)))
		}
	}

	return sb.String()
}

type item struct {
	company domain.CompanyProfile
	id      int64
	latest  int64
}

// buildData mirrors the filtering and sorting (without pinning) of the website company list,
// see Controller.companies and Controller.sortCompanies in internal/organizer/controller.go
func buildData(organizer domain.Organizer, companies []domain.CompanyProfile) *dev.CompaniesMarkdown {
	var items []item

	for _, company := range companies {
		if company.Ignore {
			continue
		}

		id := organizers.CompanyAliasToCodeMap[company.LinkedInProfile.Alias]
		if id == 0 {
			continue
		}

		var (
			profile = company.Languages[organizer.Language]
			show    = len(profile.Vacancies) > 0 ||
				len(profile.TechnologyUsageReferences) > 0 ||
				(organizer.Language == domain.Rust && company.RustFoundationMember)
		)
		if !show {
			continue
		}

		var latest int64
		for _, vacancy := range profile.Vacancies {
			if vacancy.Date.Unix() > latest {
				latest = vacancy.Date.Unix()
			}
		}

		items = append(items, item{company: company, id: id, latest: latest})
	}

	slices.SortFunc(items, func(a, b item) int {
		if a.latest != b.latest {
			if a.latest > b.latest {
				return -1
			}
			return 1
		}

		// ID descending
		if a.id > b.id {
			return -1
		}
		return 1
	})

	rows := make([]*dev.CompanyMarkdownRow, 0, len(items))
	for _, it := range items {
		company := it.company

		row := &dev.CompanyMarkdownRow{
			Name:       company.Name,
			WebsiteURL: company.BaseURL,
			ProfileURL: fmt.Sprintf("%s/%s/companies/%s", siteURL, organizer.Alias, company.LinkedInProfile.Alias),
			CareersURL: company.CareersURL,
			LinkedIn:   fmt.Sprintf("https://www.linkedin.com/company/%s/", company.LinkedInProfile.Alias),
		}
		if company.GitHubProfile.Login != "" {
			row.GitHub = "https://github.com/" + company.GitHubProfile.Login
		}

		rows = append(rows, row)
	}

	return &dev.CompaniesMarkdown{
		LanguageTitle: organizer.Title,
		LanguageAlias: organizer.Alias,
		Total:         len(rows),
		Companies:     rows,
	}
}
