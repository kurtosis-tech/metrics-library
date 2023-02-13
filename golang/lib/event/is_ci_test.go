package event

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const envVarValueForTesting = "testing"

func TestISCIWhenEnvironmentVariableIsSetSucceeds(t *testing.T) {
	for _, envVar := range ciEnvironmentVariables {
		err := os.Setenv(envVar, envVarValueForTesting)
		require.Nil(t, err)
		require.Equal(t, trueStr, isCI())
		err = os.Unsetenv(envVar)
		require.Nil(t, err)
	}
}

func TestISCIWhenEnvironmentVariableIsNotSetFails(t *testing.T) {
	for _, envVar := range ciEnvironmentVariables {
		err := os.Unsetenv(envVar)
		require.Nil(t, err)
		require.Equal(t, falseStr, isCI())
	}
}
