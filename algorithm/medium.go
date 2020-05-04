package algorithm

import "strings"

/** 1. 寻找最长回文子串 **/
// 例1
// 输入："cbbd"
// 输出："bb"
func LongestPalindrome(s string) string {
	for max := len(s); max > 0; max-- {
		for start := 0; start <= len(s)-max; start++ {
			found := true
			left := start
			right := left + max -1
			for ; left <= right; {
				if s[left] != s[right] {
					found = false
					break
				}
				left += 1
				right -= 1
			}

			if found {
				return s[start: start+max]
			}
		}
	}

	return ""
}



/** 2. Z字形变换字符串 **/
//
// 输入："LEETCODEISHIRING"
// rowNum=3时，变换为:
// L   C   I   R
// E T O E S I I G
// E   D   H   N
// 输出："LCIRETOESIIGEDHN"
//
// rowNum=4时，变换为:
// L     D     R
// E   O E   I I
// E C   I H   N
// T     S     G
// 输出："LDREOEIIECIHNTSG"

// 解法1
func ConvertAsZ1(s string, numRows int) string {
	unitLen := 2*numRows - 2
	if unitLen < 1 {
		return s
	}
	matrix := make([]string, 0)
	for i := 0; i < numRows; i++ {
		matrix = append(matrix, "")
	}

	for i := 1; i <= len(s); i++ {
		ch := s[i-1:i]
		tmp := i % unitLen
		if 0 < tmp && tmp <= numRows {
			matrix[tmp-1] = matrix[tmp-1]+ch
		} else {
			if tmp == 0 {
				tmp = unitLen
			}

			tgtRowIndex := 2*numRows - tmp - 1
			matrix[tgtRowIndex] = matrix[tgtRowIndex] + ch
		}
	}

	var res string
	for _, str := range matrix {
		res = res + str
	}

	return res
}

// 解法2
func ConvertAsZ2(s string, numRows int) string {
	if len(s) < 2 {
		return s
	}

	 matrix := make([]string, 0)
	 for i := 0; i < len(s); i++ {
	 	matrix = append(matrix, "")
	 }

	 i, flag := 0, -1
	 for n := 0; n < len(s); n++ {
	 	matrix[i] = matrix[i] + s[n:n+1]
	 	if i == 0 || i == numRows - 1 {
	 		flag = -flag
		}
		i += flag
	 }

	 return strings.Join(matrix, "")
}


/** 3. 水壶问题 **/
//有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
//如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
//
//你允许：
//装满任意一个水壶
//清空任意一个水壶
//从一个水壶向另外一个水壶倒水，直到装满或者倒空
//
//示例 1:
//输入: x = 3, y = 5, z = 4
//输出: True
//
//示例 2:
//输入: x = 2, y = 6, z = 5
//输出: False

func CanMeasureWater(x int, y int, z int) bool {
	if x == y && x != z && x*2 != z {
		return false
	}


	//over := false
	xc, yc := x, 0

	found := false
	for !found {
		if isOk(xc, yc, z) {
			found = true
			break
		}

		// xc ——> yc
		if xc > 0 && yc < y {
			if xc <= y-yc {
				yc += xc
				xc = 0
			} else {
				xc -= y-yc
				yc = y
			}

			if isOk(xc, yc, z) {
				found = true
				break
			}
		}

		// yc ——> 0
		if yc == y {
			yc = 0
			if isOk(xc, yc, z) {
				found = true
				break
			}
		}

		if xc == 0 && yc == 0 {
			//over = true
			break
		}

		// xc ——> x
		if xc == 0 {
			xc = x
			if isOk(xc, yc, z) {
				found = true
				break
			}
		}
	}

	return found
}
func isOk(xc, yc, z int) bool {
	if xc == z || yc == z || xc + yc == z {
		return true
	}

	return false
}


/** 4. 判断字符串是否有效 **/
var stackTop  = -1
var options = []string{"(", ")", " "}
func CheckValidString(s string) bool {
	if stackTop < -1 {
		return false
	}

	if s == "" {
		if stackTop == -1 {
			return true
		} else {
			return false
		}
	}

	curS := s[:1]
	nextS := ""
	if len(s) == 1 {
		nextS = ""
	} else {
		nextS = s[1:]
	}

	if curS == "(" {
		stackTop += 1
	} else if curS == ")" {
		stackTop -= 1
	} else if curS == "*" {
		top := stackTop
		for _, opt := range options {
			stackTop = top
			if opt == "(" {
				stackTop += 1
			} else if opt == ")" {
				stackTop -= 1
			}

			if CheckValidString(nextS) {
				return true
			}
		}
		return false
	}

	return CheckValidString(nextS)
}