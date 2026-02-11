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
	Date              time.Time
}

type LinkedInProfileShortResponse struct {
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

type CompanyLogoResponse struct {
	MainSize string `json:"72x72"`
}

type LinkedInProfileResponse struct {
	ID        int    `json:"id"`
	IDs       []int  `json:"ids"`
	Alias     string `json:"alias"` // vanity name
	Name      string `json:"name"`
	Employees string `json:"employees"`
	Verified  bool   `json:"verified"`
}

type GitHubProfileResponse struct {
	Login     string `json:"login"`
	Followers string `json:"followers"`
	Verified  bool   `json:"verified"`
}

type GlassdoorProfileResponse struct {
	OverviewURL string `json:"overview_url"`
	ReviewsURL  string `json:"reviews_url"`
	ReviewsRate string `json:"reviews_rate"`
	Verified    bool   `json:"verified"`
}

type CompanyResponse struct {
	ID                        int64                    `json:"id"`
	Type                      CompanyType              `json:"type"`
	Logo                      CompanyLogoResponse      `json:"logo"`
	Name                      string                   `json:"name"`
	BaseURL                   string                   `json:"base_url"`
	CareersURL                string                   `json:"careers_url"`
	AboutURL                  string                   `json:"about_url"`
	BlogURL                   string                   `json:"blog_url"`
	LinkedInProfile           LinkedInProfileResponse  `json:"linkedin_profile"`
	GitHubProfile             GitHubProfileResponse    `json:"github_profile"`
	GlassdoorProfile          GlassdoorProfileResponse `json:"glassdoor_profile"`
	ShortDescription          string                   `json:"short_description"`
	Industries                []Industry               `json:"industries"`
	CloudProviders            []CloudProvider          `json:"cloud_providers"`
	HasEmployeesFromCountries []Country                `json:"has_employees_from_countries"`
	RustFoundationMember      bool                     `json:"rust_foundation_member"`
	PinnedUntil               *time.Time               `json:"pinned_until"`
	Remote                    bool                     `json:"remote"`
	LatestVacancyDate         *time.Time               `json:"latest_vacancy_date"`
	GitHubRepositoryCount     int                      `json:"github_repository_count"`
	Favorite                  bool                     `json:"favorite"`
	Hidden                    bool                     `json:"hidden"`
}

type VacancyCompanyLinkedInProfileResponse struct {
	Alias     string `json:"alias"` // vanity name
	Employees string `json:"employees"`
	Verified  bool   `json:"verified"`
}

type VacancyCompanyGlassdoorProfileResponse struct {
	ReviewsRate string `json:"reviews_rate"`
}

type VacancyCompanyResponse struct {
	ID                        int64                                  `json:"id"`
	Type                      CompanyType                            `json:"type"`
	Logo                      CompanyLogoResponse                    `json:"logo"`
	Name                      string                                 `json:"name"`
	LinkedInProfile           VacancyCompanyLinkedInProfileResponse  `json:"linkedin_profile"`
	GlassdoorProfile          VacancyCompanyGlassdoorProfileResponse `json:"glassdoor_profile"`
	Industries                []Industry                             `json:"industries"`
	HasEmployeesFromCountries []Country                              `json:"has_employees_from_countries"`
	RustFoundationMember      bool                                   `json:"rust_foundation_member"`
}

type UnsafeVacancyResponse struct {
	URL  string    `json:"url"`
	Date time.Time `json:"date"`
}

type CompanyReferenceResponse struct {
	ID int64 `json:"id"`
}

type LocationCountryResponse struct {
	Code string `json:"code"`
}

type LocationResponse struct {
	Raw     string                  `json:"raw"`
	Country LocationCountryResponse `json:"country"`
}

type VacancyResponse struct {
	ID               int64                    `json:"id"`
	Title            string                   `json:"title"`
	ShortDescription string                   `json:"short_description"`
	Location         LocationResponse         `json:"location"`
	Source           VacancySource            `json:"source"`
	CloudProviders   []CloudProvider          `json:"cloud_providers"`
	Remote           bool                     `json:"remote"`
	Date             time.Time                `json:"date"`
	PinnedUntil      *time.Time               `json:"pinned_until"`
	MonthlyViews     int64                    `json:"monthly_views"`
	Company          CompanyReferenceResponse `json:"company"`
	Favorite         bool                     `json:"favorite"`
	Hidden           bool                     `json:"hidden"`
}

type GitHubProfile struct {
	Login     string
	Followers string
	Verified  bool
	Date      time.Time
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
	Date        time.Time
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
	URL                  string
	Location             string
	CloudProviders       []CloudProvider
	Date                 time.Time
	WithSalary           bool
	PinnedUntil          time.Time
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
	GitHubRepositoryCount int
	Vacancies             []Vacancy
	PinnedUntil           time.Time
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
	BaseURL                   string // Production website
	CareersURL                string // Careers page URL
	AboutURL                  string // About URL
	BlogURL                   string // Development blog URL
	ReferralProgramURL        string // Referral program URL
	TransparencyURL           string // Transparency URL e.g. https://buffer.com/transparency
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
	CloudProviders            []CloudProvider
	HasEmployeesFromCountries []Country
	RustFoundationMember      bool // https://foundation.rust-lang.org/members/
	Ignore                    bool
	SyncSources               []CompanySyncSource
	GitHubRepositoryCount     int       // Populated on the fly from the language profile
	Remote                    bool      // Populated on the fly if at least one remote job exists in the language profile
	LatestVacancyDate         time.Time // Populated on the fly from the language profile
	PinnedUntil               time.Time // Populated on the fly from the language profile
}

type UnsafeCompaniesResponse struct {
	Companies []UnsafeCompanyResponse `json:"companies"`
}

type UnsafeVacanciesResponse struct {
	Vacancies []UnsafeVacancyResponse `json:"vacancies"`
}

type CompaniesResponse struct {
	Companies []CompanyResponse `json:"companies"`
}

type VacanciesResponse struct {
	Companies []*VacancyCompanyResponse `json:"companies"`
	Vacancies []*VacancyResponse        `json:"vacancies"`
}
