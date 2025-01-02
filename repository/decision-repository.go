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

// GetDecisionUserByID busca um usuário pelo ID
func (r *DecisionRepository) FilterJobByProductAndSubProduct(product string, subProduct string) (*model.JobFilter, error) {
	query := fmt.Sprintf(`
        SELECT *
        FROM %s
        WHERE product = $1 and sub_product = $2
    `, model.JobFilter{}.TableName())

	var jobFiltered model.JobFilter
	err := r.db.QueryRow(query, product, subProduct).Scan(
		&jobFiltered.ID,
		&jobFiltered.Product,
		&jobFiltered.SubProduct,
		&jobFiltered.PayloadValidation,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Errorf("Sem registro no banco de dados")
			return nil, nil // Retorna nil se o registro não for encontrado
		}
		return nil, fmt.Errorf("erro ao buscar o usuário no banco: %w", err)
	}

	return &jobFiltered, nil
}
