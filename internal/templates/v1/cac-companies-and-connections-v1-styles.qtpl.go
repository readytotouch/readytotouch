// Code generated by qtc from "cac-companies-and-connections-v1-styles.qtpl". DO NOT EDIT.
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

func streamcompaniesAndConnectionsV1Styles(qw422016 *qt422016.Writer) {
}

func writecompaniesAndConnectionsV1Styles(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamcompaniesAndConnectionsV1Styles(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func companiesAndConnectionsV1Styles() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writecompaniesAndConnectionsV1Styles(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
