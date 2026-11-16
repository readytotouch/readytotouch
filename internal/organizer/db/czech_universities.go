package db

import "github.com/readytotouch/readytotouch/internal/domain"

func CzechUniversities() []domain.University {
	// https://www.linkedin.com/search/results/schools/?keywords=Czech%20OR%20Prague

	return []domain.University{
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15092560,
				Alias: "czuvpraze",
				Name:  "Czech University of Life Sciences Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15091777,
				Alias: "cvutpraha",
				Name:  "Czech Technical University in Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    86287,
				Alias: "university-of-pardubice",
				Name:  "University of Pardubice",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15093979,
				Alias: "univerzita-j-e-purkyne",
				Name:  "Jan Evangelista Purkyně University in Ústí nad Labem",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    553366,
				Alias: "fit-ctu",
				Name:  "Faculty of Information Technology CTU in Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    3496049,
				Alias: "univerzita-karlova",
				Name:  "Charles University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    156185,
				Alias: "vsb-technical-university-of-ostrava",
				Name:  "VSB — Technical University of Ostrava",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15092562,
				Alias: "tomas-bata-university",
				Name:  "Tomas Bata University in Zlín",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15093419,
				Alias: "vysoké-učení-technické-v-brně",
				Name:  "Brno University of Technology",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15093978,
				Alias: "mendeluniversityinbrno",
				Name:  "Mendel University in Brno",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1113156,
				Alias: "technical-university-of-liberec",
				Name:  "Technical University of Liberec",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18219827,
				Alias: "jihoceska-univerzita-v-ceskych-budejovicich",
				Name:  "University of South Bohemia in České Budějovice",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    10126918,
				Alias: "provozne-ekonomicka-fakulta-mendelovy-univerzity-v-brne",
				Name:  "Faculty of Business and Economics, Mendel University in Brno",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    245997,
				Alias: "vysok-kola-finan-n-a-spr-vn-",
				Name:  "University of Finance and Administration, Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15091775,
				Alias: "akademie-múzických-umění-v-praze",
				Name:  "Academy of Performing Arts in Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    865041,
				Alias: "prague-city-university",
				Name:  "Prague City University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15143943,
				Alias: "cvutfel",
				Name:  "Faculty of Electrical Engineering, Czech Technical University in Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    86285,
				Alias: "university-of-new-york-in-prague",
				Name:  "University of New York in Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15092563,
				Alias: "vysoká-škola-ekonomická-v-praze",
				Name:  "Prague University of Economics and Business",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15093981,
				Alias: "vschtpraha",
				Name:  "University of Chemistry and Technology in Prague (UCT Prague)",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    365087,
				Alias: "metropolitni-univerzita-praha",
				Name:  "Metropolitan University Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15137040,
				Alias: "pefczucz",
				Name:  "Faculty of Economics and Management CZU Prague",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15131369,
				Alias: "anglo-american-university-aau",
				Name:  "Anglo-American University, Prague",
			},
		},

		// Словаччина
		// {
		// 	LinkedInProfile: domain.LinkedInProfile{
		// 		ID:    386175,
		// 		Alias: "university-of-zilina",
		// 		Name:  "University of Zilina",
		// 	},
		// },

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    881395,
				Alias: "charles-university",
				Name:  "Charles University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    484970,
				Alias: "newton-university",
				Name:  "NEWTON University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15108939,
				Alias: "msct-avadh-mahavidhyalaya-bed-college-visnagar",
				Name:  "Charles University (Univerzita Karlova)",
			},
		},
	}
}
