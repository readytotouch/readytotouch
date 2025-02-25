import UniqueName from "./unique_name";

const unique = new UniqueName();

export const COMPANY_SEARCH_QUERY = unique.unique("q");
export const COMPANY_TYPE_CRITERIA_NAME = unique.unique("type");
export const COMPANY_INDUSTRY_CRITERIA_NAME = unique.unique("industry");
export const COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME = unique.unique("rust-foundation-members");
export const COMPANY_REMOTE_CRITERIA_NAME = unique.unique("remote");
export const COMPANY_IN_FAVORITES_CRITERIA_NAME = unique.unique("in-favorites");
export const COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME = unique.unique("has-employees-from-country");
export const COMPANY_CRITERIA_NAMES = unique.names();
