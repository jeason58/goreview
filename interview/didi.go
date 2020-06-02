package interview


/*
N长的线段，切成M段，有多少种切法
举例： 长为3的线段，切成2段，一共有1 2，2 1 两种切法

len = 4, seg = 2
1-3,  2-2,  3-1
*/

func SplitLine(length, seg int) int {
	if length < seg || seg <= 0 {
		return 0
	}

	if seg == 1 {
		return 1
	}

	sum := 0
	for i := 1; i < length; i++ {
		sum += SplitLine(length-i, seg-1)
	}

	return sum
}
