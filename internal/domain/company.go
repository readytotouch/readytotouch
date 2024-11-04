package domain

import "time"

type (
	Language              int
	LanguageTitleKeywords string
	CompanyType           string
)

const (
	Go      Language = 0
	Rust    Language = 1
	Zig     Language = 2
	Scala   Language = 3
	Elixir  Language = 4
	Clojure Language = 5
	Haskell Language = 6
)

const (
	GoTitleKeywords      LanguageTitleKeywords = `"Golang Engineer" OR "Golang Software Engineer" OR "Golang Developer" OR "Go Engineer" OR "Go Software Engineer" OR "Go Developer"`
	RustTitleKeywords    LanguageTitleKeywords = `"Rust Engineer" OR "Rust Software Engineer" OR "Rust Developer"`
	ZigTitleKeywords     LanguageTitleKeywords = `"Zig Engineer" OR "Zig Software Engineer" OR "Zig Developer"`
	ScalaTitleKeywords   LanguageTitleKeywords = `"Scala Engineer" OR "Scala Software Engineer" OR "Scala Developer"`
	ElixirTitleKeywords  LanguageTitleKeywords = `"Elixir Engineer" OR "Elixir Software Engineer" OR "Elixir Developer" OR "Erlang Engineer" OR "Erlang Software Engineer" OR "Erlang Developer"`
	ClojureTitleKeywords LanguageTitleKeywords = `"Clojure Engineer" OR "Clojure Software Engineer" OR "Clojure Developer"`
	HaskellTitleKeywords LanguageTitleKeywords = `"Haskell Engineer" OR "Haskell Software Engineer" OR "Haskell Developer"`
)

const (
	CompanyTypeProduct CompanyType = "product"
	CompanyTypeStartup CompanyType = "startup"
)

type Vacancies = [7][]string
type Languages = [7]LanguageProfile

type LinkedInProfile struct {
	ID                int
	IDs               []int
	Alias             string
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

type GitHubProfile struct {
	Login             string
	GoRepositoryCount int
	Verified          bool
}

type GlassdoorProfile struct {
	OverviewURL string
	ReviewsURL  string
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
	Title            string
	ShortDescription string // proof that the vacancy is for a particular technology
	URL              string
	Date             time.Time
}

type LanguageProfile struct {
	GitHubRepositoriesCount int
	Vacancies               []Vacancy
}

type CompanyProfile struct {
	ID                        int64       // populates from the CompanyAliasMap
	Type                      CompanyType // populates from the CompanyTypeMap
	Name                      string
	URL                       string // Deprecated
	Website                   string // Production website
	Careers                   string // Careers page URL
	About                     string // About URL
	Blog                      string // Development blog URL
	LinkedInProfile           LinkedInProfile
	GitHubProfile             GitHubProfile
	BlindProfile              BlindProfile
	LevelsFyiProfile          LevelsFyiProfile
	GlassdoorProfile          GlassdoorProfile
	IndeedProfile             IndeedProfile
	OttaProfileSlug           string
	YouTubeChannelURL         string
	GoMainLanguage            bool      // Golang is the main language
	Vacancies                 Vacancies // Deprecated
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
}

type UnsafeCompaniesResponse struct {
	Companies []LinkedInProfileResponse `json:"companies"`
}
