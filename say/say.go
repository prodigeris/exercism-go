package say

func Say(n int64) (string, bool) {
	text := say(n)
	return text, text != ""
}

func say(n int64) string {
	// out of range
	if n < 0 || n > 999999999999 {
		return ""
	}

	basics := map[int64]string{
		0:  "zero",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "thirty",
		40: "forty",
		50: "fifty",
		60: "sixty",
		70: "seventy",
		80: "e	ighty",
		90: "ninety",
	}

	if v, ok := basics[n]; ok {
		return v
	}

	limits := []int64{1000000000, 1000000, 1000, 100}
	limitMap := map[int64]string{
		1000000000: "billion",
		1000000:    "million",
		1000:       "thousand",
		100:        "hundred",
	}

	for _, k := range limits {
		v := limitMap[k]
		if n < k {
			continue
		}
		remainder := n % k
		digit := (n - remainder) / k

		if remainder == 0 {
			return say(digit) + " " + v
		}

		return say(digit) + " " + v + " " + say(remainder)

	}

	// tens

	if n < 100 {
		remainder := n % 10
		whole := n - remainder
		return say(whole) + "-" + say(remainder)
	}

	return ""
}
