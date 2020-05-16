/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-05-12 11:57 
# @File : pdf.go
# @Description : 
# @Attention : 
*/
package utils

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/signintech/gopdf"
	"log"
	"testing"
)

func TestContainEmptyBlank(t *testing.T) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")

}

func TestConvertArrayStringToInt(t *testing.T) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(false)

	// Create a new input page from an URL
	page := wkhtmltopdf.NewPage("https://www.baidu.com/")
	// Set options for this page

	page.FooterRight.Set("[asddddddddd]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func TestHtml2PdfBySource(t *testing.T) {
	req := NewRemotetHtml2PdfPureSourceReq("https://www.baidu.com/", "/Users/joker/go/src/myLibrary/go-library/go/utils", "sss")
	resp, e := Html2PdfBySource(req)
	if nil != e {
		panic(e)
	}
	fmt.Println(resp)

}
