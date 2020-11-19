package java

import (
	"path/filepath"
	"testing"
)

func TestLoadJavaSourceFile(t *testing.T) {
	testFile := filepath.Join("samples", "HelloWorld.java")
	LoadJavaSourceFile(testFile)
}
