// ชัยวิทย์ชิต คนฟู ภูมิ
// 640510609
// HW02_3
// 204203 Sec 002

package main

import (
	"math"
	"strings"
)

const BIT_LEN = 16

// func main() {

// 	var x float32
// 	var line string

// 	scanner := bufio.NewScanner(os.Stdin)
// 	// var result []string
// 	for scanner.Scan() {
// 		line = scanner.Text()

// 		_, err := fmt.Sscanf(line, "%f", &x)
// 		if err != nil {
// 			if err != io.EOF {
// 				log.Fatal(err)
// 			}
// 			break
// 		}
// 		result := float16bitNormed(x)
// 		if len(result) == BIT_LEN {
// 			fmt.Println(string(result[0]) + " " + string(result[1:9]) + " " + string(result[9:]))
// 		} else {
// 			fmt.Println(result)
// 		}

// 	}
// }

const fracLen = 7
const expLen = 8

const BASE uint8 = 2
const DEBUG = false

//from Lab01_2
const MAX_INT = 64
const DEC_PLACE = 128

func float16bitNormed(n float32) string {

	// expLen = 8, fracLen = 7
	expLen := 8
	fracLen := 7

	var curr string = ""
	var bias = int(pow(2, expLen-1) - 1) // bias = 127

	if DEBUG {
		println("Bias", bias)
	}

	var minNorm float64 = 0.00000000000000000000000000000000000005 // dummy value
	var maxNorm float64 = 338953138925153547590470800371487866880  // dummy value

	sign := "0"
	if n < 0 {
		n = -n
		sign = "1"
	}

	if (float64(n) > maxNorm) || (float64(n) < minNorm) {

		if DEBUG {
			println(n, "overflow")
		}
		return ""
	}

	//หา E

	Eyai := math.Floor(math.Log2(float64(n)))
	// //println(Eyai)

	// //หา exponent
	numN := float64(n) / (pow(2, int(Eyai)))
	// //println(numN)

	// //เเปลงเป็นฐาน 2
	pangbinary := floatToBaseB(numN, 2)
	// //println(pangbinary)

	//แยกกันโดยจุด
	fracter := strings.Split(pangbinary, ".")

	// //รับ8ตัวของfrac
	frac := fracter[1][:fracLen]

	pangexp := Eyai + float64(bias)
	// แปลงเป็นเลขฐาน2
	exponent := strings.Split(floatToBaseB(pangexp, 2), ".")

	//เช็คว่าไม่เกิน 8 ตัว
	if len(exponent[0]) < expLen {
		for i := 0; i < expLen-len(exponent[0]); i++ {
			curr += "0"
		}
		exponent[0] = curr + exponent[0]
	}
	exp := exponent[0][:expLen]

	return sign + exp + frac
}

func pow(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

//------------------------------------ Lab01_2

func floatToBaseB(x float64, b uint8) string {
	sign := ""

	if x < 0 {
		x = -x
		sign = "-"
	}
	// split at the decimal point
	front := int64(x)
	back := x - float64(front)

	frontStr := posIntToBaseB(front, b)
	backStr := fractionToBaseB(back, b)
	converted := sign + frontStr + "." + backStr

	return converted

}

func fractionToBaseB(x float64, b uint8) string {

	result := []byte(strings.Repeat("x", DEC_PLACE))
	var currDigit byte

	for i := 0; i < DEC_PLACE; i++ {
		x = x * float64(b)
		front := int64(x)
		currDigit = byte((front) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		x = x - float64(front)
		result[i] = currDigit

	}

	return string(result[:])
}

func posIntToBaseB(x int64, b uint8) string {

	if x == 0 {
		return "0"
	}

	result := []byte(strings.Repeat("x", MAX_INT))
	k := MAX_INT - 1
	var currDigit byte

	for x > 0 {
		currDigit = byte((x % int64(b)) + int64('0'))
		if currDigit > '9' {
			currDigit = 'A' + currDigit - '9' - 1
		}
		result[k] = currDigit
		x = x / int64(b)
		k--
	}

	return string(result[k+1:])
}
