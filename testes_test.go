package main

import (
	"app/validator"
	"testing"
)

func TestHaveContent(t *testing.T) {
	valid := validator.ValidatorTest{}
	valid.ValidatorTest = append(valid.ValidatorTest,
		validator.Validator{Value: "", Esperado: false},
		validator.Validator{Value: "alguma coisa", Esperado: true},
		validator.Validator{Value: "         ", Esperado: false},
		validator.Validator{Value: "outra_coisa", Esperado: true},
	)

	if ok := valid.ExecValidate(validator.HaveContent); !ok {
		t.Error("Falha no teste de validação de conteudo de strings")
		t.Fail()
	}
}
