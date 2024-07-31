package services

import (
    "errors"
    "library_management/models"
)

// LibraryManager interface
type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int) error
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
}

// Contains a map of books and a map of members
type Library struct {
    books   map[int]models.Book
    members map[int]models.Member
}

// Returns a new instance of the Library struct
func NewLibrary() *Library {
    return &Library{
        books:   make(map[int]models.Book),
        members: make(map[int]models.Member),
    }
}

// Adding a new book to the library
func (l *Library) AddBook(book models.Book) {
    l.books[book.ID] = book
}

// Removing a book from the library
func (l *Library) RemoveBook(bookID int) error {
    if _, exists := l.books[bookID]; exists {
        delete(l.books, bookID)
        return nil
    }
    return errors.New("book not found")
}

// Borrowing a book from the library
func (l *Library) BorrowBook(bookID int, memberID int) error {
    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }

    if book.Status == "Borrowed" {
        return errors.New("book is already borrowed")
    }

    member, exists := l.members[memberID]
    if !exists {
        return errors.New("member not found")
    }

    book.Status = "Borrowed"
    l.books[bookID] = book

    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.members[memberID] = member

    return nil
}

// Returning a book to the library
func (l *Library) ReturnBook(bookID int, memberID int) error {
    member, exists := l.members[memberID]
    if !exists {
        return errors.New("member not found")
    }

    book, exists := l.books[bookID]
    if !exists {
        return errors.New("book not found")
    }

    if book.Status == "Available" {
        return errors.New("book is not borrowed")
    }

    for i, b := range member.BorrowedBooks {
        if b.ID == bookID {
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
            break
        }
    }

    book.Status = "Available"
    l.books[bookID] = book
    l.members[memberID] = member

    return nil
}

// List of available books
func (l *Library) ListAvailableBooks() []models.Book {
    var availableBooks []models.Book
    for _, book := range l.books {
        if book.Status == "Available" {
            availableBooks = append(availableBooks, book)
        }
    }
    return availableBooks
}

// List of borrowed books
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, exists := l.members[memberID]
    if !exists {
        return nil
    }
    return member.BorrowedBooks
}
