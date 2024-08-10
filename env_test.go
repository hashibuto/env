package env

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ApplicationConfig struct {
	StringVariableDefault   string `env:"CONFIG_MY_STRING_DEFAULT,hello"`
	NonVariable             string
	StringVariableNoDefault string        `env:"CONFIG_MY_STRING"`
	IntVar                  int           `env:"CONFIG_MY_INT"`
	IntVarDef               int           `env:"CONFIG_MY_INT_DEFAULT,44"`
	FloatVar                float32       `env:"CONFIG_MY_FLOAT,32.3"`
	BoolVal                 bool          `env:"CONFIG_MY_BOOL"`
	BoolValDefault          bool          `env:"CONFIG_MY_BOOL_DEFAULT,t"`
	DurationM               time.Duration `env:"CONFIG_MY_MINUTES,10m,duration"`
	DurationS               time.Duration `env:"CONFIG_MY_SECONDS,10s,duration"`
	SliceOfStrings          []string      `env:"CONFIG_NAMES,james johnson|robin ramsey|silly squirrels"`
	EmptySlice              []string      `env:"CONFIG_NO_NAMES"`
}

func TestEnv(t *testing.T) {
	os.Setenv("CONFIG_NO_NAMES", "")
	err := os.Setenv("CONFIG_MY_STRING", "funny")
	assert.Nil(t, err)
	err = os.Setenv("CONFIG_MY_INT", "122")
	assert.Nil(t, err)
	err = os.Setenv("CONFIG_MY_BOOL", "true")
	assert.Nil(t, err)
	config := &ApplicationConfig{}
	err = Initialize(config)
	assert.Nil(t, err)
	assert.Equal(t, "funny", config.StringVariableNoDefault)
	assert.Equal(t, "hello", config.StringVariableDefault)
	assert.Equal(t, "", config.NonVariable)
	assert.Equal(t, 122, config.IntVar)
	assert.Equal(t, 44, config.IntVarDef)
	assert.Equal(t, float32(32.3), config.FloatVar)
	assert.Equal(t, true, config.BoolVal)
	assert.Equal(t, true, config.BoolValDefault)
	assert.Equal(t, time.Minute*10, config.DurationM)
	assert.Equal(t, time.Second*10, config.DurationS)
	assert.Len(t, config.SliceOfStrings, 3)
	assert.Equal(t, config.SliceOfStrings[0], "james johnson")
	assert.Equal(t, config.SliceOfStrings[1], "robin ramsey")
	assert.Len(t, config.EmptySlice, 0)
}
