package paciente

import (
	"clinicaodontologica/internal/domain"
)

type IService interface {
	GetById(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(p domain.Paciente, id int) (domain.Paciente, error)
	Delete(id int) error
}

type Service struct {
	Repository IRepository
}

func NewService(Repository IRepository) IService {
	return &Service{Repository}
}

func (s *Service) GetById(id int) (domain.Paciente, error) {
	pac, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	return pac, nil
}

func (s *Service) Create(p domain.Paciente) (domain.Paciente, error) {
	pac, err := s.Repository.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return pac, nil
}

func (s *Service) Update(p domain.Paciente, id int) (domain.Paciente, error) {
	pac, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if p.Nombre != "" {
		pac.Nombre = p.Nombre
	}
	if p.Apellido != "" {
		pac.Apellido = p.Apellido
	}
	if p.Domicilio != "" {
		pac.Domicilio = p.Domicilio
	}
	if p.DNI != 0 {
		pac.DNI = p.DNI
	}
	if p.FechaAlta != "" {
		pac.FechaAlta = p.FechaAlta
	}

	pac, err = s.Repository.Update(pac, id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return pac, nil
}

func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
