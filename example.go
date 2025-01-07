package gomutcheck

type ExampleStruct struct {
	ExampleField string
}

func (s ExampleStruct) MutateField() {
	s.ExampleField = "new value"
}
