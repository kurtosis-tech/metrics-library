package event

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

const (
	someValueToTestHash     = "some value some value"
	expectedHashedSomeValue = "f9b668fec50b7eebb1319e0a2cce08ee56e5b123ace456428342d4b60dc19968"

	anotherSentenceToTestHash = "It's not a simple random text."
	expectedHashedSentence    = "029f21eeab056e2e98933d2b872d2ff365e5e837876534227d534fcdba248ac7"

	exampleContainerImageOrgAndRepo = "docker/getting-started"
	exampleContainerImageTag        = "1.2.3"
	wrongExampleContainerImage      = "docker/getting-started:latest:3.2.5"
)

var (
	exampleContainerImage = strings.Join([]string{exampleContainerImageOrgAndRepo, exampleContainerImageTag}, ":")
)

func TestHashString(t *testing.T) {
	hashedValue := hashString(someValueToTestHash)

	require.Equal(t, expectedHashedSomeValue, hashedValue)

	hashedSentence := hashString(anotherSentenceToTestHash)

	require.Equal(t, expectedHashedSentence, hashedSentence)
}

func TestBestEffortSplitContainerImageIntoOrgRepoAndVersion_ImageAndTag(t *testing.T) {
	actualContainerImageName, actualContainerImageVersion := bestEffortSplitContainerImageIntoOrgRepoAndVersion(exampleContainerImage)
	require.Equal(t, exampleContainerImageOrgAndRepo, actualContainerImageName)
	require.Equal(t, exampleContainerImageTag, actualContainerImageVersion)
}

func TestBestEffortSplitContainerImageIntoOrgRepoAndVersion_InvalidImage(t *testing.T) {
	orgAndRepo, tag := bestEffortSplitContainerImageIntoOrgRepoAndVersion(wrongExampleContainerImage)
	require.Equal(t, "", orgAndRepo)
	require.Equal(t, "", tag)
}

func TestBestEffortSplitContainerImageIntoOrgRepoAndVersion_NoTag(t *testing.T) {
	orgAndRepo, tag := bestEffortSplitContainerImageIntoOrgRepoAndVersion(exampleContainerImageOrgAndRepo)
	require.Equal(t, exampleContainerImageOrgAndRepo, orgAndRepo)
	require.Equal(t, dockerDefaultImageTag, tag)
}
