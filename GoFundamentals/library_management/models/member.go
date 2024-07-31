package models

// Member model
type Member struct {
    ID           int
    Name         string
    BorrowedBooks []Book // a slice to hold borrowed books
}
