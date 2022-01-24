package models

type Node struct {
	Address     string
	PublicKey   []byte
	FreeStorage int64
}

func (n Node) Add(Node) {}
func (n Node) Del()     {}
