package model

import "time"

type Patient struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	Address     string    `json:"addrress,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	ConsultBy   string    `json:"consult_by,omitempty"`
	DateOfVisit time.Time `json:"date_of_visit,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
