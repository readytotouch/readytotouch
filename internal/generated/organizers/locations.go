package organizers

var (
	LocationCodeMap = map[string]string{
		"United States":                            "us", // 134
		"London, England, United Kingdom":          "gb", // 104
		"Bengaluru, Karnataka, India":              "in", // 103
		"Berlin, Berlin, Germany":                  "de", // 101
		"London Area, United Kingdom":              "gb", //  50
		"Paris, Île-de-France, France":             "fr", //  48
		"Warsaw, Mazowieckie, Poland":              "pl", //  45
		"Cracow, Małopolskie, Poland":              "pl", //  44
		"Ho Chi Minh City, Vietnam":                "vn", //  41
		"United Kingdom":                           "gb", //  41
		"New York, NY":                             "us", //  40
		"Barcelona, Catalonia, Spain":              "es", //  36
		"Poland":                                   "pl", //  35
		"Prague, Prague, Czechia":                  "cz", //  34
		"Toronto, ON":                              "ca", //  30
		"Spain":                                    "es", //  29
		"Chicago, IL":                              "us", //  27
		"Hamburg, Hamburg, Germany":                "de", //  25
		"Hyderabad, Telangana, India":              "in", //  25
		"Pune, Maharashtra, India":                 "in", //  25
		"Tel Aviv-Yafo, Tel Aviv District, Israel": "il", //  25
		"Canada":                                            "ca", //  24
		"McLean, VA":                                        "us", //  24
		"Berlin, Germany":                                   "de", //  21
		"Stockholm, Stockholm County, Sweden":               "se", //  21
		"Germany":                                           "de", //  20
		"Sofia, Sofia City, Bulgaria":                       "bg", //  20
		"Budapest, Budapest, Hungary":                       "hu", //  19
		"Santa Clara, CA":                                   "us", //  19
		"Singapore, Singapore":                              "sg", //  19
		"Austin, TX":                                        "us", //  18
		"Ho Chi Minh City Metropolitan Area":                "vn", //  17
		"Munich, Bavaria, Germany":                          "de", //  17
		"Vilnius, Vilniaus, Lithuania":                      "lt", //  17
		"Istanbul, Türkiye":                                 "tr", //  16
		"Madrid, Community of Madrid, Spain":                "es", //  16
		"San Jose, CA":                                      "us", //  16
		"São Paulo, São Paulo, Brazil":                      "br", //  16
		"Czechia":                                           "cz", //  15
		"Dublin, County Dublin, Ireland":                    "ie", //  15
		"Tallinn, Harjumaa, Estonia":                        "ee", //  15
		"Chennai, Tamil Nadu, India":                        "in", //  14
		"Plano, TX":                                         "us", //  14
		"San Francisco, CA":                                 "us", //  14
		"Sydney, New South Wales, Australia":                "au", //  14
		"Amsterdam, North Holland, Netherlands":             "nl", //  13
		"Brno, South Moravia, Czechia":                      "cz", //  13
		"Montreal, QC":                                      "ca", //  13
		"Milan, Lombardy, Italy":                            "it", //  12
		"Palo Alto, CA":                                     "us", //  12
		"Richmond, VA":                                      "us", //  12
		"Atlanta, GA":                                       "us", //  11
		"Belfast, Northern Ireland, United Kingdom":         "gb", //  11
		"Belgrade, Serbia":                                  "rs", //  11
		"Brazil":                                            "br", //  11
		"Helsinki, Uusimaa, Finland":                        "fi", //  11
		"Taipei, Taipei City, Taiwan":                       "tw", //  11
		"Vienna, Vienna, Austria":                           "at", //  11
		"Wrocław, Dolnośląskie, Poland":                     "pl", //  11
		"Bengaluru East, Karnataka, India":                  "in", //  10
		"Gurugram, Haryana, India":                          "in", //  10
		"India":                                             "in", //  10
		"Ireland":                                           "ie", //  10
		"Jersey City, NJ":                                   "us", //  10
		"Ottawa, ON":                                        "ca", //  10
		"Philippines":                                       "ph", //  10
		"Romania":                                           "ro", //  10
		"Seattle, WA":                                       "us", //  10
		"Istanbul, Istanbul, Türkiye":                       "tr", //   9
		"Mumbai, Maharashtra, India":                        "in", //   9
		"Noida, Uttar Pradesh, India":                       "in", //   9
		"Sofia Metropolitan Area":                           "bg", //   9
		"Türkiye":                                           "tr", //   9
		"Boston, MA":                                        "us", //   8
		"Cary, NC":                                          "us", //   8
		"European Union":                                    "",   //   8
		"Lyon, Auvergne-Rhône-Alpes, France":                "fr", //   8
		"Bellevue, WA":                                      "us", //   7
		"Bucharest, Bucharest, Romania":                     "ro", //   7
		"EMEA":                                              "",   //   7
		"Greater Munich Metropolitan Area":                  "de", //   7
		"Herzliya, Tel Aviv District, Israel":               "il", //   7
		"Kaunas, Kaunas, Lithuania":                         "lt", //   7
		"Leicester, England, United Kingdom":                "gb", //   7
		"Mumbai Metropolitan Region":                        "in", //   7
		"New York City Metropolitan Area":                   "us", //   7
		"Porto, Porto, Portugal":                            "pt", //   7
		"Portugal":                                          "pt", //   7
		"Pune Division, Maharashtra, India":                 "in", //   7
		"San Francisco Bay Area":                            "us", //   7
		"Singapore":                                         "sg", //   7
		"Texas, United States":                              "us", //   7
		"Utrecht, Utrecht, Netherlands":                     "nl", //   7
		"Australia":                                         "au", //   6
		"Boulder, CO":                                       "us", //   6
		"Bristol, England, United Kingdom":                  "gb", //   6
		"Cyprus":                                            "cy", //   6
		"Dedham, MA":                                        "us", //   6
		"Dhaka, Dhaka, Bangladesh":                          "bd", //   6
		"European Economic Area":                            "",   //   6
		"Glasgow, Scotland, United Kingdom":                 "gb", //   6
		"Greater Buenos Aires":                              "ar", //   6
		"Isleworth, England, United Kingdom":                "gb", //   6
		"Limassol, Cyprus":                                  "cy", //   6
		"Lisbon, Lisbon, Portugal":                          "pt", //   6
		"Los Angeles, CA":                                   "us", //   6
		"Mexico City, Mexico":                               "mx", //   6
		"Pune/Pimpri-Chinchwad Area":                        "in", //   6
		"Redlands, CA":                                      "us", //   6
		"Rome, Latium, Italy":                               "it", //   6
		"Slovakia":                                          "sk", //   6
		"Tbilisi, Georgia":                                  "ge", //   6
		"Tel Aviv District, Israel":                         "il", //   6
		"Ukraine":                                           "ua", //   6
		"Allen, TX":                                         "us", //   5
		"Costa Mesa, CA":                                    "us", //   5
		"Cracow Metropolitan Area":                          "pl", //   5
		"Georgia":                                           "ge", //   5
		"Greater Barcelona Metropolitan Area":               "es", //   5
		"Greater London, England, United Kingdom":           "gb", //   5
		"Grünheide, Brandenburg, Germany":                   "de", //   5
		"Israel":                                            "il", //   5
		"Lisboa, Lisbon, Portugal":                          "pt", //   5
		"Lithuania":                                         "lt", //   5
		"New York, United States":                           "us", //   5
		"Riga, Riga, Latvia":                                "lv", //   5
		"Sofia City, Bulgaria":                              "bg", //   5
		"Sunnyvale, CA":                                     "ca", //   5
		"Sweden":                                            "se", //   5
		"Thiruvananthapuram, Kerala, India":                 "in", //   5
		"Tokyo, Japan":                                      "jp", //   5
		"Budapest Metropolitan Area":                        "hu", //   4
		"Budapest, Hungary":                                 "hu", //   4
		"Columbus, OH":                                      "us", //   4
		"Community of Madrid, Spain":                        "es", //   4
		"Delhi, Delhi, India":                               "in", //   4
		"Fort Worth, TX":                                    "us", //   4
		"Greater Colorado Springs Area":                     "us", //   4
		"Greater Kolkata Area":                              "in", //   4
		"Hanoi, Hanoi, Vietnam":                             "vn", //   4
		"Jacksonville, FL":                                  "us", //   4
		"Leuven, Flemish Region, Belgium":                   "be", //   4
		"Lexington, MA":                                     "us", //   4
		"Limassol, Limassol, Cyprus":                        "cy", //   4
		"Lisbon, Portugal":                                  "pt", //   4
		"Long Beach, CA":                                    "us", //   4
		"Manchester, England, United Kingdom":               "gb", //   4
		"Melbourne, Victoria, Australia":                    "au", //   4
		"Mexico":                                            "mx", //   4
		"Mulhuddart, Fingal, Ireland":                       "ie", //   4
		"Netanya, Center District, Israel":                  "il", //   4
		"Netherlands":                                       "nl", //   4
		"Oslo, Oslo, Norway":                                "no", //   4
		"Philadelphia, PA":                                  "us", //   4
		"Phoenix, AZ":                                       "us", //   4
		"Prague, Czechia":                                   "cz", //   4
		"Szeged, Csongrád, Hungary":                         "hu", //   4
		"Valbonne, Provence-Alpes-Côte d'Azur, France":      "fr", //   4
		"Valencia, Valencian Community, Spain":              "es", //   4
		"Valenzano, Apulia, Italy":                          "it", //   4
		"Wilmington, DE":                                    "us", //   4
		"Łódź, Łódzkie, Poland":                             "pl", //   4
		"Alpharetta, GA":                                    "us", //   3
		"Argentina":                                         "ar", //   3
		"Arlington, VA":                                     "in", //   3
		"Aveiro, Aveiro, Portugal":                          "pt", //   3
		"Bangalore Rural, Karnataka, India":                 "in", //   3
		"Bangkok, Bangkok City, Thailand":                   "th", //   3
		"Bengaluru South, Karnataka, India":                 "in", //   3
		"Bratislava, Bratislava, Slovakia":                  "sk", //   3
		"Cape Town, Western Cape, South Africa":             "za", //   3
		"City Of London, England, United Kingdom":           "gb", //   3
		"Cluj-Napoca, Cluj, Romania":                        "ro", //   3
		"Colombo, Western Province, Sri Lanka":              "lk", //   3
		"Creve Coeur, MO":                                   "us", //   3
		"Dallas, TX":                                        "us", //   3
		"Denver, CO":                                        "us", //   3
		"Dubai, United Arab Emirates":                       "ae", //   3
		"Estonia":                                           "ee", //   3
		"Finland":                                           "fi", //   3
		"France":                                            "fr", //   3
		"Geneva, Switzerland":                               "ch", //   3
		"Glendale, CA":                                      "us", //   3
		"Greater Hyderabad Area":                            "in", //   3
		"Greater Madrid Metropolitan Area":                  "es", //   3
		"Greater Seattle Area":                              "us", //   3
		"Greece":                                            "gr", //   3
		"Italy":                                             "it", //   3
		"Leiden, South Holland, Netherlands":                "nl", //   3
		"Lund, Skåne County, Sweden":                        "se", //   3
		"Mississauga, ON":                                   "ca", //   3
		"Montevideo, Montevideo, Uruguay":                   "uy", //   3
		"Naperville, IL":                                    "us", //   3
		"Nicosia, Nicosia, Cyprus":                          "cy", //   3
		"Ontario, Canada":                                   "ca", //   3
		"Petaling Jaya, Selangor, Malaysia":                 "my", //   3
		"Riga, Latvia":                                      "lv", //   3
		"Rožňava, Kosice, Slovakia":                         "sk", //   3
		"San Diego, CA":                                     "us", //   3
		"Santa Barbara, CA":                                 "us", //   3
		"Tampa, FL":                                         "us", //   3
		"Thane, Maharashtra, India":                         "in", //   3
		"Vancouver, BC":                                     "ca", //   3
		"Vietnam":                                           "vn", //   3
		"Warsaw Metropolitan Area":                          "pl", //   3
		"Zurich, Switzerland":                               "ch", //   3
		"APAC":                                              "",   //   2
		"Addis Ababa, Addis Ababa, Ethiopia":                "et", //   2
		"Ahmedabad, Gujarat, India":                         "in", //   2
		"Athens, Attiki, Greece":                            "gr", //   2
		"Auckland, Auckland, New Zealand":                   "nz", //   2
		"Bangalore Urban, Karnataka, India":                 "in", //   2
		"Bedford, NS":                                       "ca", //   2
		"Birmingham, England, United Kingdom":               "gb", //   2
		"Brussels, Brussels Region, Belgium":                "be", //   2
		"Bucharest, Romania":                                "ro", //   2
		"Calgary, AB":                                       "ca", //   2
		"Cambridge, England, United Kingdom":                "gb", //   2
		"Carlsbad, CA":                                      "us", //   2
		"Chantilly, VA":                                     "us", //   2
		"Chesterfield, MO":                                  "us", //   2
		"Ciudad Bolivia, Barinas State, Venezuela":          "ve", //   2
		"Collingwood, Victoria, Australia":                  "au", //   2
		"Cologne, North Rhine-Westphalia, Germany":          "de", //   2
		"Copenhagen, Capital Region of Denmark, Denmark":    "dk", //   2
		"Dresden, Saxony, Germany":                          "de", //   2
		"Edinburgh, Scotland, United Kingdom":               "gb", //   2
		"Fort Collins, CO":                                  "us", //   2
		"Frankfurt, Hesse, Germany":                         "de", //   2
		"Gdańsk, Pomorskie, Poland":                         "pl", //   2
		"Geneva, Geneva, Switzerland":                       "ch", //   2
		"Ghent, Flemish Region, Belgium":                    "be", //   2
		"Greater Chennai Area":                              "in", //   2
		"Greater Istanbul":                                  "tr", //   2
		"Greater Manchester, England, United Kingdom":       "uk", //   2
		"Greater Ottawa Metropolitan Area":                  "ca", //   2
		"Greater Porto Alegre":                              "br", //   2
		"Greater Toronto Area, Canada":                      "ca", //   2
		"Greater Vancouver Metropolitan Area":               "ca", //   2
		"Ho Chi Minh City, Ho Chi Minh City, Vietnam":       "vn", //   2
		"Hong Kong, Hong Kong SAR":                          "hk", //   2
		"Hopkins, MN":                                       "us", //   2
		"Hungary":                                           "hu", //   2
		"Krakowski, Małopolskie, Poland":                    "pl", //   2
		"Kyiv Metropolitan Area":                            "ua", //   2
		"Latin America":                                     "",   //   2
		"Leipzig, Saxony, Germany":                          "de", //   2
		"Limassol Municipality, Limassol, Cyprus":           "cy", //   2
		"Ljubljana, Ljubljana, Slovenia":                    "si", //   2
		"Los Angeles Metropolitan Area":                     "us", //   2
		"Mandaluyong, National Capital Region, Philippines": "ph", //   2
		"Manila, National Capital Region, Philippines":      "ph", //   2
		"Merrimack, NH":                                     "us", //   2
		"Metro Manila":                                      "ph", //   2
		"Monett, MO":                                        "us", //   2
		"NAMER":                                             "",   //   2
		"Navi Mumbai, Maharashtra, India":                   "in", //   2
		"New Jersey, United States":                         "us", //   2
		"New Zealand":                                       "nz", //   2
		"Newton, MA":                                        "us", //   2
		"North Sydney, New South Wales, Australia":          "au", //   2
		"Oakville, ON":                                      "ca", //   2
		"Offenburg, Baden-Württemberg, Germany":             "de", //   2
		"Pune District, Maharashtra, India":                 "in", //   2
		"Querétaro, Querétaro, Mexico":                      "mx", //   2
		"Raanana, Center District, Israel":                  "il", //   2
		"Raleigh, NC":                                       "us", //   2
		"Ramat Gan, Tel Aviv District, Israel":              "il", //   2
		"Reston, VA":                                        "us", //   2
		"Rosh HaAyin, Center District, Israel":              "il", //   2
		"Roswell, GA":                                       "us", //   2
		"Rotterdam, South Holland, Netherlands":             "nl", //   2
		"Sahibzada Ajit Singh Nagar, Punjab, India":         "in", //   2
		"Sandy, UT":                                         "us", //   2
		"Santiago, Santiago Metropolitan Region, Chile":     "cl", //   2
		"Schöppingen, North Rhine-Westphalia, Germany":      "de", //   2
		"Serbia":                             "rs", //   2
		"Stockholm, NY":                      "us", //   2
		"Sunrise, FL":                        "us", //   2
		"Switzerland":                        "ch", //   2
		"Taunton, MA":                        "us", //   2
		"Thu Đuc, Vietnam":                   "vn", //   2
		"Tokyo, Tokyo, Japan":                "jp", //   2
		"Tottenham, England, United Kingdom": "gb", //   2
		"Vindafjord, Rogaland, Norway":       "no", //   2
		"Waltham, MA":                        "us", //   2
		"Warrenton, VA":                      "us", //   2
		"Yerevan, Yerevan, Armenia":          "am", //   2
		"York, England, United Kingdom":      "gb", //   2
		"Zagreb, Croatia":                    "hr", //   2
		"Zurich, Zurich, Switzerland":        "ch", //   2
	}
)
