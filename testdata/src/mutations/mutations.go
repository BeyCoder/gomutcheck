package mutations

type MyStruct struct {
	TestField string
}

func (s *MyStruct) MutateWithPointer() {
	s.TestField = "new value"
}

func (s MyStruct) CallWithPointerMethod() {
	sPtr := &s
	sPtr.MutateWithPointer()
}

func (s MyStruct) CreateNewInstance() {
	newStruct := MyStruct{TestField: "new instance"}
	_ = newStruct
}

func (s MyStruct) MutateField() {
	s.TestField = "new value" // want "struct field 'TestField' is being mutated in value receiver method"
}

func (s MyStruct) ReadOnly() {
	_ = s.TestField
}

func (s MyStruct) CorrectUsage() {
	println(s.TestField)
}
