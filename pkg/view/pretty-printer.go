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
	p.newLine()
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error) {
	err = p.print("%4s %s", emoji.GreenCircle, aurora.Cyan(subtitle))
	p.newLine()
	return
}

func (p *PrettyPrinter) newLine() {
	_ = p.print("\n")
}

func (p *PrettyPrinter) print(format string, a ...interface{}) (err error) {
	_, err = fmt.Fprintf(p.out, format, a...)
	return
}
