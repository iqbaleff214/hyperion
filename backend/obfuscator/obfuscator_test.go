package obfuscator_test

import (
	"github.com/stretchr/testify/assert"
	obf "hyperion/backend/obfuscator"
	"testing"
)

func TestNewObfuscator(t *testing.T) {
	config := obf.Config{}
	obfuscator := obf.NewObfuscator(config)

	assert.NotNil(t, obfuscator)
	assert.IsType(t, obf.Obfuscator{}, obfuscator)
}

func Test_obfuscator_Obfuscate(t *testing.T) {

}
