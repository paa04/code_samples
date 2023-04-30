package uniq

import (
	. "awesomeProject/uniq_project/ansStruct"
	"strings"
)

func sNum(s string, num int) string {
	if num > len(s) {
		num = len(s)
	}
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}
	return s[i+num:]
}

func fNum(s string, num int) string {
	i := 0
	cur := 0
	for ; i < len(s) && cur < num; i++ {
		if s[i] == ' ' && (i == 0 || s[i-1] != ' ') {
			cur += 1
		}
	}
	if i >= len(s) {
		return ""
	}
	return s[i:]
}

func chooseComp(unicKey int) func(int) bool {
	var comp func(int) bool
	switch unicKey {
	case 0:
		comp = func(a int) bool {
			return a != 0
		}
	case 1:
		comp = func(a int) bool {
			return a == 1
		}
	case -1:
		comp = func(a int) bool {
			return a > 1
		}
	}
	return comp
}

func Uniq(arrOr []string, size int, s int, f int, unicKey int, iKey bool) []AnsStruct {
	arr := make(map[string]int)
	for i := 0; i < size; i++ {
		str := sNum(fNum(arrOr[i], f), s)
		if iKey {
			str = strings.ToLower(str)
		}
		arr[str] += 1
	}
	ret := make([]AnsStruct, 0)
	for i := 0; i < size; i++ {
		str := sNum(fNum(arrOr[i], f), s)
		if iKey {
			str = strings.ToLower(str)
		}
		comp := chooseComp(unicKey)
		if comp(arr[str]) {
			ret = append(ret, AnsStruct{arrOr[i], arr[str]})
			delete(arr, str)
		}
	}
	return ret
}
