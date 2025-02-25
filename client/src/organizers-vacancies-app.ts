import {organizersWelcome} from "./welcome";

import urlStateContainer from "./framework/vacancy_url_state_container";
import {
    VACANCY_SEARCH_QUERY,
    VACANCY_COMPANY_TYPE_CRITERIA_NAME,
    VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME,
    VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    VACANCY_REMOTE_CRITERIA_NAME,
    VACANCY_IN_FAVORITES_CRITERIA_NAME,
    VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
} from "./framework/vacancy_criteria_names";
import {InputCheckboxes} from "./framework/checkboxes";
import {companyTypes} from "./framework/company_types";
import {industries} from "./framework/industries";
import {hasEmployeesFromCountries} from "./framework/has_employees_from_countries";
import {htmlToNode} from "./framework/html";
import {Alias} from "./framework/alias";
import {renderSelected} from "./framework/selected_criteria";
import {firstQuerySelector} from "./framework/query_selector";
import {setStateByURLMapper} from "./framework/set_state_by_url";
import {toEnter} from "./framework/enter";

function markVacancyFavorite(vacancyId: number, favorite: boolean, callback: () => void) {
    fetch(`/api/v1/vacancies/${vacancyId}/favorite.json`, {
        method: "PATCH",
        body: JSON.stringify({
            favorite: favorite,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        callback();
    }).catch(console.error);
}

const $vacancies = document.querySelectorAll(".js-vacancy");
const $resultCount = document.getElementById("js-result-count");

$vacancies.forEach(function ($vacancy: HTMLElement) {
    const vacancyId = parseInt($vacancy.getAttribute("data-vacancy-id"));

    const $favorite = $vacancy.querySelector(".js-vacancy-favorite");
    $favorite.addEventListener("click", function () {
        const current = $favorite.classList.contains("in-favorite");
        const next = !current;

        markVacancyFavorite(vacancyId, next, function () {
            if (next) {
                $favorite.classList.add("in-favorite");

                $favorite.setAttribute("title", "Remove from favorites")
            } else {
                $favorite.classList.remove("in-favorite");

                $favorite.setAttribute("title", "Add to favorites")
            }
        });
    });
});


const $search = document.getElementById("js-vacancy-query") as HTMLInputElement;
const $typeCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-type") as any as Array<HTMLInputElement>);
const $industryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-company-industry") as any as Array<HTMLInputElement>);
const $hasEmployeesFromCountryCheckboxes = new InputCheckboxes(document.querySelectorAll("input.js-criteria-has-employees-from-country") as any as Array<HTMLInputElement>);
const $inRustFoundationMembersCheckbox = document.getElementById("js-criteria-rust-foundation-members") as HTMLInputElement;
const $remoteCheckbox = document.getElementById("js-criteria-remote") as HTMLInputElement;
const $inFavoritesCheckbox = document.getElementById("js-criteria-in-favorites") as HTMLInputElement;
const $selectedCriteria = document.getElementById("js-vacancy-selected-criteria");
const $reset = document.getElementById("js-criteria-reset");

$typeCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_TYPE_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

$industryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

$hasEmployeesFromCountryCheckboxes.onChange(function (state: Array<string>) {
    urlStateContainer.setArrayCriteria(VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME, state);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

if ($inRustFoundationMembersCheckbox) {
    $inRustFoundationMembersCheckbox.addEventListener("change", function () {
        urlStateContainer.setBoolCriteria(VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME, $inRustFoundationMembersCheckbox.checked);
        urlStateContainer.setPage(1);
        urlStateContainer.storeCurrentState();

        renderSelectedCriteriaByURL();

        search();
    });
}

$remoteCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(VACANCY_REMOTE_CRITERIA_NAME, $remoteCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
});

$inFavoritesCheckbox.addEventListener("change", function () {
    urlStateContainer.setBoolCriteria(VACANCY_IN_FAVORITES_CRITERIA_NAME, $inFavoritesCheckbox.checked);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    renderSelectedCriteriaByURL();

    search();
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
    $reset.style.visibility = visibility;
    $selectedCriteria.parentElement.style.visibility = visibility;
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

    search();
}

$search.addEventListener("keyup", toEnter(handleSearch));
$search.addEventListener("change", function () {
    urlStateContainer.setStringCriteria(VACANCY_SEARCH_QUERY, $search.value);
});
$search.addEventListener("search", handleSearch);

$reset.addEventListener("click", function () {
    urlStateContainer.keep(VACANCY_SEARCH_QUERY);
    urlStateContainer.setPage(1);
    urlStateContainer.storeCurrentState();

    updatePageState();
});

function updatePageState() {
    setStateByURL();

    renderSelectedCriteriaByURL();

    search();
}

function search() {
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
            $vacancy.style.display = "block";

            total++;

            return;
        }

        $vacancy.style.display = "none";
    });

    $resultCount.innerHTML = total.toString();
}

updatePageState();
