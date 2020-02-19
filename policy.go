package password

import (
	"fmt"
)

const (
	// DefaultPasswordCustomSpecial defines the default character set
	// that is used to define special characters
	DefaultPasswordCustomSpecial string = SpecialCharacters

	// DefaultPasswordMinimumLowercaseCount defines the default for
	// the number of lower-cased characters required
	DefaultPasswordMinimumLowercaseCount int = 0

	// DefaultPasswordMinimumUppercaseCount defines the default for
	// the number of upper-cased characters required
	DefaultPasswordMinimumUppercaseCount int = 0

	// DefaultPasswordMaximumLength defines the default for the length
	// of the password
	DefaultPasswordMaximumLength int = 64

	// DefaultPasswordMinimumLength defines the default for the minimum
	// length of the password
	DefaultPasswordMinimumLength int = 8

	// DefaultPasswordMinimumNumericCount defines the default
	// number of numeric characters in the password
	DefaultPasswordMinimumNumericCount int = 0

	// DefaultPasswordMinimumSpecialCount defines the default
	// number of special characters in the password
	DefaultPasswordMinimumSpecialCount int = 0
)

// Policy defines possible configurations for password
// requirements
type Policy struct {
	MaximumLength         int
	MinimumLength         int
	MinimumLowercaseCount int
	MinimumUppercaseCount int
	MinimumNumericCount   int
	MinimumSpecialCount   int
	CustomSpecial         []byte
}

// GetDefaultPolicy returns a Policy with
// its values set to the default
func GetDefaultPolicy() Policy {
	return Policy{
		MaximumLength:         DefaultPasswordMaximumLength,
		MinimumLength:         DefaultPasswordMinimumLength,
		MinimumLowercaseCount: DefaultPasswordMinimumLowercaseCount,
		MinimumUppercaseCount: DefaultPasswordMinimumUppercaseCount,
		MinimumNumericCount:   DefaultPasswordMinimumNumericCount,
		MinimumSpecialCount:   DefaultPasswordMinimumSpecialCount,
		CustomSpecial:         []byte(DefaultPasswordCustomSpecial),
	}
}

// Validate validates a provided plaintext password using the
// default PasswordPolicy or a custom policy if it's provided
func Validate(plaintext string, customPolicy ...Policy) error {
	policy := GetDefaultPolicy()
	if len(customPolicy) > 0 {
		policy = customPolicy[0]
	}

	passwordMetadata := GetStringMetadata(plaintext, policy.CustomSpecial)
	switch true {
	case passwordMetadata.Length < policy.MinimumLength:
		return fmt.Errorf("provided password requires at least %v characters", policy.MinimumLength)
	case passwordMetadata.Length > policy.MaximumLength:
		return fmt.Errorf("provided password exceeds the maximum length of %v characters", policy.MaximumLength)
	case passwordMetadata.Lowercases.Len() < policy.MinimumLowercaseCount:
		return fmt.Errorf("provided password requires at least %v lower-cased characters", policy.MinimumLowercaseCount)
	case passwordMetadata.Uppercases.Len() < policy.MinimumUppercaseCount:
		return fmt.Errorf("provided password requires at least %v upper-cased characters", policy.MinimumUppercaseCount)
	case passwordMetadata.Numerics.Len() < policy.MinimumNumericCount:
		return fmt.Errorf("provided password requires at least %v numeric characters", policy.MinimumNumericCount)
	case passwordMetadata.Specials.Len() < policy.MinimumSpecialCount:
		return fmt.Errorf("provided password requires at least %v special characters", policy.MinimumSpecialCount)
	}
	return nil
}
