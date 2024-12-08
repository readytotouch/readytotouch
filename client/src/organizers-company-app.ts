import {organizersWelcome} from "./welcome";
import {TimeCountStats} from "./domain";
import apexchartsCommonOptions from "./apexcharts-common-options";
import {formatDay} from "./format";

function markCompanyFavorite(companyId: number, favorite: boolean, callback: () => void) {
    fetch(`/api/v1/companies/${companyId}/favorite.json`, {
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

const $companies = document.querySelectorAll(".js-company");

$companies.forEach(function ($company: HTMLElement) {
    const companyId = parseInt($company.getAttribute("data-company-id"));

    const $favorite = $company.querySelector(".js-company-favorite");
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

    fetch(`/api/v1/companies/${companyId}/views/stats/daily.json`)
        .then(function (response) {
            return response.json();
        })
        .then(renderViewsStats)
        .catch(console.error);
});

function renderViewsStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-views-statistics"), 'Views', "#5484FF", data);
}

function render($element: Element, name, color: string, data: Array<TimeCountStats>) {
    const options = {
        ...apexchartsCommonOptions,
        chart: {...apexchartsCommonOptions.chart, height: 254},
        labels: data.map(item => formatDay(item)),
        series: [{name: name, data: data.map(item => item.count),}],
        colors: [color],
        fill: {...apexchartsCommonOptions.fill, colors: [color]},
        markers: {...apexchartsCommonOptions.markers, strokeColors: color},
        yaxis: {...apexchartsCommonOptions.yaxis, stepSize: 10},
    };

    const chart = new ApexCharts($element, options);
    chart.render();
}

const $vacancies = document.querySelectorAll(".js-vacancy");

$vacancies.forEach(function ($vacancy: HTMLElement) {
    const vacancyId = parseInt($vacancy.getAttribute("data-vacancy-id"));

    const $favorite = $vacancy.querySelector(".js-vacancy-favorite");
    $favorite.addEventListener("click", function () {
        const current = $favorite.classList.contains("in-favorite");
        const next = !current;

        markVacancyFavorite(vacancyId, next, function () {
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
