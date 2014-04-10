/* Application : MeritWiki API
  Programming Language: Go
  FILE NAME: CreateUserAccount.go
  Created By: Robert Gosnell
  FUNCTIONS: createUserAccount - pointer to the User Struct
		   createPage - pointer to the Page Struct
		   userHandler - get the user values
		   read - read the inputted user values
		   stringHandler - handles string constraints in the api per the db constraints
  STRUCTS: User struct - defines the user information to used to create a user account

  PURPOSE of FILE and FUNCTIONS:
  *****************************************************
  MODIFICATIONS
  ****************************************************
  MODIFED BY: Rhonda B. Reece
  DATE of MODIFICATION: 3/1/2014
  REASON FOR MODIFICATION: added header to the file to explain functions in this library
*/
package main

import (
	"../dbfiles/meritdb" // import from the meritdb.go file of the MeritWiki Database functions
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"unicode/utf8"
)

//To use for Database Queries
var (
	userName string
)

type User struct {
	Firstname string
	LastName  string
	Email     string
	UserName  string
	Password  string
}

type Page struct {
	Title string
	Body  string
	Url   string
}

func createUserAccount(v []byte) {
	var u User
	err := json.Unmarshal(v, &u)
	if err != nil {
		log.Fatal(err)
	}
	response := meritdb.CalladdUserAccountSP(u.Firstname, u.LastName, u.Email, u.UserName, u.Password)
	fmt.Println(response)
	/*fmt.Println("User Created!")
	fmt.Printf(u.Firstname)
	fmt.Printf(u.LastName)
	fmt.Printf(u.Email)
	fmt.Printf(u.UserName)
	fmt.Printf(u.Password)*/
}

func userhandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		return
	}
	v, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	r.Body.Close()
	//fmt.Println(v)
	createUserAccount(v)
}

func read(prompt string) string {
	in := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)

	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return line
}

func stringHandler(s string) string {
	slength := (utf8.RuneCountInString(s) - 1)
	//fmt.Println(slength)
	//n:=0
	for slength == 0 {
		fmt.Println("Name Required")
		s = read("First Name: ")
		slength = (utf8.RuneCountInString(s) - 1)
		if slength > 0 {
			break
		}
	}
	return s
}

func main() {

	//UNCOMMENT NEXT TWO LINES FOR HTTP FORM INTERFACE
	//http.HandleFunc("/createuser", userhandler)
	//http.ListenAndServe(":8080", nil)

	u := User{"Rob", "Gosnell", "gosnellrd@gmail.com", "RobDaMan", "12344321"}
	v, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error:", err)
	}
	createUserAccount(v)
}
