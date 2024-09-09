package domain

import "github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"

type OrganizerFeature struct {
	Organizer Organizer
	Feature   dbs.FeatureWait
	Path      string
	Title     string
}

type Organizer struct {
	Alias string
	Title string
	Logo  string
}

var (
	OrganizerGolang = Organizer{
		Alias: "golang",
		Title: "Golang",
		Logo:  "golang-organizer.svg",
	}
	OrganizerRust = Organizer{
		Alias: "rust",
		Title: "Rust",
		Logo:  "rust-organizer.svg",
	}
	OrganizerZig = Organizer{
		Alias: "zig",
		Title: "Zig",
		Logo:  "zig-organizer.svg",
	}
	OrganizerScala = Organizer{
		Alias: "scala",
		Title: "Scala",
		Logo:  "scala-organizer.svg",
	}
	OrganizerElixir = Organizer{
		Alias: "elixir",
		Title: "Elixir",
		Logo:  "elixir-organizer.svg",
	}
	OrganizerClojure = Organizer{
		Alias: "clojure",
		Title: "Clojure",
		Logo:  "clojure-organizer.svg",
	}
)
