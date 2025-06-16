package db

import (
	"clinic-management/pkg/model"
	"fmt"
	"log/slog"
	"time"
)

func CreatePatient(data model.Patient) error {

	query := `INSERT INTO PATIENTS(id,name,contact,address,reason,consult_by,date_of_visit) VALUES ($1,$2,$3,$4,$5,$6,$7)`

 
	_, err := conn.Exec(query, data.Id, data.Name, data.Contact, data.Address, data.Reason, data.ConsultBy, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func ListAll(limit, offset int, search string) ([]model.Patient, error) {
	query := fmt.Sprintf("SELECT id, name,contact,address,reason,consult_by,date_of_visit FROM patients limit %d offset %d", limit, offset)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []model.Patient

	for rows.Next() {
		var data model.Patient
		err := rows.Scan(&data.Id, &data.Name, &data.Contact, &data.Address, &data.Reason, &data.ConsultBy, &data.DateOfVisit)
		if err != nil {
			return nil, err
		}
		patients = append(patients, data)
	}

	return patients, nil

}

func GetPatient(idparam string) (model.Patient, error) {
	var patient model.Patient

	query := `SELECT id, name, contact, address, reason, consult_by, date_of_visit, created_at, updated_at 
	          FROM patients WHERE id = $1 LIMIT 1`
	err := conn.QueryRow(query, idparam).Scan(
		&patient.Id,
		&patient.Name,
		&patient.Contact,
		&patient.Address,
		&patient.Reason,
		&patient.ConsultBy,
		&patient.DateOfVisit,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)

	if err != nil {
		slog.Error("Failed to fetch patient", "error", err.Error())
		return patient, err
	}

	return patient, nil
}
func Update(id string, p model.Patient) error {
	query := `
		UPDATE patients
		SET name = $1,
		    contact = $2,
		    address = $3,
		    consult_by = $4,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $5;
	`

	_, err := conn.Exec(query, p.Name, p.Contact, p.Address, p.ConsultBy, id)
	if err != nil {
		slog.Error("Failed to update patient", "id", id, "error", err.Error())
		return err
	}

	slog.Info("Patient updated successfully", "id", id)
	return nil
}
func Delete(id string) error {
	query := `
		DELETE FROM patients
		WHERE id = $1;
	`

	_, err := conn.Exec(query, id)
	if err != nil {
		slog.Error("Failed to delete patient", "id", id, "error", err.Error())
		return err
	}

	slog.Info("Patient deleted successfully", "id", id)
	return nil
}
