package sub

import "fmt"

type mydata struct {
	num int
	str string
}

func ClassPrint() {
	var x mydata
	x.num = 10
	x.str = "something"
	fmt.Printf("x.num=%d, x.str=%s\n", x.num, x.str)
}
