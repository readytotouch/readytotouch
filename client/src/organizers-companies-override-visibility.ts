import {organizersWelcome} from "./welcome";
import {CompanyResponse} from "./organizers-companies-v3-models";

function addOverrideCompanyVisibilityEvent(
    $companyHiddenPlaceholder: HTMLElement,
    $company: HTMLElement,
    company: CompanyResponse,
) {
    const $button = $companyHiddenPlaceholder.querySelector(".js-override-company-visibility-button");
    if ($button === null) {
        return;
    }

    $button
        .addEventListener("click", function () {
            $companyHiddenPlaceholder.hidden = true;
            $company.hidden = false;
            company.hidden = false;
        });
}

function addOverrideIndustryVisibilityEvent(
    $companyHiddenPlaceholder: HTMLElement,
    $company: HTMLElement,
    company: CompanyResponse,
    sourceCompanies: Array<CompanyResponse>,
) {
    const $button = $companyHiddenPlaceholder.querySelector(".js-override-industry-visibility-button");
    if ($button === null) {
        return
    }

    $button
        .addEventListener("click", function () {
            $companyHiddenPlaceholder.hidden = true;
            $company.hidden = false;
            company.hidden = false;
        });
}

export function addCompanyOverrideVisibilityEvents(
    $companyHiddenPlaceholder: HTMLElement,
    $company: HTMLElement,
    company: CompanyResponse,
    sourceCompanies: Array<CompanyResponse>,
) {
    addOverrideCompanyVisibilityEvent($companyHiddenPlaceholder, $company, company);
    addOverrideIndustryVisibilityEvent($companyHiddenPlaceholder, $company, company, sourceCompanies);
}
