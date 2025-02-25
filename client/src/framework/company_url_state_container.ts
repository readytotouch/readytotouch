import UrlStateContainer from "./url_state_container"
import {
    CheckedCriteriaConverter,
    IdentityCriteriaConverter,
    MultiSelectCriteriaConverter,
} from "./criteria_converter";
import {
    COMPANY_SEARCH_QUERY,
    COMPANY_TYPE_CRITERIA_NAME,
    COMPANY_INDUSTRY_CRITERIA_NAME,
    COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    COMPANY_REMOTE_CRITERIA_NAME,
    COMPANY_IN_FAVORITES_CRITERIA_NAME,
    COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
    COMPANY_CRITERIA_NAMES,
} from "./company_criteria_names";

const multiSelectCriteriaConverter = new MultiSelectCriteriaConverter();
const identityCriteriaConverter = new IdentityCriteriaConverter();
const checkedCriteriaConverter = new CheckedCriteriaConverter();

const companyUrlStateContainer = new UrlStateContainer(COMPANY_CRITERIA_NAMES, {
    [COMPANY_SEARCH_QUERY]: identityCriteriaConverter,
    [COMPANY_TYPE_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [COMPANY_INDUSTRY_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME]: checkedCriteriaConverter,
    [COMPANY_REMOTE_CRITERIA_NAME]: checkedCriteriaConverter,
    [COMPANY_IN_FAVORITES_CRITERIA_NAME]: checkedCriteriaConverter,
});

export default companyUrlStateContainer;
