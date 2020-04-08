package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

const (
	bannerHt = 94.0
	xIndent  = 40.0
	taxRate  = 0.9
)

// LineItem ...
type LineItem struct {
	UnitName       string
	PricePerUnit   int
	UnitsPurchased int
}

func main() {
	lineItems := []LineItem{
		{
			UnitName:       "Monitor",
			PricePerUnit:   100,
			UnitsPurchased: 50,
		},
		{
			UnitName:       "Keyboard",
			PricePerUnit:   10,
			UnitsPurchased: 50,
		},
		{
			UnitName:       "Mouse",
			PricePerUnit:   5,
			UnitsPurchased: 50,
		},
		{
			UnitName:       "Just a very long line for testing if our program can handle wrapping properly",
			PricePerUnit:   100,
			UnitsPurchased: 50,
		},
		{
			UnitName:       "Monitor",
			PricePerUnit:   100,
			UnitsPurchased: 50,
		},
		{
			UnitName:       "Monitor",
			PricePerUnit:   100,
			UnitsPurchased: 50,
		},
	}
	subtotal := 0
	for _, line := range lineItems {
		subtotal += line.PricePerUnit * line.UnitsPurchased
	}
	tax := int(float64(subtotal) * taxRate)
	total := subtotal + tax
	totalUSD := toUSD(total)

	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	pdf.SetFillColor(103, 60, 79)
	pdf.Polygon([]gofpdf.PointType{
		{0, 0},
		{w, 0},
		{w, bannerHt},
		{0, bannerHt * 0.9},
	}, "F")
	pdf.Polygon([]gofpdf.PointType{
		{0, h},
		{0, h - (bannerHt * 0.2)},
		{w, h - (bannerHt * 0.1)},
		{w, h},
	}, "F")

	pdf.SetFont("Arial", "B", 40)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt := pdf.GetFontSize()
	pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3.1, "INVOICE")

	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-2.0*124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "(123) 456-789\nlalala@lala.com\nlala.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

	pdf.SetFont("arial", "", 12)
	pdf.SetTextColor(255, 255, 255)
	_, lineHt = pdf.GetFontSize()
	pdf.MoveTo(w-xIndent-124.0, (bannerHt-(lineHt*1.5*3.0))/2.0)
	pdf.MultiCell(124.0, lineHt*1.5, "Jl Buntu\nabcd, JKT, ID\n12345", gofpdf.BorderNone, gofpdf.AlignRight, false)

	_, sy := summaryBlock(pdf, xIndent, bannerHt+lineHt*2.0, "Billed To", "Client Name", "Client Address", "City, State, Country", "Postal Code")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*2.0, "Invoice Number", "00000000123")
	summaryBlock(pdf, xIndent*2.0+lineHt*12.5, bannerHt+lineHt*6.25, "Date of Issue", "05/29/2018")

	x, y := w-xIndent-124.0, bannerHt+lineHt*2.25
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.CellFormat(124.0, lineHt, "Invoice Total", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x, y = x+2.0, y+lineHt*1.5
	pdf.MoveTo(x, y)
	pdf.SetFont("times", "", 48)
	_, lineHt = pdf.GetFontSize()
	alpha := 58
	pdf.SetTextColor(72+alpha, 42+alpha, 55+alpha)
	pdf.CellFormat(124.0, lineHt, totalUSD, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")

	if sy > y {
		y = sy
	}

	x, y = xIndent-20.0, y+30.0
	pdf.Rect(x, y, w-(xIndent*2.0)+40.0, 3.0, "F")

	pdf.SetFont("times", "", 14)
	_, lineHt = pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	x, y = xIndent-2.0, y+lineHt
	pdf.MoveTo(x, y)
	pdf.CellFormat(w/2.65+1.5, lineHt, "Description", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignLeft, false, 0, "")
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, "Price Per Unit", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, "Quantity", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, "Amount", gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	// drawGrid(pdf)
	y = y + lineHt
	for _, line := range lineItems {
		x, y = lineItem(pdf, x, y, line)
	}

	x, y = w/1.75, y+lineHt*2.25
	x, y = trailerLine(pdf, x, y, "Subtotal", subtotal)
	x, y = trailerLine(pdf, x, y, "Tax", tax)
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(x+20.0, y, x+220.0, y)
	y = y + lineHt*0.5
	x, y = trailerLine(pdf, x, y, "Total", total)

	err := pdf.OutputFileAndClose("p4.pdf")
	if err != nil {
		panic(err)
	}
}

func trailerLine(pdf *gofpdf.Fpdf, x, y float64, label string, amount int) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(180, 180, 180)
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, label, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(119.5, lineHt, toUSD(amount), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	y = y + lineHt*1.5
	return origX, y
}

func toUSD(cents int) string {
	centsStr := fmt.Sprintf("%d", cents%100)
	if len(centsStr) < 2 {
		centsStr = "0" + centsStr
	}
	return fmt.Sprintf("$%d.%s", cents/100, centsStr)
}

func lineItem(pdf *gofpdf.Fpdf, x, y float64, lineItem LineItem) (float64, float64) {
	origX := x
	w, _ := pdf.GetPageSize()
	pdf.SetFont("times", "", 14)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(50, 50, 50)
	pdf.MoveTo(x, y)
	x, y = xIndent-2.0, y+lineHt*.75
	pdf.MoveTo(x, y)
	pdf.MultiCell(w/2.65+1.5, lineHt, lineItem.UnitName, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	tmp := pdf.SplitLines([]byte(lineItem.UnitName), w/2.65+1.5)
	maxY := y + float64(len(tmp)-1)*lineHt
	x = x + w/2.65 + 1.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(100.0, lineHt, toUSD(lineItem.PricePerUnit), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = x + 100.0
	pdf.MoveTo(x, y)
	pdf.CellFormat(80.0, lineHt, fmt.Sprintf("%d", lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	x = w - xIndent - 2.0 - 119.5
	pdf.MoveTo(x, y)
	pdf.CellFormat(119.5, lineHt, toUSD(lineItem.PricePerUnit*lineItem.UnitsPurchased), gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignRight, false, 0, "")
	if maxY > y {
		y = maxY
	}
	y = y + lineHt*1.75
	pdf.SetDrawColor(180, 180, 180)
	pdf.Line(xIndent-10.0, y, w-xIndent+10.0, y)
	return origX, y
}

func summaryBlock(pdf *gofpdf.Fpdf, x, y float64, title string, data ...string) (float64, float64) {
	pdf.SetFont("times", "", 14)
	pdf.SetTextColor(180, 180, 180)
	_, lineHt := pdf.GetFontSize()
	y += lineHt
	pdf.Text(x, y, title)
	y += lineHt * 0.25
	pdf.SetTextColor(50, 50, 50)
	for _, str := range data {
		y += lineHt * 1.25
		pdf.Text(x, y, str)
	}
	return x, y
}

func drawGrid(pdf *gofpdf.Fpdf) {
	w, h := pdf.GetPageSize()
	pdf.SetFont("courier", "", 12)
	pdf.SetTextColor(80, 80, 80)
	pdf.SetDrawColor(200, 200, 200)
	for x := 0.0; x < w; x = x + (w / 20.0) {
		pdf.SetTextColor(200, 200, 200)
		pdf.Line(x, 0, x, h)
		_, lineHt := pdf.GetFontSize()
		pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
	}
	for y := 0.0; y < h; y = y + (w / 20.0) {
		if y < bannerHt*0.9 {
			pdf.SetTextColor(200, 200, 200)
		} else {
			pdf.SetTextColor(80, 80, 80)
		}
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
