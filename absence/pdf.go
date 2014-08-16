package absence

import (
	"code.google.com/p/gofpdf"
	"fmt"
	"log"
	"os"
)

const (
	GOFPDF_DIR = "."
	FONT_DIR   = GOFPDF_DIR + "/font"
	IMG_DIR    = GOFPDF_DIR + "/image"
	TEXT_DIR   = GOFPDF_DIR + "/text"
)

type pdfWriter struct {
	pdf *gofpdf.Fpdf
	fl  *os.File
	idx int
}

func GeneratePdf(absence Absence) {

	layout := "01.02.2006"

	x := 0.0
	y := 0.0

	pdf := gofpdf.New("P", "mm", "A4", FONT_DIR)

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 5, "Abwesenheitsmeldung")
	y += 20
	pdf.SetY(y)
	pdf.SetFillColor(152, 251, 152)
	pdf.CellFormat(190, 5, "Abwesenheitsdaten", "", 1, "C", true, 0, "")
	pdf.SetFont("Arial", "", 11)

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Name Mitarbeiter/in:")

	//y += 10
	//pdf.SetY(y)
	//pdf.Cell(40, 5, absence.CreatedAt.String())

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Art der Abwesenheit:")

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Datum der Abwesenheit:")

	pdf.SetX(x + 60)
	pdf.Cell(40, 5, "von:")

	if absence.From != absence.To {
		pdf.SetX(x + 70)
		pdf.Cell(40, 5, absence.From.Format(layout))
	}

	pdf.SetX(x + 120)
	pdf.Cell(40, 5, "bis:")

	if absence.From != absence.To {
		pdf.SetX(x + 130)
		pdf.Cell(40, 5, absence.To.Format(layout))
	}

	y += 5
	pdf.SetY(y)
	pdf.SetX(x + 60)
	pdf.Cell(40, 5, "oder am:")

	if absence.From == absence.To {
		pdf.SetX(x + 80)
		pdf.Cell(40, 5, absence.From.Format(layout))
	}

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Grund der Abwesenheit:")

	y += 10
	pdf.SetY(y)
	pdf.SetTextColor(255, 0, 0)
	pdf.MultiCell(190, 5, "Ausser im Krankheitsfall muessen Antraege auf Berurlaubung in der Regel 14 Tage vor dem ersten Tag der Abwesenheit eingereicht werden.", "", "L", false)

	pdf.SetTextColor(0, 0, 0)

	y += 15
	pdf.SetY(y)
	pdf.Cell(190, 5, "( ) Ich habe den Termin meiner Abwesenheit mit dem/der zustaendigen Abegordneten geklaert.")

	y += 15
	pdf.SetY(y)
	pdf.CellFormat(190, 5, "Datum                                        Unterschrift MdL", "T", 1, "L", false, 0, "")

	y += 10
	pdf.SetY(y)
	pdf.Cell(190, 5, "( ) Ich habe den Termin meiner Abwesenheit mit meinem/meiner Vertreter/in geklaert.")

	y += 15
	pdf.SetY(y)
	pdf.CellFormat(190, 5, "Datum                                        Unterschrift des Mitarbeiters", "T", 1, "L", false, 0, "")

	pdf.SetTextColor(0, 0, 255)
	pdf.SetFont("Arial", "B", 11)
	y += 15
	pdf.SetY(y)
	pdf.CellFormat(190, 5, "Im Falle der Inanspruchnahme von Urlaubstagen bitte ausfuellen:", "", 0, "L", false, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 11)

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Anspruch an Urlaubstagen:")
	pdf.SetX(x + 80)
	pdf.Cell(40, 5, "PLACEHOLDER")

	y += 5
	pdf.SetY(y)
	pdf.Cell(40, 5, "Bisher genommene Urlaubstage:")
	pdf.SetX(x + 80)
	pdf.Cell(40, 5, "PLACEHOLDER")

	y += 5
	pdf.SetY(y)
	pdf.Cell(40, 5, "Beantragte Urlaubstage:")
	pdf.SetX(x + 80)
	pdf.Cell(40, 5, "PLACEHOLDER")

	y += 5
	pdf.SetY(y)
	pdf.Cell(40, 5, "Verbleibende Urlaubstage:")
	pdf.SetX(x + 80)
	pdf.Cell(40, 5, "PLACEHOLDER")

	pdf.SetFont("Arial", "B", 12)
	y += 15
	pdf.SetY(y)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(152, 251, 152)
	pdf.CellFormat(190, 5, "Genehmigung durch die Geschaeftsfuehrung", "", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 11)

	y += 10
	pdf.SetY(y)
	pdf.SetX(x + 20)
	pdf.Cell(40, 5, "( ) Genehmigt")

	y += 10
	pdf.SetY(y)
	pdf.SetX(x + 20)
	pdf.Cell(40, 5, "( ) Abgelehnt")

	y += 10
	pdf.SetY(y)
	pdf.Cell(40, 5, "Kommentare:")

	y += 40
	pdf.SetY(y)
	pdf.CellFormat(190, 5, "Datum                                         Unterschrift", "T", 1, "L", false, 0, "")

	pdf.OutputAndClose(docWriter(pdf, 1))
}

func (pw *pdfWriter) Write(p []byte) (n int, err error) {
	if pw.pdf.Ok() {
		return pw.fl.Write(p)
	}

	return
}

func (pw *pdfWriter) Close() (err error) {
	if pw.fl != nil {
		pw.fl.Close()
		pw.fl = nil
	}

	if pw.pdf.Ok() {
		log.Println("PDF successfully generated")
	} else {
		log.Println(pw.pdf.Error())
	}

	return
}

func docWriter(pdf *gofpdf.Fpdf, idx int) *pdfWriter {
	pw := new(pdfWriter)
	pw.pdf = pdf
	pw.idx = idx

	if pdf.Ok() {
		var err error
		fileStr := fmt.Sprintf("%s/pdf/absence%02d.pdf", GOFPDF_DIR, idx)
		pw.fl, err = os.Create(fileStr)
		if err != nil {
			pdf.SetErrorf("Error opening output file %s", fileStr)
		}
	}

	return pw
}
