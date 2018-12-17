package Contract

import "github.com/ontio/ontology-test/testframework"

func TestContract() {
	testframework.TFramework.RegTestCase("TestGetStorageContext", TestGetStorageContext)
	testframework.TFramework.RegTestCase("TestContractDestroy", TestContractDestroy)
}
