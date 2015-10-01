package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

type Search struct {
	SearchString string `json:"searchString"`
}

func main() {
	http.HandleFunc("/api/search", query)
	http.ListenAndServe(":9002", nil)
}

func query(w http.ResponseWriter, r *http.Request) {
	var results []Data
	var temp1 Search
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println("body is ", string(body))
	json.Unmarshal(body, &temp1)
	fmt.Println("temp1 is", temp1.SearchString)
	session, _ := mgo.Dial("192.168.22.128:27017")
	c := session.DB("test").C("weekly")
	//temp := new([]Data)
	regexSearch := temp1.SearchString + ".*"
	fmt.Println(regexSearch)
	c.Find(bson.M{"hash": bson.M{"$regex": bson.RegEx{regexSearch, "i"}}}).All(&results)
	//results = append(results, *temp...)
	resultsJSON, _ := json.Marshal(results)
	fmt.Println("results are", results)
	fmt.Fprint(w, string(resultsJSON))

}
