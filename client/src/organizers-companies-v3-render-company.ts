import {CompanyResponse, Industry} from "./organizers-companies-v3-models";

const organizerGitHubAlias = "go";

export function renderCompany(company: CompanyResponse, favorite: boolean): string {
    const companyURL = `/golang/companies/${company.linkedin_profile.alias}`;
    const linkedinURL = `https://www.linkedin.com/company/${company.linkedin_profile.alias}/`;

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

    <div class="card__info-group">
        <button type="button" class="button button--small button--black card__info-group-button js-mobile-company-show-more-info-button">
            <span class="card__info-group-button-text js-toggle-text">Show more info</span>
            <img
                class="card__info-group-button-image"
                src="/assets/images/pages/common/double-arrow-down.svg"
                width="13"
                height="14"
                alt="icon double-arrow-down"
            />
        </button>
        <div class="card__info-group-content">
            <p class="card__text">${company.short_description}</p>
            <div class="card__links">
                <ul class="card__links-group">
                    <li class="card__links-item card__links-item--title">
                        <div class="card__links-item-group">
                            <img
                                class="card__links-icon"
                                alt="linkedin icon"
                                width="32"
                                height="32"
                                src="/assets/images/pages/organizer/linkedin.svg"
                            />
                            <a href="${linkedinURL}" target="_blank" class="card__links-link">LinkedIn</a>
                            ${renderLinkedInVerified(company)}
                        </div>
                    </li>
                    <li class="card__links-item">
                        <a href="${linkedinURL}" target="_blank" class="button-link card__links-link">Overview</a>
                    </li>
                    <li class="card__links-item card__links-item--inner">
                        <p class="card__links-item-text">Connections (employees):</p>
                        <ul class="card__links-group-inner">
                            <li class="card__links-item">
                                <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">Global</a>
                                <img
                                    class="card__links-icon"
                                    alt="language icon"
                                    width="20"
                                    height="20"
                                    src="/assets/images/pages/organizer/language.svg"
                                />
                            </li>
                            <li class="card__links-item">
                                <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">UA</a>
                                <img
                                    class="card__links-icon"
                                    alt="language icon"
                                    width="20"
                                    height="20"
                                    src="/assets/images/pages/common/flags/4x3/ua.svg"
                                />
                            </li>
                            <li class="card__links-item">
                                <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">CZ</a>
                                <img
                                    class="card__links-icon"
                                    alt="language icon"
                                    width="20"
                                    height="20"
                                    src="/assets/images/pages/common/flags/4x3/cz.svg"
                                />
                            </li>
                            <li class="card__links-item">
                                <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">Former (All)</a>
                            </li>
                        </ul>
                    </li>
                    <li class="card__links-item">
                        <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">Employees' posts</a>
                    </li>
                    <li class="card__links-item">
                        <a href="javascript:void(0);" target="_blank" class="button-link card__links-link">Jobs</a>
                    </li>
                </ul>
                ${renderGitHub(company)}
                <ul class="card__links-group">
                {% if company.GlassdoorProfile.OverviewURL == "" %}
                <li class="card__links-item card__links-item--title card__links-item--disabled">
                  <div class="card__links-item-group">
                    <img
                      class="card__links-icon"
                      alt="Glassdoor icon"
                      width="32"
                      height="32"
                      src="/assets/images/pages/organizer/glassdoor.svg"
                    />
                    <span class="card__links-link">Glassdoor</span>
                  </div>
                </li>
                <li class="card__links-item card__links-item--disabled">
                  <span class="button-link card__links-link">Overview</span>
                  <a href="{%s googleSearchGlassdoor(company.Name) %}" target="_blank" class="card__links-link card__links-link--google">
                    <img class="card__links-icon--google" alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg">
                  </a>
                </li>
                <li class="card__links-item card__links-item--disabled">
                  <span class="button-link card__links-link">Reviews</span>
                  <span class="card__links-link-star">?.? ★</span>
                </li>
                {% else %}
                <li class="card__links-item card__links-item--title">
                  <div class="card__links-item-group">
                    <img
                      class="card__links-icon"
                      alt="Glassdoor icon"
                      width="32"
                      height="32"
                      src="/assets/images/pages/organizer/glassdoor.svg"
                    />
                    <a href="{%s company.GlassdoorProfile.OverviewURL %}" target="_blank" class="card__links-link">Glassdoor</a>
                    {% if company.GlassdoorProfile.Verified %}
                    <span class="card__links-link card__links-link--verify">
                      <img
                        class="card__links-icon"
                        alt="Glassdoor verified icon"
                        src="/assets/images/pages/organizer/verified-icon-2.png"
                      />
                    </span>
                    {% endif %}
                  </div>
                </li>
                <li class="card__links-item">
                  <a href="{%s company.GlassdoorProfile.OverviewURL %}" target="_blank" class="button-link card__links-link">Overview</a>
                </li>
                <li class="card__links-item">
                  <a href="{%s company.GlassdoorProfile.ReviewsURL %}" target="_blank" class="button-link card__links-link">Reviews</a>
                  <span class="card__links-link-star">{%s formatGlassdoorReviewsRate(company.GlassdoorProfile.ReviewsRate) %} ★</span>
                </li>
                {% endif %}
                </ul>
            </div>
        </div>
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

function renderLinkedInVerified(company: CompanyResponse): string {
    if (company.linkedin_profile.verified === false) {
        return "";
    }

    return `<a href="https://www.linkedin.com/company/${company.linkedin_profile.alias}/about" target="_blank" class="card__links-link card__links-link--verify">
    <img
        class="card__links-icon"
        alt="LinkedIn verified icon"
        src="/assets/images/pages/organizer/verified-icon.png"
    />
</a>`;
}

function renderGitHub(company: CompanyResponse): string {
    if (company.github_profile.login === "") {
        return renderGitHubEmpty(company);
    }

    const githubProfileURL = `https://github.com/${company.github_profile.login}`;

    return `<ul class="card__links-group">
    <li class="card__links-item card__links-item--title">
        <div class="card__links-item-group">
            <img
                class="card__links-icon"
                alt="GitHub icon"
                width="32"
                height="32"
                src="/assets/images/pages/organizer/github.svg"
            />
            <a href="${githubProfileURL}" target="_blank" class="card__links-link">GitHub</a>
            ${renderGitHubVerified(company)}
        </div>
    </li>
    <li class="card__links-item">
        <a href="${githubProfileURL}" target="_blank" class="button-link card__links-link">Overview</a>
    </li>
    <li class="card__links-item">
        <a href="https://github.com/orgs/${company.github_profile.login}/repositories?q=lang:${organizerGitHubAlias}" target="_blank" class="button-link card__links-link">Repositories ({%d fetchGitHubRepositoriesCount(company, organizerFeature.Organizer.Language) %})</a>
    </li>
    <li class="card__links-item">
        <a href="https://github.com/orgs/${company.github_profile.login}/followers" target="_blank" class="button-link card__links-link">Followers ({%s fetchGitHubFollowers(company) %})</a>
    </li>
</ul>`;
}

function renderGitHubEmpty(company: CompanyResponse): string {
    return `<ul class="card__links-group">
    <li class="card__links-item card__links-item--title card__links-item--disabled">
        <div class="card__links-item-group">
            <img
                class="card__links-icon"
                alt="GitHub icon"
                width="32"
                height="32"
                src="/assets/images/pages/organizer/github.svg"
            />
            <span class="card__links-link">GitHub</span>
        </div>
    </li>
    <li class="card__links-item card__links-item--disabled">
        <span class="button-link card__links-link">Overview</span>
        <a href="${googleSearchGitHub(company.name)}" target="_blank" class="card__links-link card__links-link--google">
            <img class="card__links-icon--google" alt="google icon" width="20" height="20" src="/assets/images/pages/organizer/google.svg">
        </a>
    </li>
    <li class="card__links-item card__links-item--disabled">
        <span class="button-link card__links-link">Repositories (?)</span>
    </li>
    <li class="card__links-item card__links-item--disabled">
        <span class="button-link card__links-link">Followers (?)</span>
    </li>
</ul>`;
}

function renderGitHubVerified(company: CompanyResponse): string {
    if (company.github_profile.verified === false) {
        return "";
    }

    return `<span class="card__links-link card__links-link--verify">
    <img
        class="card__links-icon"
        alt="GitHub verified icon"
        src="/assets/images/pages/organizer/verified.png"
    />
</span>`;
}

function renderGlassdoor(company: CompanyResponse): string {
    if (company.glassdoor_profile.overview_url === "") {
        return renderGlassdoorEmpty(company);
    }

    return ``;
}

function renderGlassdoorEmpty(company: CompanyResponse): string {
    return ``;
}

function renderGlassdoorVerified(company: CompanyResponse): string {
    return ``;
}

function googleSearchGitHub(companyName: string): string {
    const params = new URLSearchParams({
        q: `site:github.com ${companyName}`,
    });

    return `https://www.google.com/search?${params.toString()}`;
}

function googleSearchGlassdoor(companyName: string): string {
    const params = new URLSearchParams({
        q: `site:glassdoor.com ${companyName}`,
    });

    return `https://www.google.com/search?${params.toString()}`;
}
