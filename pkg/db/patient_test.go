package db

import (
	"clinic-management/config"
	"clinic-management/pkg/model"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.LoadConfig("../../.env")
	err := ConnectDB()
	if err != nil {
		slog.Error("error while connection database", "error", err.Error())
		return
	}
	m.Run()

}

func TestCreatePatient_Positive(t *testing.T) {
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "John Doe",
		Contact:   "1234567890",
		Address:   "123 Main Street",
		Reason:    "Fever",
		ConsultBy: "Dr. Smith",
	}

	err := CreatePatient(patient)
	assert.NoError(t, err)

	err = Delete(patient.Id)
	assert.NoError(t, err)
}

func TestListAll(t *testing.T) {
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "John Doe",
		Contact:   "1234567890",
		Address:   "123 Main Street",
		Reason:    "Fever",
		ConsultBy: "Dr. Smith",
	}

	err := CreatePatient(patient)
	assert.NoError(t, err)

	p, err := ListAll(10, 0, "")
	assert.NoError(t, err)
	assert.Len(t, p, 1)

	err = Delete(patient.Id)
	assert.NoError(t, err)
}

func TestGetPatient(t *testing.T){
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "John Doe",
		Contact:   "1234567890",
		Address:   "123 Main Street",
		Reason:    "Fever",
		ConsultBy: "Dr. Smith",
	}

	err := CreatePatient(patient)
	assert.NoError(t, err)

	p, err := GetPatient(patient.Id)
	assert.NoError(t, err)

	assert.Equal(t, patient.Name, p.Name)
	assert.Equal(t, patient.Contact, p.Contact)
	assert.Equal(t, patient.Address, p.Address)
	assert.Equal(t, patient.Reason, p.Reason)
	assert.Equal(t, patient.ConsultBy, p.ConsultBy)


	err = Delete(patient.Id)
	assert.NoError(t, err)


}

func TestUpdatePatient(t *testing.T) {
	
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "Original Name",
		Contact:   "1234567890",
		Address:   "Old Address",
		Reason:    "Cold",
		ConsultBy: "Dr. Smith",
	}

	err := CreatePatient(patient)
	assert.NoError(t, err)

	
	updatedPatient := model.Patient{
		Name:      "Updated Name",
		Contact:   "0987654321",
		Address:   "New Address",
		ConsultBy: "Dr. Adams",
	}

	
	err = Update(patient.Id, updatedPatient)
	assert.NoError(t, err)

	
	err = Delete(patient.Id)
	assert.NoError(t, err)
}

func TestCreatePatient_invalidContact(t *testing.T) {
	patient := model.Patient{
		Id:        uuid.NewString(),
		Name:      "veena",
		Contact:   "99999999999588888888888888888888888888888888",
		Address:   " Address",
		Reason:    "gas",
		ConsultBy: "Dr. reena",
	}
	err := CreatePatient(patient)
	assert.Error(t, err)
}
