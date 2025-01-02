package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB é uma implementação personalizada para lidar com o tipo JSONB do PostgreSQL
type JSONB map[string]interface{}

// Scan implementa a interface sql.Scanner
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	// Convertendo o valor recebido para []byte
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("JSONB: Scan source is not []byte")
	}

	// Desserializando JSON para o mapa
	return json.Unmarshal(bytes, j)
}

// Value implementa a interface driver.Valuer
func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	// Serializando o mapa para JSON
	return json.Marshal(j)
}
