package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	userId := 1
	fmt.Println("Enter user Id:")
	// get user Id from user
	fmt.Scanln(&userId)
	// pass the user id to the function``
	makeGetRequest(userId)
}

func makeGetRequest(userId int) {
	//convert int to string for concatinationg it to a string
	t := strconv.Itoa(userId)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/" + t + "/todos")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
