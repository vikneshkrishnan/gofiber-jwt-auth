package model

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"not null" json:"title"`
	Author string `gorm:"not null" json:"author"`
}

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password_hash,omitempty"`
}
