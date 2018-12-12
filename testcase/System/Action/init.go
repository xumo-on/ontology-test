package Action

import "github.com/ontio/ontology-test/testframework"

func TestAction() {
	testframework.TFramework.RegTestCase("TestRegisterAction", TestRegisterAction)
}
