const CACHE_DURATION = 5 * 60 * 1000;

let companiesCache = null;
let cacheTimestamp = 0;

function fetchCompanyList() {
    return fetch("https://readytotouch.com/api/v1/unsafe/companies.json")
        .then(response => response.json())
        .then(data => {
            companiesCache = data.companies;
            cacheTimestamp = Date.now();
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

    const exists = companiesCache.any(function (c) {
        return c.name.toLowerCase() === companyName || c.id.toString() === companyName;
    });
    if (exists) {
        return "exists";
    }

    return "not_found";
}

function updateCompanyColor(status) {
    const $companyName = getCompanyNameElement();

    if ($companyName === null) {
        return;
    }

    if (status === "exists") {
        $companyName.style.color = "#28a745";
    } else if (status === "not_found") {
        $companyName.style.color = "#dc3545";
    } else if (status === "in_progress") {
        $companyName.style.color = "#007bff";
    }
}

function monitorCompanyNameChange() {
    const $company = getCompanyNameElement();

    if ($company) {
        const currentCompanyName = $company.textContent.trim();

        const status = checkCompanyStatusFromCache(currentCompanyName);

        updateCompanyColor(status);
    }
}

function getCompanyNameElement() {
    const url = window.location.href;

    // Визначаємо селектор в залежності від сайту
    if (url.includes("linkedin.com")) {
        return document.querySelector(".job-details-jobs-unified-top-card__company-name a");
    }

    if (url.includes("welcometothejungle.com")) {
        return null;
    }

    return null;
}

function initializeCompanyCache() {
    if (companiesCache === null || Date.now() - cacheTimestamp > CACHE_DURATION) {
        fetchCompanyList();
    }
}

setInterval(initializeCompanyCache, CACHE_DURATION);

initializeCompanyCache();

setInterval(monitorCompanyNameChange, 250);
