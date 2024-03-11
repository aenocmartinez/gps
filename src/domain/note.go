package domain

import "abix/src/infraestructure/util"

type Note struct {
	id         int64
	comment    string
	createdAt  string
	registerBy string
}

func NewNote(comment string, registerBy string) *Note {
	return &Note{
		comment:    comment,
		registerBy: registerBy,
		createdAt:  util.CurrentDateFormatTimestamp(),
	}
}

func (n *Note) SetId(id int64) {
	n.id = id
}

func (n *Note) Id() int64 {
	return n.id
}

func (n *Note) SetComment(comment string) {
	n.comment = comment
}

func (n *Note) Comment() string {
	return n.comment
}

func (n *Note) SetCreatedAt(createdAt string) {
	n.createdAt = createdAt
}

func (n *Note) CreatedAt() string {
	return n.createdAt
}

func (n *Note) SetRegisterBy(registerBy string) {
	n.registerBy = registerBy
}

func (n *Note) RegisterBy() string {
	return n.registerBy
}

func (n *Note) Exists() bool {
	return n.id > 0
}
