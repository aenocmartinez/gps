package dto

type NoteDto struct {
	Id         int64  `json:"id,omitempty"`
	Comment    string `json:"comment,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	RegisterBy string `json:"register_by,omitempty"`
}
