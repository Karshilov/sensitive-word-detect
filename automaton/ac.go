package automaton

type ACAutomaton struct {
	nxt  [](map[rune]int)
	fail []int
	tag  []int
	ncnt int
}

func (ac *ACAutomaton) Reserve(n int) {
	ac.nxt = make([]map[rune]int, n)
	ac.fail = make([]int, n)
	ac.tag = make([]int, n)
	ac.ncnt = 1
}

func (ac *ACAutomaton) Insert(keyword string) {
	u := 1
	for _, v := range keyword {
		if _, ok := ac.nxt[u][v]; !ok {
			if ac.nxt[u] == nil {
				ac.nxt[u] = make(map[rune]int)
			}
			ac.ncnt = ac.ncnt + 1
			ac.nxt[u][v] = ac.ncnt
		}
		u = ac.nxt[u][v]
	}
	ac.tag[u]++
}

func (ac *ACAutomaton) Build() {
	q := NewArrayQueue()

	for _, v := range ac.nxt[1] {
		q.Push(v)
		ac.fail[v] = 1
	}

	ac.fail[1] = 1

	for !q.Empty() {
		u := q.Front()
		q.Pop()

		for k, v := range ac.nxt[u] {
			faiu := ac.fail[u]
			for faiu > 1 && ac.nxt[faiu][k] == 0 {
				faiu = ac.fail[faiu]
			}

			if ac.nxt[faiu][k] != 0 {
				ac.fail[v] = ac.nxt[faiu][k]
			} else {
				ac.fail[v] = 1
			}

			q.Push(v)
		}
	}
}

func (ac *ACAutomaton) Query(str string) int {
	ans := 0

	u := 1
	for _, v := range str {
		for u > 1 && ac.nxt[u][v] == 0 {
			u = ac.fail[u]
		}

		if ac.nxt[u][v] != 0 {
			u = ac.nxt[u][v]
			for m := u; m > 1 && ac.tag[m] != -1; {
				ans += ac.tag[m]
				ac.tag[m] = -1
				m = ac.fail[m]
			}
		}
	}
	return ans
}
