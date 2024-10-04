import {toEnter} from "./framework/enter";
import {organizersWelcome} from "./welcome";

let latestKeywords = '';
let latestLocation = '';
let latestCompanies = [];

fetch('/api/v1/companies-and-connections/companies.json')
    .then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return [];
        }

        return response.json();
    }).then(function (companies: Array<ResponseCompany>) {
    latestCompanies = companies;

    renderCompanies(companies, false);
}).catch(console.error);

function addCompany(companyVanityName: string, callback: () => void) {
    fetch(`/api/v1/companies-and-connections/companies.json`, {
        method: "POST",
        body: JSON.stringify({
            alias: companyVanityName,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        return response.json();
    }).then(function (companies: Array<ResponseCompany>) {
        latestCompanies = companies;

        renderCompanies(companies, true);

        callback();
    }).catch(console.error);
}

function deleteCompany(id: number) {
    fetch(`/api/v1/companies-and-connections/companies.json`, {
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

// Manage Connections toggle
const $manageConnectionsHeader = document.getElementById('manage-connections-header');
const $manageConnectionsContent = document.getElementById('manage-connections-content');
$manageConnectionsHeader.addEventListener('click', () => {
    $manageConnectionsContent.classList.toggle('hidden');
});

// Add Company toggle
const $addCompanyHeader = document.getElementById('add-company-header');
const $addCompanyContent = document.getElementById('add-company-content');
$addCompanyHeader.addEventListener('click', () => {
    $addCompanyContent.classList.toggle('hidden');
});

// Add company functionality and show Company List when there's at least one company
const $companyList = document.getElementById('company-list');
const $companyListBlock = document.getElementById('company-list-block') as HTMLElement;
const $companyUrlInput = document.getElementById('company-url') as HTMLInputElement
const $addCompanyButton = document.getElementById('add-company') as HTMLButtonElement;

const $location = document.getElementById('location') as HTMLInputElement;
const $keywords = document.getElementById('keywords') as HTMLInputElement;
const $updateConnectionsButton = document.getElementById('update-connections') as HTMLButtonElement;

// Enable/Disable Add button based on input
$companyUrlInput.addEventListener('input', function () {
    if ($companyUrlInput.value.trim() !== '') {
        $addCompanyButton.removeAttribute('disabled');
    } else {
        $addCompanyButton.setAttribute('disabled', 'disabled');
    }
});

$companyUrlInput.addEventListener("keyup", toEnter(submitCompany));

// Add Company button click event
$addCompanyButton.addEventListener('click', submitCompany);

$updateConnectionsButton.addEventListener('click', function () {
    latestKeywords = $keywords.value.trim();
    latestLocation = $location.value;

    renderCompanies(latestCompanies, false)
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

    addCompany(vanityName, function () {
        // Clear input field and disable Add button
        $companyUrlInput.value = '';
        $addCompanyButton.setAttribute('disabled', 'disabled');
    });
}

function renderCompany(company: Company) {
    // Create a new company card
    const $card = document.createElement('div');
    $card.className = 'company-card';

    // Create company name container
    const $companyNameContainer = document.createElement('div');
    $companyNameContainer.className = 'company-name';
    const $companyLink = document.createElement('a');
    $companyLink.href = `https://www.linkedin.com/company/${company.vanityName}`;
    $companyLink.target = '_blank';
    $companyLink.innerText = company.name;
    $companyNameContainer.appendChild($companyLink);
    $card.appendChild($companyNameContainer);

    // Create connections container
    const companyConnectionsContainer = document.createElement('div');
    companyConnectionsContainer.className = 'company-connections';
    {
        const $link = document.createElement('a');
        $link.href = company.connections1stURL;
        $link.innerText = `Connections 1st`;
        companyConnectionsContainer.appendChild($link);
    }
    {
        const $link = document.createElement('a');
        $link.href = company.connections2ndURL;
        $link.innerText = `Connections 2nd`;
        companyConnectionsContainer.appendChild($link);
    }
    {
        const $link = document.createElement('a');
        $link.href = company.connections1stXURL;
        $link.innerText = `Connections 1st X`;
        companyConnectionsContainer.appendChild($link);
    }
    {
        const $link = document.createElement('a');
        $link.href = company.connections2ndXURL;
        $link.innerText = `Connections 2nd X`
        companyConnectionsContainer.appendChild($link);
    }
    $card.appendChild(companyConnectionsContainer);

    // Create actions container (delete button)
    const companyActionsContainer = document.createElement('div');
    companyActionsContainer.className = 'company-actions';
    const $deleteButton = document.createElement('button');
    $deleteButton.innerText = 'Delete';
    $deleteButton.addEventListener('click', function () {
        deleteCompany(company.id);

        $card.remove();

        $companyListBlock.classList.toggle('hidden', $companyList.children.length === 0);
    });
    companyActionsContainer.appendChild($deleteButton);
    $card.appendChild(companyActionsContainer);

    return $card;
}

// Function to temporarily highlight the new company card
function highlightCard(card) {
    card.classList.add('highlight');
    setTimeout(() => {
        card.classList.remove('highlight');
    }, 2000); // Highlight lasts for 2 seconds
}

function renderCompanies(companies: Array<ResponseCompany>, added: boolean) {
    $companyList.innerHTML = '';

    companies.forEach((company, index) => {
        const prepared = prepareConnections(company, companies, [], '');

        const $card = renderCompany(new Company(
            company.id,
            company.alias,
            company.name,
            prepared.connections1stURL,
            prepared.connections2ndURL,
            prepared.connections1stXURL,
            prepared.connections2ndXURL,
        ));

        $companyList.appendChild($card);

        if (added && index === 0) {
            // Highlight the new company $card
            highlightCard($card);
        }
    })

    $companyListBlock.classList.toggle('hidden', companies.length === 0);
}

function prepareConnections(
    currentCompany: ResponseCompany,
    companies: Array<ResponseCompany>,
    universities: Array<number>,
): Connections {

    const pastCompanies = [];
    for (const company of companies) {
        if (currentCompany === company) {
            continue;
        }

        pastCompanies.push(company.id.toString());
    }

    // const universitiesIds = [];
    // for (const university of universities) {
    //     universitiesIds.push(university.toString());
    // }

    const currentCompanyQueryParam = `["${currentCompany.id}"]`;

    const universitiesQueryParam = JSON.stringify(["818029","850102","364340","496320","1198954","1257361","15250306","15251128","15099424","782774","15101979","15101061","80424966","6261241","658198","11443062","15099425","15099711","15101057","15102004","18080249","15143861","15101046","1599158","15101060","15100187","9029417","7991636","15101074","27066401","18144134","15101998","15149751","18691495","15099038"]);

    let connections1stURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections1stXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');

    connections1stURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stURL.searchParams.append('network', `["F"]`);
    connections1stURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections2ndURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndURL.searchParams.append('network', `["S"]`);
    connections2ndURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections1stXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stXURL.searchParams.append('network', `["F"]`);
    connections1stXURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections2ndXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndXURL.searchParams.append('network', `["S"]`);
    connections2ndXURL.searchParams.append('schoolFilter', universitiesQueryParam);

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
