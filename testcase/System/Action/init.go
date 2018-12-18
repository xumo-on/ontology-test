package Action

import "github.com/xumo-on/ontology-test/testframework"

func TestAction() {
	testframework.TFramework.RegTestCase("TestRegisterAction", TestRegisterAction)
}
