package handler

import (
	"clinicaodontologica/internal/dentista"
	"clinicaodontologica/internal/domain"
	"clinicaodontologica/pkg/web"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DentistaHandler struct {
	DentistaService dentista.IService
}

func NewHandlerDentista(service dentista.IService) *DentistaHandler {
	return &DentistaHandler{
		DentistaService: service,
	}
}

// @Summary Get dentista by id
// @Description Get dentista by id
// @Tags Dentistas
// @Produce json
// @Param id path int true "Dentista ID"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [get]
func (h *DentistaHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		dentistaFound, errFound := h.DentistaService.GetById(id)
		if errFound != nil {
			log.Printf("ERROR: %s\n", errFound)
			web.Failure(ctx, http.StatusNotFound, errors.New("Dentista not found"))
			return
		}
		web.Success(ctx, http.StatusOK, dentistaFound)
	}
}

// @Summary Create dentista
// @Description Create a new dentista in repository
// @Tags Dentistas
// @Produce json
// @Param token header string true "token"
// @Param body body domain.Dentista true "Dentista"
// @Success 201 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas [post]
func (h *DentistaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentista domain.Dentista
		err := ctx.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		dentistaCreated, err := h.DentistaService.Create(dentista)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, dentistaCreated)
	}
}

// @Summary Update dentista
// @Description Update dentista in repository
// @Tags Dentistas
// @Produce json
// @Param id path int true "Dentista ID"
// @Param token header string true "token"
// @Param body body domain.Dentista true "Dentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [put]
func (h *DentistaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}

		var dentista domain.Dentista
		err = ctx.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid json"))
			return
		}

		dentistaUpdated, err := h.DentistaService.Update(dentista, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, dentistaUpdated)
	}
}

// @Summary Modify dentista fields
// @Description Modify dentista fields in repository
// @Tags Dentistas
// @Produce json
// @Param id path int true "Dentista ID"
// @Param token header string true "token"
// @Param body body domain.RequestDentista true "RequestDentista"
// @Success 200 {object} domain.Dentista
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [patch]
func (h *DentistaHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var r domain.RequestDentista
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
		update := domain.Dentista{
			Apellido:  r.Apellido,
			Nombre:    r.Nombre,
			Matricula: r.Matricula,
		}
		d, err := h.DentistaService.Update(update, id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, d)
	}
}

// @Summary Delete dentista
// @Description Delete dentista in repository
// @Tags Dentistas
// @Produce json
// @Param id path int true "Dentista ID"
// @Param token header string true "token"
// @Success 204
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Router /dentistas/{id} [delete]
func (h *DentistaHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("Invalid id"))
			return
		}
		err = h.DentistaService.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusNoContent, nil)
	}
}
