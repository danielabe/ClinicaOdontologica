package turno

import (
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/pkg/store"
	"errors"
)

type IRepository interface {
	GetById(id int) (domain.Turno, error)
	Create(t domain.TurnoWithIds) (domain.TurnoWithIds, error)
	Update(t domain.TurnoWithIds, id int) (domain.TurnoWithIds, error)
	Delete(id int) error
	CreateByDNIAndLicense(t domain.TurnoForDNIAndLicense, dni int, matricula string) (domain.TurnoWithIds, error)
	GetByDNI(dni int) ([]domain.Turno, error)
	ValidatePacienteId(pacienteId int) bool
	ValidateDentistaId(dentistaId int) bool
}

type Repository struct {
	Store store.StoreInterface
}

func NewRepository(Store store.StoreInterface) IRepository {
	return &Repository{Store}
}

func (r *Repository) GetById(id int) (domain.Turno, error) {
	tur, err := r.Store.ReadTurno(id)
	if err != nil {
		return domain.Turno{}, errors.New("Turno not found")
	}

	return tur, nil
}

func (r *Repository) Create(t domain.TurnoWithIds) (domain.TurnoWithIds, error) {
	if !r.ValidatePacienteId(t.PacienteId) {
		return domain.TurnoWithIds{}, errors.New("Paciente id does not exist")
	}
	if !r.ValidateDentistaId(t.DentistaId) {
		return domain.TurnoWithIds{}, errors.New("Dentista id does not exist")
	}
	tur, err := r.Store.CreateTurno(t)
	if err != nil {
		return domain.TurnoWithIds{}, errors.New("Error creating turno")
	}
	return tur, nil
}

func (r *Repository) Update(t domain.TurnoWithIds, id int) (domain.TurnoWithIds, error) {
	tur, err := r.Store.UpdateTurno(t, id)
	if err != nil {
		return domain.TurnoWithIds{}, errors.New("Error updating turno")
	}
	return tur, nil
}

func (r *Repository) Delete(id int) error {
	if !r.Store.ExistsTurnoId(id) {
		return errors.New("Turno id does not exist")
	}
	err := r.Store.DeleteTurno(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateByDNIAndLicense(t domain.TurnoForDNIAndLicense, dni int, matricula string) (domain.TurnoWithIds, error) {
	idPaciente, err := r.Store.ReadIdPacienteByDNI(dni)
	if err != nil {
		return domain.TurnoWithIds{}, errors.New("Paciente dni does not exist")
	}
	idDentista, err := r.Store.ReadDentistaByMatricula(matricula)
	if err != nil {
		return domain.TurnoWithIds{}, errors.New("Dentista matricula does not exist")
	}

	tur := domain.TurnoWithIds{
		PacienteId:  idPaciente,
		DentistaId:  idDentista,
		Fecha:       t.Fecha,
		Hora:        t.Hora,
		Descripcion: t.Descripcion,
	}

	tur, err = r.Store.CreateTurno(tur)
	if err != nil {
		return domain.TurnoWithIds{}, errors.New("Error creating turno")
	}
	return tur, nil
}

func (r *Repository) GetByDNI(dni int) ([]domain.Turno, error) {
	if !r.Store.ExistsPacienteDNI(dni) {
		return nil, errors.New("Paciente dni does not exist")
	}
	turnos, err := r.Store.ReadTurnosByDNI(dni)
	if err != nil {
		return []domain.Turno{}, errors.New("Turnos not found")
	}

	return turnos, nil
}

func (r *Repository) ValidatePacienteId(pacienteId int) bool {
	if !r.Store.ExistsPacienteId(pacienteId) {
		return false
	}
	return true
}

func (r *Repository) ValidateDentistaId(dentistaId int) bool {
	if !r.Store.ExistsDentistaId(dentistaId) {
		return false
	}
	return true
}
