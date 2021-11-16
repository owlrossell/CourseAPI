package model

import "github.com/kamva/mgm/v3"

type Course struct {
	mgm.IDField
	Name string
	Description string
}

func NewCourse( name string, description string) *Course  {
	return &Course{
		Name: name,
		Description: description,
	}
}