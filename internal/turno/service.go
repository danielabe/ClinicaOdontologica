package turno

import (
	"clinicaodontologica/internal/domain"
	"errors"
)

type IService interface {
	GetById(id int) (domain.Turno, error)
	Create(t domain.TurnoWithIds) (domain.TurnoWithIds, error)
	Update(p domain.TurnoWithIds, id int) (domain.TurnoWithIds, error)
	Delete(id int) error
	CreateByDNIAndLicense(turno domain.TurnoForDNIAndLicense, dni int, matricula string) (domain.TurnoWithIds, error)
	GetByDNI(dni int) ([]domain.Turno, error)
}

type Service struct {
	Repository IRepository
}

func NewService(Repository IRepository) IService {
	return &Service{Repository}
}

func (s *Service) GetById(id int) (domain.Turno, error) {
	tur, err := s.Repository.GetById(id)
	if err != nil {
		return domain.Turno{}, err
	}

	return tur, nil
}

func (s *Service) Create(t domain.TurnoWithIds) (domain.TurnoWithIds, error) {
	tur, err := s.Repository.Create(t)
	if err != nil {
		return domain.TurnoWithIds{}, err
	}
	return tur, nil
}

func (s *Service) Update(t domain.TurnoWithIds, id int) (domain.TurnoWithIds, error) {
	if !s.Repository.ValidatePacienteId(t.PacienteId) {
		return domain.TurnoWithIds{}, errors.New("Paciente id does not exist")
	}
	if !s.Repository.ValidateDentistaId(t.DentistaId) {
		return domain.TurnoWithIds{}, errors.New("Dentista id does not exist")
	}

	tur, err := s.Repository.GetById(id)
	if err != nil {
		return domain.TurnoWithIds{}, err
	}
	turno := domain.TurnoWithIds{
		Id:          tur.Id,
		PacienteId:  tur.Paciente.Id,
		DentistaId:  tur.Dentista.Id,
		Fecha:       tur.Fecha,
		Hora:        tur.Hora,
		Descripcion: tur.Descripcion,
	}
	if t.PacienteId != 0 {
		turno.PacienteId = t.PacienteId
	}
	if t.DentistaId != 0 {
		turno.DentistaId = t.DentistaId
	}
	if t.Fecha != "" {
		turno.Fecha = t.Fecha
	}
	if t.Hora != "" {
		turno.Hora = t.Hora
	}
	if t.Descripcion != "" {
		turno.Descripcion = t.Descripcion
	}

	turno, err = s.Repository.Update(turno, id)
	if err != nil {
		return domain.TurnoWithIds{}, err
	}
	return turno, nil
}

func (s *Service) Delete(id int) error {
	err := s.Repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CreateByDNIAndLicense(t domain.TurnoForDNIAndLicense, dni int, matricula string) (domain.TurnoWithIds, error) {
	tur, err := s.Repository.CreateByDNIAndLicense(t, dni, matricula)
	if err != nil {
		return domain.TurnoWithIds{}, err
	}
	return tur, nil
}

func (s *Service) GetByDNI(dni int) ([]domain.Turno, error) {
	turnos, err := s.Repository.GetByDNI(dni)
	if err != nil {
		return []domain.Turno{}, err
	}

	return turnos, nil
}
