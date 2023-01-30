package examples

import "github.com/deemount/gobpmn/utils"

// Hash ...
type Hash struct {
	string
}

func (h *Hash) Hash() string {
	if h.string == "" {
		h.string = utils.GenerateHash()
	}
	return h.string
}
