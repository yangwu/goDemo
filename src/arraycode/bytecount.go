package arraycode

import (
	"fmt"
)

var pc [256]byte

func init() {
	fmt.Printf("init pc = %d\n", pc)
	//以空间换时间， 初始化 0-2^8 的数据中(byte的数据范围)，每个数据中包含的1的个数。计算公式位：pc[i] = pc[i/2] + i%2
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("i=%b and value=%d\n", i, pc[i])
	}
	fmt.Printf("after init pc = %d\n", pc)

	fmt.Printf("pc value=%b\n", pc)
}

func Pcount(x uint64) int {
	fmt.Printf("x %d value=%b\n", x, x)
	//byte长度为8位,因此每次右移8位，计算该8位数据；
	count := int(pc[byte(x>>0)] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

	/*for k:=uint32(0);k<8;k++{
		index := byte(x>>(k*8))
		fmt.Printf("k=%d,pc[index],index=%b,value=%d\n", k,index,pc[index])
	}*/
	fmt.Printf("x %d count of 1=%d\n", x, count)
	return count
}

func Itcount(x uint64) int {
	var count int
	for x > 0 {
		count += (int)(x & 1)
		x = x >> 1
	}
	return count
}

func Ishex(x uint64) bool {
	return ((x != 0) && ((x & (x - 1)) == 0))
}
