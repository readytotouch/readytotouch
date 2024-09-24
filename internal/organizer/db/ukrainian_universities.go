package db

import "github.com/readytotouch/readytotouch/internal/domain"

func UkrainianUniversities() []domain.University {
	// https://dou.ua/lenta/articles/ukrainian-universities-2023/
	// УКУ
	// ХНЕУ
	// КНУ ім. Шевченка
	// НаУКМА
	// ОНУ ім. Мечникова
	// ДУТ
	// СумДУ
	// НТУУ «КПІ ім. І. Сікорського»
	// ЛНУ ім. Франка
	// ХНУРЕ
	// НУ «Львівська політехніка»
	// ДНУ ім. Гончара
	// НТУ «Дніпровська політехніка»
	// ХНУ [м. Хмельницький]
	// КНЕУ
	// НУ «Одеська політехніка» (ОНПУ)
	// ХАІ
	// ЧНУ ім. Федьковича
	// ХНУ ім. Каразіна
	// НТУ «ХПІ»

	// https://dou.ua/lenta/articles/ukrainian-universities-2020/
	// НАУ
	// ВНТУ

	// https://www.linkedin.com/search/results/schools/?keywords=Ukraine
	// https://www.linkedin.com/search/results/schools/?keywords=Ukraine%20OR%20Kyiv%20OR%20%20Kharkiv%20OR%20Lviv%20OR%20Dnipro%20OR%20Vinnytsia

	return []domain.University{
		// УКУ | Український Католицький Університет
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    818029,
				Alias: "ukrainian-catholic-university",
				Name:  "Ukrainian Catholic University",
			},
		},

		// ХНЕУ | Харківський національний економічний університет імені Семена Кузнеця
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    850102,
				Alias: "kharkiv-national-university-of-economics/",
				Name:  "Simon Kuznets Kharkiv National University of Economics",
			},
		},

		// КНУ ім. Шевченка | Київський національний університет імені Тараса Шевченка
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    364340,
				Alias: "taras-shevchenko-national-university-of-kyiv",
				Name:  "Taras Shevchenko National University of Kyiv",
			},
		},

		// НаУКМА | Національний університет «Києво-Могилянська академія»
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    496320,
				Alias: "national-university-of-kyiv-mohyla-academy",
				Name:  "National University of Kyiv-Mohyla Academy",
			},
		},

		// ОНУ ім. Мечникова | Одеський національний університет імені І. І. Мечникова
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1198954,
				Alias: "odessa-national-university",
				Name:  "Odessa I.I.Mechnikov National University",
			},
		},

		// ДУТ | Державний університет телекомунікацій
		// https://www.linkedin.com/company/duikt/

		// СумДУ | Сумський державний університет
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1257361,
				Alias: "sumy-state-university",
				Name:  "Sumy State University",
			},
		},

		// НТУУ «КПІ ім. І. Сікорського» | Київський політехнічний інститут імені Ігоря Сікорського
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15250306,
				Alias: "national-technical-university-of-ukraine-'kyiv-pol",
				Name:  "National Technical University of Ukraine 'Kyiv Polytechnic Institute'",
			},
		},

		// ЛНУ ім. Франка | Львівський національний університет імені Івана Франка
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15251128,
				Alias: "national-university-'ivan-franko'%E2%80%8B-lviv",
				Name:  "Ivan Franko National University of Lviv",
			},
		},

		// ХНУРЕ | Харківський національний університет радіоелектроніки
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15099424,
				Alias: "kharkiv-national-university-of-radioelectronics",
				Name:  "Kharkiv National University of Radioelectronics",
			},
		},

		// НУ «Львівська політехніка» | Національний університет «Львівська політехніка»
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    782774,
				Alias: "lviv-polytechnic-national-university",
				Name:  "Lviv Polytechnic National University",
			},
		},

		// ДНУ ім. Гончара | Дніпровський національний університет імені Олеся Гончара
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101979,
				Alias: "dnipropetrovs'kij-nacional'nij-universitet",
				Name:  "Dnipropetrovs'kij Nacional'nij Universitet",
			},
		},

		// НТУ «Дніпровська політехніка» | Національний технічний університет «Дніпровська політехніка»
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101061,
				Alias: "dnipro-university-of-technology",
				Name:  "Dnipro University of Technology",
			},
		},

		// ХНУ [м. Хмельницький] | Хмельницький національний університет
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    80424966,
				Alias: "khnu",
				Name:  "Khmelnitsky National University",
			},
		},

		// КНЕУ | Київський національний економічний університет імені Вадима Гетьмана
		// https://www.linkedin.com/company/%D1%96%D0%BD%D1%81%D1%82%D0%B8%D1%82%D1%83%D1%82-%D0%B4%D0%B8%D1%81%D1%82%D0%B0%D0%BD%D1%86%D1%96%D0%B9%D0%BD%D0%B8%D1%85-%D1%82%D0%B5%D1%85%D0%BD%D0%BE%D0%BB%D0%BE%D0%B3%D1%96%D0%B9-%D0%BD%D0%B0%D0%B2%D1%87%D0%B0%D0%BD%D0%BD%D1%8F/

		// НУ «Одеська політехніка» (ОНПУ) | Національний університет «Одеська політехніка»
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    6261241,
				Alias: "odessa-national-polytechnic-university",
				Name:  "Odessa National Polytechnic University",
			},
		},

		// ХАІ | Харківський авіаційний інститут
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    658198,
				Alias: "national-aerospace-university",
				Name:  "National Aerospace University -'Kharkiv Aviation Institute'",
			},
		},

		// ЧНУ ім. Федьковича | Чернівецький національний університет імені Юрія Федьковича
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    11443062,
				Alias: "chnu",
				Name:  "Yuriy Fedkovych Chernivtsi National University",
			},
		},

		// ХНУ ім. Каразіна | Харківський національний університет імені В. Н. Каразіна
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15099425,
				Alias: "v.-n.-karazin-kharkiv-national-university",
				Name:  "V. N. Karazin Kharkiv National University",
			},
		},

		// НТУ «ХПІ» | Харківський політехнічний інститут
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15099711,
				Alias: "national-technical-university-kharkiv-polytechnic-institute-",
				Name:  `National Technical University "Kharkiv Polytechnic Institute"`,
			},
		},

		// НАУ | Національний авіаційний університет
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101057,
				Alias: "national-aviation-university",
				Name:  "National Aviation University",
			},
		},

		// ВНТУ | Вінницький національний технічний університет
		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15102004,
				Alias: "vinnytsia-national-technical-university",
				Name:  "Vinnytsia National Technical University",
			},
		},

		// From https://www.linkedin.com/search/results/schools/?keywords=Ukraine

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18080249,
				Alias: "kyivnationaluniversityoftechnologiesanddesignknutd",
				Name:  "Kyiv National University of Technologies and Design (KNUTD)",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15143861,
				Alias: "kyiv-national-university-of-construction-and-archi",
				Name:  "Kyiv National University of Construction and Architecture",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101046,
				Alias: "kyiv-national-economics-university",
				Name:  "Kyiv National Economics University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    1599158,
				Alias: "donetsk-national-university",
				Name:  "Donetsk National University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101060,
				Alias: "national-medical-university-'o.o.-bogomolec'-kyiv",
				Name:  "Bogomolets National Medical University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15100187,
				Alias: "ivano-frankivsk-national-technical-university-of-oil-and-gas",
				Name:  "Ivano-Frankivsk National Technical University of Oil and Gas",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    9029417,
				Alias: "zaporizhzhya-national-technical-university",
				Name:  "Zaporizhzhya National Technical University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    7991636,
				Alias: "donetsk-national-technical-university",
				Name:  "Donetsk National Technical University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101074,
				Alias: "west-ukrainian-national-university",
				Name:  "Ternopil Academy of National Economy",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    27066401,
				Alias: "national-pedagogical-dragomanov-university",
				Name:  "National Pedagogical Dragomanov University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18144134,
				Alias: "national-university-of-food-technologies",
				Name:  "National University of Food Technologies",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15101998,
				Alias: "vasyl-stefanyk-precarpathian-national-university",
				Name:  "Vasyl Stefanyk Precarpathian National University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15149751,
				Alias: "nlu1804",
				Name:  "Yaroslav Mudryi National Law University",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    18691495,
				Alias: "national-university-of-life-and-environmental-sciences-of-ukraine",
				Name:  "National University of Life and Environmental Sciences of Ukraine",
			},
		},

		{
			LinkedInProfile: domain.LinkedInProfile{
				ID:    15099038,
				Alias: "state-university-of-trade-and-economics",
				Name:  "State University of Trade and Economics",
			},
		},

		//{
		//	LinkedInProfile: domain.LinkedInProfile{
		//		ID:    0,
		//		Alias: "",
		//		Name:  "",
		//	},
		//},
	}
}
