// Code generated by qtc from "private-unstable-wip-data-population-lists.qtpl". DO NOT EDIT.
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

func StreamDataPopulationLists(qw422016 *qt422016.Writer) {
	qw422016.N().S(`<!DOCTYPE html>
<html lang="en">

<head>
    <title>Data Population Lists</title>
    <meta name="description" content="Data Population Lists">

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

    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f9f9f9;
            padding: 20px;
            max-width: 600px;
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
        }

        a {
            color: #0073e6;
            text-decoration: none;
        }

        a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
<h1>Data Population Lists</h1>

<ul>
    <li><a href="/private/unstable/wip/organizers/data-population-lists/careers-and-about-and-blog" target="_blank">Populate Careers & About & Blog</a></li>
    <li><a href="/private/unstable/wip/organizers/data-population-lists/linkedin" target="_blank">Populate LinkedIn</a></li>
    <li><a href="/private/unstable/wip/organizers/data-population-lists/glassdoor" target="_blank">Populate Glassdoor</a></li>
</ul>

</body>
</html>
`)
}

func WriteDataPopulationLists(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	StreamDataPopulationLists(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func DataPopulationLists() string {
	qb422016 := qt422016.AcquireByteBuffer()
	WriteDataPopulationLists(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
