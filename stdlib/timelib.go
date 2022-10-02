package stdlib

import (
	"fmt"
	"pankti/object"
	"time"
)

func UnixTimeFunc(args []object.Obj) object.Obj {

	return &object.String{Value: fmt.Sprintf("%d", time.Now().Unix())}
}
