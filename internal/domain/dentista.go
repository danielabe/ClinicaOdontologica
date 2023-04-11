package domain

type Dentista struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido" binding:"required"`
	Nombre    string `json:"nombre" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}

type RequestDentista struct {
	Apellido  string `json:"apellido,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Matricula string `json:"matricula,omitempty"`
}

type DentistaInTurno struct {
	Id        int    `json:"id"`
	Apellido  string `json:"apellido,omitempty"`
	Nombre    string `json:"nombre,omitempty"`
	Matricula string `json:"matricula,omitempty"`
}
