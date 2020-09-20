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
		if message != "" {
			panic(fmt.Sprintf("%s %s\n%s\n", emoji.RedCircle, aurora.White(m), aurora.Yellow(err.Error())))
		}
		panic(fmt.Sprintf("%s %s\n", emoji.RedCircle, aurora.Yellow(err.Error())))
	}
}

//Message write message to output on the screen. Nothing else...
func Message(message string, a ...interface{}) {
	m := fmt.Sprintf(message, a...)
	fmt.Printf("%s %s\n", emoji.YellowSquare, aurora.Yellow(m))
}
