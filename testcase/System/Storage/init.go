package Storage

import "github.com/ontio/ontology-test/testframework"

func TestStorage() {
	testframework.TFramework.RegTestCase("TestGetContext", TestGetContext)
	testframework.TFramework.RegTestCase("TestPut", TestPut)
	testframework.TFramework.RegTestCase("TestGet", TestGet)
	testframework.TFramework.RegTestCase("TestDelete", TestDelete)
}
