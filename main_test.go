package main

import "testing"

// {테스트할 소스코드 파일 명}_test 파일에 정의된 Example + Method 명 -> Example 메소드 생성됨. go test 실행시 같이 실행됨.
// "Output:" 뒤의 output 과 실제 output 이 일치하는지 체크해줌.
// import 없이도 main 패키지 메소드 호출 가능.
func ExampleHello() {
	Hello()
	// Output: Let's study golang!
}

func TestSomePrivateMethod(t *testing.T) {
	answer := "Private Method Call Example"
	got := somePrivateMethod()
	if got != answer {
		t.Errorf("somePrivateMethod(): %s, want %s", got, answer)
	}
}
