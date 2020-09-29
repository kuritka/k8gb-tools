package k8sctx

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

type ListItem struct {
	Namespace string
	Name      string
	Cluster   string
	GeoTag    string
	Context   string
}

func (l ListItem) String() string {
	return fmt.Sprintf("%s %s (%s) -> geoTag: %s, context: %s",
		emoji.LightBulb, aurora.Green(l.Name), aurora.BrightCyan(l.Namespace), aurora.BrightCyan(l.GeoTag), aurora.BrightCyan(l.Context))
}
