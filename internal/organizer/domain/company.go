package domain

const (
	Go      int = 0
	Rust    int = 1
	Zig     int = 2
	Scala   int = 3
	Elixir  int = 4
	Clojure int = 5
	Haskell int = 6
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
	Code              struct{} // @TODO extract this as company_code.proto
	Name              string
	URL               string
	LinkedInProfile   LinkedInProfile
	GitHubProfile     GitHubProfile
	GlassdoorProfile  GlassdoorProfile
	OttaProfileSlug   string
	YouTubeChannelURL string
	GoMainLanguage    bool // Golang is the main language
	Vacancies         Vacancies
}
