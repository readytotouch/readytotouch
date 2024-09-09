// Code generated by qtc from "organizers-welcome.qtpl". DO NOT EDIT.
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

func StreamOrganizersWelcome(qw422016 *qt422016.Writer, organizer Organizer, authQueryParams string) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>Yet another anonymous work search</title>
	<meta name="description" content="Yet another anonymous work search">

	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersFonts(qw422016)
	qw422016.N().S(`
    `)
	streamorganizersWelcomeStyles(qw422016)
	qw422016.N().S(`
</head>

<body class="organizer-welcome">
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/" class="header__logo">
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
    <div class="header__stars">
      <iframe
        src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch.github.io&type=star&count=true&size=large"
        width="110"
        height="33"
        title="GitHub"
      ></iframe>
    </div>
  </div>
</header>



<section class="hero">
  <div class="hero__wrapper">
    <div class="hero__main">
      <h1 class="hero__title">Welcome!</h1>
      <p class="hero__subtitle">Our projects are ReadyToTouch, u8views, Organizer, and others. Log in to support.</p>
    </div>
    <div class="hero__buttons-group">
      <a href="/auth/github`)
	qw422016.E().S(authQueryParams)
	qw422016.N().S(`" class="button button--bordered-gray hero__button">
        <img
          width="20"
          height="20"
          src="/assets/images/pages/online-new/github-black.svg"
          alt="GitHub logo"
          class="button-icon"
        />Log in with GitHub
      </a>
      <a href="/auth/gitlab`)
	qw422016.E().S(authQueryParams)
	qw422016.N().S(`" class="button button--bordered-gray hero__button">
        <img
          width="20"
          height="20"
          src="/assets/images/pages/online-new/gitlab.svg"
          alt="GitLab logo"
          class="hero__button-icon"
        />Log in with GitLab
      </a>
      <a href="/auth/bitbucket`)
	qw422016.E().S(authQueryParams)
	qw422016.N().S(`" class="button button--bordered-gray hero__button">
        <img
          width="21"
          height="20"
          src="/assets/images/pages/online-new/bitbucket.svg"
          alt="Bitbucket logo"
          class="hero__button-icon"
        />Log in with Bitbucket
      </a>
    </div>
  </div>
</section>


</main>

</body>
</html>
`)
}

func WriteOrganizersWelcome(qq422016 qtio422016.Writer, organizer Organizer, authQueryParams string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersWelcome(qw422016, organizer, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersWelcome(organizer Organizer, authQueryParams string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersWelcome(qb422016, organizer, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}