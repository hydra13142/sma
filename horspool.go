package sma

// horspool算法的核心是利用字符在模式串中最后出现的位置来挪动模式串，模式串的移动是从左向右，
// 但匹配时却是从右向左。匹配过程中遇到不匹配字符时，移动模式串使文本串该字符与模式串中最后一个
// 该字符对齐，对齐最后一个的目的是防止漏解。如果模式串中不存在该字符，则移动模式串对齐该字符之
// 后的那个字符。算法关键是生成一个某字符在模式串最后出现位置的数组。注意根据该规则进行对齐可能
// 导致模式串反向移动，此时应将模式串强制右移一位。horspool算法与sunday算法类似，区别在于用于
// 对齐模式串的指示字符不同。
type Horspool struct {
	char [256]int
	temp string
}

// 创建一个Horspool对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0
func NewHorspool(s string) *Horspool {
	t := new(Horspool)
	t.temp = s
	for i := 0; i < len(s); i++ {
		t.char[int(s[i])] = i + 1
	}
	return t
}

// 搜索第一个匹配的位置，没有返回-1
func (this *Horspool) Index(s []byte) int {
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
		if k += p - this.char[int(s[k])]; i < k {
			i = k
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *Horspool) Find(s []byte) (o []int) {
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
		if k += p - this.char[int(s[k])]; i < k {
			i = k
		}
	}
	return
}
