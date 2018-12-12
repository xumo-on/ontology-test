package Transaction

import "github.com/ontio/ontology-test/testframework"

func TestTransaction() {
	testframework.TFramework.RegTestCase("TestGetType", TestGetType)
	//testframework.TFramework.RegTestCase("TestGetAttributes", TestGetAttributes)
}
