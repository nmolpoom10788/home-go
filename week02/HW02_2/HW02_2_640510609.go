// ชัยวิทย์ชิต คนฟู ภูมิ
// 640510609
// HW02_2
// 204203 Sec 002

package main

import (
	"strings"
)

const MAX = 200

// func main() {

// 	var line string
// 	var s string = ""
// 	var bp int = -1
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for scanner.Scan() {
// 		line = scanner.Text()

// 		var err error
// 		if s == "" {
// 			_, err = fmt.Sscanf(line, "%s", &s)

// 		} else {
// 			_, err = fmt.Sscanf(line, "%d", &bp)
// 		}

// 		if err != nil {
// 			if err != io.EOF {
// 				log.Fatal(err)
// 			}
// 			break
// 		}

// 		if bp != -1 {
// 			fmt.Println(roundToEven(s, uint8(bp)))
// 			s = ""
// 			bp = -1
// 		}
// 	}
// }

func roundToEven(x string, bPlace uint8) string {
	//tese ces
	//1101.101
	//2
	// println(x)      //2
	// println(bPlace) //4
	//results =
	//เเยกตำเเหน่ง หน้าจุด หลังจุด
	decPoint := strings.Index(x, ".")
	// println(decPoint)
	front := x[0:decPoint] //หน้าจุด
	back := x[decPoint+1:] //หลังจุด
	numfront := len(front)
	numback := len(back)
	result := ""
	n := ""
	backpoint := ""

	// println(front)
	// println(back)
	//เอามาปัด

	//ถ้าตัวหน้าเป็น 0
	if bPlace == 0 {

		//กรณีที่ตัวหน้าเป็น 0

		if int((front[numfront]-1))-'0' == 0 && int((back[0])-'0') == 1 {

			if uint8(numback)-bPlace > 1 {
				pangback := back[bPlace:]
				if int64((pangback[1])-'0') == 1 {
					result = AddBinary(front, "1")
				} else {
					if strings.Count(back[bPlace:], "1") > 1 {
						result = AddBinary(front, "1")
					} else {
						result = front
					}
				}

			} else {
				result = front
			}

		} else if int64((front[numfront]-1))-'0' == 1 && int64((back[0])-'0') == 1 {
			result = AddBinary(front, "1")

		} else {
			result = front
		}

	} else if uint8(numback) > bPlace {

		if int64((back[bPlace-1])-'0') == 0 && int64((back[bPlace])-'0') == 1 {

			if uint8(numback)-bPlace > 1 {
				pangback := back[bPlace:]
				if int64((pangback[1])-'0') == 1 {
					back = back[:bPlace]
					n = front + "." + back
					curry := Addcurr(len(front), int(bPlace))
					result = AddBinary(n, curry)
				} else if strings.Count(pangback, "1") > 1 {
					back = back[:bPlace]
					n = front + "." + back
					carry := Addcurr(len(front), int(bPlace))
					result = AddBinary(n, carry)
				} else {
					back = back[:bPlace]
					result = front + "." + back
				}

			} else {
				back = back[:bPlace]
				result = front + "." + back
			}

		} else if int64((back[bPlace-1])-'0') == 1 && int64((back[bPlace])-'0') == 1 {
			back = back[:bPlace]
			n = front + "." + back
			curry := Addcurr(len(front), int(bPlace))
			result = AddBinary(n, curry)

		} else {
			back = back[:bPlace]
			result = front + "." + back
		}

	} else if uint8(numback) < bPlace {
		for i := 0; i < int(bPlace)-numback; i++ {
			backpoint += "0"
		}
		x += backpoint
		result = x

	} else {
		result = x
	}

	return result
}

//function ************************************************************************************************************
func addition(n1, n2 string, base int) string { //จาก lab02_2

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

//function ************************************************************************************************************
func AddBinary(n1 string, n2 string) string { //นาย พชร บรรจง สอนการเขียนฟังก์ชั้นนี้
	var curry int = 0
	var result string = ""
	if strings.Count(n1, ".") != 0 && strings.Count(n2, ".") != 0 {
		frontPos1 := strings.Index(n1, ".")
		frontPos2 := strings.Index(n2, ".")
		front1 := n1[0:frontPos1]
		front2 := n2[0:frontPos2]
		back1 := n1[frontPos1+1:]
		back2 := n2[frontPos2+1:]

		if len(addition(back1, back2, curry)) > len(back1) {
			curry = 1
			result = addition(front1, front2, curry) + "." + string(addition(back1, back2, 0)[1:])
			return result
		}
		result = addition(front1, front2, curry) + "." + addition(back1, back2, curry)
	} else {
		result = addition(n1, n2, curry)
		return result
	}

	return result
}

//function ************************************************************************************************************
func Addcurr(lend int, lenb int) string { //นาย พชร บรรจง สอนการเขียนฟังก์ชั้นนี้
	var front string = ""
	var back string = ""
	var carry string = ""
	for i := 0; i < lend; i++ {
		front += "0"
	}
	for i := 0; i < lenb-1; i++ {
		back += "0"
	}
	back += "1"
	carry = front + "." + back
	return carry
}
