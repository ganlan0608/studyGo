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
	fmt.Scanln(&b.title)
	print("author:")
	fmt.Scanln(&b.author)
	print("price:")
	fmt.Scanln(&b.price)
	print("publish:")
	fmt.Scanln(&b.publish)
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
	fmt.Scanln(&title)
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
