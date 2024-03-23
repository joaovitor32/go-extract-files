package main

import (
	helper "go-extract-files/internal/helper/zip"
)

func main() {
	helper.UnzipAll("testdata/teste.zip", "output")
}
