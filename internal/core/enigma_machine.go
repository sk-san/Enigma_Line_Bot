package core

import "strings"

var alphabets = []string{"ekmflgdqvzntowyhxuspaibrcj",
	"ajdksiruxblhwtmcqgznpyfvoe",
	"bdfhjlcprtxvznyeiwgakmusqo",
}

type Enigma_machine struct {
	reflector Reflector
	rotors    []*Rotor
	plugb     PlugBoard
}

func (e *Enigma_machine) SetDefault(initChars string) {
	e.rotors = []*Rotor{}
	initChars = strings.ToLower(initChars)
	for i := 0; i < 3; i++ {
		r := &Rotor{}
		r.SetDefaults(alphabets[i])
		e.rotors = append(e.rotors, r)
	}

	e.reflector.SetDefault()
	e.plugb.SetDefault()

	for i := 0; i < len(initChars); i++ {
		rotor := e.rotors[i]
		rotor.SetInitialPosition(string(initChars[i]))
	}
}

func (e *Enigma_machine) RotateRotors() {
	for i := len(e.rotors) - 1; i >= 0; i-- {
		e.rotors[i].RotateChar()
		if i >= 0 && e.rotors[i].CurrentIndex != e.rotors[i].StartPosition {
			break
		}
	}
}

func (e Enigma_machine) Encrypt(OriginalText string) string {
	result := ""
	OriginalText = strings.ToLower(OriginalText)
	for i := range OriginalText {
		text := string(OriginalText[i])
		if text == " " {
			result += " "
			continue
		}
		e.RotateRotors()
		index, _ := e.plugb.GetIndex(text)
		for i := 2; i >= 0; i-- {
			index, _ = e.rotors[i].GetLeftCharIndex(index)
		}

		index = e.reflector.Reflect(index)

		for i := 0; i < 3; i++ {
			index, _ = e.rotors[i].GetRightCharIndex(index)
		}
		result += string(e.plugb.boardAlphabet[index])
	}
	return result
}

func (e Enigma_machine) Decrypt(EncryptedText string) string {
	result := ""
	EncryptedText = strings.ToLower(EncryptedText)
	for i := range EncryptedText {
		text := string(EncryptedText[i])
		if text == " " {
			result += " "
			continue
		}
		e.RotateRotors()
		index, _ := e.plugb.GetIndex(text)
		for i := 2; i >= 0; i-- {
			index, _ = e.rotors[i].GetLeftCharIndex(index)
		}

		index = e.reflector.Reflect(index)
		for i := 0; i < 3; i++ {
			index, _ = e.rotors[i].GetRightCharIndex(index)
		}
		result += string(e.plugb.boardAlphabet[index])
	}
	return result
}

func VerifyValidity(text string) bool {

	return false
}
