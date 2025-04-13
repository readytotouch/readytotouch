const VACANCY_CACHE_DURATION = 5 * 60 * 1000;

let vacanciesCache = null;
let vacanciesCacheTimestamp = 0;

function fetchVacancyList() {
    return fetch("https://readytotouch.com/api/v1/unsafe/vacancies.json")
        .then(response => response.json())
        .then(data => {
            vacanciesCache = data.vacancies;
            vacanciesCacheTimestamp = Date.now();
            console.log("Fetched and cached vacancy list.");
        })
        .catch(error => {
            console.error("Error fetching vacancy list:", error);
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

    searchLinkedInCompanyProfile($vacancyTitle.querySelector("a"))

    const url = window.location.origin + window.location.pathname;

    for (const vacancy of vacanciesCache) {
        if (vacancy.url === url) {
            $vacancyTitle.style.color = "#28a745";

            renderVacancyDate(vacancy.date)

            return;
        }
    }
}

function renderVacancyDate(date) {
    const $date = document.createElement("div");
    $date.innerText = date;

    Object.assign($date.style, {
        position: "fixed",
        bottom: "10px",
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
