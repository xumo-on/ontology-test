package Runtime

import "github.com/xumo-on/ontology-test/testframework"

func TestRuntime() {
	testframework.TFramework.RegTestCase("TestBase58ToAddress", TestBase58ToAddress)
	testframework.TFramework.RegTestCase("TestAddressToBase58", TestAddressToBase58)
	testframework.TFramework.RegTestCase("TestGetRandomHash", TestGetRandomHash)
}
