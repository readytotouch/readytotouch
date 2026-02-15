package domain

type Industry struct {
	ID    int64  `json:"id"`
	Alias string `json:"alias"`
	Name  string `json:"name"`
}

var (
	IndustryAutomotive = Industry{
		ID:    1,
		Alias: "automotive",
		Name:  "Automotive",
	}
	IndustryCyberSecurity = Industry{
		ID:    2,
		Alias: "cybersecurity",
		Name:  "CyberSecurity",
	}
	IndustryEdTech = Industry{
		ID:    3,
		Alias: "edtech",
		Name:  "EdTech",
	}
	IndustryECommerce = Industry{
		ID:    4,
		Alias: "e-commerce",
		Name:  "eCommerce",
	}
	IndustryHealthTech = Industry{
		ID:    5,
		Alias: "healthtech",
		Name:  "HealthTech",
	}
	IndustryMedTech = Industry{
		ID:    6,
		Alias: "medtech",
		Name:  "MedTech",
	}
	IndustryFinTech = Industry{
		ID:    7,
		Alias: "fintech",
		Name:  "FinTech",
	}
	IndustryGameDev = Industry{
		ID:    8,
		Alias: "gamedev",
		Name:  "GameDev",
	}
	IndustryIoT = Industry{
		ID:    9,
		Alias: "iot",
		Name:  "IoT",
	}
	IndustryAdTech = Industry{
		ID:    10,
		Alias: "adtech",
		Name:  "AdTech",
	}
	IndustryAgroTech = Industry{
		ID:    11,
		Alias: "agrotech",
		Name:  "AgroTech",
	}
	IndustryMarTech = Industry{
		ID:    12,
		Alias: "martech",
		Name:  "MarTech",
	}
	IndustryDevOps = Industry{
		ID:    13,
		Alias: "devops",
		Name:  "DevOps",
	}
	IndustryCloudComputing = Industry{
		ID:    14,
		Alias: "cloud-computing",
		Name:  "Cloud Computing",
	}
	IndustryBigData = Industry{
		ID:    15,
		Alias: "big-data",
		Name:  "Big Data",
	}
	IndustrySocialMedia = Industry{
		ID:    16,
		Alias: "social-media",
		Name:  "Social Media",
	}
	IndustryEntertainment = Industry{
		ID:    17,
		Alias: "entertainment",
		Name:  "Entertainment",
	}
	IndustryPropTech = Industry{
		ID:    18,
		Alias: "proptech",
		Name:  "PropTech",
	}
	IndustryInsurTech = Industry{
		ID:    19,
		Alias: "insurtech",
		Name:  "InsurTech",
	}
	IndustryLogisticsTech = Industry{
		ID:    20,
		Alias: "logisticstech",
		Name:  "LogisticsTech",
	}
	IndustryTelecom = Industry{
		ID:    21,
		Alias: "telecom",
		Name:  "Telecom",
	}
	IndustryDefenseTech = Industry{
		ID:    22,
		Alias: "defensetech",
		Name:  "DefenseTech",
	}
	IndustryTravelTech = Industry{
		ID:    23,
		Alias: "traveltech",
		Name:  "TravelTech",
	}
	IndustryCryptoCurrency = Industry{
		ID:    24,
		Alias: "cryptocurrency",
		Name:  "Cryptocurrency",
	}
	hiddenIndustries = []Industry{
		IndustryCryptoCurrency,
	}
)

func FindHiddenIndustryByAlias(alias string) (Industry, bool) {
	for _, industry := range hiddenIndustries {
		if industry.Alias == alias {
			return industry, true
		}
	}

	return Industry{}, false
}
