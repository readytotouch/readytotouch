package domain

type (
	Language    int
	CompanyType string
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
	CompanyTypeProduct CompanyType = "product"
	CompanyTypeStartup CompanyType = "startup"
)

type Vacancies = [7][]string

type LinkedInProfile struct {
	ID    int
	IDs   []int
	Alias string
	Name  string
}

type GitHubProfile struct {
	Login             string
	GoRepositoryCount int
}

type GlassdoorProfile struct {
	OverviewURL string
	ReviewsURL  string
}

type Company struct {
	ID                int64       // populates from the CompanyAliasMap
	Type              CompanyType // populates from the CompanyTypeMap
	Name              string
	URL               string
	LinkedInProfile   LinkedInProfile
	GitHubProfile     GitHubProfile
	GlassdoorProfile  GlassdoorProfile
	OttaProfileSlug   string
	YouTubeChannelURL string
	GoMainLanguage    bool // Golang is the main language
	Vacancies         Vacancies
	ShortDescription  string // Only for understanding what the company does
}
