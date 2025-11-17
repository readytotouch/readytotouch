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
	GitHubAlias           string
	Title                 string
	Logo                  string
	Keywords              string // for SEO
	Description           string // for SEO
}

var (
	OrganizerGolang = Organizer{
		Language:              Go,
		LanguageTitleKeywords: GoTitleKeywords,
		Alias:                 "golang",
		GitHubAlias:           "go",
		Title:                 "Golang",
		Logo:                  "go.svg", // go-original.svg
		Keywords:              "Go, Golang, GoLand, Companies using Go, Companies using Golang, Companies using GoLand, Go companies, Golang companies, GoLand companies, Go connections, Golang connections, GoLand connections",
		Description:           "ReadyToTouch helps you find companies that use Golang. Improve your chances of getting a job by connecting with Go developers and receiving further recommendations.",
	}
	OrganizerRust = Organizer{
		Language:              Rust,
		LanguageTitleKeywords: RustTitleKeywords,
		Alias:                 "rust",
		GitHubAlias:           "rust",
		Title:                 "Rust",
		Logo:                  "rust.svg",
		Keywords:              "Rust, RustRover, Companies using Rust, Companies using RustRover, Rust companies, RustRover companies, Rust connections, RustRover connections",
		Description:           "ReadyToTouch helps you find companies that use Rust. Improve your chances of getting a job by connecting with Rust developers and receiving further recommendations.",
	}
	OrganizerZig = Organizer{
		Language:              Zig,
		LanguageTitleKeywords: ZigTitleKeywords,
		Alias:                 "zig",
		GitHubAlias:           "zig",
		Title:                 "Zig",
		Logo:                  "zig.svg",
		Keywords:              "Zig, Companies using Zig, Zig companies, Zig connections",
		Description:           "ReadyToTouch helps you find companies that use Zig. Improve your chances of getting a job by connecting with Zig developers and receiving further recommendations.",
	}
	OrganizerScala = Organizer{
		Language:              Scala,
		LanguageTitleKeywords: ScalaTitleKeywords,
		Alias:                 "scala",
		GitHubAlias:           "scala",
		Title:                 "Scala",
		Logo:                  "scala.svg",
		Keywords:              "Scala, Companies using Scala, Scala companies, Scala connections",
		Description:           "ReadyToTouch helps you find companies that use Scala. Improve your chances of getting a job by connecting with Scala developers and receiving further recommendations.",
	}
	OrganizerElixir = Organizer{
		Language:              Elixir,
		LanguageTitleKeywords: ElixirTitleKeywords,
		Alias:                 "elixir",
		GitHubAlias:           "elixir",
		Title:                 "Elixir",
		Logo:                  "elixir.svg",
		Keywords:              "Elixir, Erlang, Companies using Elixir, Companies using Erlang, Elixir companies, Erlang companies, Elixir connections, Erlang connections",
		Description:           "ReadyToTouch helps you find companies that use Elixir. Improve your chances of getting a job by connecting with Elixir developers and receiving further recommendations.",
	}
	OrganizerClojure = Organizer{
		Language:              Clojure,
		LanguageTitleKeywords: ClojureTitleKeywords,
		Alias:                 "clojure",
		GitHubAlias:           "clojure",
		Title:                 "Clojure",
		Logo:                  "clojure.svg",
		Keywords:              "Clojure, Companies using Clojure, Clojure companies, Clojure connections",
		Description:           "ReadyToTouch helps you find companies that use Clojure. Improve your chances of getting a job by connecting with Clojure developers and receiving further recommendations.",
	}
	OrganizerHaskell = Organizer{
		Language:              Haskell,
		LanguageTitleKeywords: HaskellTitleKeywords,
		Alias:                 "haskell",
		GitHubAlias:           "haskell",
		Title:                 "Haskell",
		Logo:                  "haskell.svg",
		Keywords:              "Haskell, Companies using Haskell, Haskell companies, Haskell connections",
		Description:           "ReadyToTouch helps you find companies that use Haskell. Improve your chances of getting a job by connecting with Haskell developers and receiving further recommendations.",
	}
	OrganizerErlang = Organizer{
		Language:              Erlang,
		LanguageTitleKeywords: ErlangTitleKeywords,
		Alias:                 "erlang",
		GitHubAlias:           "erlang",
		Title:                 "Erlang",
		Logo:                  "erlang.svg",
		Keywords:              "Erlang, Companies using Erlang, Erlang companies, Erlang connections",
		Description:           "ReadyToTouch helps you find companies that use Erlang. Improve your chances of getting a job by connecting with Erlang developers and receiving further recommendations.",
	}
	OrganizerFSharp = Organizer{
		Language:              FSharp,
		LanguageTitleKeywords: FSharpTitleKeywords,
		Alias:                 "fsharp",
		GitHubAlias:           "f#",
		Title:                 "F#",
		Logo:                  "fsharp.svg",
		Keywords:              "F#, Companies using F#, F# companies, F# connections",
		Description:           "ReadyToTouch helps you find companies that use F#. Improve your chances of getting a job by connecting with F# developers and receiving further recommendations.",
	}
	OrganizerOCaml = Organizer{
		Language:              OCaml,
		LanguageTitleKeywords: OCamlTitleKeywords,
		Alias:                 "ocaml",
		GitHubAlias:           "ocaml",
		Title:                 "OCaml",
		Logo:                  "ocaml.svg",
		Keywords:              "OCaml, Companies using OCaml, OCaml companies, OCaml connections",
		Description:           "ReadyToTouch helps you find companies that use OCaml. Improve your chances of getting a job by connecting with OCaml developers and receiving further recommendations.",
	}
	OrganizerGleam = Organizer{
		Language:              Gleam,
		LanguageTitleKeywords: GleamTitleKeywords,
		Alias:                 "gleam",
		GitHubAlias:           "gleam",
		Title:                 "Gleam",
		Logo:                  "gleam.svg",
		Keywords:              "Gleam, Companies using Gleam, Gleam companies, Gleam connections",
		Description:           "ReadyToTouch helps you find companies that use Gleam. Improve your chances of getting a job by connecting with Gleam developers and receiving further recommendations.",
	}
	OrganizerMojo = Organizer{
		Language:              Mojo,
		LanguageTitleKeywords: MojoTitleKeywords,
		Alias:                 "mojo",
		GitHubAlias:           "mojo",
		Title:                 "Mojo",
		Logo:                  "mojo.svg",
		Keywords:              "Mojo, Companies using Mojo, Mojo companies, Mojo connections",
		Description:           "ReadyToTouch helps you find companies that use Mojo. Improve your chances of getting a job by connecting with Mojo developers and receiving further recommendations.",
	}
)
