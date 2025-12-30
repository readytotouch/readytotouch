import {parseCurrentProgrammingLanguage} from "./pl";
import {organizersWelcome} from "./welcome";
import {CompanyResponse} from "./organizers-companies-v3-models";
import {htmlToNode} from "./framework/html";
import {renderCompany} from "./organizers-companies-v3-render-company";
import {addCompanyFavoriteEvent} from "./organizers-companies-favorite";
import {responsiveHeaderProfileWidget} from "./responsive-header-profile-widget";
import {responsiveFilterWidget} from "./responsive-filter-widget";
import {responsiveCompanyShowMoreWidget} from "./responsive-company-show-more-widget";
import {githubStarsWidget} from "./github-stars-widget";

const currentProgrammingLanguage = parseCurrentProgrammingLanguage(window.location.pathname);

function fetchCompanies(callback: (companies: Array<CompanyResponse>) => void) {
    fetch(`/api/v1/unsafe/${currentProgrammingLanguage}/companies.json`, {
        method: "GET",
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        return response.json();
    }).then(function (data) {
        callback(data.companies);
    }).catch(console.error);
}

const $companiesContainer = document.getElementById("js-companies-container");

function renderCompanies(companies: Array<CompanyResponse>) {
    const length = Math.min(companies.length, 10);

    const $companies = new Array<HTMLElement>(length);

    for (let i = 0; i < length; i++) {
        const $company = htmlToNode(renderCompany(companies[i], false));

        addCompanyFavoriteEvent($company);

        $companies[i] = $company;
    }

    $companiesContainer.append(...$companies);
}

fetchCompanies(renderCompanies);

responsiveHeaderProfileWidget();

responsiveFilterWidget();

responsiveCompanyShowMoreWidget();

githubStarsWidget();
