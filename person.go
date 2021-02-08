package quickfix

var Name string

// Age defines the age var
var Age uint
var Hair string

type Person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
	Hair string `json:"hair"`
}

func GetHair(p Person) string {
	return p.Hair
}

func (p *Person) GetAge() uint {
	return p.Age
}

func (p Person) ValidateName() bool {
	return true
}
