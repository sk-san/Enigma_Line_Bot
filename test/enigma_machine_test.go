package test

import (
	"Enigma/internal/core"
	"testing"
)

func TestEnigma_machine_Decrypt(t *testing.T) {
	tests := []struct {
		name          string
		EncryptedText string
		DecryptedText string
	}{
		//{name: "TEST1", EncryptedText: "", DecryptedText: ""},
		{name: "TEST1", EncryptedText: "dltbb dprej", DecryptedText: "hello japan"},
		{name: "TEST2", EncryptedText: "o rh ss o uzemhnp gnvm f eim yqav cxp e orp wv", DecryptedText: "i am on a seafood diet i see food and i eat it"},
		{name: "TEST3", EncryptedText: "t", DecryptedText: "a"},
		{name: "TEST4", EncryptedText: "iaoujy", DecryptedText: "orange"},
		{name: "TEST5", EncryptedText: "zlhssfpcejj", DecryptedText: "lemonbanana"},
		{name: "TEST6", EncryptedText: "jvvoh", DecryptedText: "Kyoto"},
		{name: "TEST7", EncryptedText: "gojcao", DecryptedText: "nvidia"},
		{name: "TEST8", EncryptedText: "tsnchbi", DecryptedText: "abcdefg"},
		{name: "TEST9", EncryptedText: "lcexlnlegetvruoihkaf", DecryptedText: "zzzzzzzzzzzzzzzzzzzz"},
		{name: "TEST10", EncryptedText: "xfgpqcdbao gmutfv", DecryptedText: "sidewinder missle"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := core.Enigma_machine{}
			e.SetDefault("mck")

			if got := e.Decrypt(tt.EncryptedText); got != tt.DecryptedText {
				t.Errorf("Decrypt() = %v, want %v", got, tt.DecryptedText)
			}
		})
	}
}

func TestEnigma_machine_Encrypt(t *testing.T) {
	tests := []struct {
		name      string
		inputText string
		want      string
	}{
		{name: "TEST1", inputText: "hello japan", want: "dltbb dprej"},
		{name: "TEST2", inputText: "o rh ss o uzemhnp gnvm f eim yqav cxp e orp wv", want: "i am on a seafood diet i see food and i eat it"},
		{name: "TEST3", inputText: "t", want: "a"},
		{name: "TEST4", inputText: "iaoujy", want: "orange"},
		{name: "TEST5", inputText: "zlhssfpcejj", want: "lemonbanana"},
		{name: "TEST6", inputText: "jvvoh", want: "Kyoto"},
		{name: "TEST7", inputText: "gojcao", want: "nvidia"},
		{name: "TEST8", inputText: "tsnchbi", want: "abcdefg"},
		{name: "TEST9", inputText: "lcexlnlegetvruoihkaf", want: "zzzzzzzzzzzzzzzzzzzz"},
		{name: "TEST10", inputText: "xfgpqcdbao gmutfv", want: "sidewinder missle"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := core.Enigma_machine{}
			e.SetDefault("mck")
			if got := e.Encrypt(tt.inputText); got != tt.want {
				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
