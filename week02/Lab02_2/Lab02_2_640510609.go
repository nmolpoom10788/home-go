// ชัยวิทย์ชิต คนฟู ภูมิ
// 640510609
// Lab02_2
// 204203 Sec 002

package main

import (
	"strings"
)

const MAX = 100 // really? 36

// 11.01 1.1 2
func baseNAddition(r1, r2 string, base int) string {
	decPointPos1 := strings.Index(r1, ".")
	decPointPos2 := strings.Index(r2, ".")
	// println(decPointPos1) //2
	// println(decPointPos2) //1
	// println(r1)           //11.01 string
	// println(r2)           //1.1

	//summ := r1 + r1
	// println(summ)

	//เเยกคิดเป็นหน้าจุดกับหลังจุด
	//หน้าจุด
	front1 := r1
	front2 := r2
	//หลังจุด
	back1 := ""
	back2 := ""
	//เช็คว่า มีทศนิยมหรือไม่
	if decPointPos1 != -1 {
		front1 = r1[0:decPointPos1]
		back1 = r1[decPointPos1+1:]
	}
	if decPointPos2 != -1 {
		front2 = r2[0:decPointPos2]
		back2 = r2[decPointPos2+1:]
	}
	// back1 = r1[decPointPos1+1:]
	// back2 = r2[decPointPos2+1:]
	//การบวกกันของหน้าทศนิยม
	sumfront := addition(front1, front2, base)
	// println(sumfront)

	lenback1 := len(back1)
	lenback2 := len(back2)
	max_back := 0

	//หาความยาวของทศนิยม
	if lenback1 > max_back {
		max_back = lenback1
	}
	if lenback2 > max_back {
		max_back = lenback2
	}
	//เพิ่มความยาวของจุดทศนิยม
	back1 += strings.Repeat("0", max_back-lenback1)
	back2 += strings.Repeat("0", max_back-lenback2)
	sumback := ""

	//ดูว่ามีจุดทศนิยม หรือไม่มี
	if max_back != 0 {
		//การบวกกันของหลังทศนิยม
		sumback = addition(back1, back2, base)
	}
	// println(sumback)

	if len(sumback) > max_back {
		sumfront = addition(sumfront, "1", base)
		sumback = sumback[1:]
	}
	// println(sumfront + "." + sumback)
	finished := sumfront + "." + sumback
	if len(sumback) == 0 {
		finished = sumfront
	}

	return finished
}

func addition(n1, n2 string, base int) string {

	result := []byte(strings.Repeat("x", MAX))
	len1 := len(n1)
	len2 := len(n2)
	carry := 0

	i, j, k := len1-1, len2-1, MAX-1

	for ; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		temp := carry
		carry = 0

		if i >= 0 {
			temp += int(n1[i]) - int('0')
		}
		if j >= 0 {
			temp += int(n2[j]) - int('0')
		}

		result[k] = byte((temp % base) + int('0'))
		carry = temp / base
		//println(result)
	}

	if carry > 0 {
		result[k] = '1'
	} else {
		k++
	}

	return string(result[k:])

}
