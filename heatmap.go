package main

const (
	heatmapPage = `
<!DOCTYPE html>
<meta charset="utf-8">
<html>
  <head>
  <meta charset="utf-8">
      <title>vWeekly Reporter</title>
          <meta name="viewport" content="width=device-width, initial-scale=1">
      <style type="text/css">
      .entry:not(:first-of-type)
  {
      margin-top: 10px;
  }
  .glyphicon
  {
      font-size: 20px;
  }
      </style>
      <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
      <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
      <script src="http://www.appelsiini.net/download/jquery.jeditable.mini.js"></script>

      <style>



      #chart{
        width: 800px;
        margin: 0 auto;
      }
      .background {
        fill: #eee;
      }

      line {
        stroke: #fff;
      }

      text.active {
        fill: red;
      }

      .day {
        fill: #fff;
        stroke: #ccc;
      }

      .month {
        fill: none;
        stroke: #fff;
        stroke-width: 4px;
      }
      .year-title {
        font-size: 1.5em;
      }

      /* color ranges */
      .RdYlGn .q0-11{fill:rgb(165,0,38)}
      .RdYlGn .q1-11{fill:rgb(215,48,39)}
      .RdYlGn .q2-11{fill:rgb(244,109,67)}
      .RdYlGn .q3-11{fill:rgb(253,174,97)}
      .RdYlGn .q4-11{fill:rgb(254,224,139)}
      .RdYlGn .q5-11{fill:rgb(255,255,191)}
      .RdYlGn .q6-11{fill:rgb(217,239,139)}
      .RdYlGn .q7-11{fill:rgb(166,217,106)}
      .RdYlGn .q8-11{fill:rgb(102,189,99)}
      .RdYlGn .q9-11{fill:rgb(26,152,80)}
      .RdYlGn .q10-11{fill:rgb(0,104,55)}

      /* hover info */
      #tooltip {
        background-color: #fff;
        border: 2px solid #ccc;
        padding: 10px;
      }

    </style>
  </head>
    <body>
    <nav class="navbar navbar-default">
    <div class="container-fluid">
      <!-- Brand and toggle get grouped for better mobile display -->
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">VMware Weekly Reporter</a>
      </div>

      <!-- Collect the nav links, forms, and other content for toggling -->
      <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
        <ul class="nav navbar-nav">
          <li class="active"><a href="#">Help <span class="sr-only">(current)</span></a></li>
          <li><a href="#">About</a></li>
          <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Tools <span class="caret"></span></a>
            <ul class="dropdown-menu">
              <li><a href="#">Admin</a></li>
              <li><a href="#">Help</a></li>
              <li><a href="/dashboard">Analytics</a></li>
              <li><a href="#">Analytics 2</a></li>
              <li role="separator" class="divider"></li>
              <li><a href="#">Send to Socialcast</a></li>
            </ul>
          </li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><span class="sr-only">(current)</span></a></li>
          <li><a href="#"><b>September 30th, 2015</b>: Tim Callaghan</a></li>
          <li><img src="https://n2.cdn.socialcast.com/801245/socialcast.s3.amazonaws.com/tenants/5258/profile_photos/732278/tim_callaghan_square140.jpg?AWSAccessKeyId=AKIAISVYYXCGCXLJL2TQ&Expires=1445169600&Signature=TlZgwT7FtBEr2E1qCextWCpNMfc%3D" height=50px></li>
        </ul>

      </div><!-- /.navbar-collapse -->

    </div><!-- /.container-fluid -->

    </nav>
      <div id="chart" class="clearfix"></div>

  <script src="http://d3js.org/d3.v3.js"></script>
  <script>
    var width = 960,
        height = 750,
        cellSize = 25; // cell size

    var no_months_in_a_row = Math.floor(width / (cellSize * 7 + 50));
    var shift_up = cellSize * 3;

    var day = d3.time.format("%w"), // day of the week
        day_of_month = d3.time.format("%e") // day of the month
        day_of_year = d3.time.format("%j")
        week = d3.time.format("%U"), // week number of the year
        month = d3.time.format("%m"), // month number
        year = d3.time.format("%Y"),
        percent = d3.format(".1%"),
        format = d3.time.format("%Y-%m-%d");

    var color = d3.scale.quantize()
        .domain([-.05, .05])
        .range(d3.range(11).map(function(d) { return "q" + d + "-11"; }));

    var svg = d3.select("#chart").selectAll("svg")
        .data(d3.range(2014, 2016))
      .enter().append("svg")
        .attr("width", width)
        .attr("height", height)
        .attr("class", "RdYlGn")
      .append("g")

    var rect = svg.selectAll(".day")
        .data(function(d) {
          return d3.time.days(new Date(d, 0, 1), new Date(d + 1, 0, 1));
        })
      .enter().append("rect")
        .attr("class", "day")
        .attr("width", cellSize)
        .attr("height", cellSize)
        .attr("x", function(d) {
          var month_padding = 1.2 * cellSize*7 * ((month(d)-1) % (no_months_in_a_row));
          return day(d) * cellSize + month_padding;
        })
        .attr("y", function(d) {
          var week_diff = week(d) - week(new Date(year(d), month(d)-1, 1) );
          var row_level = Math.ceil(month(d) / (no_months_in_a_row));
          return (week_diff*cellSize) + row_level*cellSize*8 - cellSize/2 - shift_up;
        })
        .datum(format);

    var month_titles = svg.selectAll(".month-title")  // Jan, Feb, Mar and the whatnot
          .data(function(d) {
            return d3.time.months(new Date(d, 0, 1), new Date(d + 1, 0, 1)); })
        .enter().append("text")
          .text(monthTitle)
          .attr("x", function(d, i) {
            var month_padding = 1.2 * cellSize*7* ((month(d)-1) % (no_months_in_a_row));
            return month_padding;
          })
          .attr("y", function(d, i) {
            var week_diff = week(d) - week(new Date(year(d), month(d)-1, 1) );
            var row_level = Math.ceil(month(d) / (no_months_in_a_row));
            return (week_diff*cellSize) + row_level*cellSize*8 - cellSize - shift_up;
          })
          .attr("class", "month-title")
          .attr("d", monthTitle);

    var year_titles = svg.selectAll(".year-title")  // Jan, Feb, Mar and the whatnot
          .data(function(d) {
            return d3.time.years(new Date(d, 0, 1), new Date(d + 1, 0, 1)); })
        .enter().append("text")
          .text(yearTitle)
          .attr("x", function(d, i) { return width/2 - 100; })
          .attr("y", function(d, i) { return cellSize*5.5 - shift_up; })
          .attr("class", "year-title")
          .attr("d", yearTitle);


    //  Tooltip Object
    var tooltip = d3.select("body")
      .append("div").attr("id", "tooltip")
      .style("position", "absolute")
      .style("z-index", "10")
      .style("visibility", "hidden")
      .text("a simple tooltip");

    d3.csv("/scripts/dji.csv", function(error, csv) {
      var data = d3.nest()
        .key(function(d) { return d.Date; })
        .rollup(function(d) { return (d[0].Close - d[0].Open) / d[0].Open; })
        .map(csv);

      rect.filter(function(d) { return d in data; })
          .attr("class", function(d) { return "day " + color(data[d]); })
        .select("title")
          .text(function(d) { return d + ": " + percent(data[d]); });

      //  Tooltip
      rect.on("mouseover", mouseover);
      rect.on("mouseout", mouseout);
      function mouseover(d) {
        tooltip.style("visibility", "visible");
        var percent_data = (data[d] !== undefined) ? percent(data[d]) : percent(0);
        var purchase_text = d + ": " + percent_data;

        tooltip.transition()
                    .duration(200)
                    .style("opacity", .9);
        tooltip.html(purchase_text)
                    .style("left", (d3.event.pageX)+30 + "px")
                    .style("top", (d3.event.pageY) + "px");
      }
      function mouseout (d) {
        tooltip.transition()
                .duration(500)
                .style("opacity", 0);
        var $tooltip = $("#tooltip");
        $tooltip.empty();
      }

    });

    function dayTitle (t0) {
      return t0.toString().split(" ")[2];
    }
    function monthTitle (t0) {
      return t0.toLocaleString("en-us", { month: "long" });
    }
    function yearTitle (t0) {
      return t0.toString().split(" ")[3];
    }
  </script>

  </body>
</html>`
)
