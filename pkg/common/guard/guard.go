//Package guard panics when error occurs
package guard

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/logrusorgru/aurora"
)

//FailOnError panics when error occurs.
func FailOnError(err error, message string, a ...interface{}) {
	if err != nil {
		m := fmt.Sprintf(message, a...)
		msg := fmt.Sprintf("%s %s \n%s\n", emoji.RedCircle, aurora.White(m), aurora.Yellow(err.Error()))
		panic(msg)
	}
}

//Message write message to output on the screen. Nothing else...
func Message(message string, a ...interface{}) {
	m := fmt.Sprintf(message, a...)
	fmt.Printf("%s %s \n", emoji.OpenBook, aurora.Yellow(m))
}
