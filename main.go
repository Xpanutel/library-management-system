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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		mu.Lock()
		userStore[userID] = &models.User{ID: userID, Username: username, Password: password}
		userID++
		mu.Unlock()

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		for _, user := range userStore {
			if user.Username == username && user.Password == password {
				session, _ := sessionStore.Get(r, "session")
				session.Values["userID"] = user.ID
				session.Save(r, w)
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
				return
			}
		}
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	tmpl, _ := template.ParseFiles("templates/login.html")
	tmpl.Execute(w, nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "session")
	userID, ok := session.Values["userID"].(int)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Здесь логика для отображения информации о пользователе и его книгах
	tmpl, _ := template.ParseFiles("templates/dashboard.html")
	tmpl.Execute(w, nil)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь будет логика для администраторов
	tmpl, _ := template.ParseFiles("templates/admin.html")
	tmpl.Execute(w, nil)
}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		author := r.FormValue("author")
		genre := r.FormValue("genre")

		mu.Lock()
		bookStore[bookID] = &models.Book{ID: bookID, Title: title, Author: author, Genre: genre, IsLoaned: false}
		bookID++
		mu.Unlock()

		http.Redirect(w, r, "/books", http.StatusSeeOther)
		return
	}

	tmpl, _ := template.ParseFiles("templates/books.html")
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/books", booksHandler)
	
	http.ListenAndServe(":8080", nil)
}