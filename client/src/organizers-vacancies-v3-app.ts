import urlStateContainer from "./framework/vacancy_url_state_container";
import {
    VACANCY_SEARCH_QUERY,
    VACANCY_LOCATION_CRITERIA_NAME,
    VACANCY_COMPANY_TYPE_CRITERIA_NAME,
    VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME,
    VACANCY_GLASSDOOR_RATING_CRITERIA_NAME,
    VACANCY_LINKEDIN_COMPANY_SIZE_CRITERIA_NAME,
    VACANCY_CLOUD_PROVIDER_CRITERIA_NAME,
    VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    VACANCY_REMOTE_CRITERIA_NAME,
    VACANCY_IN_FAVORITES_CRITERIA_NAME,
    VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
} from "./framework/vacancy_criteria_names";
import {InputCheckboxes, RadioCheckboxes} from "./framework/checkboxes";
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

import {parseCurrentOrganizerAlias} from "./organizer";
import {organizersWelcome} from "./welcome";
import {VacancyResponse} from "./organizers-vacancies-v3-models
import {renderVacancy} from "./organizers-vacancies-v3-render-vacancy";
import {Pager, TotalPages} from "./framework/pager";
import Pagination from "./framework/pagination";
import {addVacancyFavoriteEvent} from "./organizers-vacancies-favorite";

const LIMIT = 20;

const currentOrganizerAlias = parseCurrentOrganizerAlias(window.location.pathname);

let sourceVacancies: Array<VacancyResponse> = [];
let currentStateVacancies: Array<VacancyResponse> = [];

const $vacanciesContainer = document.getElementById("js-vacancies-container");
const $pagination = document.getElementById("js-pagination-pages");
const $resultCount = document.getElementById("js-result-count");
const $paginationShowMoreButton = document.getElementById("js-pagination-show-more-button") as HTMLButtonElement;
const $paginationShowAllButton = document.getElementById("js-pagination-show-all-button") as HTMLButtonElement;

const $form = document.getElementById("js-vacancy-search-form");
const $search = document.getElementById("js-vacancy-query") as HTMLInputElement;
const $locationCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-location") as any as Array<HTMLInputElement>);
const $typeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-type") as any as Array<HTMLInputElement>);
const $industryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-industry") as any as Array<HTMLInputElement>);
const $glassdoorRatingCheckboxes = new RadioCheckboxes(document.querySelectorAll("input.js-criteria-glassdoor-rating") as any as Array<HTMLInputElement>);
const $linkedinCompanySizeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-linkedin-company-size") as any as Array<HTMLInputElement>);
const $hasEmployeesFromCountryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-has-employees-from-country") as any as Array<HTMLInputElement>);
const $remoteCheckbox = document.getElementById("js-criteria-remote") as HTMLInputElement;
const $inFavoritesCheckbox = document.getElementById("js-criteria-in-favorites") as HTMLInputElement;
const $inRustFoundationMembersCheckbox = document.getElementById("js-criteria-rust-foundation-members") as HTMLInputElement;
const $selectedCriteria = document.getElementById("js-vacancy-selected-criteria");
const $optionalMobileSelectedCriteriaCount = document.getElementById("js-mobile-selected-criteria-count");
const $resetButtons = document.querySelectorAll(".js-criteria-reset") as any as Array<HTMLElement>;

const pager = new Pager(LIMIT);
const pagination = new Pagination($pagination, setPage);

$typeCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_TYPE_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$industryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$hasEmployeesFromCountryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

if ($inRustFoundationMembersCheckbox) {
    $inRustFoundationMembersCheckbox.addEventListener("change", function () {
        urlStateContainer.setBoolCriteria(VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, $inRustFoundationMembersCheckbox.checked);
        urlStateContainer.setPage(1);
        urlStateContainer.storeCurrentState();

        renderSelectedCriteriaByURL();

        search(true, true);
    });
}

$remoteCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(VACANCY_REMOTE_CRITERIA_NAME, $remoteCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search(true, true);
});

$inFavoritesCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(VACANCY_IN_FAVORITES_CRITERIA_NAME, $inFavoritesCheckbox.checked);
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
    setInputStateByURL($search, VACANCY_SEARCH_QUERY);

    setCheckboxesStateByURL($typeCheckboxes, VACANCY_COMPANY_TYPE_CRITERIA_NAME);
    setCheckboxesStateByURL($industryCheckboxes, VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME);
    setCheckboxesStateByURL($hasEmployeesFromCountryCheckboxes, VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME);

    if ($inRustFoundationMembersCheckbox) {
        setCheckboxStateByURL($inRustFoundationMembersCheckbox, VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME);
    }
    setCheckboxStateByURL($remoteCheckbox, VACANCY_REMOTE_CRITERIA_NAME);
    setCheckboxStateByURL($inFavoritesCheckbox, VACANCY_IN_FAVORITES_CRITERIA_NAME);
}

function renderSelectedCriteriaByURL() {
    const $views: Array<HTMLElement> = [];

    renderSelectedCheckboxes($views, VACANCY_COMPANY_TYPE_CRITERIA_NAME, companyTypes);
    renderSelectedCheckboxes($views, VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME, industries);
    renderSelectedCheckboxes($views, VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, hasEmployeesFromCountries);
    renderSelectedCheckbox($views, VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, "Rust Foundation Members");
    renderSelectedCheckbox($views, VACANCY_REMOTE_CRITERIA_NAME, "Remote");
    renderSelectedCheckbox($views, VACANCY_IN_FAVORITES_CRITERIA_NAME, "Favorites");

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

    urlStateContainer.setStringCriteria(VACANCY_SEARCH_QUERY, $search.value);

    handleSearch();
});

