package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oporas/golang-practice/golang-web-training/21_package/controllers"
)

func main() {
	uc := controllers.NewUserController()
	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/:id", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
