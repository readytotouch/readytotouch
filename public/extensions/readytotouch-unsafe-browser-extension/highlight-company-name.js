const COMPANY_CACHE_DURATION = 5 * 60 * 1000;

let companiesCache = null;
let companiesCacheTimestamp = 0;

function fetchCompanyList() {
    return fetch("https://readytotouch.com/api/v1/unsafe/companies.json")
        .then(response => response.json())
        .then(data => {
            companiesCache = data.companies;
            companiesCacheTimestamp = Date.now();
            console.log("Fetched and cached company list.");
        })
        .catch(error => {
            console.error("Error fetching company list:", error);
        });
}

function checkCompanyStatusFromCache(companyName) {
    if (companiesCache === null) {
        return "in_progress";
    }

    const company = companiesCache.find(function (c) {
        return c.name.toLowerCase() === companyName || c.id.toString() === companyName;
    });

    if (company) {
        if (company.ignore) {
            return "ignore";
        }

        return "exists";
    }

    return "not_found";
}

function updateCompanyColor($companyName, status) {
    if (status === "exists") {
        $companyName.style.color = "#28a745";
    } else if (status === "not_found") {
        $companyName.style.color = "#007bff";
    } else if (status === "in_progress") {
        $companyName.style.color = "#ffc107";
    } else if (status === "ignore") {
        $companyName.style.color = "#dc3545";
    }
}

function monitorCompanyNameChange() {
    const $companyName = getCompanyNameElement();

    if ($companyName) {
        const currentCompanyName = $companyName.textContent.trim().toLowerCase();

        const status = checkCompanyStatusFromCache(currentCompanyName);

        updateCompanyColor($companyName, status);
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

function initializeCompanyCache() {
    if (companiesCache === null || Date.now() - companiesCacheTimestamp > COMPANY_CACHE_DURATION) {
        fetchCompanyList();
    }
}

initializeCompanyCache();

setInterval(initializeCompanyCache, COMPANY_CACHE_DURATION);

setInterval(monitorCompanyNameChange, 250);
