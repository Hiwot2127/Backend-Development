package models

// Member model
type Member struct {
    ID           int
    Name         string
    BorrowedBooks []Book 
}
