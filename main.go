package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Data struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `bson:"name"`
	Manager string        `bson:"manager"`
	Title   string        `bson:"title"`
	Content string        `bson:"content"`
	Hash    string        `bson:"hash"`
	Group   string        `bson:"group"`
	Week    string        `bson:"week"`
}

type Page struct {
	Form string
}

func main() {
	http.HandleFunc("/heatmap", heatmap)
	http.HandleFunc("/jcham", homePage)
	http.HandleFunc("/tim", admin)
	http.HandleFunc("/favicon.ico", punt)
	http.HandleFunc("/api", post)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/analytics", comp)
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/dashboard2", dashboard2)
	http.HandleFunc("/search", search)
	http.HandleFunc("/", index)
	http.HandleFunc("/api/manager/tim", apiHandler)
	http.HandleFunc("/scripts/", scripts)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func comp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, compPage)
}
func heatmap(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, heatmapPage)
}
func dashboard2(w http.ResponseWriter, r *http.Request) {
	return
}
func scripts(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func search(w http.ResponseWriter, r *http.Request) {
	var searchResults []Data
	tempPage := new(Page)

	r.ParseForm()
	search := r.FormValue("search")
	fmt.Println("search input is", search)
	if search != "" {
		fmt.Println("doing a search")
		searchResults = query(search)
	}

	buf := new(bytes.Buffer)
	for _, value := range searchResults {
		if value.Content == "" {
			continue
		}
		//fmt.Println("this is reading from the database the different lines", value)
		t := template.Must(template.New("form").Parse(form))
		t.Execute(buf, value)
	}
	//fmt.Println("this is the string of template", buf.String())
	tempPage.Form = buf.String()
	t := template.Must(template.New("page").Parse(adminPage))
	t.Execute(w, tempPage)

	//t := template.Must(template.New("form").Parse(form))
	//t.Execute(w, Tom)
	fmt.Println("finished search page")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	values := readDB()
	results, _ := json.Marshal(values)

	w.Header().Set("Content-Type", "application/json")
	w.Write(results)
}

func analytics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, charts)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, dashboardPage)
}

func index(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, r.URL.Path[1:])
	fmt.Fprint(w, loginPage)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("inputEmail")
	fmt.Println("user from  is", user)
	if user == "jcham@vmware.com" {
		http.Redirect(w, r, "/jcham", 302)
	} else if user == "tim@vmware.com" {
		http.Redirect(w, r, "/tim", 302)
	}

	http.Redirect(w, r, "/", 302)
}
func admin(w http.ResponseWriter, r *http.Request) {

	temp := new(Page)

	values := readDB()
	buf := new(bytes.Buffer)
	for _, value := range values {
		if value.Content == "" {
			continue
		}
		//fmt.Println("this is reading from the database the different lines", value)
		t := template.Must(template.New("form").Parse(form))
		t.Execute(buf, value)
	}
	//fmt.Println("this is the string of template", buf.String())
	temp.Form = buf.String()
	t := template.Must(template.New("page").Parse(adminPage))
	t.Execute(w, temp)

	//t := template.Must(template.New("form").Parse(form))
	//t.Execute(w, Tom)
	fmt.Println("finished admin page")
}
func punt(w http.ResponseWriter, r *http.Request) {
	return
}

func post(w http.ResponseWriter, r *http.Request) {
	text, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(text))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
	fmt.Println("homepage executed")
	values := processForm(w, r)
	fmt.Println("values from process form", values)
	db(values)

}

func query(query string) (results []Data) {
	session, _ := mgo.Dial("192.168.22.128:27017")
	c := session.DB("test").C("weekly")
	//temp := new([]Data)

	c.Find(bson.M{"hash": query}).All(&results)
	//results = append(results, *temp...)

	fmt.Println("readDB")
	fmt.Println("these are the search results")

	if len(results) == 0 {
		fmt.Println("no results from search")
	} else {
		for _, name := range results {
			fmt.Println("From search results", name.Manager, name.Title, name.Content, name.Hash)
		}
	}
	return
}

func processForm(w http.ResponseWriter, r *http.Request) (values []Data) {
	r.ParseForm()
	x := r.Form["fields[]"]

	search := r.FormValue("search")
	fmt.Println("search input is", search)
	if search != "" {
		fmt.Println("doing a search")
		query(search)
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
			stuff.Week = "Week 39"
			stuff.Group = r.FormValue("inlineRadioOptions" + strconv.Itoa(z))
			fmt.Println("stuff.Group is", stuff.Group, r.FormValue("inlineRadioOptions"))
			values[z] = *stuff
			z++
		}
	}

	fmt.Println("these are the values from form input", values)
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

	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
		err = c.Insert(values[i])
	}
	fmt.Println("after insert")

	if err != nil {
		fmt.Println(err)
	}

}
