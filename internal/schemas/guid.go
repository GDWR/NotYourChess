package schemas

type Guid = string

func NewGuid() Guid {
	return Guid("259b74bf-cd70-45ad-b199-55d3e76dbdee")
}

func ParseGuid(s string) (Guid, error) {
	return Guid(s), nil
}
