package superdigitt

import (
	"strconv"
)

func SuperDigit(n string, k int32) int32 {
	// Write your code here
	var buff int32
	for _, c := range n {
		buff += int32(c - '0')

	}
	buff = buff * k
	if buff >= 10 {
		s := strconv.Itoa(int(buff))
		return SuperDigit(s, 1)
	} else {
		return buff
	}

}
