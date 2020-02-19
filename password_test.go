package password

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type PasswordTests struct {
	suite.Suite
}

func TestHashes(t *testing.T) {
	suite.Run(t, &PasswordTests{})
}

func (s *PasswordTests) TestEndToEnd() {
	expectedPlaintext := "hihihi"
	unexpectedPlaintexts := []string{
		"hihihh",
		"hihihj",
		"hihih",
		"hihihih",
	}
	hash, salt, _ := Hash(expectedPlaintext)
	err := Verify(expectedPlaintext, hash, salt)
	s.Nil(err)
	for _, otherPlaintext := range unexpectedPlaintexts {
		err = Verify(otherPlaintext, hash, salt)
		s.NotNil(err)
	}
}
