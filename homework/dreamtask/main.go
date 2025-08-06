package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func singleNumberMy(nums []int) int {
	//定义一个map,用于存储元素的出现的次数
	elCountMap := make(map[int]int)

	//循环nums,将元素的个数添加到map
	for _, v := range nums {
		//如果elCountMap中找不到元素，说明元素第一次出现,如果找到了，就+1
		if _, ok := elCountMap[v]; ok {
			elCountMap[v]++
		} else {
			elCountMap[v] = 1
		}
	}
	//取出elCountMap中只出现一次的元素
	var res int
	for k, v := range elCountMap {
		if v == 1 {
			res = k
			break
		}
	}
	return res
}

func singleNumber(nums []int) int {
	var res int = 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func isPalindromeMy(x int) bool {
	//先转成string
	xStr := strconv.Itoa(x)
	var reserveNumStr string
	for i := len(xStr) - 1; i >= 0; i-- {
		reserveNumStr += string(xStr[i])
	}
	reserveNum, _ := strconv.Atoi(reserveNumStr)
	return reserveNum == x
}

func isPalindrome(x int) bool {
	reserverNum := 0
	for x > reserverNum {
		reserverNum = reserverNum*10 + x%10
		x = x / 10
	}
	return reserverNum == x || x == reserverNum/10
}

func isValid(s string) bool {
	sMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != sMap[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = compare(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func compare(prefix string, s string) string {
	//比较下2个字符串哪个长
	length := 0
	if len(prefix) < len(s) {
		length = len(prefix)
	} else {
		length = len(s)
	}
	index := 0
	for index < length && prefix[index] == s[index] {
		index++
	}
	return prefix[:index]
}

func plusOne_my(digits []int) []int {
	length := len(digits)
	sum := 0
	//将digits中的数按个个十百的数进行相加
	for i := 0; i < length; i++ {
		sum += digits[i] * int(math.Pow(10, float64(length-i-1)))
	}
	sum += 1
	sum_len := len(strconv.Itoa(sum))
	res := make([]int, sum_len)
	var sumStr = strconv.Itoa(sum)
	for i := 0; i < sum_len; i++ {
		res[i] = int(sumStr[i] - '0')
	}
	return res
}

func plusOne(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < length; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, length+1)
	digits[0] = 1
	return digits
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	slow := 1
	for fast := 1; fast < length; fast++ {
		if nums[fast] != nums[slow] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	prefix := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if prefix[1] < curr[0] {
			res = append(res, prefix)
			prefix = curr
		} else {
			prefix[1] = max(prefix[1], curr[1])
		}
	}
	res = append(res, prefix)
	return res
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	//只出现一次的数字
	//res := singleNumberMy([]int{2, 2, 1})
	//fmt.Println(res)

	//回文数
	//isPalindromeRes := isPalindromeMy(121)
	//fmt.Println(isPalindromeRes)
	//palindrome := isPalindrome(123321)
	//fmt.Println(palindrome)

	//有效的括号
	//fmt.Println(isValid("(]"))
	//fmt.Println(isValid("()[]{}"))
	//fmt.Println(isValid("([])"))

	//最长公共前缀
	//str := []string{"flower", "flow", "flight"}
	//prefix := longestCommonPrefix(str)
	//fmt.Println(prefix)

	//加一
	//digits := []int{1, 2, 3, 9, 9}
	//one := plusOne_my(digits)
	//fmt.Println(one)
	//官方解析
	//ints := plusOne(digits)
	//fmt.Println(ints)

	//删除有序数组中的重复项
	//nums := []int{1, 1, 3}
	//duplicates := removeDuplicates(nums)
	//fmt.Println(duplicates)

	//合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	mergeRes := merge(intervals)
	fmt.Println(mergeRes)

	//两数之和
	//nums := []int{2, 7, 11, 15}
	//target := 10
	//sumRes := twoSum(nums, target)
	//fmt.Println(sumRes)
}
