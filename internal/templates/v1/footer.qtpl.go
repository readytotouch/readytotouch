// Code generated by qtc from "footer.qtpl". DO NOT EDIT.
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

func streamfooter(qw422016 *qt422016.Writer) {
	qw422016.N().S(`
<footer class="footer">
    <div class="footer__wrapper main-container">
        <div class="footer__group-u8">
            <div class="footer__info">
                <a href="/" class="footer__title">
                    <h3 class="footer__title-link">ReadyToTouch</h3>
                </a>
                <p class="footer__subtitle">Anonymous work search</p>
                <div class="footer__map">
                    <img class="footer__map-Ukraine" src="/assets/images/pages/online/map-of-Ukraine.png"
                         alt="Map of Ukraine">
                </div>
            </div>
            <div class="footer__middle-section">
                <iframe class="footer__stars" src="https://ghbtns.com/github-btn.html?user=readytotouch&repo=readytotouch&type=star&count=true&size=large"
                        frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>
                <a href="https://github.com/readytotouch/readytotouch" target="_blank" class="footer__statistics-link">View the project on GitHub</a>
                <a href="https://readytotouch.com/" target="_blank" class="footer__github-link">
                    readytotouch.com
                </a>
            </div>
            <div class="footer__support">
                <a href="https://savelife.in.ua/en/" target="_blank" class="footer__link">
                    <figure class="footer__figure">
                        <img width="60" height="24" src="https://savelife.in.ua/wp-content/themes/savelife/assets/images/new-logo-en.svg" alt="support">
                        <figcaption class="footer__caption">Support</figcaption>
                        <img src="/assets/images/pages/online/arrow-up.svg" alt="arrow">
                    </figure>
                </a>
                <a href="https://war.ukraine.ua/" target="_blank" class="footer__link">
                    <figure class="footer__figure">
                        <figcaption class="footer__caption">war.ukraine.ua</figcaption>
                        <img src="/assets/images/pages/online/arrow-up.svg"  alt="arrow">
                    </figure>
                </a>
            </div>
        </div>
        <div class="footer__copyrights">
            <span>© 2025 Yaroslav Podorvanov</span>
            <img class="footer__flag-UA" src="/assets/images/pages/online/flag-of-Ukraine.svg" alt="Flag of Ukraine">
        </div>
    </div>
</footer>
`)
}

func writefooter(qq422016 qtio422016.Writer) {
	qw422016 := qt422016.AcquireWriter(qq422016)
	streamfooter(qw422016)
	qt422016.ReleaseWriter(qw422016)
}

func footer() string {
	qb422016 := qt422016.AcquireByteBuffer()
	writefooter(qb422016)
	qs422016 := string(qb422016.B)
	qt422016.ReleaseByteBuffer(qb422016)
	return qs422016
}
