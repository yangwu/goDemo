package arraycode

import (
	"fmt"
	"testing"
)

func TestPcount(t *testing.T) {
	//fmt.Printf("count(4)=%d\n",Pcount(4))
	fmt.Printf("Pcount count(25)=%d\n", Pcount(1088825))
	//fmt.Printf("Itcount count(25)=%d\n", Itcount(25))
}

func TestIshex(t *testing.T) {
	fmt.Println("TestIshex")
	for i := uint64(100);i > 0;i-- {
		if(Ishex(i)){
			fmt.Printf("i:%d,and Ishex(i) = %t, %b\n", i, Ishex(i),i)
		}
	}
}
