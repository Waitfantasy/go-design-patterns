package main

type BookShelfSlice struct {
	count int
	books []*Book
}

func NewBookShelfSlice() *BookShelfSlice {
	return &BookShelfSlice{}
}

func (bs *BookShelfSlice) GetLength() int {
	if bs.count == 0 {
		bs.count = len(bs.books)
	}
	return bs.count
}

func (bs *BookShelfSlice) AppendNewBook(book *Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookShelfSlice) GetBookAt(index int) *Book {
	return bs.books[index]
}

func (bs *BookShelfSlice) Iterator() IteratorInterface {
	return NewBookIterator(bs)
}
