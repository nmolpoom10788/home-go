// 	chaiwitchit konfoo
// 640510609
// Lab02_1
// 204203 Sec 002

package main

func powerOfTwo(x uint64) bool {
	if x & (x - 1) == 0  {
		return true
	} else {
		return false
	}
	
}
