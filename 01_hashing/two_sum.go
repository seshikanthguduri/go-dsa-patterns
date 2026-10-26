// Pattern: Hashing — trade space for time. One pass, store complements.
//
// Problem (LeetCode 1, Two Sum): given nums and target, return the indices
// of the two numbers that add up to target. Exactly one solution exists.
//
// Idea: for each value x at index i, the partner we need is target-x.
// If we have already seen it, we are done; otherwise record x -> i.
// Time O(n), space O(n).
package hashing

func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}
