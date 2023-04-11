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
