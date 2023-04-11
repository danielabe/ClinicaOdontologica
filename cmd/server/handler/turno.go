package handler

import (
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/internal/turno"
	"clinicaodontologica/pkg/web"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TurnoHandler struct {
	TurnoService turno.IService
}

func NewHandlerTurno(service turno.IService) *TurnoHandler {
	return &TurnoHandler{
		TurnoService: service,
	}
}

// @Summary Get turno by id
// @Description Get turno by id
// @Tags Turnos
// @Produce json
// @Param id path int true "Turno ID"
// @Success 200 {object} domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [get]
func (h *TurnoHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		turnoFound, errFound := h.TurnoService.GetById(id)
		if errFound != nil {
			log.Printf("ERROR: %s\n", errFound)
			web.Failure(ctx, http.StatusNotFound, errors.New("Turno not found"))
			return
		}
		web.Success(ctx, http.StatusOK, turnoFound)
	}
}

// @Summary Create turno
// @Description Create a new turno in repository
// @Tags Turnos
// @Produce json
// @Param token header string true "token"
// @Param body body domain.TurnoWithIds true "TurnoWithIds"
// @Success 201 {object} domain.TurnoWithIds
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos [post]
func (h *TurnoHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var turno domain.TurnoWithIds
		err := ctx.ShouldBindJSON(&turno)
		if err != nil {
			fmt.Println(err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		turnoCreated, err := h.TurnoService.Create(turno)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, turnoCreated)
	}
}

// @Summary Update turno
// @Description Update turno in repository
// @Tags Turnos
// @Produce json
// @Param id path int true "Turno ID"
// @Param token header string true "token"
// @Param body body domain.TurnoWithIds true "TurnoWithIds"
// @Success 200 {object} domain.TurnoWithIds
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [put]
func (h *TurnoHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		var turno domain.TurnoWithIds
		err = ctx.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		turnoUpdated, err := h.TurnoService.Update(turno, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, turnoUpdated)
	}
}

// @Summary Modify turno fields
// @Description Modify turno fields in repository
// @Tags Turnos
// @Produce json
// @Param id path int true "Turno ID"
// @Param token header string true "token"
// @Param body body domain.RequestTurno true "RequestTurno"
// @Success 200 {object} domain.TurnoWithIds
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [patch]
func (h *TurnoHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r domain.RequestTurno
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&r)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}
		update := domain.TurnoWithIds{
			PacienteId:  r.PacienteId,
			DentistaId:  r.DentistaId,
			Fecha:       r.Fecha,
			Hora:        r.Hora,
			Descripcion: r.Descripcion,
		}
		t, err := h.TurnoService.Update(update, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, t)
	}
}

// @Summary Delete turno
// @Description Delete turno in repository
// @Tags Turnos
// @Produce json
// @Param id path int true "Turno ID"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/{id} [delete]
func (h *TurnoHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}
		err = h.TurnoService.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusNoContent, nil)
	}
}

// @Summary Create turno by paciente dni and dentista matricula
// @Description Create a new turno in repository
// @Tags Turnos
// @Produce json
// @Param token header string true "token"
// @Param dni query int true "Paciente dni"
// @Param matricula query string true "Dentista matricula"
// @Param body body domain.TurnoForDNIAndLicense true "TurnoForDNIAndLicense"
// @Success 201 {object} domain.TurnoWithIds
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos/dniymatricula [post]
func (h *TurnoHandler) PostByDNIAndLicense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		matriculaParam := ctx.Query("dni")
		dni, err := strconv.Atoi(matriculaParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid dni"))
			return
		}

		matricula := ctx.Query("matricula")
		var turno domain.TurnoForDNIAndLicense
		err = ctx.ShouldBindJSON(&turno)
		if err != nil {
			fmt.Println(err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		turnoCreated, err := h.TurnoService.CreateByDNIAndLicense(turno, dni, matricula)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, turnoCreated)
	}
}

// @Summary Get turnos by dni
// @Description Get turnos by paciente dni
// @Tags Turnos
// @Produce json
// @Param dni query int true "Paciente dni"
// @Success 200 {object} []domain.Turno
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /turnos [get]
func (h *TurnoHandler) GetByDNI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dniParam := ctx.Query("dni")
		dni, err := strconv.Atoi(dniParam)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid dni"))
			return
		}

		turnosFound, errFound := h.TurnoService.GetByDNI(dni)
		if errFound != nil {
			log.Printf("ERROR: %s\n", errFound)
			web.Failure(ctx, http.StatusNotFound, errFound)
			return
		}
		web.Success(ctx, http.StatusOK, turnosFound)
	}
}
