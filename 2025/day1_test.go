package main

import (
	"testing"
)

func TestDay1Pt1(t *testing.T) {
	const start = 50
	t.Run("Test password 1", func(t *testing.T) {
		const expect = 1

		instructions := make([]string, 1)
		instructions[0] = "L50"

		actual := decodePt1(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test password 2", func(t *testing.T) {
		const expect = 2

		instructions := make([]string, 4)
		instructions[0] = "L68"
		instructions[1] = "R18"
		instructions[2] = "R5"
		instructions[3] = "L5"

		actual := decodePt1(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test example from web site", func(t *testing.T) {
		const expect = 3

		instructions := make([]string, 0)
		instructions = append(
			instructions,
			"L68",
			"L30",
			"R48",
			"L5",
			"R60",
			"L55",
			"L1",
			"L99",
			"R14",
			"L82",
		)

		actual := decodePt1(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test turn right more than 100 steps", func(t *testing.T) {
		const expect = 2
		instructions := make([]string, 0)
		instructions = append(instructions, "R800", "R50", "R500")

		actual := decodePt1(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test turn left more than 100 steps", func(t *testing.T) {
		const expect = 2
		instructions := make([]string, 0)
		instructions = append(instructions, "L800", "L50", "L500")

		actual := decodePt1(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})
}

func TestDay1Pt2(t *testing.T) {
	const start = 50
	t.Run("Test password 1", func(t *testing.T) {
		const expect = 1

		instructions := make([]string, 1)
		instructions[0] = "L50"

		actual := decodePt2(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test password 2", func(t *testing.T) {
		const expect = 3

		instructions := make([]string, 4)
		instructions[0] = "L68" // 82, +1
		instructions[1] = "R18" // 0, +1
		instructions[2] = "R5"  //, 5, +0
		instructions[3] = "L5"  //, 0, +1

		actual := decodePt2(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test example from web site", func(t *testing.T) {
		const expect = 6

		instructions := make([]string, 0)
		instructions = append(
			instructions,
			"L68", // 82, +1
			"L30", // 52, +0
			"R48", // 0, +1
			"L5",  // 95, +0
			"R60", // 55, +1
			"L55", // 0, +1
			"L1",  // 99, +0
			"L99", // 0, +1
			"R14", // 14, +0
			"L82", // 32, +1
		)

		actual := decodePt2(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test turn right more than 100 steps", func(t *testing.T) {
		const expect = 14
		instructions := make([]string, 0)
		instructions = append(instructions, "R800", "R50", "R500")

		actual := decodePt2(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test turn left more than 100 steps", func(t *testing.T) {
		const expect = 14
		instructions := make([]string, 0)
		instructions = append(instructions, "L800", "L50", "L500")

		actual := decodePt2(start, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test rotate left entire circle 5 times", func(t *testing.T) {
		const expect = 5
		instructions := make([]string, 0)
		instructions = append(instructions, "L500")

		actual := decodePt2(0, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})

	t.Run("Test rotate right entire circle 5 times", func(t *testing.T) {
		const expect = 5
		instructions := make([]string, 0)
		instructions = append(instructions, "R500")

		actual := decodePt2(0, instructions)

		if expect != actual {
			t.Errorf("got %v expect %v", actual, expect)
		}
	})
}
