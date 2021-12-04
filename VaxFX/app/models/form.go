package models

type Form struct {
	ID 			int
	Name 		string
	Vac1		string
	Effects1	[]string
	Vac2		string
	Effects2	[]string
	Booster		string
	Effects3	[]string
	GreenPass 	string // this should be a link
}