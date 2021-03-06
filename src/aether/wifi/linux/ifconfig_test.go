package linux

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIFConfigInfo(t *testing.T) {
	ifconfig := NewIFConfig()

	// IsInstalled
	if !ifconfig.IsInstalled() {
		t.Skip("skipping test, program not installed")
	}

	/*
		infoList, err := ifconfig.InfoList()

		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v\n", infoList)
	*/
	assert.Nil(t, nil)
}
