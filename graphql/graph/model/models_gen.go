// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountInput struct {
	Name string `json:"name"`
}

type PaginationInput struct {
	Skip *int `json:"skip"`
	Take *int `json:"take"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Timestamp   string  `json:"timestamp"`
	Price       float64 `json:"price"`
}

type ProductInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
