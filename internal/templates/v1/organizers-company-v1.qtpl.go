// Code generated by qtc from "organizers-company-v1.qtpl". DO NOT EDIT.
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

func StreamOrganizersCompanyV1(qw422016 *qt422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	company Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	favorite bool,
	stats CompanyStats,
	authQueryParams string,
) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
	<title>`)
	qw422016.E().S(company.Name)
	qw422016.N().S(` – company using `)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` | ReadyToTouch</title>
	<meta name="description" content="Improve your chances of getting a job by connecting with `)
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(` developers from `)
	qw422016.E().S(company.Name)
	qw422016.N().S(` and receiving further recommendations.">

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
	streamorganizersCompanyV1Styles(qw422016)
	qw422016.N().S(`
    <script src="https://cdn.jsdelivr.net/npm/apexcharts@3.52.0"></script>
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
	qw422016.E().S(organizerFeature.Organizer.Logo)
	qw422016.N().S(`"
        alt="organizer logo"
      />
    </a>
    <ul class="header__nav">
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/companies" class="header__nav-link active">Companies</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/vacancies" class="header__nav-link">Vacancies</a>
      </li>
      <li class="header__nav-item">
        <a href="/organizers/`)
	qw422016.E().S(organizerFeature.Organizer.Alias)
	qw422016.N().S(`/communities" class="header__nav-link">Communities</a>
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
		qw422016.E().S(organizerFeature.Organizer.Alias)
		qw422016.N().S(`/welcome`)
		qw422016.E().S(authQueryParams)
		qw422016.N().S(`" class="button button--small-padding button--black header__login-button">Log in</a>
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
	qw422016.E().S(organizerFeature.Organizer.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <a class="breadcrumb__link" href="`)
	qw422016.E().S(organizerFeature.Path)
	qw422016.N().S(`">`)
	qw422016.E().S(organizerFeature.Title)
	qw422016.N().S(`</a>
      </li>
      <li class="breadcrumb__item">
        <span class="breadcrumb__page" aria-current="page">`)
	qw422016.E().S(company.Name)
	qw422016.N().S(`</span>
      </li>
    </ul>
  </nav>
</div>

<section class="companies-cards">
  <div class="container companies-cards__container">
<div class="js-company card"
     data-company-id="`)
	qw422016.N().DL(company.ID)
	qw422016.N().S(`"
     data-company-name="`)
	qw422016.E().S(company.Name)
	qw422016.N().S(`"
     data-company-type="`)
	qw422016.E().S(string(company.Type))
	qw422016.N().S(`"
>
  <aside class="card__action">
    `)
	if favorite {
		qw422016.N().S(`
      <button class="js-company-favorite favorite card__action-button button-group__item in-favorite" title="Remove from favorites">
        <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
          <path
            d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
          />
        </svg>
      </button>
    `)
	} else {
		qw422016.N().S(`
      <button class="js-company-favorite favorite card__action-button button-group__item" title="Add to favorite">
        <svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
          <path
            d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
          />
        </svg>
      </button>
    `)
	}
	qw422016.N().S(`

    `)
	qw422016.N().S(`
  </aside>
  <figure class="card__header">
    <div class="card__image-overlay">
      <img
        width="18"
        height="18"
        class="card__image card__image--preview"
        alt="card image preview icon"
        src="/assets/images/pages/common/image-preview.svg"
      />
    </div>
    <figcaption class="card__header-caption">
      <a href="`)
	qw422016.E().S(company.Website)
	qw422016.N().S(`" class="card__headline vacancy__link">`)
	qw422016.E().S(company.Name)
	qw422016.N().S(`</a>
    </figcaption>
  </figure>
  <div class="card__info">
    <figure class="card__figure">
      <img
        class="card__icon"
        alt="card type icon"
        width="16"
        height="16"
        src="/assets/images/pages/vacancy/building.svg"
      />
      <figcaption class="card__figcaption">`)
	qw422016.E().S(companyTypeName[company.Type])
	qw422016.N().S(`</figcaption>
    </figure>
    `)
	qw422016.N().S(`
  </div>
  <p class="js-company-description card__text">`)
	qw422016.E().S(company.ShortDescription)
	qw422016.N().S(`</p>
  <div class="card__links">
    <ul class="card__links-group">
      <li class="card__links-item">
        <img
          class="card__links-icon"
          alt="linkedin icon"
          width="20"
          height="20"
          src="/assets/images/pages/organizer/linkedin.svg"
        />
        <a href="https://www.linkedin.com/company/`)
	qw422016.E().S(company.LinkedInProfile.Alias)
	qw422016.N().S(`/" class="button-link card__links-link">LinkedIn</a>
      </li>
      <li class="card__links-item">
        <a href="`)
	qw422016.E().S(linkedinConnectionsURL([]Company{company}, ukrainianUniversities))
	qw422016.N().S(`" class="button-link card__links-link">Connections             <img
                class="checkbox__content-image"
                alt="Flag of Ukraine with waves"
                width="16"
                height="16"
                src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
        /></a>
      </li>
      <li class="card__links-item">
        <a href="`)
	qw422016.E().S(linkedinConnectionsURL([]Company{company}, czechUniversities))
	qw422016.N().S(`" class="button-link card__links-link">Connections             <img
                class="checkbox__content-image"
                alt="Flag of Czechia"
                width="16"
                height="16"
                src="/assets/images/pages/online-new/cz_flag.svg"
        /></a>
      </li>
      <li class="card__links-item">
        <a href="`)
	qw422016.E().S(linkedinConnectionsURL([]Company{company}, nil))
	qw422016.N().S(`" class="button-link card__links-link">Connections</a>
      </li>
      <li class="card__links-item">
        <a href="`)
	qw422016.E().S(linkedinJobsURL([]Company{company}, string(organizerFeature.Organizer.LanguageTitleKeywords)))
	qw422016.N().S(`" class="button-link card__links-link">Jobs</a>
      </li>
    </ul>
`)
	if company.GitHubProfile.Login == "" {
		qw422016.N().S(`      <ul class="card__links-group">
        <li class="card__links-item card__links-item--disabled">
          <img
            class="card__links-icon"
            alt="github icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/github.svg"
          />
          <a href="`)
		qw422016.E().S(googleSearchGitHub(company.Name))
		qw422016.N().S(`" class="card__links-link card__links-link--google">
            <img
              class="card__links-icon card__links-icon--google"
              alt="google icon"
              width="20"
              height="20"
              src="/assets/images/pages/organizer/google.svg"
            />
          </a>
          <span class="button-link card__links-link">GitHub</span>
        </li>
      </ul>
    `)
	} else {
		qw422016.N().S(`
      <ul class="card__links-group">
        <li class="card__links-item">
          <img
            class="card__links-icon"
            alt="github icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/github.svg"
          />
          <a href="https://github.com/`)
		qw422016.E().S(company.GitHubProfile.Login)
		qw422016.N().S(`" class="button-link card__links-link">GitHub</a>
        </li>
        <li class="card__links-item">
          `)
		qw422016.N().S(`
          <a href="https://github.com/orgs/`)
		qw422016.E().S(company.GitHubProfile.Login)
		qw422016.N().S(`/repositories?q=lang:go" class="button-link card__links-link" title="`)
		qw422016.N().D(company.GitHubProfile.GoRepositoryCount)
		qw422016.N().S(` repositories">Repositories</a>
        </li>
      </ul>
`)
	}
	if company.GlassdoorProfile.OverviewURL == "" {
		qw422016.N().S(`      <ul class="card__links-group">
        <li class="card__links-item card__links-item--disabled">
          <img
            class="card__links-icon"
            alt="glassdoor icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/glassdoor.svg"
          />
          <a href="`)
		qw422016.E().S(googleSearchGlassdoor(company.Name))
		qw422016.N().S(`" class="card__links-link card__links-link--google">
            <img
              class="card__links-icon card__links-icon--google"
              alt="google icon"
              width="20"
              height="20"
              src="/assets/images/pages/organizer/google.svg"
            />
          </a>
          <span class="button-link card__links-link">Glassdoor</span>
        </li>
      </ul>
    `)
	} else {
		qw422016.N().S(`
      <ul class="card__links-group">
        <li class="card__links-item">
          <img
            class="card__links-icon"
            alt="glassdoor icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/glassdoor.svg"
          />
          <a href="`)
		qw422016.E().S(company.GlassdoorProfile.OverviewURL)
		qw422016.N().S(`" class="button-link card__links-link">Glassdoor</a>
        </li>
        <li class="card__links-item">
          <a href="`)
		qw422016.E().S(company.GlassdoorProfile.ReviewsURL)
		qw422016.N().S(`" class="button-link card__links-link">Reviews</a>
        </li>
      </ul>
`)
	}
	qw422016.N().S(`    <ul class="card__links-group card__links-group--unbordered">
      <li class="card__links-item">
        <img
          class="card__links-icon"
          alt="similarweb icon"
          width="20"
          height="20"
          src="/assets/images/pages/organizer/similarweb.svg"
        />
        <a href="`)
	qw422016.E().S(similarwebURL(company.Website))
	qw422016.N().S(`" class="button-link card__links-link">SimilarWeb</a>
      </li>
      <li class="card__links-item">
        <img
          class="card__links-icon"
          alt="whois icon"
          width="20"
          height="20"
          src="/assets/images/pages/organizer/whois.svg"
        />
        <a href="`)
	qw422016.E().S(whoisURL(company.Website))
	qw422016.N().S(`" class="button-link card__links-link">Whois</a>
      </li>

      `)
	qw422016.N().S(`
      <li class="card__links-item card__links-item--disabled">
        <img
          class="card__links-icon"
          alt="xing icon"
          width="20"
          height="20"
          src="/assets/images/pages/organizer/xing.svg"
        />
        <a href="`)
	qw422016.E().S(googleSearchXing(company.Name))
	qw422016.N().S(`" class="card__links-link card__links-link--google">
          <img
            class="card__links-icon card__links-icon--google"
            alt="google icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/google.svg"
          />
        </a>
        <span class="button-link card__links-link">XING</span>
      </li>

`)
	if company.OttaProfileSlug == "" {
		qw422016.N().S(`        <li class="card__links-item card__links-item--disabled">
          <img
            class="card__links-icon"
            alt="otta icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/otta.svg"
          />
          <a href="`)
		qw422016.E().S(googleSearchOtta(company.Name))
		qw422016.N().S(`" class="card__links-link card__links-link--google">
            <img
              class="card__links-icon card__links-icon--google"
              alt="google icon"
              width="20"
              height="20"
              src="/assets/images/pages/organizer/google.svg"
            />
          </a>
          <span class="button-link card__links-link">Otta</span>
        </li>
      `)
	} else {
		qw422016.N().S(`
        <li class="card__links-item">
          <img
            class="card__links-icon"
            alt="otta icon"
            width="20"
            height="20"
            src="/assets/images/pages/organizer/otta.svg"
          />
          <a href="https://app.otta.com/companies/`)
		qw422016.E().S(company.OttaProfileSlug)
		qw422016.N().S(`" class="button-link card__links-link">Otta</a>
        </li>
`)
	}
	qw422016.N().S(`    </ul>
  </div>
</div>
</div>
</section>

<section class="company-statistics">
  <div class="container company-statistics__container">
    <h2 class="company-statistics__headline">Statistics</h2>
    <div class="stats">
      <header class="stats__head">
        <div class="stats__counters">
          <h3 class="stats__title">Views</h3>
          <div class="stats__counters-group">
            <div class="stats__counters-item">
              <p class="stats__counters-views">Total</p>
              <p class="stats__counters-item-number">`)
	qw422016.N().DL(stats.TotalViews)
	qw422016.N().S(`</p>
            </div>
            <div class="stats__counters-item">
              <p class="stats__counters-views">Last month</p>
              <p class="stats__counters-item-number">`)
	qw422016.N().DL(stats.LastMonthViews)
	qw422016.N().S(`</p>
            </div>
          </div>
        </div>
        <div class="stats__counters">
          <h3 class="stats__title">Added to favorites</h3>
          <div class="stats__counters-group">
            <div class="stats__counters-item">
              <p class="stats__counters-views">Total</p>
              <p class="stats__counters-item-number">`)
	qw422016.N().DL(stats.TotalFavorites)
	qw422016.N().S(`</p>
            </div>
            <div class="stats__counters-item">
              <p class="stats__counters-views">Last month</p>
              <p class="stats__counters-item-number">`)
	qw422016.N().DL(stats.LastMonthFavorites)
	qw422016.N().S(`</p>
            </div>
          </div>
        </div>
      </header>
      <div class="stats__item">
        <div class="stats__header">
          <h3 class="stats__title">Company website views</h3>
          <div class="stats__period">Last month</div>
        </div>
        <div class="stats__chart stats__chart--page-views-statistics js-chart-views-statistics"></div>
      </div>
    </div>
  </div>
</section>


</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`
<script src="/assets/js/organizers-company-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersCompanyV1(qq422016 qtio422016.Writer,
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	company Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	favorite bool,
	stats CompanyStats,
	authQueryParams string,
) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersCompanyV1(qw422016, organizerFeature, headerProfiles, company, ukrainianUniversities, czechUniversities, favorite, stats, authQueryParams)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersCompanyV1(
	organizerFeature OrganizerFeature,
	headerProfiles []SocialProviderUser,
	company Company,
	ukrainianUniversities []University,
	czechUniversities []University,
	favorite bool,
	stats CompanyStats,
	authQueryParams string,
) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersCompanyV1(qb422016, organizerFeature, headerProfiles, company, ukrainianUniversities, czechUniversities, favorite, stats, authQueryParams)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
