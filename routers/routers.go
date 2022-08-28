package routers

import (
	"app/dto"
	"app/infra/repository"
	"app/usecase"
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

func RoutersConfig(e *echo.Echo, db *mongo.Client, ctx *context.Context) {
	// Todas as rotas estão sem necessidade de autenticação devido não haver evento de Login.
	e.POST("/planet/new", func(c echo.Context) error {
		request := dto.NewPlanetaRequest()
		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, fmt.Errorf("Dados inconsistentes: %s", err.Error()))
		}

		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		usecase := usecase.NewPlanetaUsecase(repo)
		r, err := usecase.NewPlaneta(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/planets/all", func(c echo.Context) error {
		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		usecase := usecase.NewPlanetaUsecase(repo)
		r, err := usecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/planets/id/:id", func(c echo.Context) error {
		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		usecase := usecase.NewPlanetaUsecase(repo)

		r, err := usecase.Filter("_id", c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/planets/name/:name", func(c echo.Context) error {
		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		usecase := usecase.NewPlanetaUsecase(repo)
		r, err := usecase.Filter("nome", c.Param("name"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	e.GET("/planets/name/:name", func(c echo.Context) error {
		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		usecase := usecase.NewPlanetaUsecase(repo)
		r, err := usecase.Filter("nome", c.Param("name"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, r)
	})

	e.DELETE("/planets/delete/:id", func(c echo.Context) error {
		repo := repository.NewPlanetaRepositoryDB(db, ctx)
		useCase := usecase.NewPlanetaUsecase(repo)
		id, err := usecase.StringToObjectID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = useCase.DeletePlanet(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, "Planeta removido com sucesso!")
	})
}
