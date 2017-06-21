package sma

// AC自动机的节点
type node struct {
	next map[byte]int
	deep int
	over int
	fail int
}

// AC算法，全名Alfred-Corasick自动机算法，这是一种多模式匹配的算法。这种算法某种程度上可说是
// KMP算法的多模式版本。事实上，如果只给AC算法提供一个字符串来生成AC自动机，就会发现，AC自动机
// 的goto链就是模式串，而fail链正是next数组的AC自动机版本。作为KMP算法的扩展版本，AC算法主要
// 用于多模式匹配。AC算法的预处理过程包括两个部分，其一是根据多个模式串构造Trie树的构建，其二是
// 在Trie树上添加fail“指针”，亦即匹配失败时改变当前活动节点为其指向的节点。类型AC表示由一系列
// 节点构成的AC自动机。
type AC struct {
	State []node
	Final []string
	Count int
}

// 创建一个AC自动机，使用模式串本身作为其命名
func NewAC(ts []string) *AC {
	this := new(AC)
	for _, t := range ts {
		this.Add(t)
	}
	this.Prepare()
	return this
}

// 清空结构体，以重新生成多模式匹配
func (this *AC) Reset() {
	this.State = nil
	this.Final = nil
	this.Count = 0
}

// 向AC自动机中添加模式串
func (this *AC) Add(temp string) {
	p := len(this.Final)
	this.Final = append(this.Final, temp)
	if this.Count == 0 {
		this.newNode()
	}
	this.incStr(p, 0, temp)
}

// 重新生成fail指针，每次添加完一到多个新的模式串之后，在匹配文本串之前，都应执行Prepare方法
func (this *AC) Prepare() {
	type cell struct {
		pre int
		chr byte
		now int
	}
	this.State[0].deep = 0
	stack := make([]cell, this.Count)
	for i, j := 0, 1; j < this.Count; i++ {
		t := stack[i].now
		p := &this.State[t]
		if p.next != nil {
			for k, v := range p.next {
				this.State[v].deep = p.deep + 1
				stack[j] = cell{t, k, v}
				j++
			}
		}
	}
	for i := 1; i < this.Count; i++ {
		v := stack[i]
		s := &this.State[v.now] // 当前节点
		if v.pre != 0 {
			p := &this.State[v.pre] // 当前节点的父节点
			t := p.fail             // 父节点的fail指针
			for {
				if t < 0 {
					s.fail = 0
					break
				}
				p = &this.State[t]
				t = p.fail
				if p.next != nil {
					if n, ok := p.next[v.chr]; ok {
						s.fail = n
						break
					}
				}
			}
		} else {
			s.fail = 0
		}
	}
	this.State[0].fail = 0
}

// 搜索第一个匹配的位置，返回起始位置下标和模式串的下标，没有返回-1，-1
func (this *AC) Index(s []byte) (int, int) {
	t := 0
	for i, l := 0, len(s); i < l; {
		p := &this.State[t]
		if p.over >= 0 {
			return i - p.deep, p.over
		}
		if j, ok := p.next[s[i]]; ok {
			i, t = i+1, j
		} else {
			t = p.fail
		}
	}
	return -1, -1
}

// 搜索所有匹配的位置，[2]int成员分别是起始位置下标和模式串的下标，没有返回nil
func (this *AC) Find(s []byte) (o [][2]int) {
	t := 0
	for i, l := 0, len(s); i < l; {
		p := &this.State[t]
		if p.over >= 0 {
			o = append(o, [2]int{i - p.deep, p.over})
		}
		if j, ok := p.next[s[i]]; ok {
			i, t = i+1, j
		} else {
			t = p.fail
		}
	}
	return
}

// 为了方便存储Trie树，将之做成了list形式，所以必须使用本方法来提供新的节点
func (this *AC) newNode() int {
	n := this.Count
	this.State = append(this.State, node{deep: -1, over: -1, fail: -1})
	this.Count++
	return n
}

// 以某个节点为根节点，向下生成Trie树，支持字符集匹配
func (this *AC) incStr(p, n int, s string) {
	st := &this.State[n]
	if s == "" {
		st.over = p
		return
	}
	head, tail := s[0], s[1:]
	if st.next == nil {
		st.next = make(map[byte]int)
	}
	i, ok := st.next[head]
	if !ok {
		i = this.newNode()           // 注意本行执行后，st指针很可能失效！
		this.State[n].next[head] = i // 所以要用回this.State[n]
	}
	this.incStr(p, i, tail)
}
