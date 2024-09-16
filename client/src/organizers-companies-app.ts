import {welcome} from "./welcome";

function markCompanyFavorite(companyId: number, favorite: boolean, callback: () => void) {
    fetch(`/api/v1/companies/${companyId}/favorite.json`, {
        method: "PATCH",
        body: JSON.stringify({
            favorite: favorite,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = welcome();

            return;
        }

        callback();
    }).catch(console.error);
}

document.querySelectorAll(".js-company").forEach(function ($element: HTMLElement) {
    const companyId = parseInt($element.getAttribute("data-company-id"));

    const $favorite = $element.querySelector(".js-company-favorite");
    $favorite.addEventListener("click", function () {
        const current = $favorite.classList.contains("in-favorite");
        const next = !current;

        markCompanyFavorite(companyId, next, function () {
            if (next) {
                $favorite.classList.add("in-favorite");

                $favorite.setAttribute("title", "Remove from favorites")
            } else {
                $favorite.classList.remove("in-favorite");

                $favorite.setAttribute("title", "Add to favorites")
            }
        });
    });
});

// <DEBUG>
/*
    function markCompanyFavoriteAll(favoriteGenerator: () => boolean) {
        document.querySelectorAll(".js-company").forEach(function ($element: HTMLElement) {
            markCompanyFavorite(
                parseInt($element.getAttribute("data-company-id")),
                favoriteGenerator(),
                function () {
                },
            );
        });
    }

    markCompanyFavoriteAll(function () {
        return Math.random() > 0.5;
    });
*/
// </DEBUG>
