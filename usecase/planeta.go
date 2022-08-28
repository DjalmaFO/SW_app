package usecase

import (
	"app/domain"
	"app/dto"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetaUsecase struct {
	PlanetaRepository domain.PlanetaRepository
}

func NewPlanetaUsecase(planetaRepository domain.PlanetaRepository) *PlanetaUsecase {
	return &PlanetaUsecase{PlanetaRepository: planetaRepository}
}

func (u *PlanetaUsecase) NewPlaneta(req *dto.PlanetaRequest) (r interface{}, err error) {
	if err = req.Sanetizar(); err != nil {
		log.Println(err.Error())
		return
	}

	planet := hydratePlanet(req)

	return u.PlanetaRepository.CreatePlanet(planet)
}

func hydratePlanet(req *dto.PlanetaRequest) (planet *domain.Planeta) {
	planet = domain.NewPlaneta()
	planet.ID = req.ID
	planet.Nome = req.Nome
	planet.Clima = req.Clima
	planet.Terreno = req.Terreno
	return
}

func (u *PlanetaUsecase) GetAll() (r interface{}, err error) {
	return u.PlanetaRepository.GetAll()
}

func (u *PlanetaUsecase) Filter(filter, value string) (r interface{}, err error) {
	return u.PlanetaRepository.GetByFilter(filter, value)
}

func (u *PlanetaUsecase) DeletePlanet(id primitive.ObjectID) error {
	return u.PlanetaRepository.DeletePlanetByID(id)
}

var ErrorInvalidID = fmt.Errorf("ID inválido")
