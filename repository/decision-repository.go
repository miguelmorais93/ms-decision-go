package repository

import (
	"database/sql"
	"fmt"
	"ms-decision-go/config"
	"ms-decision-go/model"
)

// DecisionRepository representa o repositório para a entidade DecisionUser
type DecisionRepository struct {
	db     *sql.DB
	logger *config.Logger
}

// NewDecisionRepository cria uma nova instância de DecisionRepository
func NewDecisionRepository() *DecisionRepository {
	return &DecisionRepository{
		db:     config.Getdatabase(),
		logger: config.GetLogger("DecisionRepository: "),
	}
}

// CreateDecisionUser insere um novo usuário no banco e retorna a entidade completa
func (r *DecisionRepository) CreateDecisionUser(decision model.DecisionUser) (*model.DecisionUser, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (id, document)
        VALUES ($1, $2)
        RETURNING id, document
    `, model.DecisionUser{}.TableName())

	var user model.DecisionUser
	err := r.db.QueryRow(query, decision.ID, decision.Document).Scan(&user.ID, &user.Document)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar os dados no banco: %w", err)
	}

	return &user, nil
}

// GetDecisionUserByID busca um usuário pelo ID
func (r *DecisionRepository) GetDecisionUserByID(id int64) (*model.DecisionUser, error) {
	query := fmt.Sprintf(`
        SELECT id, document
        FROM %s
        WHERE id = $1
    `, model.DecisionUser{}.TableName())

	var user model.DecisionUser
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Document)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Retorna nil se o registro não for encontrado
		}
		return nil, fmt.Errorf("erro ao buscar o usuário no banco: %w", err)
	}

	return &user, nil
}
