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
import {Pager, TotalPages} from "./framework/pager";
import Pagination from "./framework/pagination";

const LIMIT = 10;

const currentProgrammingLanguage = parseCurrentProgrammingLanguage(window.location.pathname);

let sourceCompanies: Array<CompanyResponse> = [];
let currentStateCompanies: Array<CompanyResponse> = [];

const $companiesContainer = document.getElementById("js-companies-container");
const $pagination = document.getElementById("js-pagination-pages");
const $resultCount = document.getElementById("js-result-count");
const $paginationShowMoreButton = document.getElementById("js-pagination-show-more-button") as HTMLButtonElement;
const $paginationShowAllButton = document.getElementById("js-pagination-show-all-button") as HTMLButtonElement;
const $paginationShowAllCount = document.getElementById("js-pagination-show-all-count") as HTMLElement;

const $form = document.getElementById("js-company-search-form");
const $search = document.getElementById("js-company-query") as HTMLInputElement;
const $typeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-type") as any as Array<HTMLInputElement>);
const $industryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-industry") as any as Array<HTMLInputElement>);
const $hasEmployeesFromCountryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-has-employees-from-country") as any as Array<HTMLInputElement>);
const $remoteCheckbox = document.getElementById("js-criteria-remote") as HTMLInputElement;
const $inFavoritesCheckbox = document.getElementById("js-criteria-in-favorites") as HTMLInputElement;
const $inRustFoundationMembersCheckbox = document.getElementById("js-criteria-rust-foundation-members") as HTMLInputElement;
const $selectedCriteria = document.getElementById("js-company-selected-criteria");
const $optionalMobileSelectedCriteriaCount = document.getElementById("js-mobile-selected-criteria-count");
// "#js-criteria-reset" for backward compatibility
const $resetButtons = document.querySelectorAll("#js-criteria-reset, .js-criteria-reset") as any as Array<HTMLElement>;

const pager = new Pager(LIMIT);
const pagination = new Pagination($pagination, setPage);

$typeCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_TYPE_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$industryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_INDUSTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$hasEmployeesFromCountryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

if ($inRustFoundationMembersCheckbox) {
    $inRustFoundationMembersCheckbox.addEventListener("change", function () {
        urlStateContainer.setBoolCriteria(COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, $inRustFoundationMembersCheckbox.checked);
        urlStateContainer.setPage(1);
        urlStateContainer.storeCurrentState();

        renderSelectedCriteriaByURL();

        search(true, true);
    });
}

$remoteCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(COMPANY_REMOTE_CRITERIA_NAME, $remoteCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$inFavoritesCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(COMPANY_IN_FAVORITES_CRITERIA_NAME, $inFavoritesCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

const {
    setInputStateByURL,
    setCheckboxesStateByURL,
    setCheckboxStateByURL,
} = setStateByURLMapper(urlStateContainer);

function setStateByURL() {
    setInputStateByURL($search, COMPANY_SEARCH_QUERY);

    setCheckboxesStateByURL($typeCheckboxes, COMPANY_TYPE_CRITERIA_NAME);
    setCheckboxesStateByURL($industryCheckboxes, COMPANY_INDUSTRY_CRITERIA_NAME);
    setCheckboxesStateByURL($hasEmployeesFromCountryCheckboxes, COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME);

    if ($inRustFoundationMembersCheckbox) {
        setCheckboxStateByURL($inRustFoundationMembersCheckbox, COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME);
    }
    setCheckboxStateByURL($remoteCheckbox, COMPANY_REMOTE_CRITERIA_NAME);
    setCheckboxStateByURL($inFavoritesCheckbox, COMPANY_IN_FAVORITES_CRITERIA_NAME);
}

function renderSelectedCriteriaByURL() {
    const $views: Array<HTMLElement> = [];

    renderSelectedCheckboxes($views, COMPANY_TYPE_CRITERIA_NAME, companyTypes);
    renderSelectedCheckboxes($views, COMPANY_INDUSTRY_CRITERIA_NAME, industries);
    renderSelectedCheckboxes($views, COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, hasEmployeesFromCountries);
    renderSelectedCheckbox($views, COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, "Rust Foundation Members");
    renderSelectedCheckbox($views, COMPANY_REMOTE_CRITERIA_NAME, "Remote");
    renderSelectedCheckbox($views, COMPANY_IN_FAVORITES_CRITERIA_NAME, "Favorites");

    $selectedCriteria.innerHTML = "";
    $selectedCriteria.append(...$views);

    const visibility = $views.length === 0 ? "hidden" : "";
    for (const $resetButton of $resetButtons) {
        $resetButton.style.visibility = visibility;
    }
    $selectedCriteria.parentElement.style.visibility = visibility;
    if ($optionalMobileSelectedCriteriaCount) {
        $optionalMobileSelectedCriteriaCount.innerHTML = $views.length.toString();
        $optionalMobileSelectedCriteriaCount.style.visibility = visibility;
    }
}

function renderSelectedCheckboxes(
    $views: Array<HTMLElement>,
    criteriaName: string,
    toAliases: (aliases: Array<string>) => Array<Alias>,
) {
    const aliases = urlStateContainer.getCriteria(criteriaName, []);

    toAliases(aliases).forEach(function (alias: Alias) {
        const $view = htmlToNode(renderSelected(alias.name, alias.image));

        firstQuerySelector($view, "button").addEventListener("click", function () {
            urlStateContainer.removeAlias(criteriaName, alias.alias);
            urlStateContainer.setPage(1);
            urlStateContainer.storeCurrentState();

            updatePageState();
        });

        $views.push($view);
    });
}

function renderSelectedCheckbox($views: Array<HTMLElement>, criteriaName: string, title: string) {
    const checked = urlStateContainer.getCriteria(criteriaName, false);

    if (checked) {
        const $view = htmlToNode(renderSelected(title));

        firstQuerySelector($view, "button").addEventListener("click", function () {
            urlStateContainer.remove(criteriaName);
            urlStateContainer.setPage(1);
            urlStateContainer.storeCurrentState();

            updatePageState();
        });

        $views.push($view);
    }
}

const handleSearch = function () {
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    search(true, true);
}

$form.addEventListener("submit", function (event) {
    event.preventDefault();

    urlStateContainer.setStringCriteria(COMPANY_SEARCH_QUERY, $search.value);

    handleSearch();
});

for (const $resetButton of $resetButtons) {
    $resetButton.addEventListener("click", function () {
        urlStateContainer.keep(COMPANY_SEARCH_QUERY);
        urlStateContainer.setPage(1);
        urlStateContainer.storeCurrentState();

        updatePageState();
    });
}

function updatePageState() {
    setStateByURL();

    renderSelectedCriteriaByURL();

    search(true, true);
}

function updateMoreButtonsVisibility() {
    const moreCount = currentStateCompanies.length - (pager.getOffset() + LIMIT);
    const hide = pager.getPage() > 1 || moreCount <= 0;

    $paginationShowMoreButton.hidden = hide;
    $paginationShowAllButton.hidden = hide;

    if (hide) {
        return;
    }

    $paginationShowAllCount.innerHTML = moreCount.toString();
}

function setPage(page: number) {
    pager.setPage(page);

    urlStateContainer.setPage(page);

    urlStateContainer.storeCurrentState();

    // Faster to just render from current state than re-searching
    {
        const offset = pager.getOffset();
        const nextPage = pager.getPage();
        const urlByPageBuilder = urlStateContainer.createUrlByPageBuilder();

        renderCompanies(currentStateCompanies.slice(offset, offset + LIMIT), true);
        pagination.render(nextPage, TotalPages(currentStateCompanies.length, LIMIT), urlByPageBuilder);

        updateMoreButtonsVisibility();
    }
}

$paginationShowMoreButton.addEventListener("click", function () {
    $paginationShowAllButton.disabled = true;
    $paginationShowMoreButton.disabled = true;

    pager.increment();

    // Faster to just render from current state than re-searching
    {
        const offset = pager.getOffset();

        renderCompanies(currentStateCompanies.slice(offset), false);
        pagination.reset();

        $paginationShowMoreButton.hidden = true;
        $paginationShowAllButton.hidden = true;
    }

    $paginationShowAllButton.disabled = false;
    $paginationShowMoreButton.disabled = false;
});

$paginationShowAllButton.addEventListener("click", function () {
    $paginationShowAllButton.disabled = true;
    $paginationShowMoreButton.disabled = true;

    {
        renderCompanies(currentStateCompanies.slice(pager.getOffset() + LIMIT), false);
        pagination.reset();

        $paginationShowMoreButton.hidden = true;
        $paginationShowAllButton.hidden = true;
    }

    $paginationShowAllButton.disabled = false;
    $paginationShowMoreButton.disabled = false;
});

function search(replaceHTML: boolean, resetPager: boolean) {
    if (resetPager) {
        pager.reset();
    }

    const query = $search.value.trim().toLowerCase();
    const types = urlStateContainer.getCriteria(COMPANY_TYPE_CRITERIA_NAME, []);
    const industries = urlStateContainer.getCriteria(COMPANY_INDUSTRY_CRITERIA_NAME, []);
    const hasEmployeesFromCountries = urlStateContainer.getCriteria(COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, []);
    const isRustFoundationMembers = urlStateContainer.getCriteria(COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, false);
    const remote = urlStateContainer.getCriteria(COMPANY_REMOTE_CRITERIA_NAME, false);
    const inFavorites = urlStateContainer.getCriteria(COMPANY_IN_FAVORITES_CRITERIA_NAME, false);

    const matchQuery = function (company: CompanyResponse): boolean {
        if (query.length === 0) {
            return true;
        }

        if (company.name.toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        if (company.short_description.toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        return false;
    }

    const matchIndustry = function (company: CompanyResponse): boolean {
        if (industries.length === 0) {
            return true;
        }

        if (!company.industries) {
            return false;
        }

        for (const industry of industries) {
            for (const companyIndustry of company.industries) {
                if (companyIndustry.alias === industry) {
                    return true;
                }
            }
        }

        return false;
    }

    const matchHasEmployeesFromCountry = function (company: CompanyResponse): boolean {
        if (hasEmployeesFromCountries.length === 0) {
            return true;
        }

        if (!company.has_employees_from_countries) {
            return false;
        }

        for (const country of hasEmployeesFromCountries) {
            for (const companyCountry of company.has_employees_from_countries) {
                if (companyCountry.alias === country) {
                    return true;
                }
            }
        }

        return false;
    }

    const match = function (company: CompanyResponse): boolean {
        if (!matchQuery(company)) {
            return false;
        }

        if (types.length > 0 && types.indexOf(company.type) === -1) {
            return false;
        }

        if (!matchIndustry(company)) {
            return false;
        }

        if (!matchHasEmployeesFromCountry(company)) {
            return false;
        }

        if (isRustFoundationMembers && !company.rust_foundation_member) {
            return false;
        }

        if (remote && !company.remote) {
            return false;
        }

        if (inFavorites) {
            // @TODO: replace with real favorites check
            return false;
        }

        return true;
    }

    const offset = pager.getOffset();
    const nextPage = pager.getPage();
    const urlByPageBuilder = urlStateContainer.createUrlByPageBuilder();

    // hide container to prevent multiple reflows
    {
        // debug time measurement
        const start = performance.now();

        currentStateCompanies = sourceCompanies.filter(match);

        // debug time measurement
        const end = performance.now();

        console.log(`Search took ${end - start} milliseconds.`);
    }

    renderCompanies(currentStateCompanies.slice(offset, offset + LIMIT), replaceHTML);
    pagination.render(nextPage, TotalPages(currentStateCompanies.length, LIMIT), urlByPageBuilder);
    $resultCount.innerHTML = currentStateCompanies.length.toString();

    updateMoreButtonsVisibility();
}

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

function renderCompanies(companies: Array<CompanyResponse>, clear: boolean = true) {
    const length = Math.min(companies.length, 10);

    const $companies = new Array<HTMLElement>(length);

    for (let i = 0; i < length; i++) {
        const $company = htmlToNode(renderCompany(companies[i], false));

        addCompanyFavoriteEvent($company);

        $companies[i] = $company;
    }

    if (clear) {
        $companiesContainer.innerHTML = "";
    }
    $companiesContainer.append(...$companies);
}

function init(companies: Array<CompanyResponse>) {
    sourceCompanies = companies;

    search(true, false);
}

pager.setPage(urlStateContainer.getPage());

fetchCompanies(init);

setStateByURL();

renderSelectedCriteriaByURL();

responsiveHeaderProfileWidget();

responsiveFilterWidget();

responsiveCompanyShowMoreWidget();

githubStarsWidget();
