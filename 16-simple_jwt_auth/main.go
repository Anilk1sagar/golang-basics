package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {

			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}

		} else {

			log.Println("Not Authorized")
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Anil Sagar"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")
	fmt.Println("Endpoint Hit: homePage")

}

func getToken(w http.ResponseWriter, r *http.Request) {

	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	fmt.Println("Valid Token is:", validToken)

	// Send Response
	fmt.Fprintf(w, validToken)

	fmt.Println("Endpoint Hit: GetToken")
}

func handleRequests() {

	http.Handle("/api/home", isAuthorized(homePage)) // Protected Path
	http.HandleFunc("/api/getToken", getToken)       // Get jwt Token

	HOST := "localhost:9110"
	fmt.Println("Server is Listning on:", HOST)
	log.Fatal(http.ListenAndServe(HOST, nil))
}

func main() {
	handleRequests()
}
