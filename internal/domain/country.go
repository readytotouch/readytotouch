package domain

type Country = AliasName

var (
	Ukraine = Country{
		Alias: "ukraine",
		Name:  "Ukraine",
	}
	Czechia = Country{
		Alias: "czechia",
		Name:  "Czechia",
	}
)
