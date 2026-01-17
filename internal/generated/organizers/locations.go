package organizers

var (
	LocationCodeMap = map[string]string{
		"United States":                             "us", //  85
		"Bengaluru, Karnataka, India":               "in", //  70
		"London, England, United Kingdom":           "gb", //  66
		"Berlin, Berlin, Germany":                   "de", //  53
		"Warsaw, Mazowieckie, Poland":               "pl", //  32
		"Cracow, Małopolskie, Poland":               "pl", //  31
		"London Area, United Kingdom":               "gb", //  29
		"Prague, Prague, Czechia":                   "cz", //  27
		"Ho Chi Minh City, Vietnam":                 "vn", //  26
		"Tel Aviv-Yafo, Tel Aviv District, Israel":  "il", //  25
		"United Kingdom":                            "gb", //  25
		"Paris, Île-de-France, France":              "fr", //  24
		"Poland":                                    "pl", //  23
		"Barcelona, Catalonia, Spain":               "es", //  22
		"New York, NY":                              "us", //  22
		"Pune, Maharashtra, India":                  "in", //  20
		"Chicago, IL":                               "us", //  19
		"McLean, VA":                                "us", //  19
		"Germany":                                   "de", //  18
		"Spain":                                     "es", //  18
		"Canada":                                    "ca", //  15
		"Toronto, ON":                               "ca", //  15
		"Hyderabad, Telangana, India":               "in", //  13
		"Hamburg, Hamburg, Germany":                 "de", //  12
		"Munich, Bavaria, Germany":                  "de", //  12
		"Plano, TX":                                 "us", //  12
		"Richmond, VA":                              "us", //  12
		"Sofia, Sofia City, Bulgaria":               "bg", //  12
		"Sydney, New South Wales, Australia":        "au", //  12
		"São Paulo, São Paulo, Brazil":              "br", //  12
		"Berlin, Germany":                           "de", //  11
		"Budapest, Budapest, Hungary":               "hu", //  11
		"Chennai, Tamil Nadu, India":                "in", //  11
		"Czechia":                                   "cz", //  11
		"Amsterdam, North Holland, Netherlands":     "nl", //  10
		"Dublin, County Dublin, Ireland":            "ie", //  10
		"Ho Chi Minh City Metropolitan Area":        "vn", //  10
		"Madrid, Community of Madrid, Spain":        "es", //  10
		"Montreal, QC":                              "ca", //  10
		"San Jose, CA":                              "us", //  10
		"Singapore, Singapore":                      "sg", //  10
		"Tallinn, Harjumaa, Estonia":                "ee", //  10
		"Austin, TX":                                "us", //   9
		"Brno, South Moravia, Czechia":              "cz", //   9
		"San Francisco, CA":                         "us", //   9
		"Stockholm, Stockholm County, Sweden":       "se", //   9
		"Vienna, Vienna, Austria":                   "at", //   9
		"Belgrade, Serbia":                          "rs", //   8
		"Gurugram, Haryana, India":                  "in", //   8
		"Istanbul, Türkiye":                         "tr", //   8
		"Milan, Lombardy, Italy":                    "it", //   8
		"Palo Alto, CA":                             "us", //   8
		"Seattle, WA":                               "us", //   8
		"Wrocław, Dolnośląskie, Poland":             "pl", //   8
		"Atlanta, GA":                               "us", //   7
		"Belfast, Northern Ireland, United Kingdom": "gb", //   7
		"Brazil":                                  "br", //   7
		"European Union":                          "",   //   7
		"Herzliya, Tel Aviv District, Israel":     "il", //   7
		"India":                                   "in", //   7
		"Istanbul, Istanbul, Türkiye":             "tr", //   7
		"Taipei, Taipei City, Taiwan":             "tw", //   7
		"Vilnius, Vilniaus, Lithuania":            "lt", //   7
		"Bengaluru East, Karnataka, India":        "in", //   6
		"Cyprus":                                  "cy", //   6
		"Jersey City, NJ":                         "us", //   6
		"Leicester, England, United Kingdom":      "gb", //   6
		"Lyon, Auvergne-Rhône-Alpes, France":      "fr", //   6
		"Noida, Uttar Pradesh, India":             "in", //   6
		"Philippines":                             "ph", //   6
		"Romania":                                 "ro", //   6
		"Slovakia":                                "sk", //   6
		"Dedham, MA":                              "us", //   5
		"Dhaka, Dhaka, Bangladesh":                "bd", //   5
		"EMEA":                                    "",   //   5
		"Greater London, England, United Kingdom": "gb", //   5
		"Helsinki, Uusimaa, Finland":              "fi", //   5
		"Ireland":                                 "ie", //   5
		"Kaunas, Kaunas, Lithuania":               "lt", //   5
		"Los Angeles, CA":                         "us", //   5
		"Mumbai, Maharashtra, India":              "in", //   5
		"Ottawa, ON":                              "ca", //   5
		"Tel Aviv District, Israel":               "il", //   5
		"Texas, United States":                    "us", //   5
		"Thiruvananthapuram, Kerala, India":       "in", //   5
		"Türkiye":                                 "tr", //   5
		"Utrecht, Utrecht, Netherlands":           "nl", //   5
		"Boston, MA":                              "us", //   4
		"Boulder, CO":                             "us", //   4
		"Delhi, Delhi, India":                     "in", //   4
		"Fort Worth, TX":                          "us", //   4
		"Georgia":                                 "ge", //   4
		"Greater Colorado Springs Area":           "us", //   4
		"Greater Munich Metropolitan Area":        "de", //   4
		"Jacksonville, FL":                        "us", //   4
		"Lisbon, Portugal":                        "pt", //   4
		"Mexico City, Mexico":                     "mx", //   4
		"Mulhuddart, Fingal, Ireland":             "ie", //   4
		"Mumbai Metropolitan Region":              "in", //   4
		"New York City Metropolitan Area":         "us", //   4
		"New York, United States":                 "us", //   4
		"Phoenix, AZ":                             "us", //   4
		"Portugal":                                "pt", //   4
		"Prague, Czechia":                         "cz", //   4
		"Rome, Latium, Italy":                     "it", //   4
		"San Francisco Bay Area":                  "us", //   4
		"Santa Clara, CA":                         "us", //   4
		"Singapore":                               "sg", //   4
		"Ukraine":                                 "ua", //   4
		"Valbonne, Provence-Alpes-Côte d'Azur, France": "fr", //   4
		"Alpharetta, GA":                                    "us", //   3
		"Argentina":                                         "ar", //   3
		"Australia":                                         "au", //   3
		"Bellevue, WA":                                      "us", //   3
		"Bratislava, Bratislava, Slovakia":                  "sk", //   3
		"Bristol, England, United Kingdom":                  "gb", //   3
		"Budapest Metropolitan Area":                        "hu", //   3
		"Cary, NC":                                          "us", //   3
		"Community of Madrid, Spain":                        "es", //   3
		"Costa Mesa, CA":                                    "us", //   3
		"Estonia":                                           "ee", //   3
		"Finland":                                           "fi", //   3
		"Glendale, CA":                                      "us", //   3
		"Greater Barcelona Metropolitan Area":               "es", //   3
		"Isleworth, England, United Kingdom":                "gb", //   3
		"Lithuania":                                         "lt", //   3
		"Manchester, England, United Kingdom":               "gb", //   3
		"Netanya, Center District, Israel":                  "il", //   3
		"Oslo, Oslo, Norway":                                "no", //   3
		"Petaling Jaya, Selangor, Malaysia":                 "my", //   3
		"Porto, Porto, Portugal":                            "pt", //   3
		"Rožňava, Kosice, Slovakia":                         "sk", //   3
		"Sofia Metropolitan Area":                           "bg", //   3
		"Sweden":                                            "se", //   3
		"Tbilisi, Georgia":                                  "ge", //   3
		"Tokyo, Japan":                                      "jp", //   3
		"Valenzano, Apulia, Italy":                          "it", //   3
		"Wilmington, DE":                                    "us", //   3
		"Ahmedabad, Gujarat, India":                         "in", //   2
		"Athens, Attiki, Greece":                            "gr", //   2
		"Aveiro, Aveiro, Portugal":                          "pt", //   2
		"Bangalore Rural, Karnataka, India":                 "in", //   2
		"Bengaluru South, Karnataka, India":                 "in", //   2
		"Birmingham, England, United Kingdom":               "gb", //   2
		"Brussels, Brussels Region, Belgium":                "be", //   2
		"Budapest, Hungary":                                 "hu", //   2
		"Cambridge, England, United Kingdom":                "gb", //   2
		"Chantilly, VA":                                     "us", //   2
		"City Of London, England, United Kingdom":           "gb", //   2
		"Collingwood, Victoria, Australia":                  "au", //   2
		"Colombo, Western Province, Sri Lanka":              "lk", //   2
		"Columbus, OH":                                      "us", //   2
		"Cracow Metropolitan Area":                          "pl", //   2
		"Dallas, TX":                                        "us", //   2
		"Denver, CO":                                        "us", //   2
		"Dubai, United Arab Emirates":                       "ae", //   2
		"Edinburgh, Scotland, United Kingdom":               "gb", //   2
		"France":                                            "fr", //   2
		"Glasgow, Scotland, United Kingdom":                 "gb", //   2
		"Greater Buenos Aires":                              "ar", //   2
		"Greater Chennai Area":                              "in", //   2
		"Greater Istanbul":                                  "tr", //   2
		"Greater Kolkata Area":                              "in", //   2
		"Greater Ottawa Metropolitan Area":                  "ca", //   2
		"Greater Porto Alegre":                              "br", //   2
		"Greater Seattle Area":                              "us", //   2
		"Greece":                                            "gr", //   2
		"Grünheide, Brandenburg, Germany":                   "de", //   2
		"Hong Kong, Hong Kong SAR":                          "hk", //   2
		"Israel":                                            "il", //   2
		"Italy":                                             "it", //   2
		"Kyiv Metropolitan Area":                            "ua", //   2
		"Latin America":                                     "",   //   2
		"Leiden, South Holland, Netherlands":                "nl", //   2
		"Leuven, Flemish Region, Belgium":                   "be", //   2
		"Limassol Municipality, Limassol, Cyprus":           "cy", //   2
		"Limassol, Cyprus":                                  "cy", //   2
		"Limassol, Limassol, Cyprus":                        "cy", //   2
		"Lisboa, Lisbon, Portugal":                          "pt", //   2
		"Long Beach, CA":                                    "us", //   2
		"Mandaluyong, National Capital Region, Philippines": "ph", //   2
		"Melbourne, Victoria, Australia":                    "au", //   2
		"Merrimack, NH":                                     "us", //   2
		"Metro Manila":                                      "ph", //   2
		"Montevideo, Montevideo, Uruguay":                   "uy", //   2
		"Naperville, IL":                                    "us", //   2
		"Netherlands":                                       "nl", //   2
		"New Jersey, United States":                         "us", //   2
		"New Zealand":                                       "nz", //   2
		"Newton, MA":                                        "us", //   2
		"Nicosia, Nicosia, Cyprus":                          "cy", //   2
		"Oakville, ON":                                      "ca", //   2
		"Philadelphia, PA":                                  "us", //   2
		"Querétaro, Querétaro, Mexico":                      "mx", //   2
		"Raanana, Center District, Israel":                  "il", //   2
		"Redlands, CA":                                      "us", //   2
		"Reston, VA":                                        "us", //   2
		"Riga, Riga, Latvia":                                "lv", //   2
		"Rotterdam, South Holland, Netherlands":             "nl", //   2
		"Sahibzada Ajit Singh Nagar, Punjab, India":         "in", //   2
		"San Diego, CA":                                     "us", //   2
		"Sandy, UT":                                         "us", //   2
		"Santa Barbara, CA":                                 "us", //   2
		"Schöppingen, North Rhine-Westphalia, Germany":      "de", //   2
		"Serbia":                               "rs", //   2
		"Sunrise, FL":                          "us", //   2
		"Switzerland":                          "ch", //   2
		"Taunton, MA":                          "us", //   2
		"Thane, Maharashtra, India":            "in", //   2
		"Tokyo, Tokyo, Japan":                  "jp", //   2
		"Tottenham, England, United Kingdom":   "gb", //   2
		"Valencia, Valencian Community, Spain": "es", //   2
		"Vietnam":                              "vn", //   2
		"Waltham, MA":                          "us", //   2
		"Warrenton, VA":                        "us", //   2
		"Zurich, Switzerland":                  "ch", //   2
		"Łódź, Łódzkie, Poland":                "pl", //   2

	}
)
