package main

import (
	"fmt"
)

type book struct {
	title   string
	author  string
	price   float32
	publish bool
}

func newBook(bs *[]*book) {
	b := new(book)
	print("title:")
	_, err := fmt.Scanln(&b.title)
	if err != nil {
		return
	}
	print("author:")
	_, err = fmt.Scanln(&b.author)
	if err != nil {
		return
	}
	print("price:")
	_, err = fmt.Scanln(&b.price)
	if err != nil {
		return
	}
	print("publish:")
	_, err = fmt.Scanln(&b.publish)
	if err != nil {
		return
	}
	*bs = append(*bs, b)
	showBooks(*bs)
	return
}
func showBooks(books []*book) {
	for _, i := range books {
		fmt.Printf("%v\n", *i)
	}
}
func delBook(bs *[]*book) {
	var title string
	print("title:")
	_, err := fmt.Scanln(&title)
	if err != nil {
		return
	}
	for index, i := range *bs {
		if (*i).title == title {
			*bs = append((*bs)[:index], (*bs)[index+1:]...)
		}
	}
	showBooks(*bs)
	return
}

func main() {
	var books []*book
	newBook(&books)
	newBook(&books)
	delBook(&books)

}
