package handler

import (
	"fmt"
	"ms-decision-go/handler/requests"
	"ms-decision-go/model"

	"github.com/gin-gonic/gin"
)

func CreateDecisionHandler(ctx *gin.Context) {

	var requestFrom requests.CreateDecisionRequest

	err := ctx.BindJSON(&requestFrom)
	logger.Infof("Body: %v", err)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		logger.Errorf("Erro ao processar requisição: %v", err)
		return
	}

	decision := model.DecisionUser{
		ID:       requestFrom.ID,
		Document: requestFrom.Document,
	}

	response, err := CreateDecisionUser(decision)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao salvar os dados no banco"})
		logger.Errorf("Erro ao salvar os dados no banco: %v", err)
		return
	}

	// Salva os dados no banco usando GORM
	/*	if err := db.Create(&decision).Error; err != nil {
			ctx.JSON(500, gin.H{"error": "Erro ao salvar os dados no banco"})
			logger.Errorf("Erro ao salvar os dados no banco: %v", err)
			return
		}
	*/
	// Retorna sucesso
	ctx.JSON(200, gin.H{
		"message": "Dados salvos com sucesso!",
		"data":    response,
	})
}

func CreateDecisionUser(decision model.DecisionUser) (*model.DecisionUser, error) {
	query := fmt.Sprintf(`
        INSERT INTO %s (id, document)
        VALUES ($1, $2)
        RETURNING id, document
    `, model.DecisionUser{}.TableName())

	var user model.DecisionUser
	err := db.QueryRow(query, decision.ID, decision.Document).Scan(&user.ID, &user.Document)
	if err != nil {
		logger.Errorf("Erro ao salvar os dados no banco: %v", err)
		return nil, err
	}

	return &user, nil
}
