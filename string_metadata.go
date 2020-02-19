package password

import (
	"bytes"
	"strings"
)

const (
	// StringTypeNumeric defines a numeric character type
	StringTypeNumeric = "NUMERIC"
	// StringTypeSpecial defines a special character type
	StringTypeSpecial = "SPECIAL"
	// StringTypeLowercase defines a lowercase character type
	StringTypeLowercase = "LOWERCASE"
	// StringTypeUppercase defines a uppercase character type
	StringTypeUppercase = "UPPERCASE"
	// StringTypeUnknown defines an unknown character type
	StringTypeUnknown = "UNKNOWN"
)

// StringMetadata provides metadata about a string, use GetStringMetadata to
// generate the metadata
type StringMetadata struct {
	Length     int
	Lowercases strings.Builder
	Uppercases strings.Builder
	Numerics   strings.Builder
	Specials   strings.Builder
	Unknowns   strings.Builder
	PrefixType string
	SuffixType string
}

// GetStringMetadata returns a populated StringMetadata structure
// that provides meta-data about the provided plaintext string for
// further processing by validators
func GetStringMetadata(plaintext string, customSpecial ...[]byte) StringMetadata {
	stringMetadata := StringMetadata{
		Length:     len(plaintext),
		Lowercases: strings.Builder{},
		Uppercases: strings.Builder{},
		Numerics:   strings.Builder{},
		Specials:   strings.Builder{},
		Unknowns:   strings.Builder{},
	}
	for i := 0; i < len(plaintext); i++ {
		currentCharacter := plaintext[i]
		specials := []byte(SpecialCharacters)
		if len(customSpecial) > 0 && len(customSpecial[0]) > 0 {
			specials = customSpecial[0]
		}
		var characterType string
		var builder *strings.Builder
		if bytes.Contains([]byte(specials), []byte{currentCharacter}) {
			characterType = StringTypeSpecial
			builder = &stringMetadata.Specials
		} else if bytes.Contains([]byte(LowercaseCharacters), []byte{currentCharacter}) {
			characterType = StringTypeLowercase
			builder = &stringMetadata.Lowercases
		} else if bytes.Contains([]byte(UppercaseCharacters), []byte{currentCharacter}) {
			characterType = StringTypeUppercase
			builder = &stringMetadata.Uppercases
		} else if bytes.Contains([]byte(NumericCharacters), []byte{currentCharacter}) {
			characterType = StringTypeNumeric
			builder = &stringMetadata.Numerics
		} else {
			characterType = StringTypeUnknown
			builder = &stringMetadata.Unknowns
		}
		if i == 0 {
			stringMetadata.PrefixType = characterType
		} else if i == len(plaintext)-1 {
			stringMetadata.SuffixType = characterType
		}
		builder.WriteByte(currentCharacter)
	}
	return stringMetadata
}
