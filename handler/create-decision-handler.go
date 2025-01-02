package handler

import (
	"ms-decision-go/config"
	"ms-decision-go/handler/requests"
	"ms-decision-go/service"

	"github.com/gin-gonic/gin"
)

type CreateDecisionHandler struct {
	serv   *service.DecisionService
	logger *config.Logger
}

func NewCreateDecisionHandler() *CreateDecisionHandler {
	return &CreateDecisionHandler{
		serv:   service.NewDecisionService(),
		logger: config.GetLogger("CreateDecisionHandler"),
	}
}

func (hdl *CreateDecisionHandler) CreateDecision(ctx *gin.Context) {

	var requestFrom requests.DecisionRequest

	err := ctx.BindJSON(&requestFrom)
	hdl.logger.Infof("Body: %v", err)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		hdl.logger.Errorf("Erro ao processar requisição: %v", err)
		return
	}

	response, err := hdl.serv.CreateDecision(requestFrom)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Erro ao salvar os dados no banco"})
		hdl.logger.Errorf("Erro ao salvar os dados no banco: %v", err)
		return
	}

	// Retorna sucesso
	ctx.JSON(200, gin.H{
		"message": "Dados salvos com sucesso!",
		"data":    response,
	})
}
