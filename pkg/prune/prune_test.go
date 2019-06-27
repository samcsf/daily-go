package prune

import (
	"testing"
	"os"
)

func TestPrune(t *testing.T) {
	os.Chdir("../../")
	root, _ := os.Getwd()
	Prune(root)
}
