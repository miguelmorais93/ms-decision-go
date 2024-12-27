package model

import (
	"gorm.io/gorm"
)

// ExampleTable representa a entidade da tabela example_table no banco de dados
type DecisionUserResponse struct {
	ID       int64  `json:"id,omitempty"`
	Document string `json:"document"`
}

type DecisionUser struct {
	gorm.Model
	ID       int64
	Document string
}
