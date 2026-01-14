package organizers

var (
	LocationCodeMap = map[string]string{
		"United States":                                     "us", // 78
		"Bengaluru, Karnataka, India":                       "in", // 67
		"London, England, United Kingdom":                   "gb", // 64
		"Berlin, Berlin, Germany":                           "de", // 51
		"Cracow, Małopolskie, Poland":                       "pl", // 30
		"Warsaw, Mazowieckie, Poland":                       "pl", // 28
		"London Area, United Kingdom":                       "gb", // 26
		"Prague, Prague, Czechia":                           "cz", // 26
		"Ho Chi Minh City, Vietnam":                         "vn", // 25
		"Tel Aviv-Yafo, Tel Aviv District, Israel":          "il", // 25
		"Paris, Île-de-France, France":                      "fr", // 23
		"United Kingdom":                                    "gb", // 23
		"Poland":                                            "pl", // 22
		"Barcelona, Catalonia, Spain":                       "es", // 21
		"New York, NY":                                      "us", // 20
		"Pune, Maharashtra, India":                          "in", // 20
		"Chicago, IL":                                       "us", // 18
		"Germany":                                           "de", // 18
		"Spain":                                             "es", // 18
		"Toronto, ON":                                       "ca", // 15
		"Canada":                                            "ca", // 14
		"Hyderabad, Telangana, India":                       "in", // 13
		"Hamburg, Hamburg, Germany":                         "de", // 12
		"Sofia, Sofia City, Bulgaria":                       "bg", // 12
		"Sydney, New South Wales, Australia":                "au", // 12
		"São Paulo, São Paulo, Brazil":                      "br", // 12
		"Berlin, Germany":                                   "de", // 11
		"Budapest, Budapest, Hungary":                       "hu", // 11
		"Munich, Bavaria, Germany":                          "de", // 11
		"Amsterdam, North Holland, Netherlands":             "nl", // 10
		"Chennai, Tamil Nadu, India":                        "in", // 10
		"Czechia":                                           "cz", // 10
		"Montreal, QC":                                      "ca", // 10
		"Plano, TX":                                         "us", // 10
		"Tallinn, Harjumaa, Estonia":                        "ee", // 10
		"Austin, TX":                                        "us", // 9
		"Brno, South Moravia, Czechia":                      "cz", // 9
		"Dublin, County Dublin, Ireland":                    "ie", // 9
		"Ho Chi Minh City Metropolitan Area":                "vn", // 9
		"Madrid, Community of Madrid, Spain":                "es", // 9
		"Richmond, VA":                                      "us", // 9
		"San Francisco, CA":                                 "us", // 9
		"San Jose, CA":                                      "us", // 9
		"Singapore, Singapore":                              "sg", // 9
		"Stockholm, Stockholm County, Sweden":               "se", // 9
		"Vienna, Vienna, Austria":                           "at", // 9
		"Belgrade, Serbia":                                  "rs", // 8
		"Gurugram, Haryana, India":                          "in", // 8
		"Milan, Lombardy, Italy":                            "it", // 8
		"Seattle, WA":                                       "us", // 8
		"Wrocław, Dolnośląskie, Poland":                     "pl", // 8
		"Atlanta, GA":                                       "us", // 7
		"Brazil":                                            "br", // 7
		"European Union":                                    "",   // 7
		"Herzliya, Tel Aviv District, Israel":               "il", // 7
		"India":                                             "in", // 7
		"Istanbul, Istanbul, Türkiye":                       "tr", // 7
		"Istanbul, Türkiye":                                 "tr", // 7
		"Palo Alto, CA":                                     "us", // 7
		"Taipei, Taipei City, Taiwan":                       "tw", // 7
		"Vilnius, Vilniaus, Lithuania":                      "lt", // 7
		"Belfast, Northern Ireland, United Kingdom":         "gb", // 6
		"Bengaluru East, Karnataka, India":                  "in", // 6
		"Cyprus":                                            "cy", // 6
		"Jersey City, NJ":                                   "us", // 6
		"Leicester, England, United Kingdom":                "gb", // 6
		"Lyon, Auvergne-Rhône-Alpes, France":                "fr", // 6
		"McLean, VA":                                        "us", // 6
		"Noida, Uttar Pradesh, India":                       "in", // 6
		"Philippines":                                       "ph", // 6
		"Slovakia":                                          "sk", // 6
		"Dedham, MA":                                        "us", // 5
		"Dhaka, Dhaka, Bangladesh":                          "bd", // 5
		"EMEA":                                              "",   // 5
		"Helsinki, Uusimaa, Finland":                        "fi", // 5
		"Ireland":                                           "ie", // 5
		"Kaunas, Kaunas, Lithuania":                         "lt", // 5
		"Los Angeles, CA":                                   "us", // 5
		"Mumbai, Maharashtra, India":                        "in", // 5
		"Ottawa, ON":                                        "ca", // 5
		"Tel Aviv District, Israel":                         "il", // 5
		"Thiruvananthapuram, Kerala, India":                 "in", // 5
		"Türkiye":                                           "tr", // 5
		"Utrecht, Utrecht, Netherlands":                     "nl", // 5
		"Boulder, CO":                                       "us", // 4
		"Delhi, Delhi, India":                               "in", // 4
		"Fort Worth, TX":                                    "us", // 4
		"Georgia":                                           "ge", // 4
		"Greater Colorado Springs Area":                     "us", // 4
		"Greater Munich Metropolitan Area":                  "de", // 4
		"Jacksonville, FL":                                  "us", // 4
		"Lisbon, Portugal":                                  "pt", // 4
		"Mexico City, Mexico":                               "mx", // 4
		"Mulhuddart, Fingal, Ireland":                       "ie", // 4
		"Mumbai Metropolitan Region":                        "in", // 4
		"New York City Metropolitan Area":                   "us", // 4
		"New York, United States":                           "us", // 4
		"Phoenix, AZ":                                       "us", // 4
		"Portugal":                                          "pt", // 4
		"Prague, Czechia":                                   "cz", // 4
		"San Francisco Bay Area":                            "us", // 4
		"Santa Clara, CA":                                   "us", // 4
		"Singapore":                                         "sg", // 4
		"Ukraine":                                           "ua", // 4
		"Valbonne, Provence-Alpes-Côte d'Azur, France":      "fr", // 4
		"Alpharetta, GA":                                    "us", // 3
		"Argentina":                                         "ar", // 3
		"Australia":                                         "au", // 3
		"Bellevue, WA":                                      "us", // 3
		"Boston, MA":                                        "us", // 3
		"Bratislava, Bratislava, Slovakia":                  "sk", // 3
		"Bristol, England, United Kingdom":                  "gb", // 3
		"Cary, NC":                                          "us", // 3
		"Community of Madrid, Spain":                        "es", // 3
		"Costa Mesa, CA":                                    "us", // 3
		"Estonia":                                           "ee", // 3
		"Finland":                                           "fi", // 3
		"Glendale, CA":                                      "us", // 3
		"Greater Barcelona Metropolitan Area":               "es", // 3
		"Lithuania":                                         "lt", // 3
		"Manchester, England, United Kingdom":               "gb", // 3
		"Netanya, Center District, Israel":                  "il", // 3
		"Oslo, Oslo, Norway":                                "no", // 3
		"Petaling Jaya, Selangor, Malaysia":                 "my", // 3
		"Porto, Porto, Portugal":                            "pt", // 3
		"Rome, Latium, Italy":                               "it", // 3
		"Rožňava, Kosice, Slovakia":                         "sk", // 3
		"Sofia Metropolitan Area":                           "bg", // 3
		"Sweden":                                            "se", // 3
		"Tbilisi, Georgia":                                  "ge", // 3
		"Tokyo, Japan":                                      "jp", // 3
		"Valenzano, Apulia, Italy":                          "it", // 3
		"Wilmington, DE":                                    "us", // 3
		"Ahmedabad, Gujarat, India":                         "in", // 2
		"Athens, Attiki, Greece":                            "gr", // 2
		"Aveiro, Aveiro, Portugal":                          "pt", // 2
		"Bangalore Rural, Karnataka, India":                 "in", // 2
		"Bengaluru South, Karnataka, India":                 "in", // 2
		"Birmingham, England, United Kingdom":               "gb", // 2
		"Brussels, Brussels Region, Belgium":                "be", // 2
		"Budapest Metropolitan Area":                        "hu", // 2
		"Budapest, Hungary":                                 "hu", // 2
		"Cambridge, England, United Kingdom":                "gb", // 2
		"Chantilly, VA":                                     "us", // 2
		"City Of London, England, United Kingdom":           "gb", // 2
		"Collingwood, Victoria, Australia":                  "au", // 2
		"Colombo, Western Province, Sri Lanka":              "lk", // 2
		"Columbus, OH":                                      "us", // 2
		"Cracow Metropolitan Area":                          "pl", // 2
		"Denver, CO":                                        "us", // 2
		"Dubai, United Arab Emirates":                       "ae", // 2
		"Edinburgh, Scotland, United Kingdom":               "gb", // 2
		"France":                                            "fr", // 2
		"Glasgow, Scotland, United Kingdom":                 "gb", // 2
		"Greater Buenos Aires":                              "ar", // 2
		"Greater Chennai Area":                              "in", // 2
		"Greater Istanbul":                                  "tr", // 2
		"Greater Kolkata Area":                              "in", // 2
		"Greater London, England, United Kingdom":           "gb", // 2
		"Greater Ottawa Metropolitan Area":                  "ca", // 2
		"Greater Porto Alegre":                              "br", // 2
		"Greater Seattle Area":                              "us", // 2
		"Greece":                                            "gr", // 2
		"Grünheide, Brandenburg, Germany":                   "de", // 2
		"Hong Kong, Hong Kong SAR":                          "hk", // 2
		"Isleworth, England, United Kingdom":                "gb", // 2
		"Italy":                                             "it", // 2
		"Kyiv Metropolitan Area":                            "ua", // 2
		"Latin America":                                     "",   // 2
		"Leiden, South Holland, Netherlands":                "nl", // 2
		"Leuven, Flemish Region, Belgium":                   "be", // 2
		"Limassol Municipality, Limassol, Cyprus":           "cy", // 2
		"Limassol, Cyprus":                                  "cy", // 2
		"Limassol, Limassol, Cyprus":                        "cy", // 2
		"Lisboa, Lisbon, Portugal":                          "pt", // 2
		"Long Beach, CA":                                    "us", // 2
		"Mandaluyong, National Capital Region, Philippines": "ph", // 2
		"Melbourne, Victoria, Australia":                    "au", // 2
		"Merrimack, NH":                                     "us", // 2
		"Montevideo, Montevideo, Uruguay":                   "uy", // 2
		"Naperville, IL":                                    "us", // 2
		"Netherlands":                                       "nl", // 2
		"New Zealand":                                       "nz", // 2
		"Newton, MA":                                        "us", // 2
		"Nicosia, Nicosia, Cyprus":                          "cy", // 2
		"Oakville, ON":                                      "ca", // 2
		"Philadelphia, PA":                                  "us", // 2
		"Querétaro, Querétaro, Mexico":                      "mx", // 2
		"Raanana, Center District, Israel":                  "il", // 2
		"Redlands, CA":                                      "us", // 2
		"Reston, VA":                                        "us", // 2
		"Riga, Riga, Latvia":                                "lv", // 2
		"Romania":                                           "ro", // 2
		"Rotterdam, South Holland, Netherlands":             "nl", // 2
		"Sahibzada Ajit Singh Nagar, Punjab, India":         "in", // 2
		"Santa Barbara, CA":                                 "us", // 2
		"Schöppingen, North Rhine-Westphalia, Germany":      "de", // 2
		"Serbia":                               "rs", // 2
		"Sunrise, FL":                          "us", // 2
		"Switzerland":                          "ch", // 2
		"Taunton, MA":                          "us", // 2
		"Texas, United States":                 "us", // 2
		"Tokyo, Tokyo, Japan":                  "jp", // 2
		"Tottenham, England, United Kingdom":   "gb", // 2
		"Valencia, Valencian Community, Spain": "es", // 2
		"Vietnam":                              "vn", // 2
		"Waltham, MA":                          "us", // 2
		"Warrenton, VA":                        "us", // 2
		"Zurich, Switzerland":                  "ch", // 2
		"Łódź, Łódzkie, Poland":                "pl", // 2
	}
)
