package structsinterfaces

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

type Printable interface {
	Info() string
	PageNum() int
}

type Book struct {
	Author string
	Title  string
	Pages  int
}

type Magazine struct {
	Title string
	Issue string
	Pages int
}

func NewBook(Author, Title string, Pages int) Book {
	return Book{
		Author: Author,
		Title:  Title,
		Pages:  Pages,
	}
}

func NewMagazine(Title, Issue string, Pages int) Magazine {
	return Magazine{
		Title: Title,
		Issue: Issue,
		Pages: Pages,
	}
}

func (b Book) Info() string {
	return fmt.Sprintf("%s, %s", b.Author, b.Title)
}

func (b Book) PageNum() int {
	return b.Pages
}

func (m Magazine) Info() string {
	return fmt.Sprintf("%s, %s", m.Title, m.Issue)
}

func (m Magazine) PageNum() int {
	return m.Pages
}
