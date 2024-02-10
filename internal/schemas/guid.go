package schemas

import "encoding/json"

type Guid struct{}

func NewGuid() Guid {
	return Guid{}
}

func (g Guid) ToString() string {
	return "259b74bf-cd70-45ad-b199-55d3e76dbdee"
}

func (g Guid) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.ToString())
}
