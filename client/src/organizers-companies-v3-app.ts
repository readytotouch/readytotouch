import urlStateContainer from "./framework/company_url_state_container";
import {
    COMPANY_SEARCH_QUERY,
    COMPANY_TYPE_CRITERIA_NAME,
    COMPANY_INDUSTRY_CRITERIA_NAME,
    COMPANY_GLASSDOOR_RATING_CRITERIA_NAME,
    COMPANY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME,
    COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
    COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    COMPANY_REMOTE_CRITERIA_NAME,
    COMPANY_IN_FAVORITES_CRITERIA_NAME,
} from "./framework/company_criteria_names";
import {InputCheckboxes, RadioCheckboxes} from "./framework/checkboxes";
import {companyTypes} from "./framework/company_types";
import {industries} from "./framework/industries";
import {ratingRatesAliases} from "./framework/rating_rates";
import {linkedinCompanySizes} from "./framework/linkedin_company_sizes";
import {hasEmployeesFromCountries} from "./framework/has_employees_from_countries";
import {htmlToNode} from "./framework/html";
import {Alias} from "./framework/alias";
import {renderSelected} from "./framework/selected_criteria";
import {firstQuerySelector} from "./framework/query_selector";
import {setStateByURLMapper} from "./framework/set_state_by_url";
import {responsiveHeaderProfileWidget} from "./responsive-header-profile-widget";
import {githubStarsWidget} from "./github-stars-widget";
import {responsiveFilterWidget} from "./responsive-filter-widget";
import {addCompanyShowMoreEvent} from "./responsive-company-show-more-widget";
import {addCompanyFavoriteEvent} from "./organizers-companies-favorite";

import {parseCurrentOrganizerAlias} from "./organizer";
import {organizersWelcome} from "./welcome";
import {CompanyResponse} from "./organizers-companies-v3-models";
import {renderCompany} from "./organizers-companies-v3-render-company";
import {Pager, TotalPages} from "./framework/pager";
import Pagination from "./framework/pagination";
import {renderSponsored, sponsoredAvailable} from "./organizers-sponsored";

const LIMIT = 10;

const currentOrganizerAlias = parseCurrentOrganizerAlias(window.location.pathname);

let sourceCompanies: Array<CompanyResponse> = [];
let currentStateCompanies: Array<CompanyResponse> = [];

const $companiesContainer = document.getElementById("js-companies-container");
const $pagination = document.getElementById("js-pagination-pages");
const $resultCount = document.getElementById("js-result-count");
const $paginationShowMoreButton = document.getElementById("js-pagination-show-more-button") as HTMLButtonElement;
const $paginationShowAllButton = document.getElementById("js-pagination-show-all-button") as HTMLButtonElement;

const $form = document.getElementById("js-company-search-form");
const $search = document.getElementById("js-company-query") as HTMLInputElement;
const $typeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-type") as any as Array<HTMLInputElement>);
const $industryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-industry") as any as Array<HTMLInputElement>);
const $glassdoorRatingCheckboxes = new RadioCheckboxes(document.querySelectorAll("input.js-criteria-glassdoor-rating") as any as Array<HTMLInputElement>);
const $linkedinCompanySizeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-linkedin-company-size") as any as Array<HTMLInputElement>);
const $hasEmployeesFromCountryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-has-employees-from-country") as any as Array<HTMLInputElement>);
const $remoteCheckbox = document.getElementById("js-criteria-remote") as HTMLInputElement;
const $inFavoritesCheckbox = document.getElementById("js-criteria-in-favorites") as HTMLInputElement;
const $inRustFoundationMembersCheckbox = document.getElementById("js-criteria-rust-foundation-members") as HTMLInputElement;
const $selectedCriteria = document.getElementById("js-company-selected-criteria");
const $optionalMobileSelectedCriteriaCount = document.getElementById("js-mobile-selected-criteria-count");
const $resetButtons = document.querySelectorAll(".js-criteria-reset") as any as Array<HTMLElement>;

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

$glassdoorRatingCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_GLASSDOOR_RATING_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$linkedinCompanySizeCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(COMPANY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME, state);
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
    setCheckboxesStateByURL($glassdoorRatingCheckboxes, COMPANY_GLASSDOOR_RATING_CRITERIA_NAME);
    setCheckboxesStateByURL($linkedinCompanySizeCheckboxes, COMPANY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME);
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
    renderSelectedCheckboxes($views, COMPANY_GLASSDOOR_RATING_CRITERIA_NAME, ratingRatesAliases);
    renderSelectedCheckboxes($views, COMPANY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME, linkedinCompanySizes);
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

    $paginationShowMoreButton.classList.toggle("d-none", hide);
    $paginationShowAllButton.classList.toggle("d-none", hide);

    if (hide) {
        return;
    }

    $paginationShowAllButton.innerHTML = `Show All (+${moreCount})`;
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

        renderCompanies(currentStateCompanies.slice(offset, offset + LIMIT), true, pager.getPage() === 1);
        pagination.render(nextPage, TotalPages(currentStateCompanies.length, LIMIT), urlByPageBuilder);

        updateMoreButtonsVisibility();
    }

    requestAnimationFrame(function () {
        window.scrollTo({
            top: 0,
            behavior: "smooth",
        });
    });
}

$paginationShowMoreButton.addEventListener("click", function () {
    $paginationShowAllButton.disabled = true;
    $paginationShowMoreButton.disabled = true;

    pager.incrementOffsetOnly();

    // Faster to just render from current state than re-searching
    {
        const offset = pager.getOffset();

        renderCompanies(currentStateCompanies.slice(offset, offset + LIMIT), false, false);
        pagination.reset();

        updateMoreButtonsVisibility();
    }

    $paginationShowAllButton.disabled = false;
    $paginationShowMoreButton.disabled = false;
});

$paginationShowAllButton.addEventListener("click", function () {
    $paginationShowAllButton.disabled = true;
    $paginationShowMoreButton.disabled = true;

    pager.incrementOffsetOnly();

    {
        renderCompanies(currentStateCompanies.slice(pager.getOffset()), false, false);
        pagination.reset();

        $paginationShowMoreButton.classList.toggle("d-none", true);
        $paginationShowAllButton.classList.toggle("d-none", true);
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
    const glassdoorReviewsRates = urlStateContainer.getCriteria(COMPANY_GLASSDOOR_RATING_CRITERIA_NAME, []);
    const linkedinCompanySizes = urlStateContainer.getCriteria(COMPANY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME, []);
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
    };

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
    };

    const matchGlassdoorReviewsRate = function (company: CompanyResponse): boolean {
        if (glassdoorReviewsRates.length === 0) {
            return true;
        }

        if (!company.glassdoor_profile) {
            return false;
        }

        if (!company.glassdoor_profile.reviews_rate) {
            return false;
        }

        return company.glassdoor_profile.reviews_rate >= glassdoorReviewsRates[0];
    };

    const matchLinkedinCompanySize = function (company: CompanyResponse): boolean {
        if (linkedinCompanySizes.length === 0) {
            return true;
        }

        return linkedinCompanySizes.indexOf(company.linkedin_profile.employees) !== -1;
    };

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
    };

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

        if (!matchGlassdoorReviewsRate(company)) {
            return false;
        }

        if (!matchLinkedinCompanySize(company)) {
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

        if (inFavorites && !company.favorite) {
            return false;
        }

        return true;
    };

    const offset = pager.getOffset();
    const nextPage = pager.getPage();
    const urlByPageBuilder = urlStateContainer.createUrlByPageBuilder();

    {
        // debug time measurement
        const start = performance.now();

        currentStateCompanies = sourceCompanies.filter(match);

        // debug time measurement
        const end = performance.now();

        console.log(`Search took ${end - start} milliseconds.`);
    }

    renderCompanies(currentStateCompanies.slice(offset, offset + LIMIT), replaceHTML, pager.getPage() === 1);
    pagination.render(nextPage, TotalPages(currentStateCompanies.length, LIMIT), urlByPageBuilder);
    $resultCount.innerHTML = currentStateCompanies.length.toString();

    updateMoreButtonsVisibility();
}

function fetchCompanies(callback: (companies: Array<CompanyResponse>) => void) {
    fetch(`/api/v1/unsafe/${currentOrganizerAlias}/companies.json`, {
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

function renderCompanies(
    companies: Array<CompanyResponse>,
    clear: boolean = true,
    showSponsored: boolean,
) {
    const $elements = [];

    for (let i = 0; i < companies.length; i++) {
        const company = companies[i];
        const sponsored = i === 0 && showSponsored && sponsoredAvailable(company.pinned_until, company.linkedin_profile.verified);
        const $company = htmlToNode(renderCompany(company, sponsored));

        addCompanyFavoriteEvent($company, company);
        addCompanyShowMoreEvent($company)

        $elements.push($company);
        if (sponsored) {
            $elements.push(htmlToNode(renderSponsored("company")))
        }
    }

    if (clear) {
        $companiesContainer.innerHTML = "";
    }
    $companiesContainer.append(...$elements);
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

githubStarsWidget();
