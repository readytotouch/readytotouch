// Code generated by qtc from "organizers-online.qtpl". DO NOT EDIT.
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

func StreamOrganizersOnline(qw422016 *qt422016.Writer,
	headerProfiles []SocialProviderUser,
	socialProviderUsers []SocialProviderUser,
) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
    <title>ReadyToTouch — Simplify Job Searching by Connecting with Developers</title>
    <meta name="description" content="ReadyToTouch is a platform that simplifies your job search. Connect directly with developers and boost your chances of getting hired through recommendations.">

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
	streamorganizersOnlineStyles(qw422016)
	qw422016.N().S(`
    <script src="https://cdn.jsdelivr.net/npm/apexcharts@3.52.0"></script>
    `)
	streamplausibleAnalytics(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
	qw422016.N().S(`
</head>

<body>
<main class="main-wrapper">
	<header class="header">
  <div class="header__wrapper">
    <a href="/" class="header__logo">
      <img width="40" height="40" class="header__logo-img" src="/assets/images/pages/online/logo.svg" alt="logo" />
      <h3 class="header__logo-title">ReadyToTouch</h3>
    </a>
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
    `)
	}
	qw422016.N().S(`
  </div>
</header>

<section class="hero">
  <div class="hero__wrapper">
    <img
      class="hero__top-flag"
      alt="Flag of Ukraine with waves"
      width="24"
      height="24"
      src="/assets/images/pages/online-new/ua_flag_with_waves.svg"
    />
    `)
	if len(headerProfiles) == 0 {
		qw422016.N().S(`
    <div class="hero__main">
      <h1 class="hero__title">ReadyToTouch is a space to help you find jobs</h1>
      <p class="hero__subtitle">Our projects are u8views, Organizer, and others. Log in to support.</p>
    </div>
    <div class="hero__buttons-group">
      <a href="/auth/github" class="button button--bordered-gray hero__button" rel="nofollow">
        <img
          width="20"
          height="20"
          src="/assets/images/pages/online-new/github-black.svg"
          alt="GitHub logo"
          class="button-icon"
        />Log in with GitHub
      </a>
      <a href="/auth/gitlab" class="button button--bordered-gray hero__button" rel="nofollow">
        <img
          width="20"
          height="20"
          src="/assets/images/pages/online-new/gitlab.svg"
          alt="GitLab logo"
          class="hero__button-icon"
        />Log in with GitLab
      </a>
      <a href="/auth/bitbucket" class="button button--bordered-gray hero__button" rel="nofollow">
        <img
          width="21"
          height="20"
          src="/assets/images/pages/online-new/bitbucket.svg"
          alt="Bitbucket logo"
          class="hero__button-icon"
        />Log in with Bitbucket
      </a>
    </div>
    `)
	} else {
		qw422016.N().S(`
    <div class="hero__main">
      <h1 class="hero__title">ReadyToTouch is a space to help you find jobs</h1>
      <p class="hero__subtitle">Our projects are u8views, Organizer, and others.</p>
    </div>
    `)
	}
	qw422016.N().S(`
  </div>
</section>

<section class="top-projects">
  <div class="container top-projects__container">
    <div class="top-projects__group">
      <a class="top-projects__link" href="/organizers/golang/companies">
        <figure class="top-projects__figure">
          <img
            class="top-projects__image"
            width="129"
            height="32"
            alt="Golang organizer logo"
            src="/assets/images/pages/organizer/golang-organizer.svg"
          />
        </figure>
      </a>
      <a class="top-projects__link" href="/organizers/rust/companies">
        <figure class="top-projects__figure">
          <img
            class="top-projects__image"
            width="129"
            height="32"
            alt="Rust organizer logo"
            src="/assets/images/pages/organizer/rust-organizer.svg"
          />
        </figure>
      </a>
      <a class="top-projects__link" href="https://u8views.com/" target="_blank">
        <figure class="top-projects__figure">
          <img
            class="top-projects__image"
            width="109"
            height="35"
            alt="u8views logo"
            src="/assets/images/pages/online-new/u8views.svg"
          />
        </figure>
      </a>
      <a class="top-projects__link" href="https://json-to-proto.github.io/" target="_blank">
        <figure class="top-projects__figure">
          <img
            class="top-projects__image"
            width="157"
            height="18"
            alt="JSON-to-PROTO logo"
            src="/assets/images/pages/online-new/json_to_proto.svg"
          />
        </figure>
      </a>
      <a class="top-projects__link" href="https://xml-to-go.github.io/" target="_blank">
        <figure class="top-projects__figure">
          <img
            class="top-projects__image"
            width="117"
            height="18"
            alt="XML-to-GO logo"
            src="/assets/images/pages/online-new/xml_to_go.svg"
          />
        </figure>
      </a>
    </div>
  </div>
</section>

<section class="project-statistics">
  <div class="container project-statistics__container">
    <header class="project-statistics__header">
      <h2 class="project-statistics__headline">Project Statistics</h2>
      <p class="project-statistics__description">
        We like statistics and transparency, so we show users who logged in here to support our projects or found our
        tools useful.
      </p>
    </header>
    <section class="stats">
      <div class="stats__item">
        <div class="stats__header">
          <h3 class="stats__title">Online users</h3>
          <div class="stats__period">Last month</div>
        </div>
        <div class="stats__chart stats__chart--online-statistics js-chart-online-statistics"></div>
      </div>
      <div class="stats__item">
        <div class="stats__header">
          <h3 class="stats__title">User registrations</h3>
          <div class="stats__period">Last month</div>
        </div>
        <div class="stats__chart">
          <div class="stats__chart stats__chart--registration-statistics js-chart-registration-statistics"></div>
        </div>
      </div>
    </section>
    `)
	streamregistrationHistory(qw422016, socialProviderUsers)
	qw422016.N().S(`
  </div>
</section>

<section class="organizer-card">
  <div class="container organizer-card__container">
    <figure class="organizer-card__figure">
      <picture class="organizer-card__picture">
        <source
          width="608"
          height="538"
          type="image/webp"
          srcset="
            /assets/images/pages/online-new/organizer/organizer.webp,
            /assets/images/pages/online-new/organizer/organizer@2x.webp 2x
          "
        />
        <source
          width="608"
          height="538"
          srcset="
            /assets/images/pages/online-new/organizer/organizer.jpg,
            /assets/images/pages/online-new/organizer/organizer@2x.jpg 2x
          "
        />
        <img
          class="organizer-card__image"
          width="608"
          height="538"
          loading="lazy"
          srcset="/assets/images/pages/online-new/organizer/organizer@2x.jpg 2x"
          src="/assets/images/pages/online-new/organizer/organizer.jpg"
          alt="organizer page screen"
        />
      </picture>
      <figcaption class="organizer-card__caption">
        <header class="organizer-card__header">
          <span class="organizer-card__header-text">Service status:</span>
          <span class="badge badge--new">New</span>
          <span class="badge badge--work-in-progress">Work in progress</span>
        </header>
        <h2 class="organizer-card__headline">Direct Contract Job Search Organizer</h2>
        <p class="organizer-card__description">
          A list of useful links to simplify your job search. The basis is a list of companies searching for Golang
          developers.
        </p>
        <footer class="organizer-card__footer">
          <iframe
            src="https://ghbtns.com/github-btn.html?user=readytotouch&amp;repo=golang-companies-organizer&amp;type=star&amp;count=true&amp;size=large"
            width="120"
            height="33"
            title="GitHub"
          ></iframe>
          <a
            class="button button--black organizer-card__try-btn"
            target="_blank"
            href="/organizers"
          >
            Try Organizer
            <div class="organizer-card__try-btn-icon-box">
              <img width="12" height="12" alt="arrow top icon" src="/assets/images/pages/common/arrow-small-top.svg" />
            </div>
          </a>
        </footer>
      </figcaption>
    </figure>
  </div>
</section>

<section class="other-projects">
  <div class="container other-projects__container">
    <h2 class="other-projects__title">Other projects</h2>
    <div class="other-projects__wrapper">
      <div class="other-projects__item project big-card">
        <div class="project__info-card">
          <div class="project__stars">
            <iframe
              src="https://ghbtns.com/github-btn.html?user=u8views&repo=go-u8views&type=star&count=true&size=large"
              width="170"
              height="30"
              title="GitHub"
            ></iframe>
          </div>
          <div class="project__info-group">
            <h3 class="project__title">u8views — views counter</h3>
            <p class="project__subtitle">
              Track your GitHub profile views. Receive, view and analyze your profile views and profile performance
              statistics.
            </p>
          </div>
          <a href="https://u8views.com/" target="_blank" class="project__link">u8views.com</a>
        </div>
        <div class="project__view">
          <img src="/assets/images/pages/online/u8views.png" class="project__view-img" alt="project view" />
        </div>
      </div>
      <div class="other-projects__group">
        <div class="other-projects__item project">
          <div class="project__info-card">
            <div class="project__stars">
              <iframe
                src="https://ghbtns.com/github-btn.html?user=json-to-proto&repo=json-to-proto.github.io&type=star&count=true&size=large"
                width="170"
                height="30"
                title="GitHub"
              ></iframe>
            </div>
            <div class="project__info-group">
              <h3 class="project__title">JSON-to-Proto</h3>
              <p class="project__subtitle">Convert JSON to Protobuf</p>
            </div>
            <a href="https://json-to-proto.github.io/" target="_blank" class="project__link">json-to-proto.github.io</a>
          </div>
          <div class="project__view">
            <img src="/assets/images/pages/online/open-project.png" class="project__view-img" alt="project view" />
          </div>
        </div>
        <div class="other-projects__item project">
          <div class="project__info-card">
            <div class="project__stars">
              <iframe
                src="https://ghbtns.com/github-btn.html?user=xml-to-go&repo=xml-to-go.github.io&type=star&count=true&size=large"
                width="170"
                height="30"
                title="GitHub"
              ></iframe>
            </div>
            <div class="project__info-group">
              <h3 class="project__title">XML-to-Go</h3>
              <p class="project__subtitle">Convert XML to Go</p>
            </div>
            <a href="https://xml-to-go.github.io/" target="_blank" class="project__link">xml-to-go.github.io</a>
          </div>
          <div class="project__view">
            <img src="/assets/images/pages/online/open-project.png" class="project__view-img" alt="project view" />
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<section class="articles">
  <div class="container articles__container">
    <h2 class="articles__title">Articles</h2>
    <div class="articles__group-wrapper">
        <div class="articles__item article">
            <a href="https://dou.ua/forums/topic/35260/" target="_blank" class="article__link">
                <img src="https://s.dou.ua/img/announces/image_Yne6Ph2.jpg" alt="article preview" class="article__preview">
            </a>
            <div class="article__info">
                <div class="article__top-group">
                    <div class="article__date-group">
                        <img src="/assets/images/pages/online/calendar.svg" alt="calendar icon"
                             class="article__date-icon">
                        <span class="article__date">5 квітня 2023</span>
                    </div>
                    <div class="article__author">
                        <img src="https://avatars.githubusercontent.com/u/63663261?v=4&s=48" alt="author"
                             class="article__author-photo">
                        <a href="https://dou.ua/users/yaroslav-harakternik/" target="_blank"
                           class="article__author-name">Ярослав Характерник</a>
                    </div>
                </div>
                <a href="https://dou.ua/forums/topic/35260/" target="_blank" class="article__title">Збереження стану онлайну користувача в Redis</a>
                <p class="article__subtitle">Ярослав Характерник — про методику вибору оптимального типу даних для збереження статусу онлайну користувачів в інтернет-магазині.</p>
            </div>
        </div>

        <div class="articles__item article">
            <a href="https://dou.ua/forums/topic/44655/" target="_blank" class="article__link">
                <img src="https://s.dou.ua/img/announces/0tech_redis_image.jpg" alt="article preview" class="article__preview">
            </a>
            <div class="article__info">
                <div class="article__top-group">
                    <div class="article__date-group">
                        <img src="/assets/images/pages/online/calendar.svg" alt="calendar icon"
                             class="article__date-icon">
                        <span class="article__date">15 серпня 2023</span>
                    </div>
                    <div class="article__author">
                        <img src="https://avatars.githubusercontent.com/u/63663261?v=4&s=48" alt="author"
                             class="article__author-photo">
                        <a href="https://dou.ua/users/yaroslav-harakternik/" target="_blank"
                           class="article__author-name">Ярослав Характерник</a>
                    </div>
                </div>
                <a href="https://dou.ua/forums/topic/44655/" target="_blank" class="article__title">Hash, Set чи Sorted set. Який тип даних вибрати для збереження стану онлайну користувача в Redis</a>
                <p class="article__subtitle">В цій статті Ярослав Характерник пропонує розглянути швидкодію, скільки пам’яті займає кожен з типів даних та яку базу даних вибрати: Redis, KeyDB або DragonflyDB.</p>
            </div>
        </div>

        <div class="articles__item article">
            <a href="https://dou.ua/forums/topic/35261/" target="_blank" class="article__link">
                <img src="https://s.dou.ua/img/announces/8_Uq7sQTh.jpg" alt="article preview" class="article__preview">
            </a>
            <div class="article__info">
                <div class="article__top-group">
                    <div class="article__date-group">
                        <img src="/assets/images/pages/online/calendar.svg" alt="calendar icon"
                             class="article__date-icon">
                        <span class="article__date">14 листопада 2023</span>
                    </div>
                    <div class="article__author">
                        <img src="https://avatars.githubusercontent.com/u/63663261?v=4&s=48" alt="author"
                             class="article__author-photo">
                        <a href="https://dou.ua/users/yaroslav-harakternik/" target="_blank"
                           class="article__author-name">Ярослав Характерник</a>
                    </div>
                </div>
                <a href="https://dou.ua/forums/topic/35261/" target="_blank" class="article__title">Batch UPDATE в PostgreSQL</a>
                <p class="article__subtitle">У цій статті Ярослав Характерник продовжує розповідати про свій проєкт з анонімного пошуку роботи. Сьогодні йтиметься про те, як зробити статистику онлайну публічною, а саме: про запуск команди перенесення онлайну пачками з Redis в PostgreSQL і тестування.</p>
            </div>
        </div>
    </div>
  </div>
</section>


</main>
`)
	streamorganizersFooter(qw422016)
	qw422016.N().S(`
<script src="/assets/js/online-stats-app.js?`)
	qw422016.N().D(appVersion)
	qw422016.N().S(`"></script>
</body>
</html>
`)
}

func WriteOrganizersOnline(qq422016 qtio422016.Writer,
	headerProfiles []SocialProviderUser,
	socialProviderUsers []SocialProviderUser,
) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamOrganizersOnline(qw422016, headerProfiles, socialProviderUsers)
	qt422016.ReleaseWriter(qw422016)
}

func OrganizersOnline(
	headerProfiles []SocialProviderUser,
	socialProviderUsers []SocialProviderUser,
) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteOrganizersOnline(qb422016, headerProfiles, socialProviderUsers)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
