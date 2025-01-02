package requests

// ExampleTable representa a entidade da tabela example_table no banco de dados
type DecisionRequest struct {
	ID         int64  `json:"id"`
	Document   string `json:"document"`
	Product    string `json:"product"`
	SubProduct string `json:"sub_product"`
	Payload    JSONB  `json:"payload"`
}

// JSONB é uma implementação personalizada para lidar com o tipo JSONB do PostgreSQL
type JSONB map[string]interface{}
