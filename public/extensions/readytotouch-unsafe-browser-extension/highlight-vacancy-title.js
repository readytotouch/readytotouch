const VACANCY_CACHE_DURATION = 5 * 60 * 1000;

let vacanciesCache = null;
let vacanciesCacheTimestamp = 0;

function fetchSafeVacancyList(url) {
    return fetch(url)
        .then(response => {
            if (response.ok) {
                return response.json();
            }

            throw new Error(`HTTP error: ${response.status}`);
        })
        .catch(error => {
            console.warn(`Fetch failed for ${url}:`, error);
            return {vacancies: []};
        });
}

function fetchVacancyList() {
    return Promise.all([
        fetchSafeVacancyList("http://localhost/api/v1/unsafe/vacancies.json"),
        fetchSafeVacancyList("https://readytotouch.com/api/v1/unsafe/vacancies.json"),
    ])
        .then(([localData, remoteData]) => {
            vacanciesCache = [...localData.vacancies, ...remoteData.vacancies];
            vacanciesCacheTimestamp = Date.now();
            console.log("Fetched and cached combined vacancy list.");
        });
}

function highlightVacancyTitle() {
    if (vacanciesCache === null) {
        console.log("Vacancy list is not loaded yet.");

        return;
    }

    const $vacancyTitle = document.querySelector("h1");

    if ($vacancyTitle === null) {
        console.log("Vacancy title element not found.");

        return;
    }

    if (window.location.href.startsWith("https://app.welcometothejungle.com/jobs/")) {
        searchLinkedInCompanyProfile($vacancyTitle.querySelector("a"))
    }

    let url = "";

    if (window.location.href.startsWith("https://www.linkedin.com/jobs/search/")) {
        const params = new URLSearchParams(window.location.search);

        const id = params.get("currentJobId");

        url = normalizeURL(`https://www.linkedin.com/jobs/view/${id}/`);
    } else {
        url = normalizeURL(window.location.origin + window.location.pathname);
    }

    for (const vacancy of vacanciesCache) {
        if (vacancy.url === url) {
            $vacancyTitle.style.color = "#28a745";

            renderVacancyDate(vacancy.date, 0);

            return;
        }
    }

    const $date = document.getElementById("readytotouch-vacancy-date");
    if ($date) {
        $date.remove();
    }
}

function normalizeURL(url) {
    const prefix = "https://www.linkedin.com/jobs/view/";

    if (url.startsWith(prefix)) {
        const match = url.match(/\/jobs\/view\/.*?(\d{7,})/);
        if (match && match[1]) {
            return prefix + match[1] + "/";
        }
    }

    return url;
}

let renderVacancyDateOnce = false;

function renderVacancyDate(date, index) {
    if (renderVacancyDateOnce) {
        return;
    }

    renderVacancyDateOnce = true;

    const $date = document.createElement("div");
    $date.id = "readytotouch-vacancy-date";
    $date.innerText = date;

    Object.assign($date.style, {
        position: "fixed",
        bottom: `${10 + index * 60}px`,
        right: "10px",
        backgroundColor: "rgba(0, 0, 0, 0.7)",
        color: "white",
        padding: "10px 15px",
        borderRadius: "8px",
        fontSize: "16px",
        zIndex: "999999",
        fontFamily: "sans-serif",
        boxShadow: "0 4px 12px rgba(0, 0, 0, 0.3)",
    });

    document.body.appendChild($date);
}

function searchLinkedInCompanyProfile($company) {
    if ($company === null) {
        return;
    }

    $company.href = `https://www.google.com/search?q=site:linkedin.com/company/ ${$company.textContent.trim()}`;
}

function initializeVacancyCache() {
    if (vacanciesCache === null || Date.now() - vacanciesCacheTimestamp > VACANCY_CACHE_DURATION) {
        fetchVacancyList()
    }
}

initializeVacancyCache();

setInterval(initializeVacancyCache, VACANCY_CACHE_DURATION);

setInterval(highlightVacancyTitle, 250);

/*
    TODO:
        1. Show if company has fresh vacancies for this language
*/
