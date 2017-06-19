package sma

// sunday算法的核心是利用字符在模式串中最后出现的位置来挪动模式串，模式串的移动是从左向右，但
// 匹配时却是从右向左。匹配过程中遇到不匹配字符时，以当前对齐位置下模式串末尾字符对应的文本串字
// 符的下一个字符作为标准，移动模式串使模式串中最后一个该字符与文本串该字符对齐，对齐最后一个的
// 目的是防止漏解。如果模式串中不存在该字符，则移动模式串对齐该字符之后的那个位置。算法关键是生
// 成一个某字符在模式串最后出现位置的数组。因为以模式串末尾后的字符为标准对齐，不会出现模式串退
// 步的情况。sunday算法与horspool算法类似，区别在于用于对齐模式串的指示字符不同。
type Sunday struct {
	char [256]int
	temp string
}

// 创建一个Sunday对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0
func NewSunday(s string) *Sunday {
	t := new(Sunday)
	t.temp = s
	for i := 0; i < len(s); i++ {
		t.char[int(s[i])] = i + 1
	}
	return t
}

// 搜索第一个匹配的位置，没有返回-1
func (this *Sunday) Index(s []byte) int {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			return k + 1
		}
		if i < q {
			i = i + p - this.char[int(s[i])]
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *Sunday) Find(s []byte) (o []int) {
	var (
		p       = len(this.temp)
		q       = len(s)
		i, j, k int
	)
	for i = p - 1; i < q; {
		i, j, k = i+1, p-1, i
		for ; j >= 0; j, k = j-1, k-1 {
			if this.temp[j] != s[k] {
				break
			}
		}
		if j < 0 {
			o = append(o, k+1)
		}
		if i < q {
			i = i + p - this.char[int(s[i])]
		}
	}
	return
}
