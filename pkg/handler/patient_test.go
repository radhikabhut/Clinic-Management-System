package handler

import (
	"clinic-management/config"
	"clinic-management/pkg/db"
	"clinic-management/pkg/model"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadConfig("../../.env")
	err := db.ConnectDB()
	if err != nil {
		slog.Error("error while connection database", "error", err.Error())
		return
	}
	m.Run()

}

func TestCreatePatientHandler_Positive(t *testing.T) {
	req := PatientReq{
		Name:      "Jane Doe",
		Contact:   "9876543210",
		Address:   "456 Side Street",
		Reason:    "Checkup",
		ConsultBy: "Dr. Watson",
	}

	id, status, err := CreatePatient(req)

	assert.NoError(t, err)
	assert.Equal(t, 201, status)
	assert.NotEmpty(t, id)

	err = db.Delete(id)
	assert.NoError(t, err)
}

func TestCreatePatientHandler_LongContact(t *testing.T) {
	req := PatientReq{
		Name:      "Patient Test",
		Contact:   "123456789012345678901234567890",
		Address:   "Test Address",
		Reason:    "Routine",
		ConsultBy: "Dr. House",
	}

	id, status, err := CreatePatient(req)

	assert.Error(t, err)
	assert.Equal(t, 500, status)
	assert.Empty(t, id)
}

func TestListAllHandler_Positive(t *testing.T) {
	
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "John Handler",
		Contact:   "9876543210",
		Address:   "Test Street",
		Reason:    "Headache",
		ConsultBy: "Dr. Watson",
	}
	err := db.CreatePatient(patient)
	assert.NoError(t, err)

	results, status, err := ListAll(10, 0, "")
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.GreaterOrEqual(t, len(results), 1)


	err = db.Delete(patient.Id)
	assert.NoError(t, err)
}

func TestGetPatientHandler_Positive(t *testing.T) {
	
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "Handler Test",
		Contact:   "8888888888",
		Address:   "Test Address",
		Reason:    "Checkup",
		ConsultBy: "Dr. Handler",
	}

	err := db.CreatePatient(patient)
	assert.NoError(t, err)

	
	result, status, err := GetPatient(patient.Id)

	
	assert.NoError(t, err)
	assert.Equal(t, 200, status)
	assert.Equal(t, patient.Name, result.Name)
	assert.Equal(t, patient.Contact, result.Contact)
	assert.Equal(t, patient.Address, result.Address)
	assert.Equal(t, patient.Reason, result.Reason)
	assert.Equal(t, patient.ConsultBy, result.ConsultBy)

	
	err = db.Delete(patient.Id)
	assert.NoError(t, err)
}
func TestGetPatientHandler_InvalidID(t *testing.T) {
	
	result, status, err := GetPatient("non-existent-id-123")

	
	assert.Error(t, err)
	assert.Equal(t, 500, status)
	assert.Empty(t, result.Id)
}

func TestUpdateHandler_Positive(t *testing.T) {

	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "Old Name",
		Contact:   "1234567890",
		Address:   "Old Address",
		Reason:    "Cold",
		ConsultBy: "Dr. Smith",
	}
	err := db.CreatePatient(patient)
	assert.NoError(t, err)

	updateReq := UpdatePatientReq{
		Name:      "New Name",
		Contact:   "9876543210",
		Address:   "New Address",
		ConsultBy: "Dr. Adams",
	}
	status, err := Update(patient.Id, updateReq)

	assert.NoError(t, err)
	assert.Equal(t, 200, status)

	err = db.Delete(patient.Id)
	assert.NoError(t, err)
}

