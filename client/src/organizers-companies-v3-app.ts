import urlStateContainer from "./framework/company_url_state_container";
import {
    COMPANY_SEARCH_QUERY,
    COMPANY_TYPE_CRITERIA_NAME,
    COMPANY_INDUSTRY_CRITERIA_NAME,
    COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
    COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    COMPANY_REMOTE_CRITERIA_NAME,
    COMPANY_IN_FAVORITES_CRITERIA_NAME,
} from "./framework/company_criteria_names";
import {InputCheckboxes} from "./framework/checkboxes";
import {companyTypes} from "./framework/company_types";
import {industries} from "./framework/industries";
import {hasEmployeesFromCountries} from "./framework/has_employees_from_countries";
import {htmlToNode} from "./framework/html";
import {Alias} from "./framework/alias";
import {renderSelected} from "./framework/selected_criteria";
import {firstQuerySelector} from "./framework/query_selector";
import {setStateByURLMapper} from "./framework/set_state_by_url";
import {responsiveHeaderProfileWidget} from "./responsive-header-profile-widget";
import {githubStarsWidget} from "./github-stars-widget";
import {responsiveFilterWidget} from "./responsive-filter-widget";
import {responsiveCompanyShowMoreWidget} from "./responsive-company-show-more-widget";
import {addCompanyFavoriteEvent} from "./organizers-companies-favorite";

import {parseCurrentProgrammingLanguage} from "./pl";
import {organizersWelcome} from "./welcome";
import {CompanyResponse} from "./organizers-companies-v3-models";
import {renderCompany} from "./organizers-companies-v3-render-company";

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
