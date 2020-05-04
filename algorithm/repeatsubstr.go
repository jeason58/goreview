package algorithm

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func LengthOfLongestSubstring(s string) int {
	cMap := make(map[byte]int, 0)
	maxLen := 0
	lk, rk := 0, -1
	for ; lk < len(s); lk++ {
		for rk+1 < len(s) && cMap[s[rk+1]] == 0 {
			cMap[s[rk+1]]++
			rk++
		}

		if rk+1-lk > maxLen {
			maxLen = rk+1-lk
		}

		delete(cMap, s[lk])
	}

	if (rk+1-lk) > maxLen {
		maxLen = rk+1-lk
	}

	return maxLen
}