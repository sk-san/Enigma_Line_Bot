package core

import (
	"strings"
)

type Rotor struct {
	LeftAlphabet  string
	RightAlphabet string
	CurrentIndex  int
	StartPosition int
}

func (r *Rotor) RotateChar() {
	r.CurrentIndex = (r.CurrentIndex + 1) % 26
	r.LeftAlphabet = r.LeftAlphabet[1:] + r.LeftAlphabet[:1]
	r.RightAlphabet = r.RightAlphabet[1:] + r.RightAlphabet[:1]
}

func (r *Rotor) SetDefaults(alphabet string) {
	r.RightAlphabet = alphabet
	r.LeftAlphabet = "abcdefghijklmnopqrstuvwxyz"
	r.CurrentIndex = 0
	r.StartPosition = 0
}

func (r *Rotor) SetInitialPosition(character string) {
	r.CurrentIndex = strings.Index(r.LeftAlphabet, character)
	r.LeftAlphabet = r.LeftAlphabet[r.CurrentIndex:] + r.LeftAlphabet[:r.CurrentIndex]
	r.RightAlphabet = r.RightAlphabet[r.CurrentIndex:] + r.RightAlphabet[:r.CurrentIndex]
	r.StartPosition = r.CurrentIndex
}

func (r Rotor) GetLeftCharIndex(charIdx int) (int, string) {
	TargetIdx := 0
	TargetChar := string(r.RightAlphabet[charIdx])
	for i := range r.LeftAlphabet {
		if TargetChar == string(r.LeftAlphabet[i]) {
			TargetIdx = i
			return TargetIdx, TargetChar
		}
	}
	return -1, ""
}

func (r Rotor) GetRightCharIndex(charIdx int) (int, string) {
	TargetIdx := 0
	TargetChar := string(r.LeftAlphabet[charIdx])
	for i := range r.RightAlphabet {
		if TargetChar == string(r.RightAlphabet[i]) {
			TargetIdx = i
			return TargetIdx, TargetChar
		}
	}
	return -1, ""
}
