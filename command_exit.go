package main

import "os"

func commandExit(_ *memory) error {
	os.Exit(0)
	return nil
}
