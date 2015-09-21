package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	page = `<!DOCTYPE html>
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
	          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
	          <ul class="dropdown-menu">
	            <li><a href="#">Admin</a></li>
	            <li><a href="#">Help</a></li>
	            <li><a href="#">Something else here</a></li>
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
	        <li><a href="#"><b>September 30th, 2015</b>: Jonathan Cham</a></li>
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
	            x++; //text box increment
	            $(wrapper).append('<div><br><div class="pull-right"><a href="#" class="remove_field"><button class="btn btn-danger btn-add" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><div class="entry input-group col-xs-10"><input class="form-control" name="fields[]" type="text" placeholder="Subject" /></div><div class="entry input-group col-xs-12"><textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea></div><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Hashtag" /></div><label class="radio-inline"><input type="radio" name="inlineRadioOptions" id="inlineRadio1" value="option1">good</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions" id="inlineRadio2" value="option2">bad</label><label class="radio-inline"><input type="radio" name="inlineRadioOptions" id="inlineRadio3" value="option3">competitive</label></div>'); //add input box
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
	                    <form action="/" role="form" autocomplete="off" >
	                      <div class="group">
	                        <div class="entry input-group col-xs-10">
	                          <input class="form-control" name="fields[]" type="text" placeholder="Subject" />
	                        </div>
	                        <div class="entry input-group col-xs-12">
	                          <textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea>
	                        </div>
	                        <div class="entry input-group col-xs-12">
	                          <input class="form-control" name="fields[]" type="text" placeholder="Hashtag" />
	                        </div>
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions" id="inlineRadio1" value="option1"> good
	                        </label>
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions" id="inlineRadio2" value="option2"> bad
	                        </label>
	                        <label class="radio-inline">
	                          <input type="radio" name="inlineRadioOptions" id="inlineRadio3" value="option3"> competitive
	                        </label>
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
	        <li class="active"><a href="#">Help <span class="sr-only">(current)</span></a></li>
	        <li><a href="#">About</a></li>
	        <li class="dropdown">
	          <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span class="caret"></span></a>
	          <ul class="dropdown-menu">
	            <li><a href="#">Admin</a></li>
	            <li><a href="#">Help</a></li>
	            <li><a href="#">Something else here</a></li>
	            <li role="separator" class="divider"></li>
	            <li><a href="#">Separated link</a></li>
	            <li role="separator" class="divider"></li>
	            <li><a href="#">One more separated link</a></li>
	          </ul>
	        </li>
	      </ul>

	    </div><!-- /.navbar-collapse -->
	  </div><!-- /.container-fluid -->
	</nav>


	<div class="container">


	  <script>
	  $(document).ready(function() {
	  var max_fields      = 20; //maximum input boxes allowed
	  var wrapper         = $(".controls"); //Fields wrapper
	  var add_button      = $(".btn-add"); //Add button ID

	  var x = 1; //initlal text box count
	  $(add_button).click(function(e){ //on add input button click
	      e.preventDefault();
	      if(x < max_fields){ //max input box allowed
	          x++; //text box increment
	          $(wrapper).append('<div class="row"><br><div class="col-md-4"><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Subject" /></div><div class="entry input-group col-xs-10"><textarea class="form-control" name="fields[]" type="text" placeholder="Description" rows="3"></textarea></div><div class="entry input-group col-xs-12"><input class="form-control" name="fields[]" type="text" placeholder="Hashtag" /></div></div><div class="col-md-4"><a href="#" class="remove_field"><button class="btn btn-danger btn-add" type="button" align="right"><span class="glyphicon glyphicon-minus"></span></button></a></div><br></div>'); //add input box
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
	          <form action="/" role="form" autocomplete="off" >

	              <div class="controls">
	                {{.Form}}
	              </div> <! -- end of controls -->
	              <br>

	            <div class="btn-toolbar" role="toolbar" aria-label="...">
	              <div class="btn-group" role="group" aria-label="...">
	                <button class="btn btn-success btn-add" type="button">
	                  <span class="glyphicon glyphicon-plus"></span>
	                </button>
	              <div class="btn-group" role="group" aria-label="...">
	                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	              </div>
	              <div class="btn-group" role="group" aria-label="...">
	              <input type="submit" class="btn btn-default btn-primary" value="Submit">
	              </div>
	            </div><!--end btn-toolbar-->
	          </form><!--end of form -->

	      </div> <!-- end of control-group -->

	</div> <!-- end of row -->
	</div> <!-- end of container -->

	</body>
	</html>
`
	form = `
	<div class="row">
		<div class="col-md-4">
			<div class="entry input-group col-xs-10">
	      <div class="edit">{{.Title}}</div>
	    </div>
	    <div class="entry input-group col-xs-12">
	      <div class="edit">{{.Content}}</div>
	    </div>
	    <div class="entry input-group col-xs-12">
	      <div class="edit">{{.Hash}}</div>
	    </div>
	  </div> <!--end of col-md-4-->
	  <div class="col-md-4">
			<a href="#" class="remove_field"><button class="btn btn-danger btn-minus" type="button" align="right">
	    <span class="glyphicon glyphicon-minus"></span></button></a>
	  </div> <!--end of col-md-4-->
	</div>`
)

