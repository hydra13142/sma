package sma

// RK算法，即Robin-Karp算法，哈希检索算法。宗旨是对模式串求哈希ID，对匹配的文本也依次求哈希
// ID，在两个ID一致的进行逐字符比对的复核。该算法的关键是不能对每个哈希ID都要重新计算，这样算
// 法复杂度不会改变，而应采用由旧的ID根据新字符生成新的ID的哈希算法，比如说，累加求和，以及求
// 异或运算。复核是必须的，因为很可能存在哈希碰撞的情况。
type RK struct {
	Hash
	temp string
	sign int
}

// 用于RK算法的hash接口
type Hash interface {
	Feed(byte) // 吃进一个字符
	Free(byte) // 释放一个字符
	Reset()    // 重启hash缓存
	Sum() int  // 当前hash缓存
}

// 根据模式串和hash算法创建一个RK算法的对象
func NewRK(s string, h Hash) *RK {
	h.Reset()
	for i := 0; i < len(s); i++ {
		h.Feed(s[i])
	}
	return &RK{h, s, h.Sum()}
}

// 搜索第一个匹配的位置，没有返回-1
func (this *RK) Index(s []byte) int {
	p := len(this.temp)
	q := len(s)
	this.Reset()
	for i := 0; i < p; i++ {
		this.Feed(s[i])
	}
	// i,j分别表示生成下一个ID需要释放和接收的文本串字符的下标
	for i, j := 0, p; j < q; i, j = i+1, j+1 {
		if this.Sum() == this.sign {
			if Compare(this.temp, s[i:]) {
				return i
			}
		}
		this.Feed(s[j])
		this.Free(s[i])
	}
	return -1
}

// 搜索所有匹配的位置，没有返回nil
func (this *RK) Find(s []byte) (o []int) {
	p := len(this.temp)
	q := len(s)
	this.Reset()
	for i := 0; i < p; i++ {
		this.Feed(s[i])
	}
	// i,j分别表示生成下一个ID需要释放和接收的文本串字符的下标
	for i, j := 0, p; j < q; i, j = i+1, j+1 {
		if this.Sum() == this.sign {
			if Compare(this.temp, s[i:]) {
				o = append(o, i)
			}
		}
		this.Feed(s[j])
		this.Free(s[i])
	}
	return
}

// 利用异或运算实现了Hash接口
type ExclusiveOr int

func (this *ExclusiveOr) Feed(c byte) {
	*this ^= ExclusiveOr(c)
}

func (this *ExclusiveOr) Free(c byte) {
	*this ^= ExclusiveOr(c)
}

func (this *ExclusiveOr) Reset() {
	*this = 0
}

func (this *ExclusiveOr) Sum() int {
	return int(*this)
}

// 利用加减运算实现了Hash接口
type PlusMinus int

func (this *PlusMinus) Feed(c byte) {
	*this += PlusMinus(c)
}

func (this *PlusMinus) Free(c byte) {
	*this -= PlusMinus(c)
}

func (this *PlusMinus) Reset() {
	*this = 0
}

func (this *PlusMinus) Sum() int {
	return int(*this)
}
