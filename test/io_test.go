package tests

import (
	"testing"

	"github.com/masamerc/sevp/internal"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	contents string
	expected []string
}

func TestGetProfiles(t *testing.T) {
	testCases := []TestCase{
		{
			contents: "[default]\n",
			expected: []string{"default"},
		},
		{
			contents: "[profile default]\n[profile prod]\n",
			expected: []string{"default", "prod"},
		},
		{
			contents: "[default]\n[profile prod]\n[profile dev]\n",
			expected: []string{"default", "prod", "dev"},
		},
	}

	for _, tc := range testCases {
		actual := internal.GetProfiles(tc.contents)
		assert.Equal(t, tc.expected, actual)
	}
}
