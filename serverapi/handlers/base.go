package handlers

type Base struct {
}

func (base Base) Auth() string {
	return "success"
}
