package dto

import (
	"app/validator"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetaRequest struct {
	ID      primitive.ObjectID `form:"id" json:"id"`
	Nome    string             `form:"nome" json:"nome"`
	Clima   string             `form:"clima" json:"clima"`
	Terreno string             `form:"terreno" json:"terreno"`
}

type PlanetaResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Nome      string             `json:"nome" bson:"nome"`
	Clima     string             `json:"clima" bson:"clima"`
	Terreno   string             `json:"terreno" bson:"terreno"`
	Aparicoes int                `json:"aparicoes" bson:"aparicoes"`
}

// NewPlanetaRequest => Retona estrutura para Binding de informações aguardadas na requisição (via form-data ou JSON)
func NewPlanetaRequest() *PlanetaRequest {
	return &PlanetaRequest{}
}

func (pr *PlanetaRequest) Sanetizar() (err error) {
	if ok := validator.HaveContent(pr.Nome); !ok {
		return fmt.Errorf("Nome inválido ou indefinido")
	}

	if ok := validator.HaveContent(pr.Clima); !ok {
		return fmt.Errorf("Clima inválido ou indefinido")
	}

	if ok := validator.HaveContent(pr.Terreno); !ok {
		return fmt.Errorf("Terreno inválido ou indefinido")
	}
	return
}
