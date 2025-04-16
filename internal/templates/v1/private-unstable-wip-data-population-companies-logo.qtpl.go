// Code generated by qtc from "private-unstable-wip-data-population-companies-logo.qtpl". DO NOT EDIT.
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

func StreamDataPopulationCompaniesLogo(qw422016 *qt422016.Writer, companies []Company, title string) {
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
	streamplausibleAnalytics(qw422016)
	qw422016.N().S(`
    `)
	streamga(qw422016)
	qw422016.N().S(`
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9f9f9;
            padding: 20px;
            max-width: 800px;
            margin: auto;
        }

        h1 {
            text-align: center;
        }

        ul {
            list-style: none;
            padding: 0;
        }

        li {
            background: #ffffff;
            margin: 10px 0;
            padding: 15px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
            display: flex;
            align-items: center;
            gap: 12px;
        }

        .logo {
            width: 112px;
            height: 56px;
            background-color: #ddd;
            object-fit: contain;
        }

        .company-name {
            font-size: 1.2em;
            font-weight: bold;
        }

        .links {
            margin-top: 5px;
        }

        .links a {
            color: #0073e6;
            text-decoration: none;
            margin-right: 10px;
        }

        .links a:hover {
            text-decoration: underline;
        }

        .actions a {
            width: 21px;
            height: 20px;
            display: inline-block;
            background-size: cover;
            background-image: url('https://www.google.com/favicon.ico');
        }
    </style>
</head>
<body>
<h1>`)
	qw422016.E().S(title)
	qw422016.N().S(` (`)
	qw422016.N().D(len(companies))
	qw422016.N().S(`)</h1>
<ul>
    `)
	for _, company := range companies {
		qw422016.N().S(`
    <li>
        <img src="`)
		qw422016.E().S(logoV1(company.Logo))
		qw422016.N().S(`" class="logo" />
        <img src="`)
		qw422016.E().S(logoV0(company.Logo))
		qw422016.N().S(`" class="logo" />
        <div>
            <div class="company-name">`)
		qw422016.E().S(company.Name)
		qw422016.N().S(`</div>
            <div class="links">
                <a href="https://linkedin.com/company/`)
		qw422016.E().S(company.LinkedInProfile.Alias)
		qw422016.N().S(`" target="_blank">LinkedIn</a>
            </div>
            <div class="actions">
                <a target="_blank" href='`)
		qw422016.E().S(googleSearch(company.Name + " " + "logo"))
		qw422016.N().S(`'></a>
                <a target="_blank" href='`)
		qw422016.E().S(googleSearchOnSite(hostname(company.Website), "logo"))
		qw422016.N().S(`'></a>
            </div>
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

func WriteDataPopulationCompaniesLogo(qq422016 qtio422016.Writer, companies []Company, title string) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamDataPopulationCompaniesLogo(qw422016, companies, title)
	qt422016.ReleaseWriter(qw422016)
}

func DataPopulationCompaniesLogo(companies []Company, title string) string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteDataPopulationCompaniesLogo(qb422016, companies, title)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
