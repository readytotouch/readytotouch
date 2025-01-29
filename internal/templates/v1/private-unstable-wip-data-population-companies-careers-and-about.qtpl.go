// Code generated by qtc from "private-unstable-wip-data-population-companies-careers-and-about.qtpl". DO NOT EDIT.
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

func StreamDataPopulationCompaniesCareersAndAbout(qw422016 *qt422016.Writer, companies []Company, title string) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
    <title>`)
	qw422016.E().S(title)
	qw422016.N().S(`</title>
    <meta name="description" content="`)
	qw422016.E().S(title)
	qw422016.N().S(`">

    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="author" type="text/plain" href="https://readytotouch.com/humans.txt"/>
    <meta property="og:image" content="/assets/images/og/organizers-light.jpg">

    `)
	streamfavicon(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
	qw422016.N().S(`
    `)
	streamdataPopulationCompaniesStyles(qw422016)
	qw422016.N().S(`
</head>
<body>
<h1>`)
	qw422016.E().S(title)
	qw422016.N().S(`</h1>

<ul>
    `)
	for _, company := range companies {
		qw422016.N().S(`
    <li>
        <div class="company-name">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</div>
        <div class="links">
            `)
		if company.Careers == "" {
			qw422016.N().S(`
                <a href='`)
			qw422016.E().S(googleSearchOnSite(hostname(company.Website), "Careers"))
			qw422016.N().S(`' target="_blank">
                    <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> `)
			qw422016.E().S(hostname(company.Website))
			qw422016.N().S(` "Careers"
                </a>
                <a href='`)
			qw422016.E().S(googleSearch(company.Name + " " + "Careers"))
			qw422016.N().S(`' target="_blank">
                    <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> `)
			qw422016.E().S(company.Name)
			qw422016.N().S(` "Careers"
                </a>
            `)
		}
		qw422016.N().S(`
            `)
		if company.About == "" {
			qw422016.N().S(`
                <a href='`)
			qw422016.E().S(googleSearchOnSite(hostname(company.Website), "About"))
			qw422016.N().S(`' target="_blank">
                    <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> `)
			qw422016.E().S(hostname(company.Website))
			qw422016.N().S(` "About"
                </a>
                <a href='`)
			qw422016.E().S(googleSearch(company.Name + " " + "About"))
			qw422016.N().S(`' target="_blank">
                    <img alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg"> `)
			qw422016.E().S(company.Name)
			qw422016.N().S(` "About"
                </a>
            `)
		}
		qw422016.N().S(`
        </div>
    </li>
    `)
	}
	qw422016.N().S(`
</ul>

</body>
</html>
`)
}

func WriteDataPopulationCompaniesCareersAndAbout(qq422016 qtio422016.Writer, companies []Company, title string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamDataPopulationCompaniesCareersAndAbout(qw422016, companies, title)
	qt422016.ReleaseWriter(qw422016)
}

func DataPopulationCompaniesCareersAndAbout(companies []Company, title string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteDataPopulationCompaniesCareersAndAbout(qb422016, companies, title)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
