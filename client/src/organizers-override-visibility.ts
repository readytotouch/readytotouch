import {organizersWelcome} from "./welcome";
import {CompanyHidden} from "./organizers-companies-v3-models";
import {companyEmptyIndustries, companyFirstIndustry, CompanyHidden, Industry} from "./organizers-v3-common-models";

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

function showAllCompanies(sourceCompanies: Array<CompanyHidden>, currentCompanyFirstIndustry: Industry) {
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

function addOverrideCompanyVisibilityEvent(
    $currentHiddenPlaceholder: HTMLElement,
    $company: HTMLElement,
    currentCompany: CompanyHidden,
) {
    const $button = $currentHiddenPlaceholder.querySelector(".js-override-company-visibility-button");
    if ($button === null) {
        return;
    }

    $button
        .addEventListener("click", function () {
            overrideCompanyVisibility(currentCompany.id, function () {
                $currentHiddenPlaceholder.hidden = true;
                $company.hidden = false;
                currentCompany.hidden = false;
            });
        });
}

function addOverrideIndustryVisibilityEvent(
    $currentHiddenPlaceholder: HTMLElement,
    $current: HTMLElement,
    currentCompany: CompanyHidden,
    sourceCompanies: Array<CompanyHidden>,
) {
    const $button = $currentHiddenPlaceholder.querySelector(".js-override-industry-visibility-button");
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
                $currentHiddenPlaceholder.hidden = true;
                $current.hidden = false;
                currentCompany.hidden = false;
                showAllCompanies(sourceCompanies, currentCompanyFirstIndustry);
            });
        });
}

export function renderHiddenPlaceholder(company: CompanyHidden): string {
    // Should be unreachable, but just in case
    const industry = companyFirstIndustry(company);
    if (industry === null) {
        return "";
    }

    return `<div class="card hidden-card">
    <figure class="hidden-card__body">
        <img
            alt="Lock icon"
            width="28"
            height="36"
            src="/assets/images/pages/organizer/lock.svg"
            class="hidden-card__image"
        />
        <figcaption class="hidden-card__caption">
            <p class="hidden-card__text">Company hidden due to ${industry.name} industry</p>
            <div class="hidden-card__action">
                <button class="button button--black button--small button--x-small-padding hidden-card__action-button js-override-industry-visibility-button" type="button">Show all companies in this industry</button>
                <button class="button button--small button--x-small-padding button--bordered-black-transparent hidden-card__action-button js-override-company-visibility-button">Show just this company</button>
            </div>
        </figcaption>
    </figure>
</div>`;
}

/**
 * Add events to override company and industry visibility.
 *
 * Using for both lists: companies and vacancies.
 */
export function addOverrideVisibilityEvents(
    $currentHiddenPlaceholder: HTMLElement,
    $current: HTMLElement,
    currentCompany: CompanyHidden,
    sourceCompanies: Array<CompanyHidden>,
) {
    addOverrideCompanyVisibilityEvent($currentHiddenPlaceholder, $current, currentCompany);
    addOverrideIndustryVisibilityEvent($currentHiddenPlaceholder, $current, currentCompany, sourceCompanies);
}
