import UniqueName from "./unique_name";

const unique = new UniqueName();

export const VACANCY_SEARCH_QUERY = unique.unique("q");
export const VACANCY_COMPANY_TYPE_CRITERIA_NAME = unique.unique("type");
export const VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME = unique.unique("industry");
export const VACANCY_IN_FAVORITES_CRITERIA_NAME = unique.unique("in-favorites");
export const VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME = unique.unique("has-employees-from-country");
export const VACANCY_CRITERIA_NAMES = unique.names();
