package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

func GetID(bit int, isFirstZero bool) string {
	var result []string

	for i := 0; i < bit; i++ {
		// 生成 0-9 的随机数
		_number := rand.Intn(10)
		if _number == 0 && !isFirstZero {
			_number += 1
		}
		// strconv.Itoa() 可以把 int 转为字符串
		result = append(result, strconv.Itoa(_number))
	}
	return strings.Join(result, "")
}

func GetCode(bit int) string {
	var result []string
	_strings := [36]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i := 0; i < bit; i++ {
		// 生成 0-35 的随机数
		index := rand.Intn(36)
		if i == 0 && index < 10 {
			index += 1
		}
		result = append(result, _strings[index])
	}
	return strings.Join(result, "")
}
