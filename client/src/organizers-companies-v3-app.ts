import {parseCurrentProgrammingLanguage} from "./pl";
import {organizersWelcome} from "./welcome";

const currentProgrammingLanguage = parseCurrentProgrammingLanguage(window.location.pathname);

function fetchCompanies(callback: () => void) {
    fetch(`/api/v1/unsafe/${currentProgrammingLanguage}/companies.json`, {
        method: "GET",
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        callback();
    }).catch(console.error);
}

