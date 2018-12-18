package Cond_loop

import (
	"github.com/xumo-on/ontology-test/testframework"
)

func TestCondLoop() {
	testframework.TFramework.RegTestCase("TestFor", TestFor)
	testframework.TFramework.RegTestCase("TestIfElse", TestIfElse)
	testframework.TFramework.RegTestCase("TestSwitch", TestSwitch)
}
