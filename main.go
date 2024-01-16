package main

import (
	helper "go-extract-files/pkg/helper/zip"
)

func main() {
	helper.UnzipAll("samples/teste.zip", "output")
}
