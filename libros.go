package main
import (
    "fmt"
    "time"
)
type Book struct {
    Title       string
    Author      string
    IsAvailable bool
}
type User struct {
    Name     string
    UserType string 
}
type Loan struct {
    Book       *Book
    User       *User
    LoanDate   time.Time
    ReturnDate time.Time
}
