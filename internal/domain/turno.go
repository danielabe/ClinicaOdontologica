package domain

type Turno struct {
	Id          int             `json:"id"`
	Paciente    PacienteInTurno `json:"paciente" binding:"required"`
	Dentista    DentistaInTurno `json:"dentista" binding:"required"`
	Fecha       string          `json:"fecha" binding:"required"`
	Hora        string          `json:"hora" binding:"required"`
	Descripcion string          `json:"descripcion" binding:"required"`
}

type RequestTurno struct {
	PacienteId  int    `json:"paciente_id,omitempty"`
	DentistaId  int    `json:"dentista_id,omitempty"`
	Fecha       string `json:"fecha,omitempty"`
	Hora        string `json:"hora,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
}

type TurnoWithIds struct {
	Id          int    `json:"id"`
	PacienteId  int    `json:"paciente_id" binding:"required"`
	DentistaId  int    `json:"dentista_id" binding:"required"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

type TurnoForDNIAndLicense struct {
	Id          int    `json:"id"`
	Fecha       string `json:"fecha" binding:"required"`
	Hora        string `json:"hora" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}
