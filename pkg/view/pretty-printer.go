package view

import (
	"fmt"
	"io"
	"os"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

type PrettyPrinter struct {
	out io.Writer
}

func DefaultPrettyPrinter() *PrettyPrinter {
	return &PrettyPrinter{
		out: os.Stdout,
	}
}

func (p *PrettyPrinter) Title(title string) (err error) {
	err = p.print("%s %s", emoji.FourLeafClover, aurora.BrightCyan(title))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error) {
	err = p.print("%4s %s", emoji.GreenCircle, aurora.Cyan(subtitle))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Paragraph(paragraph string) (err error) {
	err = p.print("%8s %s"," ", aurora.White(paragraph))
	p.NewLine()
	return
}


func (p *PrettyPrinter) NewLine() {
	_ = p.print("\n")
}

func (p *PrettyPrinter) print(format string, a ...interface{}) (err error) {
	_, err = fmt.Fprintf(p.out, format, a...)
	return
}
