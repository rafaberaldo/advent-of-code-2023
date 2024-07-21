package assert

import (
	"fmt"
)

func Assert(condition bool, msg string, v ...any) {
	if !condition {
		var newMsg = fmt.Sprintf(msg, v...)
		panic(newMsg)
	}
}
