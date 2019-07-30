package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/db"

	fbConfig "gitlab.com/anilk1sagar/go_basic_practice/21-simple-firebase-test/firebaseConfig"
)

var firebaseDb *db.Client

func main() {

	fmt.Println("Simple Firebase add, get test")

	client := fbConfig.InitFirebase()
	fmt.Println("firebase db: ", client)
	firebaseDb = client

	/* Routes */
	http.HandleFunc("/api/test/firebase/add", add)
	http.HandleFunc("/api/test/firebase/getAll", getAll)

	/* Server */
	fmt.Println("Server Started on port: ", 9110)
	log.Fatal(http.ListenAndServe(":9110", nil))

}

// Test ...
type Test struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func add(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	_mDbRef := firebaseDb.NewRef("dev/tests")

	key := fbConfig.KeyGenerate()
	fmt.Println("Key: ", key)

	test := Test{
		ID:   key,
		Name: r.FormValue("name"),
	}

	dbRefKey := _mDbRef.Child(key)

	//Adding
	err := dbRefKey.Set(ctx, test)
	if err != nil {
		log.Fatalln("Save Error: ", err.Error())
	}

	fmt.Fprintf(w, "Hi there, I love csgo")
	return
}

func getAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get all runs")

	ctx := context.Background()

	_mDbRef := firebaseDb.NewRef("dev/tests")

	var dbObject map[string]interface{}
	var testList []interface{}

	fmt.Println("\n start: ", time.Now())

	err := _mDbRef.Get(ctx, &dbObject)
	if err != nil {
		log.Fatalln("Error reading value:", err)
	}

	fmt.Println("\n end: ", time.Now())

	for _, val := range dbObject {
		testList = append(testList, val)
	}

	fmt.Println("\nTests list: ", testList)

	fmt.Fprintf(w, "getAll testList")
}
