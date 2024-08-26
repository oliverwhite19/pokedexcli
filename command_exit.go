package main

import "os"

func commandExit(_ *memory, _ ...string) error {
	os.Exit(0)
	return nil
}
