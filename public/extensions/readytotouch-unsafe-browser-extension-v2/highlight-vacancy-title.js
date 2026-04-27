const VACANCY_CACHE_DURATION = 5 * 60 * 1000;

let vacanciesCache     = null;
let vacanciesCacheTime = 0;

function fetchSafeVacancyList(url) {
    return fetch(url)
        .then(r => r.ok ? r.json() : Promise.reject(r.status))
        .catch(err => {
            console.warn(`[RTT] fetch failed for ${url}:`, err);
            return { vacancies: [] };
        });
}

function fetchVacancyList() {
    return Promise.all([
        fetchSafeVacancyList("http://localhost/api/v1/unsafe/vacancies.json"),
        fetchSafeVacancyList("https://readytotouch.com/api/v1/unsafe/vacancies.json"),
    ]).then(([local, remote]) => {
        vacanciesCache     = [...local.vacancies, ...remote.vacancies];
        vacanciesCacheTime = Date.now();
        console.log("[RTT] vacancy list cached");
    });
}

function highlightVacancyTitle() {
    if (!vacanciesCache) return;

    const $title = document.querySelector("h1");
    if (!$title) return;

    if (window.location.href.startsWith("https://app.welcometothejungle.com/jobs/")) {
        const $company = $title.querySelector("a");
        if ($company) {
            $company.href = `https://www.google.com/search?q=site:linkedin.com/company/ ${$company.textContent.trim()}`;
        }
    }

    const url = resolveVacancyURL();

    const matched = vacanciesCache.find(v => v.url === url);

    if (matched) {
        $title.style.color = "#28a745";
        renderVacancyBadge(matched.date, 0);
    } else {
        $title.style.color = "";
        removeVacancyBadge();
    }
}

function resolveVacancyURL() {
    if (window.location.href.startsWith("https://www.linkedin.com/jobs/search/")) {
        const id = new URLSearchParams(window.location.search).get("currentJobId");
        return normalizeLinkedInURL(`https://www.linkedin.com/jobs/view/${id}/`);
    }
    return normalizeLinkedInURL(window.location.origin + window.location.pathname);
}

function normalizeLinkedInURL(url) {
    const prefix = "https://www.linkedin.com/jobs/view/";
    if (url.startsWith(prefix)) {
        const m = url.match(/\/jobs\/view\/.*?(\d{7,})/);
        if (m) return prefix + m[1] + "/";
    }
    return url;
}

// Badge — always recreate so it stays correct after SPA navigation.
function renderVacancyBadge(date, index) {
    removeVacancyBadge(); // clear stale badge first

    const el = document.createElement("div");
    el.id         = "readytotouch-vacancy-date";
    el.innerText  = date;
    Object.assign(el.style, {
        position:        "fixed",
        bottom:          `${10 + index * 60}px`,
        right:           "10px",
        backgroundColor: "rgba(0, 0, 0, 0.7)",
        color:           "white",
        padding:         "10px 15px",
        borderRadius:    "8px",
        fontSize:        "16px",
        zIndex:          "999999",
        fontFamily:      "sans-serif",
        boxShadow:       "0 4px 12px rgba(0, 0, 0, 0.3)",
    });
    document.body.appendChild(el);
}

function removeVacancyBadge() {
    document.getElementById("readytotouch-vacancy-date")?.remove();
}

function initVacancyCache() {
    if (!vacanciesCache || Date.now() - vacanciesCacheTime > VACANCY_CACHE_DURATION) {
        fetchVacancyList();
    }
}

initVacancyCache();
setInterval(initVacancyCache,      VACANCY_CACHE_DURATION);
setInterval(highlightVacancyTitle, 250);

/*
    TODO:
        1. Show if company has fresh vacancies for this language
*/
