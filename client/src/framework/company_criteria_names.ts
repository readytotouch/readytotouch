import UniqueName from "./unique_name";

const unique = new UniqueName();

export const COMPANY_SEARCH_QUERY = unique.unique("q");
export const COMPANY_TYPE_CRITERIA_NAME = unique.unique("type");
export const COMPANY_CRITERIA_NAMES = unique.names();
