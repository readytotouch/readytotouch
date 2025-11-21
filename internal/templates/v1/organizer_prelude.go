package v1

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
)

type (
	Company         = domain.CompanyProfile
	University      = domain.University
	PreparedVacancy = domain.PreparedVacancy

	OrganizerState struct {
		Organizer domain.Organizer
		Available bool
	}
)

const (
	golangKeywordsTitles = string(domain.GoTitleKeywords)
	keywordsCommon       = `"Developer" OR "Engineer" OR "DevOps"`
)

var (
	showLocationMap = map[string]bool{
		"United States":                            true, // 62
		"Bengaluru, Karnataka, India":              true, // 57
		"London, England, United Kingdom":          true, // 49
		"Berlin, Berlin, Germany":                  true, // 40
		"Ho Chi Minh City, Vietnam":                true, // 23
		"Warsaw, Mazowieckie, Poland":              true, // 23
		"Cracow, Małopolskie, Poland":              true, // 23
		"London Area, United Kingdom":              true, // 21
		"Poland":                                   true, // 19
		"United Kingdom":                           true, // 19
		"Prague, Prague, Czechia":                  true, // 18
		"Paris, Île-de-France, France":             true, // 18
		"Pune, Maharashtra, India":                 true, // 18
		"Tel Aviv-Yafo, Tel Aviv District, Israel": true, // 18
		"Spain":                                     true, // 17
		"New York, NY":                              true, // 15
		"Toronto, ON":                               true, // 14
		"Barcelona, Catalonia, Spain":               true, // 13
		"Germany":                                   true, // 13
		"Chicago, IL":                               true, // 13
		"Canada":                                    true, // 11
		"São Paulo, São Paulo, Brazil":              true, // 11
		"Hyderabad, Telangana, India":               true, // 10
		"San Francisco, CA":                         true, // 9
		"Czechia":                                   true, // 9
		"Richmond, VA":                              true, // 9
		"Hamburg, Hamburg, Germany":                 true, // 9
		"Budapest, Budapest, Hungary":               true, // 9
		"Sofia, Sofia City, Bulgaria":               true, // 9
		"Sydney, New South Wales, Australia":        true, // 9
		"Plano, TX":                                 true, // 9
		"San Jose, CA":                              true, // 8
		"Munich, Bavaria, Germany":                  true, // 8
		"Madrid, Community of Madrid, Spain":        true, // 8
		"Belgrade, Serbia":                          true, // 8
		"Montreal, QC":                              true, // 8
		"Amsterdam, North Holland, Netherlands":     true, // 8
		"Brno, South Moravia, Czechia":              true, // 8
		"Ho Chi Minh City Metropolitan Area":        true, // 7
		"Taipei, Taipei City, Taiwan":               true, // 7
		"Istanbul, Türkiye":                         true, // 7
		"Vienna, Vienna, Austria":                   true, // 7
		"Palo Alto, CA":                             true, // 7
		"Austin, TX":                                true, // 7
		"Berlin, Germany":                           true, // 7
		"Singapore, Singapore":                      true, // 7
		"Tallinn, Harjumaa, Estonia":                true, // 7
		"Gurugram, Haryana, India":                  true, // 7
		"Stockholm, Stockholm County, Sweden":       true, // 7
		"India":                                     true, // 7
		"Chennai, Tamil Nadu, India":                true, // 7
		"McLean, VA":                                true, // 6
		"Slovakia":                                  true, // 6
		"Milan, Lombardy, Italy":                    true, // 6
		"Brazil":                                    true, // 6
		"Istanbul, Istanbul, Türkiye":               true, // 6
		"Seattle, WA":                               true, // 6
		"Lyon, Auvergne-Rhône-Alpes, France":        true, // 6
		"Vilnius, Vilniaus, Lithuania":              true, // 6
		"Dublin, County Dublin, Ireland":            true, // 6
		"Jersey City, NJ":                           true, // 5
		"Leicester, England, United Kingdom":        true, // 5
		"Atlanta, GA":                               true, // 5
		"Noida, Uttar Pradesh, India":               true, // 5
		"Philippines":                               true, // 5
		"Türkiye":                                   true, // 5
		"Ottawa, ON":                                true, // 5
		"Wrocław, Dolnośląskie, Poland":             true, // 5
		"Belfast, Northern Ireland, United Kingdom": true, // 5
		"EMEA":                                true, // 5
		"New York, United States":             true, // 4
		"Fort Worth, TX":                      true, // 4
		"Dhaka, Dhaka, Bangladesh":            true, // 4
		"Cyprus":                              true, // 4
		"Los Angeles, CA":                     true, // 4
		"Mexico City, Mexico":                 true, // 4
		"Ukraine":                             true, // 4
		"Bengaluru East, Karnataka, India":    true, // 4
		"Herzliya, Tel Aviv District, Israel": true, // 4
		"Valbonne, Provence-Alpes-Côte d'Azur, France": true, // 4
		"Utrecht, Utrecht, Netherlands":                true, // 4
		"European Union":                               true, // 4
		"Mumbai, Maharashtra, India":                   true, // 4
		"Helsinki, Uusimaa, Finland":                   true, // 4
		"Singapore":                                    true, // 4
		"Ireland":                                      true, // 4
		"Greater Colorado Springs Area":                true, // 4
		"Dedham, MA":                                   true, // 4
		"Jacksonville, FL":                             true, // 4
		"Boulder, CO":                                  true, // 4
		"Thiruvananthapuram, Kerala, India":            true, // 4
		"Lithuania":                                    true, // 3
		"Sweden":                                       true, // 3
		"Alpharetta, GA":                               true, // 3
		"Delhi, Delhi, India":                          true, // 3
		"Netanya, Center District, Israel":             true, // 3
		"Sofia Metropolitan Area":                      true, // 3
		"Finland":                                      true, // 3
		"Boston, MA":                                   true, // 3
		"Greater Barcelona Metropolitan Area":          true, // 3
		"Bristol, England, United Kingdom":             true, // 3
		"Phoenix, AZ":                                  true, // 3
		"Petaling Jaya, Selangor, Malaysia":            true, // 3
		"Oslo, Oslo, Norway":                           true, // 3
		"Greater Munich Metropolitan Area":             true, // 3
		"Kaunas, Kaunas, Lithuania":                    true, // 3
		"Porto, Porto, Portugal":                       true, // 3
		"Georgia":                                      true, // 3
		"Mulhuddart, Fingal, Ireland":                  true, // 3
		"Tel Aviv District, Israel":                    true, // 3
		"Glendale, CA":                                 true, // 3
		"New York City Metropolitan Area":              true, // 3
		"Bellevue, WA":                                 true, // 3
		"Portugal":                                     true, // 3
		"Estonia":                                      true, // 3
		"Argentina":                                    true, // 3
		"Mumbai Metropolitan Region":                   true, // 3
		"Tokyo, Japan":                                 true, // 3
		"Manchester, England, United Kingdom":          true, // 3
		"Costa Mesa, CA":                               true, // 3
		"Cary, NC":                                     true, // 3
	}
)

