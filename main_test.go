package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, Jenkins!"
    if output := hello(); output != expected {
        t.Errorf("Expected %q but got %q", expected, output)
    }
}

func hello() string {
    return "Hello, Jenkins!"
}
