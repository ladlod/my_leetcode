package daily

// 36进制数加法
// 36进制数：0-9, a-z

func AddBase36(num1, num2 string) string {
	add := byte('0')
	res := []byte{}
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		b := transNum36Int(add)
		if i >= 0 {
			b += transNum36Int(num1[i])
		}
		if j >= 0 {
			b += transNum36Int(num2[j])
		}
		res = append([]byte{transNum36Byte(b % 36)}, res...)
		add = transNum36Byte(b / 36)
	}

	if add != '0' {
		res = append([]byte{add}, res...)
	}

	return string(res)
}

func transNum36Int(num byte) byte {
	if num >= '0' && num <= '9' {
		return num - '0'
	} else if num >= 'a' && num <= 'z' {
		return num - 'a' + 10
	} else {
		return 0
	}
}

func transNum36Byte(num byte) byte {
	if num >= 0 && num <= 9 {
		return num + '0'
	} else if num >= 10 && num <= 36 {
		return num + 'a' - 10
	} else {
		return '0'
	}
}
