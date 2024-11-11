func NewLoan(book *Book, user *User) *Loan {
    if book.IsAvailable {
        book.IsAvailable = false
        return &Loan{
            Book:       book,
            User:       user,
            LoanDate:   time.Now(),
            ReturnDate: time.Now().AddDate(0, 0, 14),
        }
    } 
    fmt.Println("Libro no disponible para préstamo")
    return nil
}
