package main

const (
	dashboardPage = `

	<html><head><meta http-equiv="Content-Type" content="text/html; charset=windows-1252">
	<script src="/scripts/scripts.js"></script>
	<link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="http://www.appelsiini.net/download/jquery.jeditable.mini.js"></script>


	<style type="text/css">

	.axis path,.axis line {fill: none;stroke:#b6b6b6;shape-rendering: crispEdges;}
	/*.tick line{fill:none;stroke:none;}*/
	.tick text{fill:#999;}
	g.journal.active{cursor:pointer;}
	text.label{font-size:12px;font-weight:bold;cursor:pointer;}
	text.value{font-size:12px;font-weight:bold;}
	</style>

	<style id="style-1-cropbar-clipper">/* Copyright 2014 Evernote Corporation. All rights reserved. */
	.en-markup-crop-options {
	    top: 18px !important;
	    left: 50% !important;
	    margin-left: -100px !important;
	    width: 200px !important;
	    border: 2px rgba(255,255,255,.38) solid !important;
	    border-radius: 4px !important;
	}

	.en-markup-crop-options div div:first-of-type {
	    margin-left: 0px !important;
	}
	</style></head>
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
						<li><a href="#">Heatmap</a></li>
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
	<body>

	<script type="text/javascript">
	function truncate(str, maxLength, suffix) {
		if(str.length > maxLength) {
			str = str.substring(0, maxLength + 1);
			str = str.substring(0, Math.min(str.length, str.lastIndexOf(" ")));
			str = str + suffix;
		}
		return str;
	}

	var margin = {top: 20, right: 200, bottom: 0, left: 20},
		width = 300,
		height = 650;

	var start_year = 2004,
		end_year = 2013;

	var c = d3.scale.category20c();

	var x = d3.scale.linear()
		.range([0, width]);

	var xAxis = d3.svg.axis()
		.scale(x)
		.orient("top");

	var formatYears = d3.format("0000");
	xAxis.tickFormat(formatYears);

	var svg = d3.select("body").append("svg")
		.attr("width", width + margin.left + margin.right)
		.attr("height", height + margin.top + margin.bottom)
		.style("margin-left", margin.left + "px")
		.append("g")
		.attr("transform", "translate(" + margin.left + "," + margin.top + ")");

	// var dataset = [[ [2002, 8], [2003, 1], [2004, 1], [2005, 1], [2006, 3], [2007, 3], [2009, 3], [2013, 3]], [ [2004, 5], [2005, 1], [2006, 2], [2010, 20], [2011, 3] ] ,[ [2001, 5], [2005, 15], [2006, 2], [2010, 20], [2012, 25] ]];
	// var dataset = [ [2001, 5], [2005, 15], [2006, 2], [2010, 20], [2012, 25] ];

	d3.json("./scripts/journals_optogenetic.json", function(data) {
		x.domain([start_year, end_year]);
		var xScale = d3.scale.linear()
			.domain([start_year, end_year])
			.range([0, width]);

		svg.append("g")
			.attr("class", "x axis")
			.attr("transform", "translate(0," + 0 + ")")
			.call(xAxis);

		for (var j = 0; j < data.length; j++) {
			var g = svg.append("g").attr("class","journal");

			var circles = g.selectAll("circle")
				.data(data[j]['articles'])
				.enter()
				.append("circle");

			var text = g.selectAll("text")
				.data(data[j]['articles'])
				.enter()
				.append("text");

			var rScale = d3.scale.linear()
				.domain([0, d3.max(data[j]['articles'], function(d) { return d[1]; })])
				.range([2, 9]);

			circles
				.attr("cx", function(d, i) { return xScale(d[0]); })
				.attr("cy", j*20+20)
				.attr("r", function(d) { return rScale(d[1]); })
				.style("fill", function(d) { return c(j); });

			text
				.attr("y", j*20+25)
				.attr("x",function(d, i) { return xScale(d[0])-5; })
				.attr("class","value")
				.text(function(d){ return d[1]; })
				.style("fill", function(d) { return c(j); })
				.style("display","none");

			g.append("text")
				.attr("y", j*20+25)
				.attr("x",width+20)
				.attr("class","label")
				.text(truncate(data[j]['name'],30,"..."))
				.style("fill", function(d) { return c(j); })
				.on("mouseover", mouseover)
				.on("mouseout", mouseout);
		};

		function mouseover(p) {
			var g = d3.select(this).node().parentNode;
			d3.select(g).selectAll("circle").style("display","none");
			d3.select(g).selectAll("text.value").style("display","block");
		}

		function mouseout(p) {
			var g = d3.select(this).node().parentNode;
			d3.select(g).selectAll("circle").style("display","block");
			d3.select(g).selectAll("text.value").style("display","none");
		}
	});

	</script>


	</body></html>
`
	loginPage = `
	<!DOCTYPE html>
	<html lang="en">
	  <head>
	    <meta charset="utf-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	    <meta name="viewport" content="width=device-width, initial-scale=1">
	    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
	    <meta name="description" content="">
	    <meta name="author" content="">
	    <link rel="icon" href="../../favicon.ico">

	    <title>vWeekly Reporter</title>

	    <!-- Bootstrap core CSS -->
	    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">

	    <!-- Custom styles for this template -->
	    <link href="http://getbootstrap.com/examples/signin/signin.css" rel="stylesheet">

	    <![endif]-->
	  </head>

	  <body>

	    <div class="container">

	      <form class="form-signin" action="/login" method="POST">
	        <h2 class="form-signin-heading">vWeekly Reporter Sign In</h2>
	        <label for="inputEmail" class="sr-only">Email address</label>
	        <input type="email" name="inputEmail" class="form-control" placeholder="Email address" required autofocus>
	        <label for="inputPassword" class="sr-only">Password</label>
	        <input type="password" name="inputPassword" class="form-control" placeholder="Password" required>
	        <div class="checkbox">
	          <label>
	            <input type="checkbox" value="remember-me"> Remember me
	          </label>
	        </div>
	        <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
	      </form>

	    </div> <!-- /container -->


	    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
	    <script src="../../assets/js/ie10-viewport-bug-workaround.js"></script>
	  </body>
	</html>
`
	page = `<!DOCTYPE html>
	<html lang="en">
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
	    <!--<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>-->
	    <script src="https://code.jquery.com/jquery-2.1.4.min.js"></script>
	    <script src="http://code.jquery.com/ui/1.11.4/jquery-ui.js"></script>
	    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	    <link rel="stylesheet" href="http://code.jquery.com/ui/1.11.4/themes/smoothness/jquery-ui.css">
	    <link href="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/css/bootstrap-editable.css" rel="stylesheet"/>
	    <script src="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/js/bootstrap-editable.min.js"></script>

	    <script>
	      $(function() {
	        var availableTags = [
	          "#vROps",
	          "#NSX",
	          "#vSphere6",
	          "#vRA",
	          "#BigData",
	          "#vCM",
	          "#Containers",
	          "#Docker",
	          "#vDS",
	          "#Dell",
	          "#EVORack",
	          "#vRAC",
	          "#vCA"
	        ];
	        $( "#tags" ).autocomplete({
	          source: availableTags
	        });
	      });
	      </script>

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
	        <li><a href="#"><b>September 30th, 2015</b>: Jonathan Cham</a></li>
					<li><img src="https://n2.cdn.socialcast.com/801245/socialcast.s3.amazonaws.com/tenants/5258/profile_photos/1750704/jc_square140.jpg?AWSAccessKeyId=AKIAISVYYXCGCXLJL2TQ&Expires=1445169600&Signature=OiWwWzyGsXa%2BcWpwbcQ2bG2XIwk%3D" height=50px></li>
	      </ul>

	    </div><!-- /.navbar-collapse -->

	  </div><!-- /.container-fluid -->

	</nav>
	  <div class="container">
	    <script>
	    $(document).ready(function() {
	    var max_fields      = 10; //maximum input boxes allowed
	    var wrapper         = $(".group"); //Fields wrapper
	    var add_button      = $(".btn-add"); //Add button ID

	    var x = 1; //initlal text box count
	    $(add_button).click(function(e){ //on add input button click
	        e.preventDefault();
	        if(x < max_fields){ //max input box allowed
	          $(wrapper).append('<div><br><div class="pull-right"><a href="#" class="remove_field"><button class="btn btn-danger btn-add" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><div class="entry input-group col-xs-10"><input class="form-control" name="fields[]" type="text" placeholder="Subject" /></div><div class="entry input-group col-xs-12"><textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea></div><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Hashtag" /></div><fieldset class="radiogroup"><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+' id="inlineRadio1" value="good">good</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio2" value="bad">bad</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio3" value="competitive">competitive</label></fieldset></div>'); //add input box
	            x++; //text box increment
	        }
	    });

	    $(wrapper).on("click",".remove_field", function(e){ //user click on remove text
	        e.preventDefault(); $(this).parent('div').parent('div').remove(); x--;
	    })
	    });
	    </script>

	  <div class="row">
	        <div class="control-group" id="fields">
	            <label class="control-label" for="field1">Weekly Report</label>

	                <div class="row">
	                  <div class="col-md-6">
	                    <form action="/jcham" role="form" autocomplete="off" >
	                      <div class="group">
	                        <div class="entry input-group col-xs-10">
	                          <input class="form-control" name="fields[]" type="text" placeholder="Subject" required/>
	                        </div>
	                        <div class="entry input-group col-xs-12">
	                          <textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3" required></textarea>
	                        </div>
	                        <div class="entry input-group col-xs-12">
	                          <input class="form-control" id="tags" name="fields[]" type="text" placeholder="Hashtag" required />
	                        </div>
	                        <fieldset class="radiogroup">
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions0" id="inlineRadio1" value="good"> good
	                        </label>
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions0" id="inlineRadio2" value="bad"> bad
	                        </label>
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions0" id="inlineRadio3" value="competitive"> competitive
	                        </label>
	                      </fieldset>
	                      </div> <!-- end of div class group -->
	                    <br>
	                      <div class="btn-toolbar" role="toolbar" aria-label="...">
	                        <div class="btn-group" role="group" aria-label="...">
	                          <button class="btn btn-success btn-add" type="button">
	                            <span class="glyphicon glyphicon-plus"></span>
	                          </button>
	                        </div>
	                        <div class="btn-group" role="group" aria-label="...">
	                          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	                        </div>
	                        <div class="btn-group" role="group" aria-label="...">
	                          <input type="submit" class="btn btn-default btn-primary" value="Submit">
	                        </div><!--end btn-group-->
	                      </div> <!-- end of button toolbar-->
	                    </form><!--end of form -->
	                    </div> <!--end of col -->

	                    <div class="col-md-2">
	                      <br>
	                    </div>
	                    <div class="col-md-4">
	                      <form class="form-inline" action="/" method="POST">
	                        <div class="form-group">
	                          <label for="inputPassword2" class="sr-only">search terms</label>
	                          <input type="text" class="form-control" name="search" placeholder="search terms">
	                        </div>
	                      <button type="submit" class="btn btn-default">Search</button>
	                      </form>
	                    </div>
	                  </div> <!--end of row-->
	                </div>
	                <br>



	        </div> <!-- end of control-group -->

	  </div> <!-- end of row -->
	</div> <!-- end of container -->

	</body>
	</html>

`
	adminPage = `
<!DOCTYPE html>
<html lang="en">
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



    <script>
    $(document).ready(function() {
     $('.edit').editable('http://localhost:9001/api', {
         type      : 'textarea',
         cancel    : 'Cancel',
         submit    : 'OK',
         indicator : '<img src="img/indicator.gif">',
         tooltip   : 'Click to edit...',
      });
 });</script>
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

  <div class="container">
    <script>
    $(document).ready(function() {
    var max_fields      = 10; //maximum input boxes allowed
    var wrapper         = $(".group"); //Fields wrapper
    var add_button      = $(".btn-add"); //Add button ID

    var x = 0; //initlal text box count
    $(add_button).click(function(e){ //on add input button click
        e.preventDefault();
        if(x < max_fields){ //max input box allowed
          $(wrapper).append('<div><br><div class="pull-right"><a href="#" class="remove_field"><button class="btn btn-danger btn-add" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><div class="entry input-group col-xs-10"><input class="form-control" name="fields[]" type="text" placeholder="Subject" /></div><div class="entry input-group col-xs-12"><textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea></div><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Hashtag" /></div><fieldset class="radiogroup"><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+' id="inlineRadio1" value="good">good</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio2" value="bad">bad</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio3" value="competitive">competitive</label></fieldset></div>'); //add input box
            x++; //text box increment
        }
    });

    $(wrapper).on("click",".remove_field", function(e){ //user click on remove text
        e.preventDefault(); $(this).parent('div').parent('div').remove(); x--;
    })
    });
    </script>

  <div class="row">
        <div class="control-group" id="fields">
            <label class="control-label" for="field1">Weekly Report - Manager Page</label>

                <div class="row">
                  <div class="col-md-6">
                    <form action="/" role="form" autocomplete="off" >
                      <div class="group">
                        {{.Form}}<!--end of loop-->

                      </div> <!-- end of div class group -->
                    <br>
                      <div class="btn-toolbar" role="toolbar" aria-label="...">
                        <div class="btn-group" role="group" aria-label="...">
                          <button class="btn btn-success btn-add" type="button">
                            <span class="glyphicon glyphicon-plus"></span>
                          </button>
                        </div>
                        <div class="btn-group" role="group" aria-label="...">
                          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                        </div>
                        <div class="btn-group" role="group" aria-label="...">
                          <input type="submit" class="btn btn-default btn-primary" value="Submit">
                        </div><!--end btn-group-->
                      </div> <!-- end of button toolbar-->
                    </form><!--end of form -->
                    </div> <!--end of col -->

                    <div class="col-md-2">
                      <br>
                    </div>
                    <div class="col-md-4">
                    <form class="form-horizontal" action="/search">
                      <div class="form-group-lg">
                        <label for="inputPassword2" class="sr-only">search terms</label>
                        <input type="text" class="form-control" name="search" placeholder="search terms">
                      </div>
                      <p>
                      <div class="form-group-sm">
                        <label class="col-sm-4 control-label" for="formGroupInputSmall">From Date</label>
                        <div class="col-sm-8"><input type="date" class="form-control" name="date1" placeholder="From Date"></div>
                      </div>
                      <div class="form-group-sm">
                        <label class="col-sm-4 control-label" for="formGroupInputSmall">To Date</label>
                        <div class="col-sm-8"><input type="date" class="form-control" name="date1" placeholder="From Date"></div>
                      </div>
                      &nbsp;
                      <div class="form group-sm">
                        <button type="submit" class="btn btn-default">Search</button>
                      </div>
                    </form>
                    </div>
                  </div> <!--end of row-->
                </div>
                <br>



        </div> <!-- end of control-group -->

  </div> <!-- end of row -->
</div> <!-- end of container -->

</body>
</html>
`
	/*adminPage = `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="utf-8">
		    <title>vWeekly Reporter</title>
		        <meta name="viewport" content="width=device-width, initial-scale=1">
		    <link href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css" rel="stylesheet" id="bootstrap-css">
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
		    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.0.0-alpha1/jquery.min.js"></script>
		    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
		    <link href="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/css/bootstrap-editable.css" rel="stylesheet"/>
		    <script src="http://cdnjs.cloudflare.com/ajax/libs/x-editable/1.5.0/bootstrap3-editable/js/bootstrap-editable.min.js"></script>
				<script src="jquery.min.js"></script>
		    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
		    <script src="http://www.appelsiini.net/download/jquery.jeditable.mini.js"></script>
		    <script>$(document).ready(function() {
		     $('.edit').editable('http://localhost:9001/api', {
		         type      : 'textarea',
		         cancel    : 'Cancel',
		         submit    : 'OK',
		         indicator : '<img src="img/indicator.gif">',
		         tooltip   : 'Click to edit...',
		      });
		 });</script>
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
		        <li class="active"><a href="#">Home <span class="sr-only">(current)</span></a></li>
		        <li><a href="#">About</a></li>
		        <li class="dropdown">
		          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
		          <ul class="dropdown-menu">
		            <li><a href="#">Analytics</a></li>
		            <li><a href="#">Dashboards</a></li>
		            <li><a href="#">Search</a></li>
		            <li role="separator" class="divider"></li>
		            <li><a href="#">Separated link</a></li>
		            <li role="separator" class="divider"></li>
		            <li><a href="#">One more separated link</a></li>
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

		  <div class="container">
		    <script>
		    $(document).ready(function() {
		    var max_fields      = 10; //maximum input boxes allowed
		    var wrapper         = $(".group"); //Fields wrapper
		    var add_button      = $(".btn-add"); //Add button ID

		    var x = 0; //initlal text box count
		    $(add_button).click(function(e){ //on add input button click
		        e.preventDefault();
		        if(x < max_fields){ //max input box allowed
		          $(wrapper).append('<div><br><div class="pull-right"><a href="#" class="remove_field"><button class="btn btn-danger btn-add" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><div class="entry input-group col-xs-10"><input class="form-control" name="fields[]" type="text" placeholder="Subject" /></div><div class="entry input-group col-xs-12"><textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea></div><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Hashtag" /></div><fieldset class="radiogroup"><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+' id="inlineRadio1" value="good">good</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio2" value="bad">bad</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions'+x+'" id="inlineRadio3" value="competitive">competitive</label></fieldset></div>'); //add input box
		            x++; //text box increment
		        }
		    });

		    $(wrapper).on("click",".remove_field", function(e){ //user click on remove text
		        e.preventDefault(); $(this).parent('div').parent('div').remove(); x--;
		    })
		    });
		    </script>

		  <div class="row">
		        <div class="control-group" id="fields">
		            <label class="control-label" for="field1">Weekly Report - Manager Page</label>

		                <div class="row">
		                  <div class="col-md-6">
		                    <form action="/" role="form" autocomplete="off" >
		                      <div class="group">
		                        {{.Form}}<!--end of loop-->

		                      </div> <!-- end of div class group -->
		                    <br>
		                      <div class="btn-toolbar" role="toolbar" aria-label="...">
		                        <div class="btn-group" role="group" aria-label="...">
		                          <button class="btn btn-success btn-add" type="button">
		                            <span class="glyphicon glyphicon-plus"></span>
		                          </button>
		                        </div>
		                        <div class="btn-group" role="group" aria-label="...">
		                          &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
		                        </div>
		                        <div class="btn-group" role="group" aria-label="...">
		                          <input type="submit" class="btn btn-default btn-primary" value="Submit">
		                        </div><!--end btn-group-->
		                      </div> <!-- end of button toolbar-->
		                    </form><!--end of form -->
		                    </div> <!--end of col -->

		                    <div class="col-md-2">
		                      <br>
		                    </div>
		                    <div class="col-md-4">
												<form class="form-horizontal" action="/search">
													<div class="form-group-lg">
														<label for="inputPassword2" class="sr-only">search terms</label>
														<input type="text" class="form-control" name="search" placeholder="search terms">
													</div>
													<p>
													<div class="form-group-sm">
														<label class="col-sm-4 control-label" for="formGroupInputSmall">From Date</label>
														<div class="col-sm-8"><input type="date" class="form-control" name="date1" placeholder="From Date"></div>
													</div>
													<div class="form-group-sm">
														<label class="col-sm-4 control-label" for="formGroupInputSmall">To Date</label>
														<div class="col-sm-8"><input type="date" class="form-control" name="date1" placeholder="From Date"></div>
													</div>
													&nbsp;
													<div class="form group-sm">
														<button type="submit" class="btn btn-default">Search</button>
													</div>
												</form>
		                    </div>
		                  </div> <!--end of row-->
		                </div>
		                <br>



		        </div> <!-- end of control-group -->

		  </div> <!-- end of row -->
		</div> <!-- end of container -->

		</body>
		</html>

	`
	*/

	form = `
	<div><!--start of loop--><br><div class="pull-right"><a href="#" class="remove_field"><button class="btn btn-danger btn-minus" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><div class="entry input-group col-xs-10"><div class="edit">{{.Title}}</div></div><div class="entry input-group col-xs-12"><div class="edit">{{.Content}}</div></div><div class="entry input-group col-xs-12"><div class="edit">{{.Hash}}</div></div><div class="entry input-group col-xs-12"><div class="edit">{{.Group}}</div></div></div>
    `
)
