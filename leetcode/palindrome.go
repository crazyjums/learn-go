package leetcode

// IsPalindrome 剑指 Offer II 018. 有效的回文
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	s = filter(s)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func filter(s string) string {
	var bs []byte
	for _, v := range s {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {
			if v >= 'A' && v <= 'Z' {
				v += 32
			}
			bs = append(bs, byte(v))
		}
	}
	return string(bs)
}

// CanPermutePalindrome 面试题 01.04. 回文排列
func CanPermutePalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	cm := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cm[s[i]]++
	}

	odd := 0
	even := 0
	for _, v := range cm {
		if v%2 == 1 {
			odd++
		} else {
			even++
		}
	}

	return odd <= 1
}
