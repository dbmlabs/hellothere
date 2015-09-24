var dashboard2 = (function () {

    "use strict";

    // Currently selected values in the dashboard
    var selectedYear = "2010";

    /* Functions to create the individual charts involved in the dashboard */

    function createLineChart(selector) {
        $(selector).dxChart({
            dataSource: summary2,
            animation: {
                duration: 350
            },
            commonSeriesSettings: {
                argumentField: "year"
            },
            series: [
                { valueField: "vRA, integration", name: "vRA, integration" },
                { valueField: "NSX, upgrades", name: "NSX, upgrades" },
                { valueField: "SDDC, integration", name: "SDDC, integration" },
                { valueField: "vCenter, SSO", name: "vCenter, SSO" },
                { valueField: "VSAN, performance", name: "VSAN, performance" }
            ],
            argumentAxis: {
                grid: {
                    visible: true
                }
            },
            tooltip: {
                enabled: true
            },
            title: {
                text: "Top 5 What's Not Working Terms",
                font: {
                    size: "24px"
                }
            },
            legend: {
                verticalAlignment: "bottom",
                horizontalAlignment: "center"
            },
            commonPaneSettings: {
                border: {
                    visible: true,
                    right: false
                }
            },
            pointClick: function(clickedPoint, clickEvent){
                selectedYear = clickedPoint.argument.substr(0,4);
                updatePieChart();
            }
        });
    }

    function createPieChart(selector) {
        $(selector).dxPieChart({
            dataSource: results2[selectedYear],
            animation: {
                duration: 350
            },
            title: {
                text: "What's Not Working Terms",
                font: {
                    size: "24px"
                }
            },
            legend: {
                horizontalAlignment: "left",
                verticalAlignment: "bottom",
                margin: 0
            },
            series: [
                {
                    type: "doughnut",
                    argumentField: "Country",
                    valueField: "Total",
                    label: {
                        visible: true,
                        connector: {
                            visible: true
                        }
                    }
                }
            ]
        });
    }

    /* Functions to update individual charts when their underlying dataset changes */

    function updatePieChart() {
        var chart2 = $("#chart2").dxPieChart("instance");
        chart2.option({
            dataSource: results2[selectedYear],
            title: "Total Medals by Country in " + selectedYear
        });
    }

    /* Render the dashboard */

    function render() {

        var html =
            '<div id="chart1" class="chart2"></div>' +
            '<div id="chart2" class="chart2"></div>';

        $("#content").html(html);

        createLineChart('#chart1');
        createPieChart('#chart2');

    }

    /* Functions to transform/format the data as required by specific charts */

    function getTypeByCountry(year, countryIdx) {
        var item = results2[year][countryIdx];
        return [
            {'type': 'Gold', 'count': item.Gold},
            {'type': 'Silver', 'count': item.Silver},
            {'type': 'Bronze', 'count': item.Bronze}
        ];
    }

    return {
        render: render
    }

}());
