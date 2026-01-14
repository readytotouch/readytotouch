package organizers

var (
	LocationCodeMap = map[string]string{
		"United States":                                "us", // 78
		"Bengaluru, Karnataka, India":                  "in", // 65
		"London, England, United Kingdom":              "gb", // 60
		"Berlin, Berlin, Germany":                      "de", // 49
		"Cracow, Małopolskie, Poland":                  "pl", // 29
		"Warsaw, Mazowieckie, Poland":                  "pl", // 27
		"Tel Aviv-Yafo, Tel Aviv District, Israel":     "il", // 25
		"London Area, United Kingdom":                  "gb", // 24
		"Prague, Prague, Czechia":                      "cz", // 24
		"Ho Chi Minh City, Vietnam":                    "vn", // 24
		"United Kingdom":                               "gb", // 23
		"Paris, Île-de-France, France":                 "fr", // 22
		"Poland":                                       "pl", // 22
		"Barcelona, Catalonia, Spain":                  "es", // 20
		"New York, NY":                                 "us", // 20
		"Pune, Maharashtra, India":                     "in", // 19
		"Germany":                                      "de", // 18
		"Spain":                                        "es", // 18
		"Toronto, ON":                                  "ca", // 15
		"Chicago, IL":                                  "us", // 15
		"Canada":                                       "ca", // 14
		"Hamburg, Hamburg, Germany":                    "de", // 12
		"Hyderabad, Telangana, India":                  "in", // 12
		"São Paulo, São Paulo, Brazil":                 "br", // 12
		"Budapest, Budapest, Hungary":                  "hu", // 11
		"Sofia, Sofia City, Bulgaria":                  "bg", // 11
		"Munich, Bavaria, Germany":                     "de", // 11
		"Sydney, New South Wales, Australia":           "au", // 11
		"Berlin, Germany":                              "de", // 10
		"Chennai, Tamil Nadu, India":                   "in", // 10
		"Amsterdam, North Holland, Netherlands":        "nl", // 10
		"Montreal, QC":                                 "ca", // 10
		"Austin, TX":                                   "us", // 9
		"Stockholm, Stockholm County, Sweden":          "se", // 9
		"Tallinn, Harjumaa, Estonia":                   "ee", // 9
		"Singapore, Singapore":                         "sg", // 9
		"Ho Chi Minh City Metropolitan Area":           "vn", // 9
		"Brno, South Moravia, Czechia":                 "cz", // 9
		"Madrid, Community of Madrid, Spain":           "es", // 9
		"Plano, TX":                                    "us", // 9
		"San Francisco, CA":                            "us", // 9
		"Dublin, County Dublin, Ireland":               "ie", // 9
		"Czechia":                                      "cz", // 9
		"San Jose, CA":                                 "us", // 9
		"Richmond, VA":                                 "us", // 9
		"Wrocław, Dolnośląskie, Poland":                "pl", // 8
		"Seattle, WA":                                  "us", // 8
		"Belgrade, Serbia":                             "rs", // 8
		"Milan, Lombardy, Italy":                       "it", // 8
		"Vienna, Vienna, Austria":                      "at", // 8
		"Gurugram, Haryana, India":                     "in", // 8
		"Palo Alto, CA":                                "us", // 7
		"Brazil":                                       "br", // 7
		"Vilnius, Vilniaus, Lithuania":                 "lt", // 7
		"Istanbul, Türkiye":                            "tr", // 7
		"Taipei, Taipei City, Taiwan":                  "tw", // 7
		"India":                                        "in", // 7
		"Atlanta, GA":                                  "us", // 7
		"Istanbul, Istanbul, Türkiye":                  "tr", // 7
		"European Union":                               "",   // 7
		"Bengaluru East, Karnataka, India":             "in", // 6
		"Herzliya, Tel Aviv District, Israel":          "il", // 6
		"Lyon, Auvergne-Rhône-Alpes, France":           "fr", // 6
		"Philippines":                                  "ph", // 6
		"Cyprus":                                       "cy", // 6
		"McLean, VA":                                   "us", // 6
		"Noida, Uttar Pradesh, India":                  "in", // 6
		"Slovakia":                                     "sk", // 6
		"Leicester, England, United Kingdom":           "gb", // 6
		"Belfast, Northern Ireland, United Kingdom":    "gb", // 6
		"Jersey City, NJ":                              "us", // 6
		"Thiruvananthapuram, Kerala, India":            "in", // 5
		"Ottawa, ON":                                   "ca", // 5
		"Türkiye":                                      "tr", // 5
		"Dedham, MA":                                   "us", // 5
		"Helsinki, Uusimaa, Finland":                   "fi", // 5
		"EMEA":                                         "",   // 5
		"Utrecht, Utrecht, Netherlands":                "nl", // 5
		"Los Angeles, CA":                              "us", // 5
		"Dhaka, Dhaka, Bangladesh":                     "bd", // 5
		"Mumbai, Maharashtra, India":                   "in", // 5
		"Tel Aviv District, Israel":                    "il", // 5
		"New York, United States":                      "us", // 4
		"Fort Worth, TX":                               "us", // 4
		"Portugal":                                     "pt", // 4
		"New York City Metropolitan Area":              "us", // 4
		"Mexico City, Mexico":                          "mx", // 4
		"Jacksonville, FL":                             "us", // 4
		"Valbonne, Provence-Alpes-Côte d'Azur, France": "fr", // 4
		"Lisbon, Portugal":                             "pt", // 4
		"Singapore":                                    "sg", // 4
		"Boulder, CO":                                  "us", // 4
		"Mulhuddart, Fingal, Ireland":                  "ie", // 4
		"Mumbai Metropolitan Region":                   "in", // 4
		"Kaunas, Kaunas, Lithuania":                    "lt", // 4
		"Greater Colorado Springs Area":                "us", // 4
		"Ireland":                                      "ie", // 4
		"Prague, Czechia":                              "cz", // 4
		"Ukraine":                                      "ua", // 4
		"Delhi, Delhi, India":                          "in", // 4
		"Sweden":                                       "se", // 3
		"Valenzano, Apulia, Italy":                     "it", // 3
		"Glendale, CA":                                 "us", // 3
		"Lithuania":                                    "lt", // 3
		"Manchester, England, United Kingdom":          "gb", // 3
		"Rožňava, Kosice, Slovakia":                    "sk", // 3
		"San Francisco Bay Area":                       "us", // 3
		"Australia":                                    "au", // 3
		"Wilmington, DE":                               "us", // 3
		"Bellevue, WA":                                 "us", // 3
		"Georgia":                                      "ge", // 3
		"Greater Barcelona Metropolitan Area":          "es", // 3
		"Sofia Metropolitan Area":                      "bg", // 3
		"Estonia":                                      "ee", // 3
		"Phoenix, AZ":                                  "us", // 3
		"Bristol, England, United Kingdom":             "gb", // 3
		"Cary, NC":                                     "us", // 3
		"Argentina":                                    "ar", // 3
		"Petaling Jaya, Selangor, Malaysia":            "my", // 3
		"Santa Clara, CA":                              "us", // 3
		"Tokyo, Japan":                                 "jp", // 3
		"Greater Munich Metropolitan Area":             "de", // 3
		"Finland":                                      "fi", // 3
		"Tbilisi, Georgia":                             "ge", // 3
		"Community of Madrid, Spain":                   "es", // 3
		"Boston, MA":                                   "us", // 3
		"Costa Mesa, CA":                               "us", // 3
		"Alpharetta, GA":                               "us", // 3
		"Oslo, Oslo, Norway":                           "no", // 3
		"Rome, Latium, Italy":                          "it", // 3
		"Netanya, Center District, Israel":             "il", // 3
		"Porto, Porto, Portugal":                       "pt", // 3
	}
)
