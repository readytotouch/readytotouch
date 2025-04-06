function fetchVacancyList(success) {
    return fetch("https://readytotouch.com/api/v1/unsafe/vacancies.json")
        .then(response => response.json())
        .then(data => {
            success(data.vacancies);
        })
        .catch(error => {
            console.error("Error fetching vacancy list:", error);
        });
}

function highlightVacancyName() {
    const $vacancyName = document.querySelector("h1");

    if ($vacancyName === null) {
        console.error("Vacancy name element not found.");

        return;
    }

    fetchVacancyList(function (vacancies) {
        const url = window.location.origin + window.location.pathname;

        for (const vacancy of vacancies) {
            if (vacancy.url === url) {
                $vacancyName.style.color = "#28a745";

                return;
            }
        }
    })
}

highlightVacancyName();
