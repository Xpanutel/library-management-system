package main

import (
	"html/template"
	"net/http"
	"sync"
	"github.com/gorilla/sessions"
	"log"
)

var (
	bookStore   = make(map[int]*models.Book)
	userStore   = make(map[int]*models.User)
	sessionStore = sessions.NewCookieStore([]byte("secret-key"))
	bookID      = 1
	userID      = 1
	mu          sync.Mutex
)

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/books", booksHandler)
	
	http.ListenAndServe(":8080", nil)
}

// Handlers и другие функции будут здесь...
