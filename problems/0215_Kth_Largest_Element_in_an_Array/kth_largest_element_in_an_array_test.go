package problem0215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_Problem0215(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{3, 2, 1, 5, 6, 4},
				two: 2,
			},
			a: ans{
				one: 5,
			},
		},
		question{
			p: para{
				one: []int{3, 3, 3, 3, 3, 3, 3, 3, 3},
				two: 1,
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findKthLargest(p.one, p.two))
	}
}

func Test_heap(t *testing.T) {
	ast := assert.New(t)

	h := new(maxHeap)

	i := 5
	h.Push(i)
	ast.Equal(i, h.Pop())

}

func Benchmark_findKthLargest(b *testing.B) {

	qs := []question{
		question{
			p: para{
				one: []int{3, 2, 1, 5, 6, 4},
				two: 2,
			},
			a: ans{
				one: 5,
			},
		},
	}

	for i := 0; i < b.N; i++ {
		for _, q := range qs {
			_, p := q.a, q.p
			findKthLargest(p.one, p.two)
		}
	}
}
