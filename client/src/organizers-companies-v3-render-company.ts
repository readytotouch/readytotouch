import {CompanyResponse, Industry} from "./organizers-companies-v3-models";

export function renderCompany(company: CompanyResponse, favorite: boolean): string {
    const companyURL = `/golang/companies/${company.linkedin_profile.alias}`;

    return `<div class="js-company card" data-company-id="{%dl company.ID %}">
    <aside class="card__action">
        ${renderFavorite(favorite)}
        <a href="${companyURL}" class="button-group__item" title="View statistics">
            <img width="20" height="20" alt="icon stats" src="/assets/images/pages/common/stats.svg" />
        </a>
    </aside>
    
    <figure class="card__header">
        <div class="card__logo-overlay">
            ${logo72x72(company)}
        </div>
        <figcaption class="card__header-caption">
            <a href="${companyURL}" class="card__headline vacancy__link">${company.name}</a>
            ${renderLocation(company)}
            ${renderLatestVacancyDate(company)}
        </figcaption>
    </figure>
    
    <div class="card__top-links">
        ${renderTopLink("Website", company.base_url)}
        ${renderTopLink("Careers", company.careers_url)}
        ${renderTopLink("About", company.about_url)}
        ${renderTopLink("Dev Blog", company.blog_url)}
    </div>
    
    <div class="card__info">
        <figure class="card__figure">
            <img
                class="card__icon"
                alt="card type icon"
                width="16"
                height="16"
                src="/assets/images/pages/vacancy/building.svg"
            />
            <figcaption class="card__figcaption">${companyTypeNameMap[company.type]}</figcaption>
        </figure>
        <figure class="card__figure">
            <img
                class="card__icon"
                alt="card type icon"
                width="16"
                height="16"
                src="/assets/images/pages/vacancy/company-type.svg"
            />
            <figcaption class="card__figcaption">${industryNames(company.industries)}</figcaption>
        </figure>
        ${renderCloudProviders(company)}
    </div>
</div>`;
}

function renderFavorite(favorite: boolean): string {
    const star = `<svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
    <path
        d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
    />
</svg>`;

    const defaultButton = `<button class="js-company-favorite favorite card__action-button button-group__item" title="Add to favorite">
    ${star}
</button>`;

    const favoriteButton = `<button class="js-company-favorite favorite card__action-button button-group__item in-favorite" title="Remove from favorites">
    ${star}
</button>`;

    return favorite ? favoriteButton : defaultButton;
}

function renderLocation(company: CompanyResponse): string {
    return "";

    // @TODO: implement location rendering when location data is available
    /*
        <p class="card__header-caption-text">
            <img
            class="card__header-caption-icon"
            alt="marker"
            width="12"
            height="12"
            src="/assets/images/pages/organizer/marker.svg"
            />Country, City (Full Remote)
        </p>
    */
}

const defaultDateFormatter = new Intl.DateTimeFormat("en-CA", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
});

const prettyDateFormatter = new Intl.DateTimeFormat("en-GB", {
    day: "2-digit",
    month: "short",
    year: "numeric",
});

function renderLatestVacancyDate(company: CompanyResponse): string {
    if (company.latest_vacancy_date === null) {
        return "";
    }

    const date = new Date(company.latest_vacancy_date);

    return `<p class="card__header-caption-text">
    <img
        class="card__header-caption-icon"
        alt="briefcase"
        width="12"
        height="12"
        src="/assets/images/pages/organizer/briefcase.svg"
    />Latest job posting: <time datetime="${defaultDateFormatter.format(date)}">${prettyDateFormatter.format(date)}</time>
</p>`;
}

function logo72x72(company: CompanyResponse): string {
    const logo = company.logo["72x72"];
    if (logo === "") {
        return "";
    }

    return `<img
    class="card__logo"
    alt="${company.name} logo"
    src="/assets/unstable/logos/72x72/${logo}"
/>`;
}

function renderTopLink(name: string, url: string): string {
    if (url === "") {
        return `<span class="card__top-link button-link disabled">${name}</span>`;
    }

    return `<a href="${url}" target="_blank" class="card__top-link button-link">${name}</a>`;
}

const companyTypeNameMap: Record<string, string> = {
    "product": "Product",
    "startup": "Startup",
}

function industryNames(industries: Industry[]): string {
    if (industries.length === 0) {
        return "Unknown industry";
    }

    return industries
        .map(function (industry) {
            return industry.name;
        })
        .join(", ");
}

function renderCloudProviders(company: CompanyResponse): string {
    return `<figure class="card__figure">
    <img
        class="card__icon"
        alt="card hosting icon"
        width="16"
        height="16"
        src="/assets/images/pages/organizer/clouds.svg"
    />
    <figcaption class="card__figcaption card__figcaption--images">Unknown cloud</figcaption>
</figure>`;

    // @TODO: implement cloud providers rendering when data is available
}
