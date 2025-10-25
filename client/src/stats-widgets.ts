import {TimeCountStats} from "./domain";
import {formatDay} from "./format";
import apexchartsCommonOptions from "./apexcharts-common-options";

function connect(url: string, render: (data: Array<TimeCountStats>) => void) {
    fetch(url)
        .then(function (response) {
            return response.json();
        })
        .then(render)
        .catch(console.error);
}

function renderRegistrationStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-registration-statistics"), "#5484FF", data);
}

function renderOnlineStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-online-statistics"), "#8D6DFF", data);
}

function render($element: Element, color: string, data: Array<TimeCountStats>) {
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

export function dailyRegistrationStatsWidget() {
    connect("/api/v1/users/registration/stats/daily.json", renderRegistrationStats);
}

export function dailyOnlineStatsWidget() {
    connect("/api/v1/users/online/stats/daily.json", renderOnlineStats);
}
