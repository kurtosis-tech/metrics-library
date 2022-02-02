package event

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const (
	someValueToTestHash = "some value some value"
	expectedHashedSomeValue = "f9b668fec50b7eebb1319e0a2cce08ee56e5b123ace456428342d4b60dc19968"

	anotherSentenceToTestHash = "It's not a simple random text."
	expectedHashedSentence = "029f21eeab056e2e98933d2b872d2ff365e5e837876534227d534fcdba248ac7"


	exampleContainerImageName    = "docker/getting-started"
	exampleContainerImageVersion = "latest"
	wrongExampleContainerImage   = "docker/getting-started:latest:3.2.5"
)

var (
	exampleContainerImage = strings.Join([]string{exampleContainerImageName, exampleContainerImageVersion} ,":")
)

func TestHashString(t *testing.T) {
	hashedValue := hashString(someValueToTestHash)

	require.Equal(t, expectedHashedSomeValue, hashedValue)

	hashedSentence := hashString(anotherSentenceToTestHash)

	require.Equal(t, expectedHashedSentence, hashedSentence)
}

func TestChekIfNotEmptyStringAndGetHashedValue(t *testing.T) {
	hashedValue, err := chekIfNotEmptyStringAndGetHashedValue(someValueToTestHash)

	require.Nil(t, err)
	require.Equal(t, expectedHashedSomeValue, hashedValue)

	_, shouldBeError := chekIfNotEmptyStringAndGetHashedValue("")

	require.Error(t, shouldBeError)
}

func TestSplitContainerImageIntoNameAndVersion(t *testing.T) {
	actualContainerImageName, actualContainerImageVersion, err := splitContainerImageIntoNameAndVersion(exampleContainerImage)

	require.Nil(t, err)
	require.Equal(t, exampleContainerImageName, actualContainerImageName)
	require.Equal(t, exampleContainerImageVersion, actualContainerImageVersion)

	_, _, shouldBeError := splitContainerImageIntoNameAndVersion(wrongExampleContainerImage)

	require.Error(t, shouldBeError)
}
