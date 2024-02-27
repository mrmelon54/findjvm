package findjvm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFindPath(t *testing.T) {
	_, err := os.Stat("/usr/lib/jvm/temurin-17-jdk-amd64/bin/java")
	assert.NoError(t, err)
	jvmPath, err := FindPath("17")
	assert.NoError(t, err)
	assert.Equal(t, "/usr/lib/jvm/temurin-17-jdk-amd64/bin/java", jvmPath)
}

func TestGetVersion(t *testing.T) {
	_, err := os.Stat("/usr/lib/jvm/temurin-17-jdk-amd64/release")
	assert.NoError(t, err)
	version, err := GetVersion("/usr/lib/jvm/temurin-17-jdk-amd64")
	assert.NoError(t, err)
	assert.Equal(t, uint64(17), version.Major())
}