func companyOrganizers(company Company) []OrganizerState {
	var (
		source = []domain.Organizer{
			domain.OrganizerGolang,
			domain.OrganizerRust,
			domain.OrganizerScala,
			domain.OrganizerElixir,
			domain.OrganizerErlang,
			domain.OrganizerClojure,
			domain.OrganizerZig,
			domain.OrganizerHaskell,
			domain.OrganizerFSharp,
			domain.OrganizerOCaml,
		}
		result = make([]OrganizerState, len(source))
	)

	for i, organizer := range source {
		// Show companies only if they have vacancies or are Rust Foundation members
		result[i] = OrganizerState{
			Organizer: organizer,
			Available: len(company.Languages[organizer.Language].Vacancies) > 0 || (organizer.Language == domain.Rust && company.RustFoundationMember),
		}
	}

	return result
}

func showLocation(s string) bool {
	if s == "" {
		return false
	}

	return showLocationMap[s]
}

func linkedinConnectionsURL(companies []Company, universities []University) string {
	companyQueryParam, _ := json.Marshal(companiesToLinkedInIDs(companies))

	values := url.Values{
		"currentCompany": {string(companyQueryParam)},
		"network":        {`["F","S"]`},
		"keywords":       {keywordsCommon},
	}

	if len(universities) > 0 {
		schoolQueryParam, _ := json.Marshal(universitiesToLinkedInIDs(universities))

		values["schoolFilter"] = []string{string(schoolQueryParam)}
	}

	return "https://www.linkedin.com/search/results/people/?" + values.Encode()
}

func linkedinConnectionsFormerEmployeesURL(companies []Company) string {
	companyQueryParam, _ := json.Marshal(companiesToLinkedInIDs(companies))

	values := url.Values{
		"pastCompany": {string(companyQueryParam)},
		"network":     {`["F","S"]`},
		"keywords":    {keywordsCommon},
	}

	return "https://www.linkedin.com/search/results/people/?" + values.Encode()
}

func linkedinConnectionsCEO(company Company) string {
	companyQueryParam, _ := json.Marshal(companiesToLinkedInIDs([]Company{company}))

	values := url.Values{
		"currentCompany": {string(companyQueryParam)},
		"network":        {`["F","S","O"]`},
		"keywords":       {`"CEO" OR "Chief Executive Officer" OR "Founder" OR "Co-Founder"`},
	}

	return "https://www.linkedin.com/search/results/people/?" + values.Encode()
}

