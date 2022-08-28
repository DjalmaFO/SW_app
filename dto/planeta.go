package dto

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlanetaRequest struct {
	ID      primitive.ObjectID `form:"id" json:"id"`
	Nome    string             `form:"nome" json:"nome"`
	Clima   string             `form:"clima" json:"clima"`
	Terreno string             `form:"terreno" json:"terreno"`
}

// NewPlanetaRequest => Retona estrutura para Binding de informações aguardadas na requisição (via form-data ou JSON)
func NewPlanetaRequest() *PlanetaRequest {
	return &PlanetaRequest{}
}

func (pr *PlanetaRequest) Sanetizar() (err error) {
	if ok := pr.ValidateName(); !ok {
		return fmt.Errorf("Nome inválido ou indefinido")
	}

	if ok := pr.ValidateClimate(); !ok {
		return fmt.Errorf("Clima inválido ou indefinido")
	}

	if ok := pr.ValidateGround(); !ok {
		return fmt.Errorf("Terreno inválido ou indefinido")
	}
	return
}

func (pr *PlanetaRequest) ValidateName() bool {
	if len(strings.TrimSpace(pr.Nome)) == 0 {
		return false
	}

	return true
}

func (pr *PlanetaRequest) ValidateClimate() bool {
	if len(strings.TrimSpace(pr.Clima)) == 0 {
		return false
	}

	return true
}
func (pr *PlanetaRequest) ValidateGround() bool {
	if len(strings.TrimSpace(pr.Terreno)) == 0 {
		return false
	}

	return true
}
