package handler

import (
	"clinic-management/pkg/db"
	"clinic-management/pkg/model"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type PatientReq struct {
	Name        string    `json:"name,omitempty"`
	Contact     string    `json:"contact,omitempty"`
	Address     string    `json:"address,omitempty"`
	Reason      string    `json:"reason,omitempty"`
	ConsultBy   string    `json:"consult_by,omitempty"`
	DateOfVisit time.Time `json:"date_of_visit,omitempty"`
}

type UpdatePatientReq struct {
	Name      string `json:"name,omitempty"`
	Contact   string `json:"contact,omitempty"`
	Address   string `json:"address,omitempty"`
	ConsultBy string `json:"consult_by,omitempty"`
}

func CreatePatient(req PatientReq) (string, int, error) {
	data := model.Patient{
		Id:          uuid.NewString(),
		Name:        req.Name,
		Contact:     req.Contact,
		Address:     req.Address,
		Reason:      req.Reason,
		ConsultBy:   req.ConsultBy,
		DateOfVisit: time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := db.CreatePatient(data)
	if err != nil {

		slog.Error("error while creating record in database", "error", err.Error())
		return "", 500, fmt.Errorf("internal server error")
	}

	slog.Info("Patient record inserted successfully")
	return data.Id, 201, nil
}

func ListAll(intLimit, intSkip int, search string) ([]model.Patient, int, error) {
	record, err := db.ListAll(intLimit, intSkip, search)
	if err != nil {
		slog.Error("error while creating record in database", "error", err.Error())
		return []model.Patient{}, 500, fmt.Errorf("internal server error")
	}
	return record, 200, err
}

func GetPatient(idparam string) (model.Patient, int, error) {
	record, err := db.GetPatient(idparam)
	if err != nil {
		return model.Patient{}, 500, fmt.Errorf("internal server error")
	}
	return record, 200, err
}

func Update(id string, req UpdatePatientReq) (int, error) {
	patient := model.Patient{

		Name:      req.Name,
		Contact:   req.Contact,
		Address:   req.Address,
		ConsultBy: req.ConsultBy,
	}
	err := db.Update(id, patient)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
func Delete(id string) (int, error) {
	err := db.Delete(id)
	if err != nil {
		return 500, err
	}
	return 200, nil
}
