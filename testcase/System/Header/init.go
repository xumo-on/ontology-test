package Header

import "github.com/ontio/ontology-test/testframework"

func TestHeader() {
	testframework.TFramework.RegTestCase("TestGetIndex", TestGetIndex)
	testframework.TFramework.RegTestCase("TestGetHash", TestGetHash)
	testframework.TFramework.RegTestCase("TestGetPrevHash", TestGetPrevHash)
	testframework.TFramework.RegTestCase("TestGetTimestamp", TestGetTimestamp)
}
