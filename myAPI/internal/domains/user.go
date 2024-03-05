package domains

type (
	User struct {
		ID       uint   `gorm:"primary_key"`
		Nama     string `gorm:"nama"`
		Email    string `gorm:"email"`
		Password string `gorm:"password"`
		Role     string `gorm:"role"`
	}
)
