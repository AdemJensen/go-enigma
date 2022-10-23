package machine

import (
	"enigma/errors"
	"enigma/machine/plug_board"
	"enigma/machine/reflector"
	"enigma/machine/rotor"
	"enigma/validations"
	"fmt"
	"unicode"
)

type EnigmaMachine struct {
	PlugBoard *plug_board.PlugBoard
	Rotors    []*rotor.Rotor // rotors from left to right
	Reflector *reflector.Reflector
}

const (
	DefaultNRotors = 3
)

func NewDefaultEnigmaMachine() *EnigmaMachine {
	var rotors []*rotor.Rotor
	for i := 0; i < DefaultNRotors; i++ {
		rotors = append(rotors, rotor.NewDefaultRotor(i+1))
	}
	return &EnigmaMachine{
		PlugBoard: plug_board.NewPlugBoard(),
		Rotors:    rotors,
		Reflector: reflector.NewDefaultReflector(),
	}
}

func (e *EnigmaMachine) IsValid() bool {
	return e.PlugBoard != nil && e.Reflector != nil
}

func (e *EnigmaMachine) TypeKey(k KeyboardKey) (rune, error) {
	if err := k.Validate(); err != nil {
		return 0, err
	}

	v := rune(k)

	// 1. spin rotors
	for i := len(e.Rotors) - 1; i >= 0; i-- {
		shouldMoveNext := e.Rotors[i].Spin()
		if !shouldMoveNext {
			break
		}
	}

	// 2. plug boards
	v = e.PlugBoard.Encode(v)

	// 3. wiring flow through rotors from right to left
	for i := len(e.Rotors) - 1; i >= 0; i-- {
		v = e.Rotors[i].EncodeRightToLeft(v)
	}

	// 4. reflector
	v = e.Reflector.Encode(v)

	// 5. wiring flow through rotors from left to right
	for i := 0; i < len(e.Rotors); i++ {
		v = e.Rotors[i].EncodeLeftToRight(v)
	}

	// 6. plug boards again
	v = e.PlugBoard.Encode(v)

	return v, nil
}

// GetRotorPosition gets rotor location from left to right.
// The position using rune 'A' to 'Z' representing 1-26
func (e *EnigmaMachine) GetRotorPosition() []rune {
	var res []rune
	for _, r := range e.Rotors {
		res = append(res, r.CurrentPosition())
	}
	return res
}

// SetRotorPosition set rotor position from left to right.
// The position using rune 'A' to 'Z' representing 1-26
func (e *EnigmaMachine) SetRotorPosition(rotorPosition []rune) error {
	if len(rotorPosition) != len(e.Rotors) {
		return fmt.Errorf("must provide %d rotor positions", len(e.Rotors))
	}
	for _, r := range rotorPosition {
		if !validations.IsValidRotorPosition(r) {
			return errors.NewInputError(errors.InputTypeRotorLocation, string(r))
		}
	}
	for i, r := range rotorPosition {
		_ = e.Rotors[i].SetPosition(r)
	}
	return nil
}

type KeyboardKey rune

func (k KeyboardKey) Validate() error {
	if !unicode.IsUpper(rune(k)) {
		return errors.NewInputError(errors.InputTypeKeyboardKey, string(k))
	}
	return nil
}
