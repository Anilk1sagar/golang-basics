package main

import (
	"fmt"
)

func main() {

	// List of endpoints that doesn't require auth
	notAuth := []string{
		"/api/test/mysql/add", "/api/test/mysql/getAll", "/api/test/mysql/get/{name}", //test mysql routes
		"/api/user/register", "/api/user/auth", //User routes
	}

	//current request path
	requestPath := "/api/test/mysql/add"

	// Check if request does not need authentication, serve the request if it doesn't need it
	for _, value := range notAuth {

		if value == requestPath {

			fmt.Println("no auth: ", value, "**==** reqPath: ", requestPath)
			return
		}

		fmt.Println("after")
	}
}
