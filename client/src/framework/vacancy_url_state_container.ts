import UrlStateContainer from "./url_state_container"
import {
    CheckedCriteriaConverter,
    IdentityCriteriaConverter,
    MultiSelectCriteriaConverter,
} from "./criteria_converter";
import {
    VACANCY_SEARCH_QUERY,
    VACANCY_COMPANY_TYPE_CRITERIA_NAME,
    VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME,
    VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME,
    VACANCY_REMOTE_CRITERIA_NAME,
    VACANCY_IN_FAVORITES_CRITERIA_NAME,
    VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME,
    VACANCY_CRITERIA_NAMES,
} from "./vacancy_criteria_names";

const multiSelectCriteriaConverter = new MultiSelectCriteriaConverter();
const identityCriteriaConverter = new IdentityCriteriaConverter();
const checkedCriteriaConverter = new CheckedCriteriaConverter();

const vacancyUrlStateContainer = new UrlStateContainer(VACANCY_CRITERIA_NAMES, {
    [VACANCY_SEARCH_QUERY]: identityCriteriaConverter,
    [VACANCY_COMPANY_TYPE_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [VACANCY_COMPANY_INDUSTRY_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [VACANCY_COMPANY_HAS_EMPLOYEES_FROM_COUNTRY_CRITERIA_NAME]: multiSelectCriteriaConverter,
    [VACANCY_COMPANY_RUST_FOUNDATION_MEMBERS_CRITERIA_NAME]: checkedCriteriaConverter,
    [VACANCY_REMOTE_CRITERIA_NAME]: checkedCriteriaConverter,
    [VACANCY_IN_FAVORITES_CRITERIA_NAME]: checkedCriteriaConverter,
});

export default vacancyUrlStateContainer;
