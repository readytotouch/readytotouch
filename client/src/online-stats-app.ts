import {TimeCountStats} from "./domain";

fetch("/api/v1/users/registration/stats/daily.json")
    .then(function (response) {
        return response.json();
    })
    .then(renderRegistrationStats)
    .catch(console.error);


fetch("/api/v1/users/online/stats/daily.json")
    .then(function (response) {
        return response.json();
    })
    .then(renderOnlineStats)
    .catch(console.error);

function renderRegistrationStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-registration-statistics"), "#5484FF", data);
}

function renderOnlineStats(data: Array<TimeCountStats>) {
    render(document.querySelector(".js-chart-online-statistics"), "#8D6DFF", data);
}

function render($element: Element, color: string, data: Array<TimeCountStats>) {
    const options = {
        series: [{
            name: "Users",
            data: data.map(item => item.count),
        }],
        chart: {
            type: "area",
            height: 340,
            zoom: {
                enabled: false,
            },
            toolbar: {
                show: false,
            },
        },
        colors: [color],
        fill: {
            type: "gradient",
            colors: [color],
            gradient: {
                shadeIntensity: 1,
                opacityFrom: 0.65,
                opacityTo: 0.3,
            },
        },
        dataLabels: {
            enabled: false,
        },
        stroke: {
            curve: "smooth",
        },
        labels: data.map(item => formatDay(item)),
        xaxis: {
            type: "datey",
        },
        yaxis: {
            labels: {
                style: {
                    colors: ["#575D65"],
                    fontSize: "14px",
                },
            },
        },
        markers: {
            size: 5,
            colors: ["#fff"],
            strokeColors: "#8D6DFF",
            strokeWidth: 2,
            strokeOpacity: 0.9,
            strokeDasharray: 1,
            fillOpacity: 1,
            shape: "circle",
            radius: 2,
        },
    };

    const chart = new ApexCharts($element, options);
    chart.render();
}

function formatDay(item: TimeCountStats) {
    return new Date(item.time).toLocaleDateString("en-us", {day: "2-digit", month: "long"});
}
