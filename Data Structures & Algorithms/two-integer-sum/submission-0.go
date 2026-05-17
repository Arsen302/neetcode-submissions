func twoSum(nums []int, target int) []int {
	result := make(map[int]int) 

	for index, num := range nums {
		compile := target - num
		
		if i, ok := result[compile]; ok {
			return []int{i, index}
		}	

		result[num] = index
	}

	return nil
}