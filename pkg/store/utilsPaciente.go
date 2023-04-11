package store

import (
	"clinicaodontologica/internal/domain"
	"log"
)

func (s *SqlStore) ExistsPacienteId(id int) bool {
	pacientes, err := s.getPacientesByQuery("SELECT * FROM pacientes;")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, p := range pacientes {
		if p.Id == id {
			return true
		}
	}
	return false
}

func (s *SqlStore) ExistsPacienteDNI(dni int) bool {
	pacientes, err := s.getPacientesByQuery("SELECT * FROM pacientes;")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, p := range pacientes {
		if p.DNI == dni {
			return true
		}
	}
	return false
}

func (s *SqlStore) getPacientesByQuery(query string, args ...interface{}) ([]domain.Paciente, error) {
	var pacientes []domain.Paciente

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			return nil, err
		}
		pacientes = append(pacientes, paciente)
	}
	if err := rows.Err(); err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}

	return pacientes, nil
}
