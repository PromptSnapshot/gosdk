package utils

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilTest struct {
	suite.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTest))
}
