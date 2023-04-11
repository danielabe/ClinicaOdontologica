package main

import (
	"clinicaodontologica/cmd/docs"
	"clinicaodontologica/cmd/server/handler"
	"clinicaodontologica/internal/dentista"
	"clinicaodontologica/internal/paciente"
	"clinicaodontologica/internal/turno"
	"clinicaodontologica/pkg/middleware"
	"clinicaodontologica/pkg/store"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Clínica Odontológica
// @version 1.0
// @description This API Handle Dentistas, Pacientes y Turnos.

// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name Daniela Berardi
// @contact.email danielaberardi@live.com.ar
// @contact.url https://github.com/danielabe

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := godotenv.Load("./.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	portDb := os.Getenv("MYSQL_PORT")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:%s)/clinica_db", user, pass, portDb)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing)
	}

	storage := store.SqlStore{db}

	repoDentista := dentista.Repository{&storage}
	serviceDentista := dentista.Service{&repoDentista}
	dentistaHandler := handler.DentistaHandler{&serviceDentista}

	repoPaciente := paciente.Repository{&storage}
	servicePaciente := paciente.Service{&repoPaciente}
	pacienteHandler := handler.PacienteHandler{&servicePaciente}

	repoTurno := turno.Repository{&storage}
	serviceTurno := turno.Service{&repoTurno}
	turnoHandler := handler.TurnoHandler{&serviceTurno}

	r := gin.New()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })
	r.Use(gin.Recovery(), middleware.Logger())

	dentistas := r.Group("dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetById())
		dentistas.POST("", middleware.Authentication(), dentistaHandler.Post())
		dentistas.PUT(":id", middleware.Authentication(), dentistaHandler.Put())
		dentistas.PATCH(":id", middleware.Authentication(), dentistaHandler.Patch())
		dentistas.DELETE(":id", middleware.Authentication(), dentistaHandler.Delete())
	}

	pacientes := r.Group("pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetById())
		pacientes.POST("", middleware.Authentication(), pacienteHandler.Post())
		pacientes.PUT(":id", middleware.Authentication(), pacienteHandler.Put())
		pacientes.PATCH(":id", middleware.Authentication(), pacienteHandler.Patch())
		pacientes.DELETE(":id", middleware.Authentication(), pacienteHandler.Delete())
	}

	turnos := r.Group("turnos")
	{
		turnos.GET(":id", turnoHandler.GetById())
		turnos.POST("", middleware.Authentication(), turnoHandler.Post())
		turnos.PUT(":id", middleware.Authentication(), turnoHandler.Put())
		turnos.PATCH(":id", middleware.Authentication(), turnoHandler.Patch())
		turnos.DELETE(":id", middleware.Authentication(), turnoHandler.Delete())
		turnos.POST("/dniymatricula", middleware.Authentication(), turnoHandler.PostByDNIAndLicense())
		turnos.GET("", turnoHandler.GetByDNI())
	}

	r.Run(":8080")
}
