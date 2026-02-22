import {organizersWelcome} from "./welcome";
import {companyEmptyIndustries, companyFirstIndustry, CompanyResponse} from "./organizers-companies-v3-models";
import {Industry} from "./organizers-v3-common-models";

function overrideCompanyVisibility(companyId: number, callback: () => void) {
    fetch(`/api/v1/companies/${companyId}/visibility.json`, {
        method: "PATCH",
        body: JSON.stringify({
            visibility: true,
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

function overrideIndustryVisibility(industryAlias: string, callback: () => void) {
    fetch(`/api/v1/industries/${industryAlias}/visibility.json`, {
        method: "PATCH",
        body: JSON.stringify({
            visibility: true,
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

function addOverrideCompanyVisibilityEvent(
    $currentCompanyHiddenPlaceholder: HTMLElement,
    $currentCompany: HTMLElement,
    currentCompany: CompanyResponse,
) {
    const $button = $currentCompanyHiddenPlaceholder.querySelector(".js-override-company-visibility-button");
    if ($button === null) {
        return;
    }

    $button
        .addEventListener("click", function () {
            overrideCompanyVisibility(currentCompany.id, function () {
                $currentCompanyHiddenPlaceholder.hidden = true;
                $currentCompany.hidden = false;
                currentCompany.hidden = false;
            });
        });
}

function showAllCompanies(sourceCompanies: Array<CompanyResponse>, currentCompanyFirstIndustry: Industry) {
    sourceCompanies.forEach(function (sourceCompany) {
        if (!sourceCompany.hidden) {
            return;
        }

        if (companyEmptyIndustries(sourceCompany)) {
            return;
        }

        for (const sourceCompanyIndustry of sourceCompany.industries) {
            if (sourceCompanyIndustry.alias === currentCompanyFirstIndustry.alias) {
                sourceCompany.hidden = false;

                return;
            }
        }
    });
}

function addOverrideIndustryVisibilityEvent(
    $currentCompanyHiddenPlaceholder: HTMLElement,
    $currentCompany: HTMLElement,
    currentCompany: CompanyResponse,
    sourceCompanies: Array<CompanyResponse>,
) {
    const $button = $currentCompanyHiddenPlaceholder.querySelector(".js-override-industry-visibility-button");
    if ($button === null) {
        return
    }

    $button
        .addEventListener("click", function () {
            const currentCompanyFirstIndustry = companyFirstIndustry(currentCompany);
            if (currentCompanyFirstIndustry === null) {
                return;
            }

            overrideIndustryVisibility(currentCompanyFirstIndustry.alias, function () {
                $currentCompanyHiddenPlaceholder.hidden = true;
                $currentCompany.hidden = false;
                currentCompany.hidden = false;
                showAllCompanies(sourceCompanies, currentCompanyFirstIndustry);
            });
        });
}

export function addCompanyOverrideVisibilityEvents(
    $currentCompanyHiddenPlaceholder: HTMLElement,
    $currentCompany: HTMLElement,
    currentCompany: CompanyResponse,
    sourceCompanies: Array<CompanyResponse>,
) {
    addOverrideCompanyVisibilityEvent($currentCompanyHiddenPlaceholder, $currentCompany, currentCompany);
    addOverrideIndustryVisibilityEvent($currentCompanyHiddenPlaceholder, $currentCompany, currentCompany, sourceCompanies);
}
