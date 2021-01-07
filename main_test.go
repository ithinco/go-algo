package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitcode := m.Run()
	os.Exit(exitcode)
}
