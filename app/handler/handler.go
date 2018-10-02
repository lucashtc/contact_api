package handler

// Contact ...
type Contact struct {
	Number string
	Person person
}

type person struct {
	Name string
}

// Create contact ...
func Create(c Contact) error {
	return nil
}

// GetContact ...
func GetContact(id string) Contact {
	var c Contact
	return c
}

// GetAll ...
func GetAll() []Contact {
	var c []Contact

	return c
}

// EditContact ...
func EditContact(c Contact) error {
	return nil
}

// DeleteContact ...
func DeleteContact() error {
	return nil
}
