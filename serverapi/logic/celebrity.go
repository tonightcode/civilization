package logic

type Celebrity struct {
}

func (celebrity Celebrity) GetAll() map[string]string {
	return map[string]string{"total": "3", "list": ""}
}
