// Code generated by qtc from "ga.qtpl". DO NOT EDIT.
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

func streamga(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<!-- Google tag (gtag.js) -->
<script async src="https://www.googletagmanager.com/gtag/js?id=G-0TET50K0HS"></script>
<script>
  window.dataLayer = window.dataLayer || [];
  function gtag(){dataLayer.push(arguments);}
  gtag('js', new Date());

  gtag('config', 'G-0TET50K0HS');
</script>
`)
}

func writega(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamga(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func ga() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writega(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
