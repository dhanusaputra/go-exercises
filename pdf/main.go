package main

import (
	"flag"
	"fmt"
	"image/color"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// PDFOption ...
type PDFOption func(*gofpdf.Fpdf)

// PDF ...
type PDF struct {
	fpdf *gofpdf.Fpdf
	x, y float64
}

func fillColor(c color.RGBA) PDFOption {
	return func(pdf *gofpdf.Fpdf) {
		r, g, b := rgb(c)
		pdf.SetFillColor(r, g, b)
	}
}

func rgb(c color.RGBA) (int, int, int) {
	alpha := float64(c.A) / 255.0
	alphaWhite := int(255 * (1.0 - alpha))
	r := int(float64(c.R)*alpha) + alphaWhite
	g := int(float64(c.G)*alpha) + alphaWhite
	b := int(float64(c.B)*alpha) + alphaWhite
	return r, g, b
}

func (p *PDF) move(xDelta, yDelta float64) {
	p.x, p.y = p.x+xDelta, p.y+yDelta
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) moveAbs(x, y float64) {
	p.x, p.y = x, y
	p.fpdf.MoveTo(p.x, p.y)
}

func (p *PDF) text(text string) {
	p.fpdf.Text(p.x, p.y, text)
}

func (p *PDF) polygon(pts []gofpdf.PointType, opts ...PDFOption) {
	for _, opt := range opts {
		opt(p.fpdf)
	}
	p.fpdf.Polygon(pts, "F")
}

func main() {
	name := flag.String("name", "", "the name of the person who completed the course")
	flag.Parse()
	fpdf := gofpdf.New(gofpdf.OrientationLandscape, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := fpdf.GetPageSize()
	fpdf.AddPage()
	pdf := PDF{
		fpdf: fpdf,
	}

	primary := color.RGBA{103, 60, 79, 255}
	secondary := color.RGBA{103, 60, 79, 220}

	pdf.polygon([]gofpdf.PointType{
		{0, 0},
		{0, h / 9.0},
		{w - (w / 6.0), 0},
	}, fillColor(secondary))
	pdf.polygon([]gofpdf.PointType{
		{w / 6.0, 0},
		{w, 0},
		{w, h / 9.0},
	}, fillColor(primary))
	pdf.polygon([]gofpdf.PointType{
		{w, h},
		{w, h - (h / 8.0)},
		{0, h},
	}, fillColor(secondary))
	pdf.polygon([]gofpdf.PointType{
		{0, h},
		{0, h - h/8.0},
		{w - (w / 6), h},
	}, fillColor(primary))

	fpdf.SetFont("times", "B", 50)
	fpdf.SetTextColor(50, 50, 50)
	pdf.moveAbs(0, 100)
	_, lineHt := fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "Certificate of Completion", gofpdf.AlignCenter)
	pdf.move(0, lineHt*2.0)

	fpdf.SetFont("arial", "", 28)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, "This certificate is awarded to", gofpdf.AlignCenter)
	pdf.move(0, lineHt*2.0)

	fpdf.SetFont("times", "B", 42)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt, *name, gofpdf.AlignCenter)
	pdf.move(0, lineHt*1.75)

	fpdf.SetFont("arial", "", 22)
	_, lineHt = fpdf.GetFontSize()
	fpdf.WriteAligned(0, lineHt*1.5, "For successfully alalalalal lalala lalalala lalalalalala alalalla alal alla alalal alalalala laalal alalla alala alalal alala alal ala laal", gofpdf.AlignCenter)
	pdf.move(0, lineHt*4.5)

	pdf.move(0, 65.0)
	fpdf.SetFillColor(100, 100, 100)
	fpdf.Rect(60.0, pdf.y, 240.0, 1.0, "F")
	fpdf.Rect(490.0, pdf.y, 240.0, 1.0, "F")

	fpdf.SetFont("arial", "", 12)
	pdf.move(0, lineHt/1.5)
	fpdf.SetTextColor(100, 100, 100)
	pdf.moveAbs(60.0+105.0, pdf.y)
	pdf.text("Date")
	pdf.moveAbs(490.0+100.0, pdf.y)
	pdf.text("Instructor")
	pdf.moveAbs(60.0, pdf.y-lineHt/1.5)
	fpdf.SetFont("times", "", 22)
	_, lineHt = fpdf.GetFontSize()
	pdf.move(0, -lineHt)
	fpdf.SetTextColor(50, 50, 50)
	yr, mo, day := time.Now().Date()
	dateStr := fmt.Sprintf("%d/%d/%d", mo, day, yr)
	fpdf.CellFormat(240.0, lineHt, dateStr, gofpdf.BorderNone, gofpdf.LineBreakNone, gofpdf.AlignCenter, false, 0, "")
	pdf.moveAbs(490.0, pdf.y)
	sig, err := gofpdf.SVGBasicFileParse("images/sig.svg")
	if err != nil {
		panic(err)
	}
	pdf.move(0, -(sig.Ht*0.45 - lineHt))
	fpdf.SVGBasicWrite(&sig, 0.5)
	// drawGrid(fpdf)

	err = fpdf.OutputFileAndClose("cert.pdf")
	if err != nil {
		panic(err)
	}
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
		pdf.Line(0, y, w, y)
		pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
	}
}
