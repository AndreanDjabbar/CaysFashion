package entities

type UserRole string

const (
    UserRoleUser   UserRole = "User"
    UserRoleAdmin  UserRole = "Admin"
)

type User struct {
	UserID   uint      `gorm:"primaryKey;autoIncrement" json:"userID"`
	Username string    `gorm:"unique;size:50;not null" json:"username"`
	Email    string    `gorm:"unique;size:50;not null" json:"email"`
	Password string    `gorm:"size:70;not null" json:"-"` 
	Role     UserRole  `gorm:"type:enum('User', 'Admin');default:'User'" json:"role"`
}
