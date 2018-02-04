package main

func cnum(s string) int {
	var x int
	x = (int(s[0]-'0'))*100 + (int(s[1]-'0'))*10 + (int(s[2] - '0'))
	return x
}

func Validip(str string) int {
	var k int
	for i := 0; i < len(str); i++ {
		if str[i] == '.' || i == len(str)-1 {
			var s string
			var num int
			if i == len(str)-1 {
				s = str[k : i+1]

			} else {
				s = str[k:i]
			}
			if len(s) < 4 {
				num = cnum(s)
				if num > 256 {
					return 0
				}
			} else {
				return 0
			}

			if i != len(str)-1 {
				k = i + 1
			}
		}
	}
	return 1
}

func Nonalpha(str string) string {
	//var str string
	//	str = "a..s($4.;'1241d135"
	//	fmt.Println(str)
	var s []byte
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= '0' && str[i] <= '9') {
			s = append(s, str[i])
			//fmt.Println(i)
		} else {
			//str = strings.Trim(str, string(str[i]))
		}

	}
	st := string(s)
	return st
}
