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

            return;
        }
    }
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
