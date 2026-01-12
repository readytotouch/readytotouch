import {VacancyCompanyResponse, VacancyResponse} from "./organizers-vacancies-v3-models";

export function renderVacancy(vacancy: VacancyResponse): string {
    return `<div class="card">
    <div class="card__vacancy">
        <aside class="card__action">
            ${renderFavorite(vacancy.favorite)}
        </aside>
        
        <figure class="card__header card__header--organizer">
            <div class="card__logo-overlay">
                ${logo72x72(vacancy.company)}
            </div>
            <figcaption class="card__header-caption">
                <a href="/v/${vacancy.id}" target="_blank" class="card__headline">${vacancy.title}</a>
            </figcaption>
        </figure>
    </div>
</div>`;
}

function renderFavorite(favorite: boolean): string {
    const star = `<svg xmlns="http://www.w3.org/2000/svg" xml:space="preserve" class="favorite__icon" viewBox="0 0 28 28">
    <path
        d="m14.5 22.1-.5-.3-.5.3-6.8 4.2c-.5.3-1.1-.1-.9-.7L7.5 18l.1-.6-.4-.4-5.9-5.2c-.3-.3-.3-.6-.2-.8.1-.2.3-.4.5-.4l7.9-.7.6-.1.2-.6 2.9-7.4c.2-.5 1-.5 1.2 0l3.1 7.3.2.5.6.1 7.9.7c.2 0 .4.2.5.5.1.3 0 .6-.2.7l-5.9 5.2-.4.4.1.6 1.8 7.7c.1.3 0 .5-.2.6-.2.1-.5.2-.8 0l-6.6-4z"
    />
</svg>`;

    const defaultButton = `<button class="js-vacancy-favorite favorite card__action-button button-group__item" title="Add to favorite">
    ${star}
</button>`;

    const favoriteButton = `<button class="js-vacancy-favorite favorite card__action-button button-group__item in-favorite" title="Remove from favorites">
    ${star}
</button>`;

    return favorite ? favoriteButton : defaultButton;
}

function logo72x72(company: VacancyCompanyResponse): string {
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
