package store

import (
	"clinicaodontologica/internal/domain"
)

type StoreInterface interface {
	// Dentistas
	ReadDentista(id int) (domain.Dentista, error)
	CreateDentista(dentista domain.Dentista) (domain.Dentista, error)
	UpdateDentista(dentista domain.Dentista, id int) (domain.Dentista, error)
	DeleteDentista(id int) error
	ReadDentistaByMatricula(matricula string) (int, error)
	ExistsMatricula(matricula string) bool
	ExistsMatriculaWithOtherId(matricula string, id int) bool
	ExistsDentistaId(id int) bool

	// Pacientes
	ReadPaciente(id int) (domain.Paciente, error)
	CreatePaciente(paciente domain.Paciente) (domain.Paciente, error)
	UpdatePaciente(paciente domain.Paciente, id int) (domain.Paciente, error)
	DeletePaciente(id int) error
	ReadIdPacienteByDNI(dni int) (int, error)
	ExistsPacienteId(id int) bool
	ExistsPacienteDNI(dni int) bool

	// Turnos
	ReadTurno(id int) (domain.Turno, error)
	CreateTurno(turno domain.TurnoWithIds) (domain.TurnoWithIds, error)
	UpdateTurno(turno domain.TurnoWithIds, id int) (domain.TurnoWithIds, error)
	DeleteTurno(id int) error
	ReadTurnosByDNI(dni int) ([]domain.Turno, error)
	ExistsTurnoId(id int) bool
}
