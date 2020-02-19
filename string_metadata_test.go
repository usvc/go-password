package password

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StringMetadataTests struct {
	suite.Suite
}

func TestStringMetadata(t *testing.T) {
	suite.Run(t, &StringMetadataTests{})
}

func (s *StringMetadataTests) TestGetStringMetadata_base() {
	metadata := GetStringMetadata("aA1!")
	s.Equal(4, metadata.Length)
	s.Equal(1, metadata.Lowercases.Len())
	s.Equal(1, metadata.Uppercases.Len())
	s.Equal(1, metadata.Numerics.Len())
	s.Equal(1, metadata.Specials.Len())
}

func (s *StringMetadataTests) TestGetStringMetadata_induction() {
	metadata := GetStringMetadata("aaaAAA111!!!")
	s.Equal(12, metadata.Length)
	s.Equal(3, metadata.Lowercases.Len())
	s.Equal(3, metadata.Uppercases.Len())
	s.Equal(3, metadata.Numerics.Len())
	s.Equal(3, metadata.Specials.Len())
}

func (s *StringMetadataTests) TestGetStringMetadata_prefixSuffixAccuracy() {
	metadata := GetStringMetadata("aa")
	s.Equal(StringTypeLowercase, metadata.PrefixType)
	s.Equal(StringTypeLowercase, metadata.SuffixType)

	metadata = GetStringMetadata("aA")
	s.Equal(StringTypeLowercase, metadata.PrefixType)
	s.Equal(StringTypeUppercase, metadata.SuffixType)

	metadata = GetStringMetadata("Aa")
	s.Equal(StringTypeUppercase, metadata.PrefixType)
	s.Equal(StringTypeLowercase, metadata.SuffixType)

	metadata = GetStringMetadata("!a")
	s.Equal(StringTypeSpecial, metadata.PrefixType)
	s.Equal(StringTypeLowercase, metadata.SuffixType)

	metadata = GetStringMetadata("a!")
	s.Equal(StringTypeLowercase, metadata.PrefixType)
	s.Equal(StringTypeSpecial, metadata.SuffixType)

	metadata = GetStringMetadata("1a")
	s.Equal(StringTypeNumeric, metadata.PrefixType)
	s.Equal(StringTypeLowercase, metadata.SuffixType)

	metadata = GetStringMetadata("a1")
	s.Equal(StringTypeLowercase, metadata.PrefixType)
	s.Equal(StringTypeNumeric, metadata.SuffixType)
}
