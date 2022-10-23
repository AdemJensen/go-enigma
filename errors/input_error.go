package errors

import "fmt"

type InputType string

const (
	InputTypeKeyboardKey   = "key"
	InputTypeRotorLocation = "rotor location"
	InputTypePlugBoardPin  = "plug board pin"
)

type InputError struct {
	inputType    InputType
	inputContent string
}

func NewInputError(typ InputType, cont string) error {
	return &InputError{
		inputType:    typ,
		inputContent: cont,
	}
}

func (e *InputError) Error() string {
	return fmt.Sprintf("'%s' is not a valid %s", e.inputContent, e.inputType)
}
