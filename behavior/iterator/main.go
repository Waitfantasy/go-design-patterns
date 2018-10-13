package main

import "fmt"

func main() {
	var bookShelf BookShelfInterface
	bookShelf = NewBookShelfSlice()
	bookShelf.AppendNewBook(NewBook("Golang"))
	bookShelf.AppendNewBook(NewBook("etcd"))
	bookShelf.AppendNewBook(NewBook("Paxos"))
	bookShelf.AppendNewBook(NewBook("Raft"))

	bookShelf = NewBookShelfMap()
	bookShelf.AppendNewBook(NewBook("test1"))
	bookShelf.AppendNewBook(NewBook("test2"))
	bookShelf.AppendNewBook(NewBook("test3"))
	it := bookShelf.Iterator()
	for it.HasNext() {
		b := it.Next().(*Book)
		fmt.Println(b.GetName())
	}
}
