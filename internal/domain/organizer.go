package domain

import "github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"

type OrganizerFeature struct {
	Organizer Organizer
	Feature   dbs.FeatureWait
	Path      string
	Title     string
}

type Organizer struct {
	Language Language
	Alias    string
	Title    string
	Logo     string
}

var (
	OrganizerGolang = Organizer{
		Language: Go,
		Alias:    "golang",
		Title:    "Golang",
		Logo:     "golang-organizer.svg",
	}
	OrganizerRust = Organizer{
		Language: Rust,
		Alias:    "rust",
		Title:    "Rust",
		Logo:     "rust-organizer.svg",
	}
	OrganizerZig = Organizer{
		Language: Zig,
		Alias:    "zig",
		Title:    "Zig",
		Logo:     "zig-organizer.svg",
	}
	OrganizerScala = Organizer{
		Language: Scala,
		Alias:    "scala",
		Title:    "Scala",
		Logo:     "scala-organizer.svg",
	}
	OrganizerElixir = Organizer{
		Language: Elixir,
		Alias:    "elixir",
		Title:    "Elixir",
		Logo:     "elixir-organizer.svg",
	}
	OrganizerClojure = Organizer{
		Language: Clojure,
		Alias:    "clojure",
		Title:    "Clojure",
		Logo:     "clojure-organizer.svg",
	}
)
