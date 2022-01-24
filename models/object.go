package models

type Object struct {
	Uid  string
	Data []byte
}

func (o Object) Add(Object) {}
func (o Object) Del()       {}
