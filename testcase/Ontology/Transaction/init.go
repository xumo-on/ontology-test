package Transaction

import "github.com/xumo-on/ontology-test/testframework"

func TestTransaction() {
	testframework.TFramework.RegTestCase("TestGetType", TestGetType)
	//testframework.TFramework.RegTestCase("TestGetAttributes", TestGetAttributes)
}
