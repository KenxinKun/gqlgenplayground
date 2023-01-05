// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	Name string `json:"name"`
}

type Book struct {
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Fractal struct {
	Level   int      `json:"level"`
	Fractal *Fractal `json:"fractal"`
}

type Library struct {
	Books   []*Book `json:"books"`
	Address string  `json:"address"`
}
