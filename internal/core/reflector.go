package core

type Reflector struct {
	alphabets string
}

func (r Reflector) Reflect(CharIdx int) int {
	TargetChar := string(r.alphabets[CharIdx])
	TargetIdx := 0
	for i := range r.alphabets {
		if TargetChar == string(r.alphabets[i]) && CharIdx != i {
			TargetIdx = i
		}
	}
	return TargetIdx
}

func (r *Reflector) SetDefault() {
	r.alphabets = "abcdefgdijkgmkmiebftcvvjat"
}
