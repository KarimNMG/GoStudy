package controllers

import (
	"net/http"
	"regexp"
)

type UserController struct {
	userIdPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from user controller"))
}

// constructor
func newUserController() *UserController {
	return &UserController{
		userIdPattern: regexp.MustCompile(`^/users/(\d)/?`),
	}
} 
