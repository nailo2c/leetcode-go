package problem0528

// Solution is a struct
type Solution struct {
    vals []int
    maxVal int
}


func Constructor(w []int) Solution {
    vals := []int{}
    current := 0
    for _, v := range w {
        current += v
        vals = append(vals, current)
    }
    return Solution{
        vals: vals,
        maxVal: current,
    }
}


func (this *Solution) PickIndex() int {
    num := rand.Intn(this.maxVal)
    return findGreater(this.vals, num)
}

func findGreater(vals []int, num int) int {
    l, r := 0, len(vals)-1
    for l < r {
        m := (r-l)/2 + l
        if vals[m] > num {
            r = m
        } else {
            l = m+1
        }
    }
    return l
}