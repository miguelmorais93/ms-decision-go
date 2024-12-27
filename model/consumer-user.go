package model

type DecisionUser struct {
	ID       int64
	Document string
}

func (DecisionUser) TableName() string {
	return "consumer_user"
}
