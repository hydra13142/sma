package sma

// Brute-Force算法，即暴力搜索算法，最基本的字符串比对算法，首先对齐模式串和文本串的0位，然后
// 进行匹配，匹配失败则模式串向右移动一位，再次重新匹配，依次进行下去。
type BF struct {
	temp string
}

// 创建一个BF对象
func NewBF(s string) *BF {
	return &BF{s}
}

// 搜索第一个匹配的位置，没有返回-1
func (this BF) Index(s []byte) int {
	p := len(this.temp)
	q := len(s)
	for i := 0; i+p < q; i++ {
		if Compare(this.temp, s[i:]) {
			return i
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this BF) Find(s []byte) (o []int) {
	p := len(this.temp)
	q := len(s)
	for i := 0; i+p < q; i++ {
		if Compare(this.temp, s[i:]) {
			o = append(o, i)
		}
	}
	return
}

// 固定位置匹配string格式的模式串和[]byte格式的文本串
func Compare(t string, s []byte) bool {
	for i, l := 0, len(t); i < l; i++ {
		if t[i] != s[i] {
			return false
		}
	}
	return true
}
