package handler

import (
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/internal/paciente"
	"clinicaodontologica/pkg/web"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PacienteHandler struct {
	PacienteService paciente.IService
}

func NewHandlerPaciente(service paciente.IService) *PacienteHandler {
	return &PacienteHandler{
		PacienteService: service,
	}
}

// @Summary Get paciente by id
// @Description Get paciente by id
// @Tags Pacientes
// @Produce json
// @Param id path int true "Paciente ID"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [get]
func (h *PacienteHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		pacienteFound, errFound := h.PacienteService.GetById(id)
		if errFound != nil {
			log.Printf("ERROR: %s\n", errFound)
			web.Failure(ctx, http.StatusNotFound, errors.New("Paciente not found"))
			return
		}
		web.Success(ctx, http.StatusOK, pacienteFound)
	}
}

// @Summary Create paciente
// @Description Create a new paciente in repository
// @Tags Pacientes
// @Produce json
// @Param token header string true "token"
// @Param body body domain.Paciente true "Paciente"
// @Success 201 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes [post]
func (h *PacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente
		err := ctx.ShouldBindJSON(&paciente)
		if err != nil {
			fmt.Println(err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		pacienteCreated, err := h.PacienteService.Create(paciente)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, pacienteCreated)
	}
}

// @Summary Update paciente
// @Description Update paciente in repository
// @Tags Pacientes
// @Produce json
// @Param id path int true "Paciente ID"
// @Param token header string true "token"
// @Param body body domain.Paciente true "Paciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [put]
func (h *PacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		var paciente domain.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		pacienteUpdated, err := h.PacienteService.Update(paciente, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, pacienteUpdated)
	}
}

// @Summary Modify paciente fields
// @Description Modify paciente fields in repository
// @Tags Pacientes
// @Produce json
// @Param id path int true "Paciente ID"
// @Param token header string true "token"
// @Param body body domain.RequestPaciente true "RequestPaciente"
// @Success 200 {object} domain.Paciente
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [patch]
func (h *PacienteHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r domain.RequestPaciente
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
		update := domain.Paciente{
			Nombre:    r.Nombre,
			Apellido:  r.Apellido,
			Domicilio: r.Domicilio,
			DNI:       r.DNI,
			FechaAlta: r.FechaAlta,
		}
		p, err := h.PacienteService.Update(update, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, p)
	}
}

// @Summary Delete paciente
// @Description Delete paciente in repository
// @Tags Pacientes
// @Produce json
// @Param id path int true "Paciente ID"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /pacientes/{id} [delete]
func (h *PacienteHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}
		err = h.PacienteService.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusNoContent, nil)
	}
}
