import {TimeCountStats} from "./domain";
import {formatDay} from "./format";
import apexchartsCommonOptions from "./apexcharts-common-options";

fetch(`/api/v1${window.location.pathname}/waitlist/stats.json`)
    .then(function (response) {
        return response.json();
    })
    .then(renderWaitlistStats)
    .catch(console.error);

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
