package App

import "github.com/ontio/ontology-test/testframework"

func TestApp() {
	testframework.TFramework.RegTestCase("TestRegisterAppCall", TestRegisterAppCall)
}
