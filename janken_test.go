package janken

func TestJanken01(t *testing.T) {
	expected := 0
	input1 := 0
	input2 := 1
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test01 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken02(t *testing.T) {
	expected := 0
	input1 := 1
	input2 := 2
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test02 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken03(t *testing.T) {
	expected := 0
	input1 := 2
	input2 := 0
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test03 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken04(t *testing.T) {
	expected := 1
	input := 2
	input := 0
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test04 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken05(t *testing.T) {
	expected := 1
	input1 := 0
	input2 := 1
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test05 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken06(t *testing.T) {
	expected := 1
	input1 := 1
	input2 := 2
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test06 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken07(t *testing.T) {
	expected := 2
	input1 := 0
	input2 := 0
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test07 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken07(t *testing.T) {
	expected := 2
	input1 := 1
	input2 := 1
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test07 fail expected: %d, result: %d\n", expected, result)
	}
}

func TestJanken08(t *testing.T) {
	expected := 2
	input1 := 2
	input2 := 2
	result := judge_win_lose(input1,input2)
	if expected != result {
		t.Errorf("Test08 fail expected: %d, result: %d\n", expected, result)
	}
}
