package module

import (
	"encoding/json"

	"github.com/wzkun/aurora/utils/decode"
)

type Relation struct {
	ID   string
	Name string
}

// ToJSON function
func (o *Relation) ToJSON() json.RawMessage {
	js, _ := decode.JSON.Marshal(o)
	return js
}
