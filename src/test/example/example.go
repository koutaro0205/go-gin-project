package example

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greeting(text string) string {
	return fmt.Sprintf("%s(%d)さん, %s", p.Name, p.Age, text)
}
