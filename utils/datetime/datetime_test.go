package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUnixTimeByYYYYMMDDhhmmss(t *testing.T) {
	var tests = []struct {
		A    string
		Want int64
	}{
		{"2022-05-13 00:00:00", 1652371200},
		{"2022-05-13 14:52:04", 1652424724},
		{"9022-05-13 23:59:59", 222551078399},
	}
	for _, test := range tests {
		timeTime, err := GetUnixTimeByYYYYMMDDhhmmss(test.A)
		assert.NoError(t, err)
		assert.Equal(t, test.Want, timeTime.Unix())
	}
}
