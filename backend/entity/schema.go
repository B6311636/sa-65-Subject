package entity

import (
	"gorm.io/gorm"
)

type Officer struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex" valid:"email"`
	Password string `json:"-"`

	Subject []Subject `gorm:"foreignKey:OfficerID"`
}

type Teacher struct {
	gorm.Model
	Level string
	Name  string
	Email string

	FacultyID *uint
	Faculty   Faculty
	Subject   []Subject `gorm:"foreignKey:TeacherID"`
}

type Faculty struct {
	gorm.Model
	Name string

	Teacher []Teacher `gorm:"foreignKey:FacultyID"`

	Subject []Subject `gorm:"foreignKey:FacultyID"`
}

type Time struct {
	gorm.Model
	Period string

	Subject []Subject `gorm:"foreignKey:TimeID"`
}

type Subject struct {
	gorm.Model
	Code    string
	Name    string
	Credit  uint
	Section uint
	Day     string
	Take    uint

	TeacherID *uint
	Teacher   Teacher

	FacultyID *uint
	Faculty   Faculty

	OfficerID *uint
	Officer   Officer

	TimeID *uint
	Time   Time
}
