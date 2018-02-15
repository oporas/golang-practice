package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oporas/golang-practice/golang-web-training/22_mongodb/controllers"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	uc := controllers.NewUserController(getSession())
	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/:id", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
