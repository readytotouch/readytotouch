const apexchartsCommonOptions = {
    chart: {
        type: 'area',
        zoom: {enabled: false},
        toolbar: {show: false},
        parentHeightOffset: null,
    },
    stroke: {
        width: 2,
        curve: 'straight',
    },
    fill: {
        type: 'gradient',
        gradient: {
            shadeIntensity: 1,
            opacityFrom: 0.65,
            opacityTo: 0.3,
        },
    },
    dataLabels: {enabled: false},
    xaxis: {
        type: 'datey',
        axisTicks: {
            show: false,
        },
        labels: {
            style: {cssClass: 'stats__labels'},
        },
        axisBorder: {
            color: '#efeff5',
        },
    },
    yaxis: {
        labels: {
            style: {cssClass: 'stats__labels'},
        },
    },
    grid: {
        borderColor: '#efeff5',
    },
    markers: {
        size: 4,
        colors: ['#fff'],
        strokeWidth: 3,
        strokeOpacity: 0.9,
        strokeDasharray: 1,
        fillOpacity: 1,
        shape: 'circle',
        radius: 2,
    },
};

export default apexchartsCommonOptions;
