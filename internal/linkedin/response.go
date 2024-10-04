package linkedin

type CompaniesResponse struct {
	Elements []CompanyResponse `json:"elements"`
}

type PreferredLocale struct {
	Country  string `json:"country"`
	Language string `json:"language"`
}

type Name struct {
	Localized       map[string]string `json:"localized"`
	PreferredLocale PreferredLocale   `json:"preferredLocale"`
}

type CompanyResponse struct {
	ID         int64  `json:"id"`
	VanityName string `json:"vanityName"`
	Name       Name   `json:"name"`
}

type Company struct {
	ID         int64
	VanityName string
	Name       string
}
