package cal

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	//调用
	res := addUpper(10)
	if res != 55 {
		//fmt.Printf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n"，55，res)
		t.Fatalf("AddUpper(10)执行错误，期望值=%v 实际值=%v\n", 55, res)
	}

	//如果正确，输出日志
	t.Logf("Addupper(18)执行正确...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被调用..")
}
