package handler

import (
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

	// Salva os dados no banco usando GORM
	if err := db.Create(&decision).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao salvar os dados no banco"})
		logger.Errorf("Erro ao salvar os dados no banco: %v", err)
		return
	}

	// Retorna sucesso
	ctx.JSON(200, gin.H{
		"message": "Dados salvos com sucesso!",
		"data":    requestFrom,
	})
}
