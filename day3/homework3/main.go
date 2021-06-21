package main

import (
	"fmt"
	"os"
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
func showMenu() string {
	var opt string
	fmt.Print(`
1.新增书籍：
2.删除书籍：
3.查看书籍：
4.退出：
`)
	_, err := fmt.Scanln(&opt)
	if err != nil {
		return ""
	}
	return opt
}

func main() {
	var books []*book

	for {
		opt := showMenu()
		switch opt {
		case "1":
			newBook(&books)
		case "2":
			delBook(&books)
		case "3":
			showBooks(books)
		case "4":
			os.Exit(0)

		}

	}

}
