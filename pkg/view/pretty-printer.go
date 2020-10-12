package view

import (
	"fmt"
	"io"
	"os"
	"strings"

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

func (p *PrettyPrinter) Title(title ...string) (err error) {
	if len(title) == 0 {
		return fmt.Errorf("missing title")
	}
	if len(title) == 1 {
		err = p.print("%s %s", emoji.FourLeafClover, aurora.BrightCyan(title))
		p.NewLine()
		return
	}
	err = p.print("%s %s (%s)", emoji.FourLeafClover, aurora.BrightCyan(title[0]),
		aurora.BrightGreen(strings.Join(title[1:], ",")))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Subtitle(subtitle string) (err error) {
	err = p.print("%4s %s", emoji.GreenCircle, aurora.Cyan(subtitle))
	p.NewLine()
	return
}

func (p *PrettyPrinter) Paragraph(property, value string, serr error) (err error) {
	var e = ""
	if serr != nil {
		e = fmt.Sprintf("%s %s", emoji.LightBulb, aurora.BrightRed(serr.Error()))
	}
	err = p.print("%8s%-10s : %v %4s %s", " ", aurora.BrightMagenta(property), aurora.BrightYellow(value), " ", e)
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
