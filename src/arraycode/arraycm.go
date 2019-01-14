package arraycode

import (
	"fmt"
	"crypto/sha256"
)

func comparecode(b1 []byte, b2 []byte) int {
	c1 := sha256.Sum256(b1)
	c2 := sha256.Sum256(b2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	
	
	return 1
}
