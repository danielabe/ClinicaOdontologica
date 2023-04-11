package store

import (
	"clinicaodontologica/internal/domain"
	"database/sql"
	"log"
)

type SqlStore struct {
	DB *sql.DB
}

// Dentista

func (s *SqlStore) ReadDentista(id int) (domain.Dentista, error) {
	var dentista domain.Dentista

	query := "SELECT * FROM dentistas WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentista.Id, &dentista.Apellido, &dentista.Nombre, &dentista.Matricula)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

func (s *SqlStore) CreateDentista(dentista domain.Dentista) (domain.Dentista, error) {
	query := "INSERT INTO dentistas (apellido, nombre, matricula) VALUES (?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	lid, _ := res.LastInsertId()
	dentista.Id = int(lid)

	return dentista, nil
}

func (s *SqlStore) UpdateDentista(dentista domain.Dentista, id int) (domain.Dentista, error) {
	stmt, err := s.DB.Prepare("UPDATE dentistas SET apellido = ?, nombre = ?, matricula = ? WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula, id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Dentista{}, err
	}

	return dentista, nil
}

func (s *SqlStore) DeleteDentista(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM dentistas WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	return nil
}

func (s *SqlStore) ReadDentistaByMatricula(matricula string) (int, error) {
	var id int

	query := "SELECT id FROM dentistas WHERE matricula = ?;"
	row := s.DB.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return 0, err
	}

	return id, nil
}

// Paciente

func (s *SqlStore) ReadPaciente(id int) (domain.Paciente, error) {
	var paciente domain.Paciente

	query := "SELECT * FROM pacientes WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&paciente.Id, &paciente.Nombre, &paciente.Apellido, &paciente.Domicilio, &paciente.DNI, &paciente.FechaAlta)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

func (s *SqlStore) CreatePaciente(paciente domain.Paciente) (domain.Paciente, error) {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaAlta)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	lid, _ := res.LastInsertId()
	paciente.Id = int(lid)

	return paciente, nil
}

func (s *SqlStore) UpdatePaciente(paciente domain.Paciente, id int) (domain.Paciente, error) {
	stmt, err := s.DB.Prepare("UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_alta = ? WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(paciente.Nombre, paciente.Apellido, paciente.Domicilio, paciente.DNI, paciente.FechaAlta, id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Paciente{}, err
	}

	return paciente, nil
}

func (s *SqlStore) DeletePaciente(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM pacientes WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	return nil
}

func (s *SqlStore) ReadIdPacienteByDNI(dni int) (int, error) {
	var id int

	query := "SELECT id FROM pacientes WHERE dni = ?;"
	row := s.DB.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return 0, err
	}

	return id, nil
}

// Turno

func (s *SqlStore) ReadTurno(id int) (domain.Turno, error) {
	var turno domain.Turno

	query := "SELECT t.id, p.*, d.*, t.fecha, t.hora, t.descripcion FROM turnos AS t INNER JOIN pacientes AS p ON t.paciente_id = p.id INNER JOIN dentistas AS d ON t.dentista_id = d.id WHERE t.id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&turno.Id,
		&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta,
		&turno.Dentista.Id, &turno.Dentista.Apellido, &turno.Dentista.Nombre, &turno.Dentista.Matricula,
		&turno.Fecha, &turno.Hora, &turno.Descripcion)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.Turno{}, err
	}

	return turno, nil
}

func (s *SqlStore) CreateTurno(turno domain.TurnoWithIds) (domain.TurnoWithIds, error) {
	query := "INSERT INTO turnos (paciente_id, dentista_id, fecha, hora, descripcion) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.TurnoWithIds{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(turno.PacienteId, turno.DentistaId, turno.Fecha, turno.Hora, turno.Descripcion)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.TurnoWithIds{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.TurnoWithIds{}, err
	}

	lid, _ := res.LastInsertId()
	turno.Id = int(lid)

	return turno, nil
}

func (s *SqlStore) UpdateTurno(turno domain.TurnoWithIds, id int) (domain.TurnoWithIds, error) {
	stmt, err := s.DB.Prepare("UPDATE turnos SET paciente_id = ?, dentista_id = ?, fecha = ?, hora = ?, descripcion = ? WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.TurnoWithIds{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(turno.PacienteId, turno.DentistaId, turno.Fecha, turno.Hora, turno.Descripcion, id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return domain.TurnoWithIds{}, err
	}

	return turno, nil
}

func (s *SqlStore) DeleteTurno(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM turnos WHERE id = ?")
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return err
	}

	return nil
}

func (s *SqlStore) ReadTurnosByDNI(dni int) ([]domain.Turno, error) {
	var turnos []domain.Turno

	query := "SELECT t.id, p.*, d.*, t.fecha, t.hora, t.descripcion FROM turnos AS t INNER JOIN pacientes AS p ON t.paciente_id = p.id INNER JOIN dentistas AS d ON t.dentista_id = d.id WHERE p.dni = ?;"
	rows, err := s.DB.Query(query, dni)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return []domain.Turno{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var turno domain.Turno
		err := rows.Scan(&turno.Id,
			&turno.Paciente.Id, &turno.Paciente.Nombre, &turno.Paciente.Apellido, &turno.Paciente.Domicilio, &turno.Paciente.DNI, &turno.Paciente.FechaAlta,
			&turno.Dentista.Id, &turno.Dentista.Apellido, &turno.Dentista.Nombre, &turno.Dentista.Matricula,
			&turno.Fecha, &turno.Hora, &turno.Descripcion)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			return []domain.Turno{}, err
		}
		turnos = append(turnos, turno)
	}
	if err = rows.Err(); err != nil {
		log.Printf("ERROR: %s\n", err)
		return []domain.Turno{}, err
	}

	return turnos, nil
}
