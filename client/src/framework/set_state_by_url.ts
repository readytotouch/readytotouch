import UrlStateContainer from "./url_state_container";
import {Checkboxes} from "./checkboxes";

export function setStateByURLMapper(container: UrlStateContainer) {
    function setInputStateByURL($input: HTMLInputElement, criteriaName: string) {
        $input.value = container.getCriteria(criteriaName, "");
    }

    function setCheckboxesStateByURL($checkboxes: Checkboxes, criteriaName: string) {
        $checkboxes.setState(container.getCriteria(criteriaName, []));
    }

    function setCheckboxStateByURL($checkbox: HTMLInputElement, criteriaName: string) {
        $checkbox.checked = container.getCriteria(criteriaName, false);
    }

    return {
        setInputStateByURL,
        setCheckboxesStateByURL,
        setCheckboxStateByURL,
    };
}
