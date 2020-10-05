package model

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

//ListItem; list command model
type ListItem struct {
	Namespace string
	Name      string
	GeoTag    string
	Context   string
	Source    string
	Error 	  error
}

//TODO: consider to move it into the view
func (l ListItem) String() string {
	if l.Error != nil {
		return fmt.Sprintf("%s ERROR: %s", emoji.Skull, aurora.BrightRed(l.Error))
	}
	return fmt.Sprintf("%s %s (%s) -> geoTag: %s, context: %s",
		emoji.LightBulb, aurora.Green(l.Name), aurora.BrightCyan(l.Namespace), aurora.BrightCyan(l.GeoTag), aurora.BrightCyan(l.Context))
}
