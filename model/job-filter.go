package model

type JobFilter struct {
	ID                int64
	Product           string
	SubProduct        string
	PayloadValidation JSONB
}

func (JobFilter) TableName() string {
	return "job_filter"
}
