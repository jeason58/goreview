package sort

//归并排序（分治思想：拆分——>合并）
func sortByMerge(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums)/2
	return merge(sortByMerge(nums[:mid]), sortByMerge(nums[mid:]))
}

func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))

	i, j, p := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res[p] = a[i]
			i++
		} else {
			res[p] = b[j]
			j++
		}
		p++
	}

	for i < len(a) {
		res[p] = a[i]
		p, i = p+1, i+1
	}
	for j < len(b) {
		res[p] = b[j]
		p, j = p+1, j+1
	}

	return res
}

