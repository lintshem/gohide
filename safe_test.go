package main

import (
	"fmt"
	"strings"
	"testing"

	"com.secret.files/safe"
)

func TestEnc(t *testing.T) {
	ops := safe.Options{
		Mode:     "enc",
		Src:      "./main.go",
		Dest:     "./main.go.enc",
		Password: "test1234",
	}
	res := safe.Run(ops)
	fmt.Println(res)
}

func TestDec(t *testing.T) {
	ops := safe.Options{
		Mode:     "dec",
		Src:      "./main.go.enc",
		Dest:     "./main2.go",
		Password: "test1234",
	}
	res := safe.Run(ops)
	fmt.Println(res)

}

func TestHide(t *testing.T) {
	ops := safe.Options{
		Mode:     "hide",
		Src:      ".",
		Dest:     ".",
		Password: "test1234",
	}
	res := safe.Run(ops)
	fmt.Println(res)

}

func TestShow(t *testing.T) {
	ops := safe.Options{
		Mode:     "show",
		Src:      ".",
		Dest:     ".",
		Password: "test1234",
	}
	res := safe.Run(ops)
	fmt.Println(res)

}

func TestPrints(t *testing.T) {
	startDir := "."

	src_files, dest_file, err := safe.GetFilteredPaths(startDir)

	fmt.Println(strings.Join(src_files, "\n"))
	safe.Consume(src_files, dest_file, err)

	// p := safe.SafeRead()

}
