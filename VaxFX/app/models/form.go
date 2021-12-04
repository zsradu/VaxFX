package models

import (
	"github.com/jinzhu/gorm"
)

type Form struct {
	gorm.Model
	Name 		string
	Vac1		string
	Effects1	[]Effects
	Vac2		string
	Effects2	[]Effects
	Booster		string
	Effects3	[]Effects
	// GreenPass 	string // this should be a link
	// deleted storage for green pass, but should find a way to store it but not as img preferably
}

type Effects struct {
	gorm.Model
	Name 		string
	FormID		uint
}