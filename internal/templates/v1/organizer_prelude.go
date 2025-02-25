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
)

const (
	golangKeywordsTitles = string(domain.GoTitleKeywords)
	keywordsCommon       = `"Developer" OR "Engineer" OR "DevOps"`
)

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

	return "https://www.linkedin.com/search/results/PEOPLE/?" + values.Encode()
}

func linkedinConnectionsFormerEmployeesURL(companies []Company) string {
	companyQueryParam, _ := json.Marshal(companiesToLinkedInIDs(companies))

	values := url.Values{
		"pastCompany": {string(companyQueryParam)},
		"network":     {`["F","S"]`},
		"keywords":    {keywordsCommon},
	}

	return "https://www.linkedin.com/search/results/PEOPLE/?" + values.Encode()
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
		"q": {"site:levels.fyi" + " " + companyName},
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
