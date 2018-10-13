package main

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{
		name: name,
	}
}

func (b *Book) GetName() string {
	return b.name
}

type BookShelfInterface interface {
	GetLength() int

	AppendNewBook(book *Book)

	GetBookAt(index int) *Book

	Iterator() IteratorInterface
}


