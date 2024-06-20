package core

import "strings"

type PlugBoard struct {
	boardAlphabet string
}

func (p *PlugBoard) SetDefault() {
	p.boardAlphabet = "abcdefghijklmnopqrstuvwxyz"
}

func (p PlugBoard) GetIndex(character string) (int, string) {
	return strings.Index(p.boardAlphabet, character), character
}
func (p PlugBoard) OutputChar(charIdx int) string {
	return string(p.boardAlphabet[charIdx])
}