for (const $resetButton of $resetButtons) {
    $resetButton.addEventListener("click", function () {
        urlStateContainer.keep(VACANCY_SEARCH_QUERY);
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
    const moreCount = currentStateVacancies.length - (pager.getOffset() + LIMIT);
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

        renderVacancies(currentStateVacancies.slice(offset, offset + LIMIT), true);
        pagination.render(nextPage, TotalPages(currentStateVacancies.length, LIMIT), urlByPageBuilder);

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

        renderVacancies(currentStateVacancies.slice(offset, offset + LIMIT), false);
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
        renderVacancies(currentStateVacancies.slice(pager.getOffset()), false);
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
    const types = urlStateContainer.getCriteria(VACANCY_COMPANY_TYPE_CRITERIA_NAME, []);
    const industries = urlStateContainer.getCriteria(VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME, []);
    const hasEmployeesFromCountries = urlStateContainer.getCriteria(VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, []);
    const isRustFoundationMembers = urlStateContainer.getCriteria(VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, false);
    const remote = urlStateContainer.getCriteria(VACANCY_REMOTE_CRITERIA_NAME, false);
    const inFavorites = urlStateContainer.getCriteria(VACANCY_IN_FAVORITES_CRITERIA_NAME, false);

    const matchQuery = function ($vacancy: HTMLElement): boolean {
        if (query.length === 0) {
            return true;
        }

        if ($vacancy.getAttribute("data-vacancy-title").toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        if ($vacancy.getAttribute("data-company-name").toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        if ($vacancy.querySelector(".js-vacancy-short-description").textContent.toLowerCase().indexOf(query) !== -1) {
            return true;
        }

        return false;
    }

    const matchIndustry = function ($vacancy: HTMLElement): boolean {
        if (industries.length === 0) {
            return true;
        }

        const companyIndustries = $vacancy.getAttribute("data-company-industries").split(",");

        for (const industry of industries) {
            if (companyIndustries.indexOf(industry) !== -1) {
                return true;
            }
        }

        return false;
    }

    const matchHasEmployeesFromCountry = function ($vacancy: HTMLElement): boolean {
        if (hasEmployeesFromCountries.length === 0) {
            return true;
        }

        const companyHasEmployeesFromCountries = $vacancy.getAttribute("data-company-has-employees-from-countries").split(",");

        for (const country of hasEmployeesFromCountries) {
            if (companyHasEmployeesFromCountries.indexOf(country) !== -1) {
                return true;
            }
        }

        return false;
    }

    const match = function ($vacancy: HTMLElement): boolean {
        if (!matchQuery($vacancy)) {
            return false;
        }

        if (types.length > 0 && types.indexOf($vacancy.getAttribute("data-company-type")) === -1) {
            return false;
        }

        if (!matchIndustry($vacancy)) {
            return false;
        }

        if (!matchHasEmployeesFromCountry($vacancy)) {
            return false;
        }

        if (isRustFoundationMembers && $vacancy.getAttribute("data-company-rust-foundation-members") !== "true") {
            return false;
        }

        if (remote && $vacancy.getAttribute("data-vacancy-remote") !== "true") {
            return false;
        }

        if (inFavorites) {
            const $favorite = $vacancy.querySelector(".js-vacancy-favorite");
            const current = $favorite.classList.contains("in-favorite");

            if (!current) {
                return false;
            }
        }

        return true;
    }

    let total = 0;

    $vacancies.forEach(function ($vacancy: HTMLElement) {
        if (match($vacancy)) {
            $vacancy.hidden = false;

            total++;

            return;
        }

        $vacancy.hidden = true;
    });

    $resultCount.innerHTML = total.toString();
}

function fetchVacancies(callback: (vacancies: Array<VacancyResponse>) => void) {
    fetch(`/api/v1/unsafe/${currentOrganizerAlias}/vacancies.json`, {
        method: "GET",
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        return response.json();
    }).then(function (data) {
        callback(data.vacancies);
    }).catch(console.error);
}

function renderVacancies(vacancies: Array<VacancyResponse>, clear: boolean = true) {
    const length = Math.min(vacancies.length, 10);

    const $vacancies = new Array<HTMLElement>(length);

    for (let i = 0; i < length; i++) {
        const vacancy = vacancies[i];
        const $vacancy = htmlToNode(renderVacancy(vacancy));

        addVacancyFavoriteEvent($vacancy, vacancy);

        $vacancies[i] = $vacancy;
    }

    if (clear) {
        $vacanciesContainer.innerHTML = "";
    }
    $vacanciesContainer.append(...$vacancies);
}

function init(vacancies: Array<VacancyResponse>) {
    sourceVacancies = vacancies;

    search(true, false);
}

pager.setPage(urlStateContainer.getPage());

fetchVacancies(init);

setStateByURL();

renderSelectedCriteriaByURL();

responsiveHeaderProfileWidget();

responsiveFilterWidget();

githubStarsWidget();
