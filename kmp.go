package sma

// KMP算法，全名Knuth-Morris-Pratt算法。其理论是，在进行字符串匹配时，如果遇到匹配失败的情
// 况，则根据已经匹配的部分挪动模式串的位置，使模式串的某个前缀与已匹配部分的某个后缀成功匹配，
// 并要求挪动距离最短即匹配长度最大，这是为了避免错失成功的匹配。注意，所谓“已匹配的部分”很明显
// 也是模式串的某个前缀。因此，KMP算法的关键是计算出模式串的某个前缀对应的最长的可作为其真后缀
// 的另一模式串前缀（显然也是该前缀的前缀），这里称其为最长双配序列。对模式串s建立一个整数数组，
// 称为n，n[i]的值为s[:i]的最长双配序列的长度。这一数组即被称为next数组。构建next数组的过程
// 和根据next数组进行匹配的过程是有异曲同工之妙的。
type KMP struct {
	next []int
	temp string
}

// 生成一个KMP对象，内部包含模式串和计算好的next数组。
func NewKMP(s string) *KMP {
	l := len(s)
	n := make([]int, l+1)

	// n[i]表示s长度为i的前缀s[:i]的最长双配序列的长度
	// 很显然n[0]、n[1]都是0，从n[2]开始算起
	for i, j := 2, 0; i <= l; {
		// 已知j=n[i-1]，亦即s[:i-1]的最长双配序列的长度
		// 则s[i-1]表示s[:i-1]后的第一个字符
		// s[j]表示s[:i-1]的最长双配序列后的第一个字符
		if s[i-1] == s[j] { // 说明s[:i]的最长双配序列的长度为j+1
			n[i] = j + 1    // 对n[i]赋值
			i, j = i+1, j+1 // 更新i、j，进入下一轮循环
		} else { // 字符不匹配，则需要寻找次长双配序列再次测试
			if j == 0 { // 空双配序列也不能满足延伸的要求s[i-1] == s[j]
				// n[i]为0，即默认值
				i++ // 注意j为0不需要赋值，进入下一轮循环
			} else { // 测试次长双配序列
				j = n[j] // 次长双配序列必然是最长双配序列的最长双配序列
			}
		}
	}
	return &KMP{n, s}
}

// 搜索第一个匹配的位置，没有返回-1
func (this *KMP) Index(s []byte) int {
	t := this.temp
	p := len(t)
	q := len(s)
	i, j := 0, 0
	for i+p < j+q {
		if s[i] == t[j] {
			i, j = i+1, j+1
			if j == p {
				return i - p
			}
		} else if j == 0 {
			i++
		} else {
			j = this.next[j]
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *KMP) Find(s []byte) (o []int) {
	t := this.temp
	p := len(t)
	q := len(s)
	i, j := 0, 0
	for i+p < j+q {
		if s[i] == t[j] {
			i, j = i+1, j+1
			if j == p {
				o = append(o, i-p)
				j = this.next[j]
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = this.next[j]
			}
		}
	}
	return
}
