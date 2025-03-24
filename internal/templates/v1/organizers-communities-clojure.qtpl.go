// Code generated by qtc from "organizers-communities-clojure.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

package v1

import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

func StreamOrganizersCommunitiesClojure(qw422016 *qt422016.Writer, organizer Organizer, headerProfiles []SocialProviderUser) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>`)
	qw422016.E().S(organizer.Title)
	qw422016.N().S(` Communities | ReadyToTouch</title>
	<meta name="title" content="`)
	qw422016.E().S(organizer.Title)
	qw422016.N().S(` Communities | ReadyToTouch">
	<meta name="description" content="">
    <meta name="keywords" content="">

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="author" type="text/plain" href="https://readytotouch.com/humans.txt"/>
	<meta property="og:image" content="/assets/images/og/organizers-light.jpg">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersCommunitiesStyles(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
	qw422016.N().S(`
</head>

<body class="organizer-companies-inner">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/organizers" class="header__logo">
      <img
        width="129"
        height="32"
        class="header__logo-img"
        src="/assets/images/pages/organizer/`)
	qw422016.E().S(organizer.Logo)
	qw422016.N().S(`"
        alt="organizer logo"
      />
    </a>
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizer.Alias)
	qw422016.N().S(`/companies" class="header__nav-link">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizer.Alias)
	qw422016.N().S(`/jobs" class="header__nav-link">Jobs</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizer.Alias)
	qw422016.N().S(`/communities" class="header__nav-link active">Communities</a>
      </li>
    </ul>

    `)
	streamorganizersHeaderStars(qw422016)
	qw422016.N().S(`
    `)
	if len(headerProfiles) > 0 {
		qw422016.N().S(`
    `)
		streamorganizersHeaderProfile(qw422016, headerProfiles)
		qw422016.N().S(`
    `)
	} else {
		qw422016.N().S(`
    <a href="/organizers/`)
		qw422016.E().S(organizer.Alias)
		qw422016.N().S(`/welcome" class="button button--small-padding button--black header__login-button">Log in</a>
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<div class="container">
  <nav aria-label="breadcrumb" aria-labelledby="navigation through the bread crumbs" class="breadcrumb">
    <ul class="breadcrumb__list">
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/">Main</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="/organizers">Organizers</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="javascript:void(0);">`)
	qw422016.E().S(organizer.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <span class="breadcrumb__page" aria-current="page">Communities</span>
      </li>
    </ul>
  </nav>
</div>

<section class="organizer-communities">
	<div class="container organizer-communities__container-block">
		<div class="hero">
			<div class="hero__fill hero__fill--clojure"></div>
			<div class="organizer-communities__headline headline">Clojure Developers Communities</div>
		</div>
	</div>
</section>
<section class="organizer-communities">
	<div class="container organizer-communities__container">
		<p class="organizer-communities__count"><strong>2</strong> Communities</p>
		<div class="organizer-communities__list">
			<div class="card">
				<figure class="card__header">
					<div class="card__image-overlay card__image-overlay--small">
						<img
							width="112"
							height="56"
							class="card__image card__image--preview"
							alt="card image preview icon"
                            src="/assets/images/pages/organizer/clojure-communities-1.png"
						/>
					</div>
					<figcaption class="card__header-caption">
						<a href="https://www.reddit.com/r/Clojure/" target="_blank" class="card__headline vacancy__link">r/Clojure</a>
					</figcaption>
					<a href="https://www.reddit.com/r/Clojure/" target="_blank" class="button button--small-padding button--black button--gap-images">
						Visit
						<img width="18" height="18" alt="arrow top icon" src="/assets/images/pages/common/external-link-white.svg">
					</a>
				</figure>
				<p class="card__sub-headline">
					Clojure is a dynamic, general-purpose programming language, combining the approachability and interactive development of a scripting language with an efficient and robust infrastructure for multithreaded programming.
				</p>
			</div>

			<div class="card">
				<figure class="card__header">
					<div class="card__image-overlay card__image-overlay--small">
						<img
							width="112"
							height="56"
							class="card__image card__image--preview"
							alt="card image preview icon"
							src="/assets/images/pages/organizer/clojure-communities-1.png"
						/>
					</div>
					<figcaption class="card__header-caption">
						<a href="http://clojurians.net/" target="_blank" class="card__headline vacancy__link">Clojurians</a>
					</figcaption>
					<a href="http://clojurians.net/" target="_blank" class="button button--small-padding button--black button--gap-images">
						Visit
						<img width="18" height="18" alt="arrow top icon" src="/assets/images/pages/common/external-link-white.svg">
					</a>
				</figure>
				<p class="card__sub-headline">
                    The Clojure(script) slack community
				</p>
			</div>

		</div>
	</div>
</section>
<section class="organizer-communities">
	<div class="container organizer-communities__container">
		<div class="contact__empty">
			<img class="contact__empty-logo" width="64" height="64" src="/assets/images/pages/organizer/share.png" alt="share">
			<h2 class="contact__empty-title">Know a great community that should be added to this list?</h2>
			<a class="button button--small-padding button--black" href="https://www.linkedin.com/in/yaroslav-podorvanov/" target="_blank">Contact me!</a>
		</div>
	</div>
</section>
`)
	qw422016.N().S(`

</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`
<script src="/assets/js/organizers-communities-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersCommunitiesClojure(qq422016 qtio422016.Writer, organizer Organizer, headerProfiles []SocialProviderUser) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersCommunitiesClojure(qw422016, organizer, headerProfiles)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersCommunitiesClojure(organizer Organizer, headerProfiles []SocialProviderUser) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersCommunitiesClojure(qb422016, organizer, headerProfiles)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
