package people

type People struct {
	Name   string
	Age    uint8
	Gender string
}

func (p People) GetName() string {
	return p.Name
}

func (p *People) SetName(text string) {
	p.Name = text
}
