package Contract

import "github.com/xumo-on/ontology-test/testframework"

func TestContract() {
	testframework.TFramework.RegTestCase("TestGetStorageContext", TestGetStorageContext)
	testframework.TFramework.RegTestCase("TestContractDestroy", TestContractDestroy)
}
