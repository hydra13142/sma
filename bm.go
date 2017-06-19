package sma

// BM算法，即Boyer-Moore算法。这是一种目前常用的字符串匹配算法。horspool算法是其简化版。该
// 算法除了使用了horspool算法里用到的根据字符最后出现位置挪动模式串的原理外，还利用已经匹配的
// 后缀序列：如果该序列在模式串出现多次，则挪动模式串对齐倒数第二次出现的位置；如果后缀在模式串
// 中只出现末尾那一次，则使用模式串的前缀去匹配该后缀序列的后缀，取其最长匹配序列对齐；如果两者
// 没有非空的匹配序列，则模式串与当前文本串匹配右侧末端位置后面一个字符对齐。BM算法利用这两个规
// 则，选择移动距离最长的来挪动模式串，从而提升匹配速度。改进算法将horspool对齐规则替换为更好
// 的sunday规则（horspool和sunday规则基本一致，只是对齐指示字符不同）。
type BM struct {
	char [256]int
	next []int
	temp string
}

// 创建一个BM对象
// char数组记录了某个字符最后一次出现在模式串时，模式串到该字符的前缀的长度，否则为0。next切片
// 储存模式串后缀s[i:]对应的内部从右往左数第二个匹配以及其前端序列的总长度；如果不存在内部匹配
// 则储存与该后缀的某个后缀匹配的最长模式串前缀的长度；如果也不存在则设置为0。next切片最后一个
// 元素表示未匹配到后缀时（即匹配的后缀为空字符）的情况，即0。
func NewBM(s string) *BM {
	l := len(s)
	t := new(BM)
	t.temp = s
	for i := 0; i < l; i++ {
		t.char[int(s[i])] = i + 1
	}
	t.next = make([]int, l+1)
	for i := l - 1; i >= 0; i-- {
		j, d := i-1, l-i
		for ; j >= 0; j-- {
			if s[j:j+d] == s[i:] {
				break
			}
		}
		if j >= 0 {
			t.next[i] = j + d
			continue
		}
		for d--; d > 0; d-- {
			if s[l-d:] == s[:d] {
				break
			}
		}
		t.next[i] = d
	}
	return t
}

// 搜索第一个匹配的位置，没有返回-1
func (this *BM) Index(s []byte) int {
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
			i += p - this.char[int(s[i])] // 采用sunday规则
		}
		k += p - this.next[j+1]
		if i < k {
			i = k
		}
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *BM) Find(s []byte) (o []int) {
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
			i += p - this.char[int(s[i])] // 采用sunday规则
		}
		k += p - this.next[j+1]
		if i < k {
			i = k
		}
	}
	return
}
