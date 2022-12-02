// ชัยวิทย์ชิต คนฟู ภูมิ
// 640510609
// HW02_2
// 204203 Sec 002

package main

import (
	"strconv"
	"strings"
)

// func main() {

// 	var numCase int
// 	var line string
// 	var n uint32

// 	fmt.Scan(&numCase)

// 	for i := 0; i < numCase; i++ {
// 		_, err := fmt.Scan(&line)
// 		if err != nil {
// 			if err != io.EOF {
// 				log.Fatal(err)
// 			}
// 			break
// 		}

// 		if strings.Contains(line, ".") {
// 			fmt.Println(ipv4Encode(line))
// 		} else {
// 			temp, _ := strconv.Atoi(line)
// 			n = uint32(temp)
// 			fmt.Println(ipv4Decode(n))
// 		}
// 	}

// }

func ipv4Encode(ipString string) uint32 {
	var ip [4]int
	var result int64
	neg := 24
	number := strings.Split(ipString, ".") //นายพชร บรรจง สอนการใช้

	for i := 0; i < 4; i++ {
		// println(i)
		ip[i], _ = strconv.Atoi(number[i])
		result += int64(ip[i]) << int64(neg)
		neg = neg - 8
	}
	return uint32(result)

}

func ipv4Decode(ipUint uint32) string {
	var decode [4]string
	neg := 24

	for i := 0; i < 4; i++ {
		decode[i] = strconv.FormatInt((int64(ipUint)>>int64(neg))&0xff, 10)
		neg = neg - 8
	}

	return decode[0] + "." + decode[1] + "." + decode[2] + "." + decode[3]

}
