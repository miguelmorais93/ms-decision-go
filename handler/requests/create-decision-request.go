package requests

type CreateDecisionRequest struct {
	ID       int64  `json:"id"`
	Document string `json:"document"`
}
