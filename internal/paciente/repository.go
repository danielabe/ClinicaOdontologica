package paciente

import (
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/pkg/store"
	"errors"
)

type IRepository interface {
	GetById(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(p domain.Paciente, id int) (domain.Paciente, error)
	Delete(id int) error
}

type Repository struct {
	Store store.StoreInterface
}

func NewRepository(Store store.StoreInterface) IRepository {
	return &Repository{Store}
}

func (r *Repository) GetById(id int) (domain.Paciente, error) {
	pac, err := r.Store.ReadPaciente(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found")
	}

	return pac, nil
}

func (r *Repository) Create(p domain.Paciente) (domain.Paciente, error) {
	if r.Store.ExistsPacienteDNI(p.DNI) {
		return domain.Paciente{}, errors.New("Dni value already exists")
	}
	pac, err := r.Store.CreatePaciente(p)
	if err != nil {
		return domain.Paciente{}, errors.New("Error creating paciente")
	}
	return pac, nil
}

func (r *Repository) Update(p domain.Paciente, id int) (domain.Paciente, error) {
	pac, err := r.Store.UpdatePaciente(p, id)
	if err != nil {
		return domain.Paciente{}, errors.New("Error updating paciente")
	}
	return pac, nil
}

func (r *Repository) Delete(id int) error {
	if !r.Store.ExistsPacienteId(id) {
		return errors.New("Paciente id does not exist")
	}
	err := r.Store.DeletePaciente(id)
	if err != nil {
		return err
	}

	return nil
}
