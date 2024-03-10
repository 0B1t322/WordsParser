package numeral_parser

// Numeral описывает числительно
type Numeral struct {
	Value        uint64
	Level        uint8
	IsMultiplier bool
}

// NewNumeral создает числительное
func NewNumeral(value uint64, level uint8, isMultiplier bool) Numeral {
	return Numeral{
		Value:        value,
		Level:        level,
		IsMultiplier: isMultiplier,
	}
}
