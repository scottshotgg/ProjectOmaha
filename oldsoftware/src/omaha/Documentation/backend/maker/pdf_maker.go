package main

import "github.com/jung-kurt/gofpdf"
import "log"

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")

	gofpdf_test.ExampleFpdf_MultiCell()

	if err != nil {
		log.Println("thats not cool")
	}
}re