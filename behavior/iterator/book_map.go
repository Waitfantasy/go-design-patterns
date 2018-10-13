package main

type BookShelfMap struct {
	count int
	books map[int]*Book
}

func NewBookShelfMap() *BookShelfMap {
	return &BookShelfMap{
		books: make(map[int]*Book),
	}
}

func (bs *BookShelfMap) GetLength() int {
	return bs.count
}

func (bs *BookShelfMap) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelfMap) AppendNewBook(book *Book) {
	bs.books[bs.count] = book
	bs.count++
}

func (bs *BookShelfMap) Iterator() IteratorInterface {
	return NewBookIterator(bs)
}