type Data struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `bson:"name"`
	Manager string        `bson:"manager"`
	Title   string        `bson:"title"`
	Content string        `bson:"content"`
	Hash    string        `bson:"hash"`
	Group   int           `bson:"group"`
}

type Page struct {
	Form string
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/favicon.ico", punt)
	http.HandleFunc("/api", post)
	http.HandleFunc("/admin", admin)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func query(query string) (results []Data) {
	session, _ := mgo.Dial("192.168.22.128:27017")
	c := session.DB("test").C("weekly")

	c.Find(bson.M{"manager": query}).All(&results)
	fmt.Println("readDB")
	for _, name := range results {
		fmt.Println("these are the results")
		fmt.Println(name.Content, name.Manager)
	}

	return
}
func admin(w http.ResponseWriter, r *http.Request) {
	/* import template
	t, _ := template.ParseFiles("../template/shotchart.html")
	t.Execute(wr io.Writer, p)
	*/

	temp := new(Page)

	values := readDB()
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value.Content == "" {
			continue
		}
		fmt.Println("this is reading from the database the different lines", value)
		t := template.Must(template.New("form").Parse(form))
		t.Execute(buf, value)
	}
	fmt.Println("this is the string of template", buf.String())
	temp.Form = buf.String()
	t := template.Must(template.New("page").Parse(adminPage))
	t.Execute(w, temp)

	//t := template.Must(template.New("form").Parse(form))
	//t.Execute(w, Tom)
	//fmt.Println("this is admin page")
}
func punt(w http.ResponseWriter, r *http.Request) {
	return
}

func post(w http.ResponseWriter, r *http.Request) {
	text, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(text))
}

func create() (x *Data) {
	x = &Data{Name: "Title", Manager: "manager", Title: "Wells Fargo"}
	return
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
	fmt.Println("homepage executed")
	values := processForm(w, r)
	db(values)

}

func processForm(w http.ResponseWriter, r *http.Request) (values []Data) {
	r.ParseForm()
	x := r.Form["fields[]"]
	mgr := r.FormValue("search")
	if mgr != "" {
		values = query(mgr)
		return
	}

	values = make([]Data, len(x)/3)
	stuff := new(Data)

	//var stuff Data

	z := 0
	fmt.Println("length of x is ", len(x))
	fmt.Println("content of x", x)

	if len(x) != 0 {
		for i := 0; i < len(x); i = i + 3 {
			stuff.Title = x[i]
			stuff.Content = x[i+1]
			stuff.Hash = x[i+2]
			stuff.ID = bson.NewObjectId()
			stuff.Manager = "Tim"

			values[z] = *stuff
			z++
		}
	}

	fmt.Println(values)
	return
}

func readDB() (results []Data) {
	session, _ := mgo.Dial("192.168.22.128:27017")
	c := session.DB("test").C("weekly")

	c.Find(bson.M{"manager": "Tim"}).All(&results)
	fmt.Println("readDB")
	for _, name := range results {
		fmt.Println("these are the results")
		fmt.Println(name.Content, name.Manager)
	}
	return
}

func db(values []Data) {
	fmt.Println("start of db")
	session, err := mgo.Dial("192.168.22.128:27017")
	fmt.Println("after dial")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//string := "Timmy"

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("weekly")

	index := mgo.Index{
		Key:        []string{"content"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		//panic(err)
	}

	y := create()
	y.ID = bson.NewObjectId()
	fmt.Println("before insert")

	for i := 0; i < len(values); i++ {
		err = c.Insert(values[i])
	}
	fmt.Println("after insert")

	if err != nil {
		//log.Fatal(err)
	}

}
