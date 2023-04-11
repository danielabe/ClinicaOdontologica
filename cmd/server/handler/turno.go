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
