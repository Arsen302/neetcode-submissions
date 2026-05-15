func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	fmt.Println(seen)
	
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{} 
		fmt.Println(num)
	}
	return false
}
