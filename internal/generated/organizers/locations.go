package organizers

var (
	LocationCodeMap = map[string]string{
		"United States":                            "us", // 106
		"Bengaluru, Karnataka, India":              "in", //  80
		"London, England, United Kingdom":          "gb", //  76
		"Berlin, Berlin, Germany":                  "de", //  64
		"Cracow, Małopolskie, Poland":              "pl", //  36
		"London Area, United Kingdom":              "gb", //  33
		"Paris, Île-de-France, France":             "fr", //  33
		"Warsaw, Mazowieckie, Poland":              "pl", //  33
		"Ho Chi Minh City, Vietnam":                "vn", //  30
		"Prague, Prague, Czechia":                  "cz", //  30
		"United Kingdom":                           "gb", //  30
		"New York, NY":                             "us", //  29
		"Barcelona, Catalonia, Spain":              "es", //  26
		"Poland":                                   "pl", //  25
		"Pune, Maharashtra, India":                 "in", //  25
		"Tel Aviv-Yafo, Tel Aviv District, Israel": "il", //  25
		"Spain":                                     "es", //  23
		"McLean, VA":                                "us", //  20
		"Chicago, IL":                               "us", //  19
		"Germany":                                   "de", //  19
		"Hyderabad, Telangana, India":               "in", //  18
		"Canada":                                    "ca", //  17
		"Toronto, ON":                               "ca", //  16
		"Berlin, Germany":                           "de", //  15
		"Budapest, Budapest, Hungary":               "hu", //  15
		"Hamburg, Hamburg, Germany":                 "de", //  15
		"Singapore, Singapore":                      "sg", //  15
		"Munich, Bavaria, Germany":                  "de", //  14
		"Plano, TX":                                 "us", //  14
		"Sofia, Sofia City, Bulgaria":               "bg", //  13
		"Sydney, New South Wales, Australia":        "au", //  13
		"Chennai, Tamil Nadu, India":                "in", //  12
		"Czechia":                                   "cz", //  12
		"Dublin, County Dublin, Ireland":            "ie", //  12
		"Istanbul, Türkiye":                         "tr", //  12
		"Richmond, VA":                              "us", //  12
		"São Paulo, São Paulo, Brazil":              "br", //  12
		"Tallinn, Harjumaa, Estonia":                "ee", //  12
		"Amsterdam, North Holland, Netherlands":     "nl", //  11
		"Austin, TX":                                "us", //  11
		"Montreal, QC":                              "ca", //  11
		"Stockholm, Stockholm County, Sweden":       "se", //  11
		"Brno, South Moravia, Czechia":              "cz", //  10
		"Ho Chi Minh City Metropolitan Area":        "vn", //  10
		"Madrid, Community of Madrid, Spain":        "es", //  10
		"Palo Alto, CA":                             "us", //  10
		"San Francisco, CA":                         "us", //  10
		"San Jose, CA":                              "us", //  10
		"Wrocław, Dolnośląskie, Poland":             "pl", //  10
		"Belgrade, Serbia":                          "rs", //   9
		"Brazil":                                    "br", //   9
		"Gurugram, Haryana, India":                  "in", //   9
		"Jersey City, NJ":                           "us", //   9
		"Romania":                                   "ro", //   9
		"Santa Clara, CA":                           "us", //   9
		"Seattle, WA":                               "us", //   9
		"Vienna, Vienna, Austria":                   "at", //   9
		"Vilnius, Vilniaus, Lithuania":              "lt", //   9
		"Atlanta, GA":                               "us", //   8
		"Belfast, Northern Ireland, United Kingdom": "gb", //   8
		"India":                                             "in", //   8
		"Istanbul, Istanbul, Türkiye":                       "tr", //   8
		"Milan, Lombardy, Italy":                            "it", //   8
		"Noida, Uttar Pradesh, India":                       "in", //   8
		"Boston, MA":                                        "us", //   7
		"European Union":                                    "",   //   7
		"Herzliya, Tel Aviv District, Israel":               "il", //   7
		"Leicester, England, United Kingdom":                "gb", //   7
		"Mumbai, Maharashtra, India":                        "in", //   7
		"Ottawa, ON":                                        "ca", //   7
		"Philippines":                                       "ph", //   7
		"Taipei, Taipei City, Taiwan":                       "tw", //   7
		"Bengaluru East, Karnataka, India":                  "in", //   6
		"Bristol, England, United Kingdom":                  "gb", //   6
		"Cyprus":                                            "cy", //   6
		"Helsinki, Uusimaa, Finland":                        "fi", //   6
		"Los Angeles, CA":                                   "us", //   6
		"Lyon, Auvergne-Rhône-Alpes, France":                "fr", //   6
		"Slovakia":                                          "sk", //   6
		"Türkiye":                                           "tr", //   6
		"Utrecht, Utrecht, Netherlands":                     "nl", //   6
		"Boulder, CO":                                       "us", //   5
		"Dedham, MA":                                        "us", //   5
		"Dhaka, Dhaka, Bangladesh":                          "bd", //   5
		"EMEA":                                              "",   //   5
		"Greater Barcelona Metropolitan Area":               "es", //   5
		"Greater London, England, United Kingdom":           "gb", //   5
		"Greater Munich Metropolitan Area":                  "de", //   5
		"Ireland":                                           "ie", //   5
		"Isleworth, England, United Kingdom":                "gb", //   5
		"Kaunas, Kaunas, Lithuania":                         "lt", //   5
		"Mexico City, Mexico":                               "mx", //   5
		"New York City Metropolitan Area":                   "us", //   5
		"New York, United States":                           "us", //   5
		"Redlands, CA":                                      "us", //   5
		"San Francisco Bay Area":                            "us", //   5
		"Tel Aviv District, Israel":                         "il", //   5
		"Texas, United States":                              "us", //   5
		"Thiruvananthapuram, Kerala, India":                 "in", //   5
		"Tokyo, Japan":                                      "jp", //   5
		"Australia":                                         "au", //   4
		"Bellevue, WA":                                      "us", //   4
		"Budapest, Hungary":                                 "hu", //   4
		"Cary, NC":                                          "us", //   4
		"Costa Mesa, CA":                                    "us", //   4
		"Delhi, Delhi, India":                               "in", //   4
		"Fort Worth, TX":                                    "us", //   4
		"Georgia":                                           "ge", //   4
		"Greater Colorado Springs Area":                     "us", //   4
		"Jacksonville, FL":                                  "us", //   4
		"Limassol, Cyprus":                                  "cy", //   4
		"Lisbon, Portugal":                                  "pt", //   4
		"Manchester, England, United Kingdom":               "gb", //   4
		"Mulhuddart, Fingal, Ireland":                       "ie", //   4
		"Mumbai Metropolitan Region":                        "in", //   4
		"Phoenix, AZ":                                       "us", //   4
		"Porto, Porto, Portugal":                            "pt", //   4
		"Portugal":                                          "pt", //   4
		"Prague, Czechia":                                   "cz", //   4
		"Rome, Latium, Italy":                               "it", //   4
		"Singapore":                                         "sg", //   4
		"Ukraine":                                           "ua", //   4
		"Valbonne, Provence-Alpes-Côte d'Azur, France":      "fr", //   4
		"Valenzano, Apulia, Italy":                          "it", //   4
		"Alpharetta, GA":                                    "us", //   3
		"Argentina":                                         "ar", //   3
		"Bratislava, Bratislava, Slovakia":                  "sk", //   3
		"Bucharest, Bucharest, Romania":                     "ro", //   3
		"Budapest Metropolitan Area":                        "hu", //   3
		"City Of London, England, United Kingdom":           "gb", //   3
		"Community of Madrid, Spain":                        "es", //   3
		"Cracow Metropolitan Area":                          "pl", //   3
		"Dallas, TX":                                        "us", //   3
		"Denver, CO":                                        "us", //   3
		"Dubai, United Arab Emirates":                       "ae", //   3
		"Estonia":                                           "ee", //   3
		"Finland":                                           "fi", //   3
		"Glendale, CA":                                      "us", //   3
		"Greater Buenos Aires":                              "ar", //   3
		"Greater Kolkata Area":                              "in", //   3
		"Israel":                                            "il", //   3
		"Leiden, South Holland, Netherlands":                "nl", //   3
		"Lisboa, Lisbon, Portugal":                          "pt", //   3
		"Lisbon, Lisbon, Portugal":                          "pt", //   3
		"Lithuania":                                         "lt", //   3
		"Melbourne, Victoria, Australia":                    "au", //   3
		"Naperville, IL":                                    "us", //   3
		"Netanya, Center District, Israel":                  "il", //   3
		"Oslo, Oslo, Norway":                                "no", //   3
		"Petaling Jaya, Selangor, Malaysia":                 "my", //   3
		"Philadelphia, PA":                                  "us", //   3
		"Pune/Pimpri-Chinchwad Area":                        "in", //   3
		"Riga, Latvia":                                      "lv", //   3
		"Rožňava, Kosice, Slovakia":                         "sk", //   3
		"Sofia City, Bulgaria":                              "bg", //   3
		"Sofia Metropolitan Area":                           "bg", //   3
		"Sweden":                                            "se", //   3
		"Tbilisi, Georgia":                                  "ge", //   3
		"Valencia, Valencian Community, Spain":              "es", //   3
		"Wilmington, DE":                                    "us", //   3
		"Ahmedabad, Gujarat, India":                         "in", //   2
		"Allen, TX":                                         "us", //   2
		"Athens, Attiki, Greece":                            "gr", //   2
		"Aveiro, Aveiro, Portugal":                          "pt", //   2
		"Bangalore Rural, Karnataka, India":                 "in", //   2
		"Bengaluru South, Karnataka, India":                 "in", //   2
		"Birmingham, England, United Kingdom":               "gb", //   2
		"Brussels, Brussels Region, Belgium":                "be", //   2
		"Calgary, AB":                                       "ca", //   2
		"Cambridge, England, United Kingdom":                "gb", //   2
		"Chantilly, VA":                                     "us", //   2
		"Cluj-Napoca, Cluj, Romania":                        "ro", //   2
		"Collingwood, Victoria, Australia":                  "au", //   2
		"Colombo, Western Province, Sri Lanka":              "lk", //   2
		"Columbus, OH":                                      "us", //   2
		"Creve Coeur, MO":                                   "us", //   2
		"Edinburgh, Scotland, United Kingdom":               "gb", //   2
		"European Economic Area":                            "",   //   2
		"France":                                            "fr", //   2
		"Ghent, Flemish Region, Belgium":                    "be", //   2
		"Glasgow, Scotland, United Kingdom":                 "gb", //   2
		"Greater Chennai Area":                              "in", //   2
		"Greater Istanbul":                                  "tr", //   2
		"Greater Madrid Metropolitan Area":                  "es", //   2
		"Greater Ottawa Metropolitan Area":                  "ca", //   2
		"Greater Porto Alegre":                              "br", //   2
		"Greater Seattle Area":                              "us", //   2
		"Greater Toronto Area, Canada":                      "ca", //   2
		"Greece":                                            "gr", //   2
		"Grünheide, Brandenburg, Germany":                   "de", //   2
		"Hong Kong, Hong Kong SAR":                          "hk", //   2
		"Hopkins, MN":                                       "us", //   2
		"Hungary":                                           "hu", //   2
		"Italy":                                             "it", //   2
		"Krakowski, Małopolskie, Poland":                    "pl", //   2
		"Kyiv Metropolitan Area":                            "ua", //   2
		"Latin America":                                     "",   //   2
		"Leuven, Flemish Region, Belgium":                   "be", //   2
		"Lexington, MA":                                     "us", //   2
		"Limassol Municipality, Limassol, Cyprus":           "cy", //   2
		"Limassol, Limassol, Cyprus":                        "cy", //   2
		"Ljubljana, Ljubljana, Slovenia":                    "si", //   2
		"Long Beach, CA":                                    "us", //   2
		"Mandaluyong, National Capital Region, Philippines": "ph", //   2
		"Merrimack, NH":                                     "us", //   2
		"Metro Manila":                                      "ph", //   2
		"Montevideo, Montevideo, Uruguay":                   "uy", //   2
		"Navi Mumbai, Maharashtra, India":                   "in", //   2
		"Netherlands":                                       "nl", //   2
		"New Jersey, United States":                         "us", //   2
		"New Zealand":                                       "nz", //   2
		"Newton, MA":                                        "us", //   2
		"Nicosia, Nicosia, Cyprus":                          "cy", //   2
		"Oakville, ON":                                      "ca", //   2
		"Querétaro, Querétaro, Mexico":                      "mx", //   2
		"Raanana, Center District, Israel":                  "il", //   2
		"Reston, VA":                                        "us", //   2
		"Riga, Riga, Latvia":                                "lv", //   2
		"Rotterdam, South Holland, Netherlands":             "nl", //   2
		"Sahibzada Ajit Singh Nagar, Punjab, India":         "in", //   2
		"San Diego, CA":                                     "us", //   2
		"Sandy, UT":                                         "us", //   2
		"Santa Barbara, CA":                                 "us", //   2
		"Santiago, Santiago Metropolitan Region, Chile":     "cl", //   2
		"Schöppingen, North Rhine-Westphalia, Germany":      "de", //   2
		"Serbia":                             "rs", //   2
		"Sunrise, FL":                        "us", //   2
		"Switzerland":                        "ch", //   2
		"Szeged, Csongrád, Hungary":          "hu", //   2
		"Tampa, FL":                          "us", //   2
		"Taunton, MA":                        "us", //   2
		"Thane, Maharashtra, India":          "in", //   2
		"Tokyo, Tokyo, Japan":                "jp", //   2
		"Tottenham, England, United Kingdom": "gb", //   2
		"Vietnam":                            "vn", //   2
		"Vindafjord, Rogaland, Norway":       "no", //   2
		"Waltham, MA":                        "us", //   2
		"Warrenton, VA":                      "us", //   2
		"York, England, United Kingdom":      "gb", //   2
		"Zagreb, Croatia":                    "hr", //   2
		"Zurich, Switzerland":                "ch", //   2
		"Zurich, Zurich, Switzerland":        "ch", //   2
		"Łódź, Łódzkie, Poland":              "pl", //   2

	}
)
