package domain

import "time"

type (
	Language              int
	LanguageTitleKeywords string
	CompanyType           string
	CompanySyncSource     int
)

const (
	Go      Language = 0
	Rust    Language = 1
	Zig     Language = 2
	Scala   Language = 3
	Elixir  Language = 4
	Clojure Language = 5
	Haskell Language = 6
	Erlang  Language = 7
	FSharp  Language = 8
	OCaml   Language = 9
	Gleam   Language = 10
	Mojo    Language = 11
)

func (l Language) String() string {
	switch l {
	case Go:
		return "Go"
	case Rust:
		return "Rust"
	case Zig:
		return "Zig"
	case Scala:
		return "Scala"
	case Elixir:
		return "Elixir"
	case Clojure:
		return "Clojure"
	case Haskell:
		return "Haskell"
	case Erlang:
		return "Erlang"
	case FSharp:
		return "F#"
	case OCaml:
		return "OCaml"
	case Gleam:
		return "Gleam"
	case Mojo:
		return "Mojo"
	default:
		return ""
	}
}

const (
	GoTitleKeywords      LanguageTitleKeywords = `"Golang Engineer" OR "Golang Software Engineer" OR "Golang Developer" OR "Go Engineer" OR "Go Software Engineer" OR "Go Developer"`
	RustTitleKeywords    LanguageTitleKeywords = `"Rust Engineer" OR "Rust Software Engineer" OR "Rust Developer"`
	ZigTitleKeywords     LanguageTitleKeywords = `"Zig Engineer" OR "Zig Software Engineer" OR "Zig Developer"`
	ScalaTitleKeywords   LanguageTitleKeywords = `"Scala Engineer" OR "Scala Software Engineer" OR "Scala Developer"`
	ElixirTitleKeywords  LanguageTitleKeywords = `"Elixir Engineer" OR "Elixir Software Engineer" OR "Elixir Developer" OR "Erlang Engineer" OR "Erlang Software Engineer" OR "Erlang Developer"`
	ClojureTitleKeywords LanguageTitleKeywords = `"Clojure Engineer" OR "Clojure Software Engineer" OR "Clojure Developer"`
	HaskellTitleKeywords LanguageTitleKeywords = `"Haskell Engineer" OR "Haskell Software Engineer" OR "Haskell Developer"`
	ErlangTitleKeywords  LanguageTitleKeywords = `"Erlang Engineer" OR "Erlang Software Engineer" OR "Erlang Developer"`
	FSharpTitleKeywords  LanguageTitleKeywords = `"F# Engineer" OR "F# Software Engineer" OR "F# Developer"`
	OCamlTitleKeywords   LanguageTitleKeywords = `"OCaml Engineer" OR "OCaml Software Engineer" OR "OCaml Developer"`
	GleamTitleKeywords   LanguageTitleKeywords = `"Gleam Engineer" OR "Gleam Software Engineer" OR "Gleam Developer"`
	MojoTitleKeywords    LanguageTitleKeywords = `"Mojo Engineer" OR "Mojo Software Engineer" OR "Mojo Developer"`
)

const (
	CompanyTypeProduct   CompanyType = "product"
	CompanyTypeStartup   CompanyType = "startup"
	CompanyTypeOutsource CompanyType = "outsource"
)

const (
	RustCompanies CompanySyncSource = 1 // https://github.com/omarabid/rust-companies
)

type Languages = [12]LanguageProfile

type LinkedInProfile struct {
	ID                int
	IDs               []int
	Alias             string
	PreviousAliases   []string
	Name              string
	Followers         string
	Employees         string
	AssociatedMembers string
	Verified          bool
}

type LinkedInProfileResponse struct {
	ID    int64  `json:"id"`
	Alias string `json:"alias"`
	Name  string `json:"name"`
}

type UnsafeCompanyLanguageStats struct {
	Language       string    `json:"language"`
	MaxVacancyDate time.Time `json:"max_vacancy_date"`
}

type UnsafeCompanyResponse struct {
	ID                 int                          `json:"id"`
	Alias              string                       `json:"alias"`
	Name               string                       `json:"name"`
	GitHubProfileAlias string                       `json:"github_profile_alias"`
	OttaProfileAlias   string                       `json:"otta_profile_alias"`
	Ignore             bool                         `json:"ignore"`
	Languages          []UnsafeCompanyLanguageStats `json:"languages"`
}

