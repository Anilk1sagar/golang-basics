package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"

// 	"github.com/ant0ine/go-json-rest/rest"
// )

// func main() {

// 	api := rest.NewApi()

// 	api.Use(rest.DefaultDevStack...)

// 	router, err := rest.MakeRouter(

// 		rest.Get("/lookup/#host", func(w rest.ResponseWriter, req *rest.Request) {

// 			ip, err := net.LookupIP(req.PathParam("host"))
// 			if err != nil {
// 				rest.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 			w.WriteJson(&ip)
// 		}),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	api.SetApp(router)
// 	fmt.Println("Server Started on port: 9110")
// 	log.Fatal(http.ListenAndServe(":9110", api.MakeHandler()))
// }
