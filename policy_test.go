package password

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PolicyTests struct {
	suite.Suite
}

func TestPolicies(t *testing.T) {
	suite.Run(t, &PolicyTests{})
}

func (s *PolicyTests) TestGetDefaultPolicy() {
	defaultPolicy := GetDefaultPolicy()
	s.Equal(DefaultPasswordMaximumLength, defaultPolicy.MaximumLength)
	s.Equal(DefaultPasswordMinimumLength, defaultPolicy.MinimumLength)
	s.Equal(DefaultPasswordMinimumLowercaseCount, defaultPolicy.MinimumLowercaseCount)
	s.Equal(DefaultPasswordMinimumUppercaseCount, defaultPolicy.MinimumUppercaseCount)
	s.Equal(DefaultPasswordMinimumNumericCount, defaultPolicy.MinimumNumericCount)
	s.Equal(DefaultPasswordMinimumSpecialCount, defaultPolicy.MinimumSpecialCount)
	s.Equal([]byte(DefaultPasswordCustomSpecial), defaultPolicy.CustomSpecial)
}

func (s *PolicyTests) TestValidate_minimumLength() {
	policy := Policy{
		MaximumLength: 2,
		MinimumLength: 2,
	}
	err := Validate("a", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "at least 2 characters")
	}
	err = Validate("aa", policy)
	s.Nil(err)
}

func (s *PolicyTests) TestValidate_maximumLength() {
	policy := Policy{MaximumLength: 2}
	err := Validate("aaa", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "maximum length of 2 characters")
	}
	err = Validate("aa", policy)
	s.Nil(err)
}

func (s *PolicyTests) TestValidate_lowercaseLength() {
	policy := Policy{
		MaximumLength:         4,
		MinimumLength:         4,
		MinimumLowercaseCount: 2,
	}
	err := Validate("aAAA", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "at least 2 lower-cased")
	}
	err = Validate("aaAA", policy)
	s.Nil(err)
}

func (s *PolicyTests) TestValidate_uppercaseLength() {
	policy := Policy{
		MaximumLength:         4,
		MinimumLength:         4,
		MinimumUppercaseCount: 2,
	}
	err := Validate("Aaaa", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "at least 2 upper-cased")
	}
	err = Validate("AAaa", policy)
	s.Nil(err)
}

func (s *PolicyTests) TestValidate_numericLength() {
	policy := Policy{
		MaximumLength:       4,
		MinimumLength:       4,
		MinimumNumericCount: 2,
	}
	err := Validate("1aaa", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "at least 2 numeric")
	}
	err = Validate("11aa", policy)
	s.Nil(err)
}

func (s *PolicyTests) TestValidate_specialLength() {
	policy := Policy{
		MaximumLength:       4,
		MinimumLength:       4,
		MinimumSpecialCount: 2,
	}
	err := Validate("!aaa", policy)
	s.NotNil(err)
	if err != nil {
		s.Contains(err.Error(), "at least 2 special")
	}
	err = Validate("!!aa", policy)
	s.Nil(err)
}
