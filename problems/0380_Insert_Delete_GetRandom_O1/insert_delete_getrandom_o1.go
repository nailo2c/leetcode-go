package problem0380

import "math/rand"

// RandomizedSet is a struct
type RandomizedSet struct {
	nums   []int
	record map[int]int
}

// Constructor : Initialize your data structure here
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, make(map[int]int)}
}

// Insert : Inserts a value to the set. Returns true if the set did not already contain the specified element.
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.record[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.record[val] = len(this.nums) - 1
	return true
}

// Remove : Removes a value from the set. Returns true if the set contained the specified element.
func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.record[val]; !ok {
		return false
	} else {
		swap := this.nums[len(this.nums)-1]
		this.record[swap] = i
		this.nums[i] = swap
		this.nums = this.nums[0 : len(this.nums)-1]

		delete(this.record, val)
	}
	return true
}

// GetRandom : Get a random element from the set.
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Int()%len(this.record)]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
