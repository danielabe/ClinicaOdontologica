package store

import (
	"clinicaodontologica/internal/domain"
	"log"
)

func (s *SqlStore) ExistsMatricula(matricula string) bool {
	dentistas, err := s.getDentistasByQuery("SELECT * FROM dentistas;")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, d := range dentistas {
		if d.Matricula == matricula {
			return true
		}
	}
	return false
}

func (s *SqlStore) ExistsMatriculaWithOtherId(matricula string, id int) bool {
	dentistas, err := s.getDentistasByQuery("SELECT * FROM dentistas WHERE id != ?;", id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, d := range dentistas {
		if d.Matricula == matricula {
			return true
		}
	}
	return false
}

func (s *SqlStore) ExistsDentistaId(id int) bool {
	dentistas, err := s.getDentistasByQuery("SELECT * FROM dentistas;")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return false
	}
	for _, d := range dentistas {
		if d.Id == id {
			return true
		}
	}
	return false
}

func (s *SqlStore) getDentistasByQuery(query string, args ...interface{}) ([]domain.Dentista, error) {
	var dentistas []domain.Dentista

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dentista domain.Dentista
		err := rows.Scan(&dentista.Id, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			return nil, err
		}
		dentistas = append(dentistas, dentista)
	}
	if err := rows.Err(); err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}

	return dentistas, nil
}
