package store

import (
	"clinicaodontologica/internal/domain"
	"log"
)

func (s *SqlStore) ExistsTurnoId(id int) bool {
	turnos, err := s.getTurnosByQuery("SELECT * FROM turnos;")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, t := range turnos {
		if t.Id == id {
			return true
		}
	}
	return false
}

func (s *SqlStore) getTurnosByQuery(query string, args ...interface{}) ([]domain.TurnoWithIds, error) {
	var turnos []domain.TurnoWithIds

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var turno domain.TurnoWithIds
		err := rows.Scan(&turno.Id, &turno.PacienteId, &turno.DentistaId, &turno.Fecha, &turno.Hora, &turno.Descripcion)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			return nil, err
		}
		turnos = append(turnos, turno)
	}
	if err := rows.Err(); err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}

	return turnos, nil
}
