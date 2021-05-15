package utils

import (
	"github.com/civet148/log"
	"testing"
)

func TestUnixMS(t *testing.T) {
	log.Infof("UNIX MS [%d]", UnixMS())
}
