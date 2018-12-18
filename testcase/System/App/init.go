package App

import "github.com/xumo-on/ontology-test/testframework"

func TestApp() {
	testframework.TFramework.RegTestCase("TestRegisterAppCall", TestRegisterAppCall)
}