func linkedinEmployeesPostsURL(companies []Company, languageTitle string) string {
	companyQueryParam, _ := json.Marshal(companiesToLinkedInIDs(companies))

	values := url.Values{
		"authorCompany": {string(companyQueryParam)},
		"datePosted":    {`"past-month"`},
		"sortBy":        {"RELEVANT", `"date_posted"`},
		"keywords":      {`"Hiring" OR ` + `"` + languageTitle + `"`}, // "Hiring" OR "Golang"
	}

	if len(companies) > 0 {
		values["f_C"] = []string{strings.Join(companiesToLinkedInIDs(companies), ",")}
	}

	return "https://www.linkedin.com/search/results/content/?" + values.Encode()
}

func linkedinJobsURL(companies []Company, keywords string) string {
	values := url.Values{
		"keywords": {keywords},
		"location": {"Worldwide"},
		"geoId":    {"92000000"}, // Worldwide
		"sortBy":   {"DD"},       // order by "Most recent
		"f_TPR":    {"r2592000"}, // filter "Past month"
		// Remote
		// f_WT => 2
	}

	if len(companies) > 0 {
		values["f_C"] = []string{strings.Join(companiesToLinkedInIDs(companies), ",")}
	}

	return "https://www.linkedin.com/jobs/search/?" + values.Encode()
}

func companiesToLinkedInIDs(companies []Company) []string {
	ids := make([]string, 0, len(companies)*2)
	for _, company := range companies {
		ids = appendLinkedInProfileIDs(ids, company.LinkedInProfile)
	}
	return ids
}

func universitiesToLinkedInIDs(universities []University) []string {
	ids := make([]string, 0, len(universities)*2)
	for _, university := range universities {
		ids = appendLinkedInProfileIDs(ids, university.LinkedInProfile)
	}
	return ids
}

func appendLinkedInProfileIDs(ids []string, profile domain.LinkedInProfile) []string {
	if profile.ID > 0 {
		ids = append(ids, strconv.FormatInt(int64(profile.ID), 10))
	}

	for _, id := range profile.IDs {
		ids = append(ids, strconv.FormatInt(int64(id), 10))
	}

	return ids
}

func similarwebURL(s string) string {
	// https://www.similarweb.com/website/readytotouch.com/

	if s == "" {
		return ""
	}

	parsedURL, err := url.Parse(s)
	if err != nil {
		// panic allowed because HTML is pre-generated

		panic(err)
	}

	return fmt.Sprintf("https://www.similarweb.com/website/%s/", parsedURL.Hostname())
}

func whoisURL(s string) string {
	// https://www.whois.com/whois/readytotouch.com

	if s == "" {
		return ""
	}

	parsedURL, err := url.Parse(s)
	if err != nil {
		// panic allowed because HTML is pre-generated

		panic(err)
	}

	return fmt.Sprintf("https://www.whois.com/whois/%s", parsedURL.Hostname())
}

func hostname(s string) string {
	if s == "" {
		return ""
	}

	parsedURL, err := url.Parse(s)
	if err != nil {
		// panic allowed because HTML is pre-generated

		panic(err)
	}

	return parsedURL.Hostname()
}

func googleSearch(q string) string {
	values := url.Values{
		"q": {q},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchLayoffsMonth(companyName string) string {
	values := url.Values{
		"q":   {companyName + " " + "layoffs"},
		"tbs": {"qdr:m"}, // month
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchLayoffsYear(companyName string) string {
	values := url.Values{
		"q":   {companyName + " " + "layoffs"},
		"tbs": {"qdr:y"}, // year
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchOnSite(site string, what string) string {
	values := url.Values{
		"q": {"site:" + site + " " + what},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchGitHub(companyName string) string {
	values := url.Values{
		"q": {"site:github.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchGlassdoor(companyName string) string {
	values := url.Values{
		"q": {"site:glassdoor.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchBlind(companyName string) string {
	values := url.Values{
		"q": {"site:teamblind.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchLevelsFyi(companyName string) string {
	values := url.Values{
		"q": {"site:levels.fyi/companies" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchIndeed(companyName string) string {
	values := url.Values{
		"q": {"site:www.indeed.com/cmp/" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchDealroom(companyName string) string {
	values := url.Values{
		"q": {"site:dealroom.co" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchCrunchbase(companyName string) string {
	values := url.Values{
		"q": {"site:crunchbase.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchPitchbook(companyName string) string {
	values := url.Values{
		"q": {"site:pitchbook.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchYahooFinance(companyName string) string {
	values := url.Values{
		"q": {"site:finance.yahoo.com/quote/" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchGoogleFinance(companyName string) string {
	values := url.Values{
		"q": {"site:google.com/finance/quote/" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchXing(companyName string) string {
	values := url.Values{
		"q": {"site:xing.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchOtta(companyName string) string {
	values := url.Values{
		"q": {"site:otta.com" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func googleSearchYCombinator(companyName string) string {
	values := url.Values{
		"q": {"site:ycombinator.com/companies/" + " " + companyName},
	}

	return "https://www.google.com/search?" + values.Encode()
}

func latestVacancies(vacancies []PreparedVacancy) []PreparedVacancy {
	if len(vacancies) >= 256 {
		return vacancies[:256]
	}

	return vacancies
}
