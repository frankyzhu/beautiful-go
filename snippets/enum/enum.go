package enum

// take from: https://godoc.org/golang.org/x/tools/cmd/stringer

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
