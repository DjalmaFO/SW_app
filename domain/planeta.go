package domain

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planeta struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Nome      string             `bson:"nome"`
	Clima     string             `bson:"clima"`
	Terreno   string             `bson:"terreno"`
	Aparicoes int                `bson:"aparicoes"`
}

func NewPlaneta() *Planeta {
	return &Planeta{}
}

type PlanetaRepository interface {
	CreatePlanet(*Planeta) (interface{}, error)
	GetAll() (interface{}, error)
	GetByFilter(string, interface{}) (interface{}, error)
	DeletePlanetByID(primitive.ObjectID) error
}

var ErrPlanetExists = fmt.Errorf("Planeta já cadastrado")
