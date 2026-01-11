import {findOrganizer, parseCurrentOrganizerAlias} from "./organizer";
import {VacancyResponse} from "./organizers-vacancies-v3-models";

const currentOrganizer = findOrganizer(parseCurrentOrganizerAlias(window.location.pathname));

export function renderVacancy(vacancy: VacancyResponse): string {
    return ``;
}
