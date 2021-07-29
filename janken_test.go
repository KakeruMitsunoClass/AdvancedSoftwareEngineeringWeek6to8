package main

import "testing"

//userの勝ち
func TestJanken01(t *testing.T) {
	expected := 0
	input1 := 0
	input2 := 1
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test01 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken02(t *testing.T) {
	expected := 0
	input1 := 1
	input2 := 2
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test02 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken03(t *testing.T) {
	expected := 0
	input1 := 2
	input2 := 0
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test03 fail expected: %d, result: %d\n", expected, result)
	}
}

//cpuの勝ち
func TestJanken04(t *testing.T) {
	expected := 1
	input1 := 1
	input2 := 0
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test04 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken05(t *testing.T) {
	expected := 0
	input1 := 0
	input2 := 1
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test05 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken06(t *testing.T) {
	expected := 1
	input1 := 0
	input2 := 2
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test06 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken07(t *testing.T) {
	expected := 2
	input1 := 0
	input2 := 0
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test07 fail expected: %d, result: %d\n", expected, result)
	}
}

//あいこ
func TestJanken08(t *testing.T) {
	expected := 2
	input1 := 1
	input2 := 1
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test08 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken09(t *testing.T) {
	expected := 2
	input1 := 2
	input2 := 2
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test09 fail expected: %d, result: %d\n", expected, result)
	}
}

//エラー
func TestJanken10(t *testing.T) {
	expected := -1
	input1 := 4
	input2 := 2
	result := judge_win_lose(input1, input2)
	if expected != result {
		t.Errorf("Test10 fail expected: %d, result: %d\n", expected, result)
	}
}

//user入力
func TestJanken11(t *testing.T) {
	expected := 0
	input := "r"
	result := get_user(input)
	if expected != result {
		t.Errorf("Test11 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken12(t *testing.T) {
	expected := 1
	input := "s"
	result := get_user(input)
	if expected != result {
		t.Errorf("Test12 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken13(t *testing.T) {
	expected := 2
	input := "p"
	result := get_user(input)
	if expected != result {
		t.Errorf("Test13 fail expected: %d, result: %d\n", expected, result)
	}
}

//cpu入力
func TestJanken14(t *testing.T) {
	expected0 := 0
	expected1 := 1
	expected2 := 2
	result := random_cpu(3)
	if !((expected0 == result) || (expected1 == result) || (expected2 == result)) {
		t.Errorf("Test14 fail expected: %d,%d,%d result: %d\n", expected0, expected1, expected2, result)
	}
}

func TestJanken15(t *testing.T) {
	expected := 0
	input1 := 5
	input2 := 3
	input3 := 2
	result := judge(input1, input2, input3)
	if expected != result {
		t.Errorf("Test15 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken16(t *testing.T) {
	expected := 1
	input1 := 5
	input2 := 7
	input3 := 2
	result := judge(input1, input2, input3)
	if expected != result {
		t.Errorf("Test16 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken17(t *testing.T) {
	expected := 2
	input1 := 5
	input2 := 5
	input3 := 2
	result := judge(input1, input2, input3)
	if expected != result {
		t.Errorf("Test17 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken18(t *testing.T) {
	expected0 := 0
	expected1 := 1
	expected2 := 2
	expected3 := 3
	result := random_cpu(4)
	if !((expected0 == result) || (expected1 == result) || (expected2 == result) || (expected3 == result)) {
		t.Errorf("Test18 fail expected: %d,%d,%d result: %d\n", expected0, expected1, expected2, result)
	}
}

func TestJanken19(t *testing.T) {
	input1 := 0
	input2 := 1
	expected := 1
	result := judge_win_lose_achimuitehoi(input1, input2) //cpu int, user int
	if expected != result {
		t.Errorf("Test19 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken20(t *testing.T) {
	input1 := 0
	input2 := 0
	expected := 0
	result := judge_win_lose_achimuitehoi(input1, input2) //cpu int, user int
	if expected != result {
		t.Errorf("Test20 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken21(t *testing.T) {
	input := "u"
	expected := 0
	result := get_user_achimuitehoi(input) //cpu int, user int
	if expected != result {
		t.Errorf("Test21 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken22(t *testing.T) {
	input := "r"
	expected := 2
	result := get_user_achimuitehoi(input) //cpu int, user int
	if expected != result {
		t.Errorf("Test22 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken23(t *testing.T) {
	input := "l"
	expected := 3
	result := get_user_achimuitehoi(input) //cpu int, user int
	if expected != result {
		t.Errorf("Test23 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken24(t *testing.T) {
	input := "d"
	expected := 1
	result := get_user_achimuitehoi(input) //cpu int, user int
	if expected != result {
		t.Errorf("Test20 fail expected: %d, result: %d\n", expected, result)
	}
}
