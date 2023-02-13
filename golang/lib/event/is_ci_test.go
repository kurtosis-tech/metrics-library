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

// This only runs on Circle
func TestISCI_PassesOnCircle(t *testing.T) {
	_, found := os.LookupEnv("CIRCLE_CI")
	if !found {
		t.Skip("Skipping as the environment isn't circle ci")
	}
	require.Equal(t, trueStr, isCI())
}
