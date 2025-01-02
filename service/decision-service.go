package service

import (
	"ms-decision-go/config"
	"ms-decision-go/handler/requests"
	"ms-decision-go/repository"
)

type DecisionService struct {
	repo   *repository.DecisionRepository
	logger *config.Logger
}

func NewDecisionService() *DecisionService {
	return &DecisionService{
		repo:   repository.NewDecisionRepository(),
		logger: config.GetLogger("CreateDecisionHandler"),
	}
}

func (service *DecisionService) CreateDecision(decisionRequest requests.DecisionRequest) (string, error) {

	decision := requests.DecisionRequest{
		ID:         decisionRequest.ID,
		Document:   decisionRequest.Document,
		Product:    decisionRequest.Product,
		SubProduct: decisionRequest.SubProduct,
		Payload:    decisionRequest.Payload,
	}

	response, err := service.repo.FilterJobByProductAndSubProduct(decision.Product, decision.SubProduct)

	if err != nil {
		return "", err
	}

	if response == nil {
		return "Produto cp nao foi encontrado!", nil
	} else if response.Product == "100200300" && response.SubProduct == "4003" {
		return "Produto cp foi encontrado com sucesso!", nil
	} else {
		return "Produto cp nao foi encontrado!", nil
	}
}