type UnsafeVacancyResponse struct {
	URL  string    `json:"url"`
	Date time.Time `json:"date"`
}

type GitHubProfile struct {
	Login     string
	Followers string
	Verified  bool
}

type GlassdoorProfile struct {
	OverviewURL string
	ReviewsURL  string
	JobsURL     string
	Jobs        string // will be mostly outdated
	Reviews     string
	Salaries    string
	ReviewsRate string
	Verified    bool
}

type BlindProfile struct {
	Alias       string
	Employees   string
	Salary      string
	Reviews     string
	ReviewsRate string
}

type LevelsFyiProfile struct {
	Alias     string
	Employees string
}

type IndeedProfile struct {
	Alias string // for now, will be more in the future
}

type Vacancy struct {
	Title                string
	SubTitle             string // @TODO, for future use, design
	ShortDescription     string // proof that the vacancy is for a particular technology
	SwitchingOpportunity string // an opportunity to apply without knowing the language, but with a willingness to learn it
	Location             string
	URL                  string
	Date                 time.Time
	WithSalary           bool
	Remote               bool // Fully remote available
	// @TODO: 401(k) match + equity
	// @TODO: team "Payments & Checkout"
}

type PreparedCompany struct {
	ID                        int64       // populates from the CompanyAliasToCodeMap
	Type                      CompanyType // populates from the CompanyStartupMap
	Logo                      CompanyLogo // populates from the CompanyAliasToLogoMap
	Name                      string
	Alias                     string // LinkedIn alias
	Industries                []Industry
	HasEmployeesFromCountries []Country
	RustFoundationMember      bool // https://foundation.rust-lang.org/members/
}

type PreparedVacancy struct {
	ID                   int64 // populates from the VacancyAliasMap
	Company              PreparedCompany
	Title                string
	ShortDescription     string // proof that the vacancy is for a particular technology
	SwitchingOpportunity string // proof that the vacancy is for a particular technology
	Location             string
	URL                  string
	Date                 time.Time
	WithSalary           bool // for future use
	Remote               bool // for future use
}

type LanguageProfile struct {
	GitHubRepositoriesCount int
	Vacancies               []Vacancy
}

type CompanyLogo struct {
	V0 string
	V1 string
	V2 string
}

type CompanyProfile struct {
	ID                        int64       // populates from the CompanyAliasToCodeMap
	Type                      CompanyType // populates from the CompanyStartupMap
	Logo                      CompanyLogo // populates from the CompanyAliasToLogoMap
	Name                      string
	Website                   string // Production website
	Careers                   string // Careers page URL
	ReferralProgram           string // Referral program URL
	About                     string // About URL
	Blog                      string // Development blog URL
	Transparency              string // Transparency URL e.g. https://buffer.com/transparency
	LinkedInProfile           LinkedInProfile
	GitHubProfile             GitHubProfile
	BlindProfile              BlindProfile
	LevelsFyiProfile          LevelsFyiProfile
	GlassdoorProfile          GlassdoorProfile
	IndeedProfile             IndeedProfile
	OttaProfileSlug           string
	YouTubeChannelURL         string
	GoMainLanguage            bool // Golang is the main language
	Languages                 Languages
	ShortDescription          string // Only for understanding what the company does
	DealroomURL               string // Investors, funding, etc.
	CrunchbaseURL             string // Investors, funding, etc.
	PitchbookURL              string // Investors, funding, etc.
	YahooFinanceURL           string // Market cap, etc.
	GoogleFinanceURL          string // Market cap, etc.
	YCombinatorURL            string // YC profile
	Industries                []Industry
	HasEmployeesFromCountries []Country
	RustFoundationMember      bool // https://foundation.rust-lang.org/members/
	Ignore                    bool
	SyncSources               []CompanySyncSource
	PinnedUntil               time.Time
	Remote                    bool      // Mark "remote" if the company has at least one remote job
	LatestVacancyDate         time.Time // Sort companies by latest vacancy date per technology
}

type UnsafeCompaniesResponse struct {
	Companies []UnsafeCompanyResponse `json:"companies"`
}

type UnsafeVacanciesResponse struct {
	Vacancies []UnsafeVacancyResponse `json:"vacancies"`
}
