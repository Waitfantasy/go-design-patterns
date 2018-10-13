package main

type BookIterator struct {
	cursor int
	bs     BookShelfInterface
}

func NewBookIterator(bookShelf BookShelfInterface) *BookIterator {
	return &BookIterator{
		cursor: 0,
		bs:     bookShelf,
	}
}

func (it *BookIterator) HasNext() bool {
	return it.cursor < it.bs.GetLength()
}

func (it *BookIterator) Next() interface{} {
	book := it.bs.GetBookAt(it.cursor)
	it.cursor++
	return book
}
