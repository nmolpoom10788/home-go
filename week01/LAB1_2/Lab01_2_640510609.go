// chaiwitchit Konfoo
// 640510609
// Lab01_2
// 204203 Sec 002

package main

import (
	"math"
	"strings"
)

const MAX_INT = 64
const DEC_PLACE = 6

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 { // turn negative numbers to positive  เป็นเลข ติดลบเป็น บวก
		x = -x
		sign = "-"
	}
	// split at the decimal point แยกจุดทศนิยม
	front := int64(x)          //ประกาศ ฟอนที่เป็นจำนวนเต็ม
	back := x - float64(front) //หลังจุด

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	// putting every part together การประกอบทุกส่วนเข้าด้วยกัน
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {
	// only need to implement this function  จำเป็นต้องใช้ฟังก์ชันนี้เท่านั้น เชียนโค้ดบรรทัดนี้นะจ้ะ
	if b == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", DEC_PLACE))
	var currDigit byte

	for i := 0; i < DEC_PLACE; i++ {
		currDigit = byte((x * float64(b)) + float64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
			//println("currDi>9:", currDigit)
		}

		result[i] = currDigit
		x *= float64(b)
		if x > 0 {
			x = x - math.Floor(x) // การปัดเศษลงเนาะ 1.5 == 1 เก็บ 1 เเล้วเหลือ 0.5
		}
	}

	return string(result[0:])

}

func posIntToBaseB(x int64, b uint8) string {
	// this function is working correctly  ฟังก์ชั่นนี้ทำงานถูกต้อง
	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1 // 64-1 = 63
	var currDigit byte

	for x > 0 {
		// calculate and convert back to char คำนวนเเล้วเเปลงเป็นฐาน
		currDigit = byte((x % int64(b)) + int64('0'))
		//println("currDigit;", (string(currDigit)))

		if currDigit > '9' {

			currDigit = 'A' + currDigit - '9' - 1
			//println("currDi>9:", currDigit)
		}
		result[k] = currDigit
		x = x / int64(b)
		//println("result ",k,":", string(result[k]))
		k--

	}
	// convert from byte array to string ไบต์อาเรย์ ทำไปเป็น สตริง
	return string(result[k+1:])
}
