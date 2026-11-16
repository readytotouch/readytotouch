package dev

import (
	"fmt"
	"net/url"
	"strings"
)

type CompanyCodePair struct {
	ID    int64
	Name  string
	Alias string
}

type VacancyCodePair struct {
	ID           int64
	URL          string
	CompanyAlias string
}

type CompanyLogoPair struct {
	Alias string
	Logo  string
}

func googleSearchLogos(companyName string) string {
	values := url.Values{
		"q": {fmt.Sprintf("%q review-logos", companyName)},
	}

	return "https://www.google.com/search?" + values.Encode()
}

type CompanyMarkdownRow struct {
	Name       string
	WebsiteURL string
	ProfileURL string
	CareersURL string
	LinkedIn   string
	GitHub     string
}

type CompaniesMarkdown struct {
	LanguageTitle string
	LanguageAlias string
	Total         int
	Companies     []*CompanyMarkdownRow
}

func markdownCell(s string) string {
	return strings.NewReplacer("|", "\\|", "[", "\\[", "]", "\\]").Replace(s)
}

func markdownLink(title, url string) string {
	if url == "" {
		return ""
	}

	return "[" + title + "](" + url + ")"
}
