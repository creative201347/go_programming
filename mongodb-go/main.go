package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github/creative201347/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserrController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
