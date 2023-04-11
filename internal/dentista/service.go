package dentista

import (
	"clinicaodontologica/internal/domain"
)

type IService interface {
	GetById(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Update(d domain.Dentista, id int) (domain.Dentista, error)
	Delete(id int) error
}

type Service struct {
	Repository IRepository
}

func NewService(Repository IRepository) IService {
	return &Service{Repository}
}

func (s *Service) GetById(id int) (domain.Dentista, error) {
	dent, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	return dent, nil
}

func (s *Service) Create(d domain.Dentista) (domain.Dentista, error) {
	dent, err := s.Repository.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dent, nil
}

func (s *Service) Update(d domain.Dentista, id int) (domain.Dentista, error) {
	dent, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if d.Apellido != "" {
		dent.Apellido = d.Apellido
	}
	if d.Nombre != "" {
		dent.Nombre = d.Nombre
	}
	if d.Matricula != "" {
		dent.Matricula = d.Matricula
	}

	dent, err = s.Repository.Update(dent, id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dent, nil
}

func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
