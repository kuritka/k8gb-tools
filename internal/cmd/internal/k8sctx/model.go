package k8sctx

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

type ListItem struct {
	Namespace string
	Name      string
	GeoTag    string
	Context   string
	Source    string
}

func (l ListItem) String() string {
	return fmt.Sprintf("%s %s (%s) -> geoTag: %s, context: %s, source: %s",
		emoji.LightBulb, aurora.Green(l.Name), aurora.BrightCyan(l.Namespace), aurora.BrightCyan(l.GeoTag), aurora.BrightCyan(l.Context), aurora.BrightMagenta(l.Source))
}
