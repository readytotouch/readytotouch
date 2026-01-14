const GoTitleKeywords =
    `"Golang Engineer" OR "Golang Software Engineer" OR "Golang Developer" OR "Go Engineer" OR "Go Software Engineer" OR "Go Developer"`;

const RustTitleKeywords =
    `"Rust Engineer" OR "Rust Software Engineer" OR "Rust Developer"`;

const ZigTitleKeywords =
    `"Zig Engineer" OR "Zig Software Engineer" OR "Zig Developer"`;

const ScalaTitleKeywords =
    `"Scala Engineer" OR "Scala Software Engineer" OR "Scala Developer"`;

const ElixirTitleKeywords =
    `"Elixir Engineer" OR "Elixir Software Engineer" OR "Elixir Developer" OR "Erlang Engineer" OR "Erlang Software Engineer" OR "Erlang Developer"`;

const ClojureTitleKeywords =
    `"Clojure Engineer" OR "Clojure Software Engineer" OR "Clojure Developer"`;

const HaskellTitleKeywords =
    `"Haskell Engineer" OR "Haskell Software Engineer" OR "Haskell Developer"`;

const ErlangTitleKeywords =
    `"Erlang Engineer" OR "Erlang Software Engineer" OR "Erlang Developer"`;

const FSharpTitleKeywords =
    `"F# Engineer" OR "F# Software Engineer" OR "F# Developer"`;

const OCamlTitleKeywords =
    `"OCaml Engineer" OR "OCaml Software Engineer" OR "OCaml Developer"`;

const GleamTitleKeywords =
    `"Gleam Engineer" OR "Gleam Software Engineer" OR "Gleam Developer"`;

const MojoTitleKeywords =
    `"Mojo Engineer" OR "Mojo Software Engineer" OR "Mojo Developer"`;

export class Organizer {
    constructor(
        public readonly alias: string,
        public readonly github_alias: string,
        public readonly title: string,
        public readonly language_title_keywords: string,
    ) {
    }
}

const OrganizerGolang = new Organizer(
    "golang",
    "go",
    "Golang",
    GoTitleKeywords,
);

const OrganizerRust = new Organizer(
    "rust",
    "rust",
    "Rust",
    RustTitleKeywords,
);

const OrganizerZig = new Organizer(
    "zig",
    "zig",
    "Zig",
    ZigTitleKeywords,
);

const OrganizerScala = new Organizer(
    "scala",
    "scala",
    "Scala",
    ScalaTitleKeywords,
);

const OrganizerElixir = new Organizer(
    "elixir",
    "elixir",
    "Elixir",
    ElixirTitleKeywords,
);

const OrganizerClojure = new Organizer(
    "clojure",
    "clojure",
    "Clojure",
    ClojureTitleKeywords,
);

const OrganizerHaskell = new Organizer(
    "haskell",
    "haskell",
    "Haskell",
    HaskellTitleKeywords,
);

const OrganizerErlang = new Organizer(
    "erlang",
    "erlang",
    "Erlang",
    ErlangTitleKeywords,
);

const OrganizerFSharp = new Organizer(
    "fsharp",
    "f#",
    "F#",
    FSharpTitleKeywords,
);

const OrganizerOCaml = new Organizer(
    "ocaml",
    "ocaml",
    "OCaml",
    OCamlTitleKeywords,
);

const OrganizerGleam = new Organizer(
    "gleam",
    "gleam",
    "Gleam",
    GleamTitleKeywords,
);

const OrganizerMojo = new Organizer(
    "mojo",
    "mojo",
    "Mojo",
    MojoTitleKeywords,
);

const organizers: Record<string, Organizer> = {
    [OrganizerGolang.alias]: OrganizerGolang,
    [OrganizerRust.alias]: OrganizerRust,
    [OrganizerZig.alias]: OrganizerZig,
    [OrganizerScala.alias]: OrganizerScala,
    [OrganizerElixir.alias]: OrganizerElixir,
    [OrganizerClojure.alias]: OrganizerClojure,
    [OrganizerHaskell.alias]: OrganizerHaskell,
    [OrganizerErlang.alias]: OrganizerErlang,
    [OrganizerFSharp.alias]: OrganizerFSharp,
    [OrganizerOCaml.alias]: OrganizerOCaml,
    [OrganizerGleam.alias]: OrganizerGleam,
    [OrganizerMojo.alias]: OrganizerMojo,
};

export function parseCurrentOrganizerAlias(pathname: string): string {
    const index = pathname.indexOf("/", 2);

    return pathname.substring(1, index);
}

export function findOrganizer(alias: string): Organizer {
    return organizers[alias];
}
