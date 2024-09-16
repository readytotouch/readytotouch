package domain

import "github.com/readytotouch/readytotouch/internal/storage/postgres/dbs"

type OrganizerFeature struct {
	Organizer Organizer
	Feature   dbs.FeatureWait
	Path      string
	Title     string
}

type Organizer struct {
	Language              Language
	LanguageTitleKeywords LanguageTitleKeywords
	Alias                 string
	Title                 string
	Logo                  string
}

var (
	OrganizerGolang = Organizer{
		Language:              Go,
		LanguageTitleKeywords: GoTitleKeywords,
		Alias:                 "golang",
		Title:                 "Golang",
		Logo:                  "golang-organizer.svg",
	}
	OrganizerRust = Organizer{
		Language:              Rust,
		LanguageTitleKeywords: RustTitleKeywords,
		Alias:                 "rust",
		Title:                 "Rust",
		Logo:                  "rust-organizer.svg",
	}
	OrganizerZig = Organizer{
		Language:              Zig,
		LanguageTitleKeywords: ZigTitleKeywords,
		Alias:                 "zig",
		Title:                 "Zig",
		Logo:                  "zig-organizer.svg",
	}
	OrganizerScala = Organizer{
		Language:              Scala,
		LanguageTitleKeywords: ScalaTitleKeywords,
		Alias:                 "scala",
		Title:                 "Scala",
		Logo:                  "scala-organizer.svg",
	}
	OrganizerElixir = Organizer{
		Language:              Elixir,
		LanguageTitleKeywords: ElixirTitleKeywords,
		Alias:                 "elixir",
		Title:                 "Elixir",
		Logo:                  "elixir-organizer.svg",
	}
	OrganizerClojure = Organizer{
		Language:              Clojure,
		LanguageTitleKeywords: ClojureTitleKeywords,
		Alias:                 "clojure",
		Title:                 "Clojure",
		Logo:                  "clojure-organizer.svg",
	}
)
