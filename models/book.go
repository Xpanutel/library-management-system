package models

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
	IsLoaned bool
}
