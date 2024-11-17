package ops

import (
	"fmt"
	"reflect"
)

func Example_cast() {
	a := NewAgent()
	fmt.Printf("test: cast_string() -> %v\n", a)

	opsT := cast(a)
	fmt.Printf("test: cast_ops() -> %v\n", reflect.TypeOf(opsT))

	//Output:
	//fail

}
