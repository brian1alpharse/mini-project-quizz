package domians

import "time"

type (
	Employee struct {
		EmployeeID int       `gorm:"type:int unsigned;primaryKey" json:"employeeID"`
		FirstName  string    `gorm:"column:first_name" json:"firstName"`
		LastName   string    `gorm:"column:last_name" json:"LastName"`
		HireDate   time.Time `gorm:"column:hire_date" json:"hireDate"`
		Salary     float32   `gorm:"column:salary" json:"salary"`
	}
)
