const COMPANY_CACHE_DURATION = 5 * 60 * 1000;

let companiesCache = null;
let companiesCacheTimestamp = 0;

function fetchSafeCompanyList(url) {
    return fetch(url)
        .then(response => {
            if (response.ok) {
                return response.json();
            }

            throw new Error(`HTTP error: ${response.status}`);
        })
        .catch(error => {
            console.warn(`Fetch failed for ${url}:`, error);
            return {companies: []};
        });
}

function fetchCompanyList() {
    return Promise.all([
        fetchSafeCompanyList("http://localhost/api/v1/unsafe/companies.json"),
        fetchSafeCompanyList("https://readytotouch.com/api/v1/unsafe/companies.json"),
    ])
        .then(([localData, remoteData]) => {
            companiesCache = [...localData.companies, ...remoteData.companies];
            companiesCacheTimestamp = Date.now();
            console.log("Fetched and cached combined company list.");
        });
}

function fetchCompany(companyName) {
    if (companiesCache === null) {
        return null;
    }

    for (const c of companiesCache) {
        if (c.name.toLowerCase() === companyName || c.id.toString() === companyName) {
            return c;
        }
    }

    return null;
}

function updateCompanyColor($companyName, company) {
    if (!companiesCache) {
        $companyName.style.color = "#ffc107";
    } else if (!company) {
        $companyName.style.color = "#007bff";
    } else if (company.ignore) {
        $companyName.style.color = "#dc3545";
    } else {
        $companyName.style.color = "#28a745";
    }
}

function monitorCompanyNameChange() {
    const $companyName = getCompanyNameElement();
    if ($companyName === null) {
        return;
    }

    const currentCompanyName = $companyName.textContent.trim().toLowerCase();

    const company = fetchCompany(currentCompanyName);

    updateCompanyColor($companyName, company);

    renderCompanyLanguages(company);
}

let renderCompanyLanguagesOnce = false;

function renderCompanyLanguages(company) {
    if (company === null || company.languages === null || company.languages.length === 0) {
        return;
    }

    if (!window.location.href.startsWith("https://www.linkedin.com/jobs/view/")) {
        return;
    }

    if (renderCompanyLanguagesOnce) {
        return;
    }

    renderCompanyLanguagesOnce = true;

    for (let i = 0; i < company.languages.length; i++) {
        const language = company.languages[i];

        renderLanguageMaxVacancyDate(language.language, language.max_vacancy_date, company.languages.length - i);
    }
}

function getCompanyNameElement() {
    const url = window.location.href;

    if (url.includes("linkedin.com/company/")) {
        return document.querySelector("h1");
    }

    if (url.includes("linkedin.com")) {
        return document.querySelector(".job-details-jobs-unified-top-card__company-name a");
    }

    if (url.includes("app.welcometothejungle.com/jobs/")) {
        return document.querySelector("h1 a");
    }

    return null;
}

function renderLanguageMaxVacancyDate(language, date, index) {
    const $date = document.createElement("div");
    $date.innerText = `${language}: ${date}`;

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

function initializeCompanyCache() {
    if (companiesCache === null || Date.now() - companiesCacheTimestamp > COMPANY_CACHE_DURATION) {
        fetchCompanyList();
    }
}

initializeCompanyCache();

setInterval(initializeCompanyCache, COMPANY_CACHE_DURATION);

setInterval(monitorCompanyNameChange, 250);
