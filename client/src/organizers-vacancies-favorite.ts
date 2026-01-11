import {organizersWelcome} from "./welcome";
import {VacancyResponse} from "./organizers-vacancies-v3-models";

function markVacancyFavorite(vacancyId: number, favorite: boolean, callback: () => void) {
    fetch(`/api/v1/vacancies/${vacancyId}/favorite.json`, {
        method: "PATCH",
        body: JSON.stringify({
            favorite: favorite,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        callback();
    }).catch(console.error);
}

export function addVacancyFavoriteEvent($vacancy: HTMLElement, vacancy: VacancyResponse) {
    const $favorite = $vacancy.querySelector(".js-vacancy-favorite");
    $favorite.addEventListener("click", function () {
        const current = $favorite.classList.contains("in-favorite");
        const next = !current;

        markVacancyFavorite(vacancy.id, next, function () {
            vacancy.favorite = next;

            if (next) {
                $favorite.classList.add("in-favorite");

                $favorite.setAttribute("title", "Remove from favorites");
            } else {
                $favorite.classList.remove("in-favorite");

                $favorite.setAttribute("title", "Add to favorites");
            }
        });
    });
}
