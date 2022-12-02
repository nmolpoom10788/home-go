// chaiwitchit Konfoo
// 640510609
// Lab01_1
// 204203 Sec 002

package main

import (
	"fmt"
	"log"
	"strings"
)

// set the maximum digit length. why 36 and not 35?
const MAX = 36

func main() {
	// why are we using string?
	var n1, n2 string

	// read in two string (can be multiple lines)
	_, err := fmt.Scan(&n1, &n2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addition(n1, n2))
}

func addition(n1, n2 string) string {
	// this is just a skeleton and will give out wrong result
	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)

	// loop from the left most digit
	i, j, k := len1-1, len2-1, MAX-1
	var p int = 0

	// loop from right most position
	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		// current digit
		currDigit := 0
		currDigit += p
		p = 0
		// add the value from the current digit of n1
		if i >= 0 {
			currDigit += int(n1[i]) - int('0')
		}

		// add the value from the current digit of n2
		if j >= 0 {
			currDigit += int(n2[j]) - int('0')
		}

		if currDigit > 9 {
			p = 1
			currDigit %= 10

		}

		// ถ้าหลักหน่วยบวกกันเเล้วมากกว่า 9 ให้ เอาหลักหน่วยเก็บไว้ เเล้วเอาหลักสิบไปทด

		// ตัวทด ให้นำมาบวกเพิ่มด้วย
		// 7+3 = 10 ให้เก็บ 0 ไว้เเล้ว เอา 1 ไปบวกกับช่องถัดไป
		// 	เราทดได้เเล้ว
		// เเต่เราเอาไปเพิ่มทหลักไม่ได้  งงง มากก วิธีการเพิ่มหลัก

		// convert back to byte (one char is called byte)
		result[k] = byte(currDigit + int('0'))

	}

	if p == 1 {
		return "1" + string(result[k+1:])
	}
	// convert array of bytes to string
	return string(result[k+1:])

}
