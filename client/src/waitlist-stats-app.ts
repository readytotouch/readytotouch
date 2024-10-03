import {TimeCountStats} from "./domain";
import {formatDay} from "./format";
import apexchartsCommonOptions from "./apexcharts-common-options";
import {organizersWelcome} from "./welcome";

fetch(`/api/v1/features/auto/waitlist/stats.json`)
    .then(function (response) {
        return response.json();
    })
    .then(renderWaitlistStats)
    .catch(console.error);

document.querySelector("a.js-feature-subscribe").addEventListener("click", function (event) {
    subscribe(true);
});

document.querySelector("a.js-feature-unsubscribe").addEventListener("click", function (event) {
    subscribe(false);
});

function subscribe(subscribed: boolean) {
    fetch(`/api/v1/features/auto/waitlist/subscribe.json`, {
        method: "POST",
        body: JSON.stringify({
            active: subscribed,
        }),
    }).then(function (response) {
        // Unauthorized
        if (response.status === 401) {
            window.location.href = organizersWelcome();

            return;
        }

        toggleSubscribeState(!subscribed);
    }).catch(console.error);
}

function toggleSubscribeState(subscribed: boolean) {
    document.querySelectorAll(".js-feature-unsubscribe").forEach(function ($element) {
        $element.classList.toggle("d-none", subscribed);
    });

    document.querySelectorAll(".js-feature-subscribe").forEach(function ($element) {
        $element.classList.toggle("d-none", !subscribed);
    });
}

function renderWaitlistStats(data) {
    renderViewsStats(data.views.daily_stats);
    renderSubscribersStats(data.subscribers.daily_stats);

    document.querySelector(".js-stats-total-subscribers").textContent = data.subscribers.total_count.toString();
    document.querySelector(".js-stats-last-month-subscribers").textContent = sum(data.subscribers.daily_stats).toString();
    document.querySelector(".js-stats-total-views").textContent = data.views.total_count.toString();
    document.querySelector(".js-stats-last-month-views").textContent = sum(data.views.daily_stats).toString();
}

function renderViewsStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-views-statistics"), 'Views', "#8D6DFF", data);
}

function renderSubscribersStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-subscribers-statistics"), 'Subscribers', "#5484FF", data);
}

function render($element: Element, name, color: string, data: Array<TimeCountStats>) {
    const options = {
        ...apexchartsCommonOptions,
        chart: {...apexchartsCommonOptions.chart, height: 340},
        labels: data.map(item => formatDay(item)),
        series: [{name: name, data: data.map(item => item.count),}],
        colors: [color],
        fill: {...apexchartsCommonOptions.fill, colors: [color]},
        markers: {...apexchartsCommonOptions.markers, strokeColors: color},
        yaxis: {...apexchartsCommonOptions.yaxis, stepSize: 20},
    };

    const chart = new ApexCharts($element, options);
    chart.render();
}

function sum(data: Array<TimeCountStats>) {
    return data.reduce((acc, item) => acc + item.count, 0);
}
