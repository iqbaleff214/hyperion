package util_test

import (
	"github.com/stretchr/testify/assert"
	"hyperion/backend/util"
	"strings"
	"testing"
)

func TestGenerateName(t *testing.T) {
	tests := []struct {
		name   string
		want   string
		prefix []string
	}{
		{"should return without prefix", "", []string{}},
		{"should return with prefix _v", "_v", []string{"_v"}},
		{"should return with prefix _c", "_c", []string{"_", "c"}},
		{"should return with prefix $", "$", []string{"$"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.GenerateName(tt.prefix...)

			assert.True(t, strings.HasPrefix(got, tt.want))
		})
	}
}
