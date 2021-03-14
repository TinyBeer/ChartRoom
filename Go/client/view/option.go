package view

type Option struct {
	content string
	method  func()
}

func NewOption(content string, method func()) *Option {
	return &Option{
		content: content,
		method:  method,
	}
}

func (op *Option) Select() {
	if op.method != nil {
		return
	}
	op.method()
	return
}
