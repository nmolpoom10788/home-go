/*
* @Author: kk
* @Date:   2022-06-26 21:47:44
* @Last Modified by:   kk
* @Last Modified time: 2022-06-26 22:22:54
 */

// chaiwitchit Konfoo
// 640510609
// HW01_1
// 204203 Sec 002

package main

import (
	"math"
)

func twosComplToInt(x string) int64 {
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
