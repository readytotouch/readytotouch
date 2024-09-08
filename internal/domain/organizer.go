package domain

import "github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"

type OrganizerFeature struct {
	Organizer Organizer
	Feature   dbs.FeatureWait
	Path      string
	Title     string
}

type Organizer struct {
	Title string
	Logo  string
}

var (
	OrganizerGolang = Organizer{
		Title: "Golang",
		Logo:  "golang-organizer.svg",
	}
	OrganizerRust = Organizer{
		Title: "Rust",
		Logo:  "rust-organizer.svg",
	}
	OrganizerZig = Organizer{
		Title: "Zig",
		Logo:  "zig-organizer.svg",
	}
	OrganizerScala = Organizer{
		Title: "Scala",
		Logo:  "scala-organizer.svg",
	}
	OrganizerElixir = Organizer{
		Title: "Elixir",
		Logo:  "elixir-organizer.svg",
	}
	OrganizerClojure = Organizer{
		Title: "Clojure",
		Logo:  "clojure-organizer.svg",
	}
)
