CREATE DATABASE  IF NOT EXISTS clinica_db;
USE clinica_db;

DROP TABLE IF EXISTS dentistas;
CREATE TABLE dentistas (
  id int NOT NULL AUTO_INCREMENT,
  apellido varchar(45) NOT NULL,
  nombre varchar(45) NOT NULL,
  matricula varchar(45) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS pacientes;
CREATE TABLE pacientes (
  id int NOT NULL AUTO_INCREMENT,
  nombre varchar(45) NOT NULL,
  apellido varchar(45) NOT NULL,
  domicilio varchar(45) NOT NULL,
  dni int NOT NULL,
  fecha_alta varchar(45) NOT NULL,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS turnos;
CREATE TABLE turnos (
  id int NOT NULL AUTO_INCREMENT,
  paciente_id int NOT NULL,
  dentista_id int NOT NULL,
  fecha varchar(10) NOT NULL,
  hora varchar(8) NOT NULL,
  descripcion text,
  PRIMARY KEY (id),
  KEY fk_turnos_dentistas_idx (dentista_id),
  KEY fk_turnos_pacientes_idx (paciente_id),
  CONSTRAINT fk_turnos_dentistas FOREIGN KEY (dentista_id) REFERENCES dentistas (id) ON DELETE RESTRICT,
  CONSTRAINT fk_turnos_pacientes FOREIGN KEY (paciente_id) REFERENCES pacientes (id) ON DELETE RESTRICT
);

INSERT INTO dentistas VALUES 
(1,'Potter','Harry','1234'),
(2,'Snape','Severus','5678'),
(3,'Lovegood','Luna','3456');

INSERT INTO pacientes VALUES 
(1,'Hermione','Granger','Sarmiento 123',12345678,'31-12-2022'),
(2,'Lord','Voldemort','Pasco 123',23456789,'13-10-2021'),
(3,'Ron','Weasley','Paraguay 234',34567890,'15-05-2020');

INSERT INTO turnos VALUES
(1,1,1,'01-05-2023','12:00','Lorem ipsum'),
(2,1,2,'03-07-2023','13:00','Lorem ipsum'),
(3,2,2,'04-07-2023','13:00','Lorem ipsum'),
(4,2,3,'09-07-2023','11:00','Lorem ipsum');