import {toEnter} from "./framework/enter";
import {welcome} from "./welcome";
import {htmlToNode} from "./framework/html";
import {firstQuerySelector} from "./framework/query_selector";

let latestKeywords = '';
let latestLocation = '';
let latestCompanies = [];

fetch('/api/v1/companies-and-connections/companies.json')
    .then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = welcome();

            return [];
        }

        return response.json();
    })
    .then(function (companies: Array<ResponseCompany>) {
        latestCompanies = companies;

        renderCompanies(companies);
    })
    .catch(console.error);

function addCompany(companyUrl: string, companyVanityName: string, callback: () => void) {
    fetch('/api/v1/companies-and-connections/companies.json', {
        method: "POST",
        body: JSON.stringify({
            alias: companyVanityName,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = welcome(companyUrl);

            return;
        }

        return response.json();
    }).then(function (companies: Array<ResponseCompany>) {
        latestCompanies = companies;

        renderCompanies(companies);

        callback();
    }).catch(console.error);
}

function deleteCompany(id: number) {
    fetch('/api/v1/companies-and-connections/companies.json', {
        method: "DELETE",
        body: JSON.stringify({
            id: id,
        }),
    })
}

class ResponseCompany {
    constructor(
        public readonly id: number,
        public readonly alias: string,
        public readonly name: string,
    ) {
    }
}

class Connections {
    constructor(
        public readonly connections1stURL: string,
        public readonly connections2ndURL: string,
        public readonly connections1stXURL: string,
        public readonly connections2ndXURL: string,
    ) {
    }
}

class Company {
    constructor(
        public readonly id: number,
        public readonly vanityName: string,
        public readonly name: string,
        public readonly connections1stURL: string,
        public readonly connections2ndURL: string,
        public readonly connections1stXURL: string,
        public readonly connections2ndXURL: string,
    ) {
    }
}

const $companies = document.getElementById('js-companies') as HTMLElement;
const $companiesCount = document.getElementById('js-companies-count') as HTMLElement;
const $companyUrlInput = document.getElementById('company-url') as HTMLInputElement
const $addCompanyButton = document.getElementById('add-company') as HTMLButtonElement;

const $location = document.getElementById('location') as HTMLInputElement;
const $keywords = document.getElementById('keywords') as HTMLInputElement;
const $updateConnectionsButton = document.getElementById('update-connections') as HTMLButtonElement;

// Enable/Disable Add button based on input
$companyUrlInput.addEventListener('input', function () {
    $addCompanyButton.disabled = $companyUrlInput.value.trim() === '';
});

$companyUrlInput.addEventListener("keyup", toEnter(submitCompany));

// Add Company button click event
$addCompanyButton.addEventListener('click', submitCompany);

$updateConnectionsButton.addEventListener('click', function () {
    latestKeywords = $keywords.value.trim();
    latestLocation = $location.value;

    renderCompanies(latestCompanies)
});

function submitCompany() {
    const companyUrl = $companyUrlInput.value.trim();
    if (companyUrl === '') {
        return;
    }

    const vanityName = parseVanityName($companyUrlInput.value);
    if (vanityName === '') {
        return;
    }

    addCompany(companyUrl, vanityName, function () {
        $companyUrlInput.value = '';
        $addCompanyButton.disabled = true;
    });
}

function renderCompany(company: Company): string {
    return `<div class="card">
    <aside class="card__action">
        <button class="button-group__item" title="Delete">
            <img width="20" height="20" alt="icon stats" src="/assets/images/pages/common/trash.svg"/>
        </button>
    </aside>
    <div class="card__header">
        <a href="https://www.linkedin.com/company/${company.vanityName}" class="button-link card__headline">${company.name}</a>
    </div>

    <div class="card__group">
        <div class="card__group-block">
            <a class="card__group-item button-link" href="${company.connections1stURL}">Connections 1st</a>
            <a class="card__group-item button-link" href="${company.connections2ndURL}">Connections 2nd</a>
        </div>
        <div class="card__group-block">
            <a class="card__group-item button-link" href="${company.connections1stXURL}">Connections 1st X</a>
            <a class="card__group-item button-link" href="${company.connections2ndXURL}">Connections 2nd X</a>
        </div>
    </div>
</div>`;
}

function renderCompanies(companies: Array<ResponseCompany>) {
    $companies.innerHTML = '';
    $companiesCount.innerHTML = companies.length.toString();

    companies.forEach((company) => {
        const prepared = prepareConnections(company, companies);

        const $card = htmlToNode(renderCompany(new Company(
            company.id,
            company.alias,
            company.name,
            prepared.connections1stURL,
            prepared.connections2ndURL,
            prepared.connections1stXURL,
            prepared.connections2ndXURL,
        )));
        firstQuerySelector($card, "button").addEventListener("click", function () {
            deleteCompany(company.id);

            latestCompanies = latestCompanies.filter(c => c.id !== company.id);

            $companiesCount.innerHTML = latestCompanies.length.toString();

            $card.remove();
        });

        $companies.appendChild($card);
    });
}

function prepareConnections(
    currentCompany: ResponseCompany,
    companies: Array<ResponseCompany>,
): Connections {

    const pastCompanies = [];
    for (const company of companies) {
        if (currentCompany === company) {
            continue;
        }

        pastCompanies.push(company.id.toString());
    }

    const currentCompanyQueryParam = `["${currentCompany.id}"]`;

    let universitiesQueryParam = '';
    if (window.location.href.indexOf('ukraine') !== -1) {
        universitiesQueryParam = JSON.stringify(["818029", "850102", "364340", "496320", "1198954", "1257361", "15250306", "15251128", "15099424", "782774", "15101979", "15101061", "80424966", "6261241", "658198", "11443062", "15099425", "15099711", "15101057", "15102004", "18080249", "15143861", "15101046", "1599158", "15101060", "15100187", "9029417", "7991636", "15101074", "27066401", "18144134", "15101998", "15149751", "18691495", "15099038"]);
    } else if (window.location.href.indexOf('brazil') !== -1) {
        universitiesQueryParam = JSON.stringify(["239895", "986104", "15171", "69715404", "10866", "28514", "17959", "312647", "1379596", "38307", "760298"]);
    }

    let connections1stURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections1stXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');

    connections1stURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stURL.searchParams.append('network', '["F"]');

    connections2ndURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndURL.searchParams.append('network', '["S"]');
    if (universitiesQueryParam !== '') {
        connections2ndURL.searchParams.append('schoolFilter', universitiesQueryParam);
    }

    connections1stXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stXURL.searchParams.append('network', '["F"]');

    connections2ndXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndXURL.searchParams.append('network', '["S"]');
    if (universitiesQueryParam !== '') {
        connections2ndXURL.searchParams.append('schoolFilter', universitiesQueryParam);
    }

    if (latestKeywords !== '') {
        connections1stURL.searchParams.append('keywords', latestKeywords);
        connections2ndURL.searchParams.append('keywords', latestKeywords);
        connections1stXURL.searchParams.append('keywords', latestKeywords);
        connections2ndXURL.searchParams.append('keywords', latestKeywords);
    }

    if (latestLocation !== '') {
        const queryParam = `["${latestLocation}"]`;
        connections1stURL.searchParams.append('geoUrn', queryParam);
        connections2ndURL.searchParams.append('geoUrn', queryParam);
        connections1stXURL.searchParams.append('geoUrn', queryParam);
        connections2ndXURL.searchParams.append('geoUrn', queryParam);
    }

    if (pastCompanies.length > 0) {
        const queryParam = JSON.stringify(pastCompanies);
        connections1stXURL.searchParams.append('pastCompany', queryParam);
        connections2ndXURL.searchParams.append('pastCompany', queryParam);
    }

    return new Connections(
        connections1stURL.toString(),
        connections2ndURL.toString(),
        connections1stXURL.toString(),
        connections2ndXURL.toString(),
    );
}

function parseVanityName(url) {
    // Your errors = your pain
    const error = 'Expected URL like https://www.linkedin.com/company/company-name/';

    let parsedUrl = null;

    try {
        parsedUrl = new URL(url);
    } catch (e) {
        alert(error);

        return '';
    }

    const prefix = '/company/';

    if (parsedUrl.pathname.indexOf(prefix) === -1) {
        alert(error);

        return '';
    }

    const end = parsedUrl.pathname.indexOf('/', prefix.length);
    if (end === -1) {
        return parsedUrl.pathname.substring(prefix.length);
    }

    return parsedUrl.pathname.substring(prefix.length, end);
}

{
    const url = new URL(window.location.href);
    const companyUrl = url.searchParams.get('company-url');
    if (companyUrl) {
        $companyUrlInput.value = companyUrl;
    }
}

$addCompanyButton.disabled = $companyUrlInput.value.trim() === '';