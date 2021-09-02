package main

import (
	"fmt"
)

type record struct {
	pos, value int
}

func trap(height []int) int {

	stack := []record{}

	if len(height) == 0 {
		return 0
	}

	sum := 0
	for i := 1; i < len(height); i++ {
		if height[i] < height[i-1] {
			// push
			fmt.Printf("slope %d\n", i)
			stack = append(stack, record{i, height[i]})
		} else if height[i] > height[i-1] {
			fmt.Printf("raise %d\n", i)
			if len(stack) > 0 {
				j := 0
				for j = len(stack) - 1; j >= 0 && stack[j].value <= height[i]; j-- {
					s := stack[j]
					lastPos := s.pos
					diff := height[i]
					if diff > height[lastPos-1] {
						diff = height[lastPos-1]
					}
					diff = diff - stack[j].value

					fmt.Printf("- sum from %d to %d diff %d\n", lastPos, i, diff)

					sum += diff * (i - lastPos)
				}
				if height[stack[j+1].pos-1] > height[i] {
					stack[j+1].value = height[i]
					j += 1
				}
				stack = stack[:j+1]

			}
		}

		fmt.Printf("- stack: %v\n", stack)
	}
	return sum
}

type Case struct {
	data     []int
	expected int
}

func main() {

	cases := []Case{
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, c := range cases {
		result := trap(c.data)
		expect := c.expected
		if result == expect {
			fmt.Println("pass")
		} else {
			fmt.Printf("fail: expected %d, was %d\n", expect, result)
		}
	}
}
