function addLanguages() {
    const path = window.location.pathname;

    const requiredLanguages = ['Go', 'Rust', 'Scala', 'Elixir', 'Clojure', 'Haskell'];
    const existingLanguages = Array.from(document.querySelectorAll('span[itemprop="programmingLanguage"]'))
        .map(span => span.textContent.trim());

    requiredLanguages.forEach(language => {
        if (existingLanguages.includes(language)) {
            return;
        }

        const content = `<a class="no-wrap color-fg-muted d-inline-block Link--muted mt-2" href="/orgs${path}/repositories?language=${language.toLowerCase()}&amp;type=all">
            <span class="ml-0 mr-2">
              <span class="repo-language-color" style="background-color: #3572A5"></span>
              <span itemprop="programmingLanguage">${language}</span>
            </span>
          </a>`

        const $language = document.createElement("div");
        $language.innerHTML = content;

        const $languages = document.querySelector("h4.f4.mb-2.text-normal").parentElement;
        $languages.append($language.firstChild);
    });
}

addLanguages();
