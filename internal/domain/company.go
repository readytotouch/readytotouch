package domain

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
