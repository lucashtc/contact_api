package app

// Contact ...
type Contact struct {
	Number string `json:"Number"`
	Person person `json:"Person`
}

type person struct {
	Name string `json:"Name"`
}
