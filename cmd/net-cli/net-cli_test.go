package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("./netcli", "ns", "--url", "google.com")

	var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

	result:= strings.Split(strings.TrimSpace(out.String()), "\n")

	t.Run("ns should return the correct nameservers for google.com", func(t *testing.T) {
		want := 4
		got := len(result)

		if got != want {
		t.Errorf("Expected %v received '%v'", want, got)
		}
	})

	t.Run("ns returned should contain the appropriate nameservers for google.com", func(t *testing.T) {
		got := result
		want := "ns1.google.com."

		fmt.Println(want, got)
		if got[2] != want {
			t.Errorf("Expected %v received '%v'", want, got)
		}
	})


}