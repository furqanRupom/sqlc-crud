package model

type Option struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"user_id" db:"user_id"`
	SubjectType string `json:"subject" db:"subject_type"`
	SubjectID   string `json:"subject_id" db:"subject_id"`
	Option      string `json:"option" db:"option"`
	Rating      string `json:"rating" db:"rating"`
}
