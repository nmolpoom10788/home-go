/*
* @Author: kk
* @Date:   2022-06-26 21:47:44
* @Last Modified by:   kk
* @Last Modified time: 2022-06-27 01:12:19
 */

// chaiwitchit Konfoo
// 640510609
// HW01_3
// 204203 Sec 002

package main

import (
	"math"
)

func addNSubtract(n1, n2 string, bitLen uint8) (int64, int64) {
	var numn1 = len(n1)
	var numn2 = len(n2)
	maxN := 0

	//เปรียบเทียบว่าตัวไหนมากกว่ากัน
	if numn1 > numn2 {
		maxN = numn1
	} else {
		maxN = numn2
	}
	maxN = int(math.Max(float64(maxN), float64(bitLen)))

	// เพิ่มจำนวนบิตให้เท่ากัน
	//1
	//1111
	//กรณีเป็น 1 ให้เติม 11111 ข้างหน้า ให้เท่ากับอีกตัว
	//กรณี เป็น 0 ให้เติม 0000 ข้างหน้า ให้เท่ากับ อีกตัว
	if maxN > numn1 {
		numn1 = maxN - numn1 //3
		k := string(n1[0])
		for i := 0; i < numn1; i++ {
			n1 = k + n1
		}
		//println(n1)
		//n1 = 1111
	}
	//0101
	//010
	if maxN > numn2 {
		numn2 = maxN - numn2 //3
		k := string(n2[0])
		for i := 0; i < numn2; i++ {
			n2 = k + n2
		}
		//println(n2)
	}

	//นำมาบวกกัน
	// 1111   n1
	// 0010   n2
	// 10001
	// 0 1 2 3 4
	result := ""
	pos := 0
	curr := 0
	for i := maxN - 1; i >= 0; i-- {
		pos = (int(n1[i]) + int(n2[i]) - int('0') - int('0')) + curr
		if pos == 2 || pos == 3 {
			//pos = 0 , 1
			pos = pos - 2
			curr = 1
		} else {
			curr = 0
		}
		result = string(byte((pos)+int('0'))) + result
		//println(pos)

	}

	//println(result)

	//นำมาลบกัน
	resultneg := ""
	pos2 := 0
	curr2 := 0
	n2 = twocom(n2) //แปลงเป็น twocom
	for i := maxN - 1; i >= 0; i-- {
		pos2 = (int(n1[i]) + int(n2[i]) - int('0') - int('0')) + curr2
		if pos2 == 2 || pos2 == 3 {
			//pos = 0 , 1
			pos2 = pos2 - 2

			curr2 = 1
		} else {
			curr2 = 0
		}
		resultneg = string(byte((pos2)+int('0'))) + resultneg // สตริง ครอบ byte  จะเป็น char  // byte ครอบ int จะเป็น จำนวนเต็ม
		//println(pos)

	}

	//println(resultneg)

	//การตัดบิท
	if bitLen < uint8(maxN) {
		bitLen = uint8(maxN) - bitLen // ทำค่ามากลบค่าน้อย จะง่ายกว่า
		//println(bitLen)
		result = result[bitLen:]
		resultneg = resultneg[bitLen:]
	}

	resultpos := decimal(result)
	resultnega := decimal(resultneg)

	return resultpos, resultnega
}

//ฟังชันเปลี่ยนเป็น twocompplement
func twocom(x string) string {
	// fmt.Println(x)
	n := len(x)
	// fmt.Println("this is n :", n)
	var bitx = ""
	// var ten int64
	var k int

	for i := n - 1; i >= 0; i-- {
		if x[i] == '1' {
			k = i
			break
		}
	}

	for i := 0; i < k; i++ {
		if x[i] == '1' {
			bitx += "0"
		} else {
			bitx += "1"
		}
	}
	bitx += x[k:]

	return bitx
}

//ฟังก์ชั่น เปลี่ยนเป็นฐาน 10
func decimal(x string) int64 {
	// เปลี่ยน x string ให้เป็น int64
	//strx, _ := strconv.ParseInt(x, 0, 64)
	//fmt.Println(strx)
	numstrx := len(x)

	//fmt.Println(numstrx)
	var result int64 = 0

	// กรณีที่ ตัวเเรก มีค่า เท่ากับ 0 0 ที่เป็น character จะมีค่าเท่ากับ 48
	if x[0] == 48 {
		for i := 0; i < numstrx; i++ {
			if x[i] == 49 {
				result += int64(math.Pow(2, float64((numstrx-1)-i))) // นำ 2 ยกกำลัง จากหน้าไปหลัง เริ่มที่ 2^0

			}

		}
	}
	// กรณีที่ ตัวเเรก เป็น 1
	if x[0] == 49 { //x[0] == 1
		//fmt.Println(x[0])
		for i := 0; i < numstrx; i++ {
			//fmt.Println(i)
			if i == 0 {
				result -= int64(math.Pow(2, float64((numstrx-1)-i)))

			} else if x[i] == 49 {
				//fmt.Println("Pass")
				result += int64(math.Pow(2, float64((numstrx-1)-i)))
			}

		}

	}

	return result
}
