import {toEnter} from "./framework/enter";

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
const $companyNameInput = document.getElementById('company-name') as HTMLInputElement
const $addCompanyButton = document.getElementById('add-company') as HTMLButtonElement;

// Enable/Disable Add button based on input
$companyNameInput.addEventListener('input', function () {
    if ($companyNameInput.value.trim() !== '') {
        $addCompanyButton.removeAttribute('disabled');
    } else {
        $addCompanyButton.setAttribute('disabled', 'disabled');
    }
});

$companyNameInput.addEventListener("keyup", toEnter(submitCompany));

// Add Company button click event
$addCompanyButton.addEventListener('click', submitCompany);

function submitCompany() {
    const companyName = $companyNameInput.value;
    if (companyName) {
        const $card = renderCompany(new Company(
            1,
            'todo',
            companyName,
            'javascript:void(0)',
            'javascript:void(0)',
            'javascript:void(0)',
            'javascript:void(0)',
        ));

        // Append $card to company list
        $companyList.appendChild($card);

        // Highlight the new company $card
        highlightCard($card);

        $companyListBlock.classList.toggle('hidden', $companyList.children.length === 0);

        // Clear input field and disable Add button
        $companyNameInput.value = '';
        $addCompanyButton.setAttribute('disabled', 'disabled');
    }
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
    const deleteButton = document.createElement('button');
    deleteButton.innerText = 'Delete';
    deleteButton.addEventListener('click', function () {
        $card.remove();

        $companyListBlock.classList.toggle('hidden', $companyList.children.length === 0);
    });
    companyActionsContainer.appendChild(deleteButton);
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

fetch('/api/v1/companies-and-connections/companies.json')
    .then(function (response) {
        return response.json();
    })
    .then(renderCompanies)
    .catch(console.error);

function renderCompanies(companies: Array<ResponseCompany>) {
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

        if (index === 0) {
            // Highlight the new company $card
            highlightCard($card);
        }
    })

    $companyListBlock.classList.toggle('hidden', companies.length === 0);
}

function prepareConnections(
    currentCompany: ResponseCompany,
    companies: Array<ResponseCompany>,
    universities: Array<ResponseCompany>,
    keywords: string,
    ): Connections {

    const pastCompanies = [];
    const universitiesIds = [];

    for (const company of companies) {
        if (currentCompany === company) {
            continue;
        }

        pastCompanies.push(company.id.toString());
    }

    for (const university of universities) {
        universitiesIds.push(university.id.toString());
    }

    const currentCompanyQueryParam = `["${currentCompany.id}"]`;
    const pastCompaniesQueryParam = JSON.stringify(pastCompanies);
    const universitiesQueryParam = JSON.stringify(universitiesIds);

    let connections1stURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections1stXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');
    let connections2ndXURL = new URL('https://www.linkedin.com/search/results/PEOPLE/');

    connections1stURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stURL.searchParams.append('network', `["F"]`);
    connections1stURL.searchParams.append('keywords', keywords);
    connections1stURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections2ndURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndURL.searchParams.append('network', `["S"]`);
    connections2ndURL.searchParams.append('keywords', keywords);
    connections2ndURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections1stXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections1stXURL.searchParams.append('network', `["F"]`);
    connections1stXURL.searchParams.append('keywords', keywords);
    connections1stXURL.searchParams.append('pastCompany', pastCompaniesQueryParam);
    connections1stXURL.searchParams.append('schoolFilter', universitiesQueryParam);

    connections2ndXURL.searchParams.append('currentCompany', currentCompanyQueryParam);
    connections2ndXURL.searchParams.append('network', `["S"]`);
    connections2ndXURL.searchParams.append('keywords', keywords);
    connections2ndXURL.searchParams.append('pastCompany', pastCompaniesQueryParam);
    connections2ndXURL.searchParams.append('schoolFilter', universitiesQueryParam);

    return new Connections(
        connections1stURL.toString(),
        connections2ndURL.toString(),
        connections1stXURL.toString(),
        connections2ndXURL.toString(),
    );
}
