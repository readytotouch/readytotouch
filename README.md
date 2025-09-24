# ReadyToTouch
A platform for simplifying job search

## Organizers:
- [Companies using Go](https://readytotouch.com/golang/companies)
- [Companies using Rust](https://readytotouch.com/rust/companies)
- [Companies using Scala](https://readytotouch.com/scala/companies)
- [Companies using Elixir](https://readytotouch.com/elixir/companies)
- [Companies using Erlang](https://readytotouch.com/erlang/companies)
- [Companies using Clojure](https://readytotouch.com/clojure/companies)
- [Companies using Zig](https://readytotouch.com/zig/companies)
- [Companies using Haskell](https://readytotouch.com/haskell/companies)
- [Companies using F#](https://readytotouch.com/fsharp/companies)
- [Companies using OCaml](https://readytotouch.com/ocaml/companies)

## Articles
- Forum | [My list of companies that use Rust](https://users.rust-lang.org/t/my-list-of-companies-that-use-rust/127300)
- Reddit | [My list of companies that use Rust](https://www.reddit.com/r/rust/comments/1jg4rrl/my_list_of_companies_that_use_rust/)
- Reddit | [My list of companies that use Golang 2.0](https://www.reddit.com/r/golang/comments/1ixglek/my_list_of_companies_that_use_golang_20/)
- Reddit | [My list of companies that use Golang](https://www.reddit.com/r/golang/comments/1fjbp1p/my_list_of_companies_that_use_golang/)

## Analytics
- https://readytotouch.com/plausible
- https://readytotouch.com/similarweb

## Research repositories:
- https://github.com/doutivity/research-online-redis-go
- https://github.com/doutivity/research-online-postgres-go

## Local development:
```bash
make env-docker-compose-development
make env-up
make pg-fixtures
make app-restart
```

## Contributions
**The project is not open for contributions for now.** 

There are several reasons for that:
1. I’m planning to monetize the project, for example by pinning one company at the top of the company list.
2. I have a team already working with me and a well-established workflow — additional help would only slow things down.
3. Currently, there’s a task to rebuild half of the project: [ReadyToTouch v3 — Responsive web design — Epic](https://github.com/readytotouch/readytotouch/issues/157)

Thank you for your interest and understanding!

## FAQ | Why am I working on this project?

I currently have a stable job in a MedTech project at DocHQ. But because there are many news about layoffs and many people say that finding a job is now harder, I thought about what I can do now, while I have time, to make my future job search easier and less stressful.

So I started making a list of companies that use Golang and adding developers from these companies to my contacts with a message like:  "Hi, I also work with Golang and am interested in ExampleCompanyName for future opportunities".

To improve the project, I organized it better and started collecting feedback from the Ukrainian Golang community. Later, I added support for Rust, Scala, and Elixir.
