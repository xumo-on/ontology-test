package Blockchain

import "github.com/xumo-on/ontology-test/testframework"

func TestBlockchain() {
	testframework.TFramework.RegTestCase("TestGetHeight", TestGetHeight)
	testframework.TFramework.RegTestCase("TestGetHeader", TestGetHeader)
	testframework.TFramework.RegTestCase("TestGetBlock", TestGetBlock)
	testframework.TFramework.RegTestCase("TestGetTransaction", TestGetTransaction)
	testframework.TFramework.RegTestCase("TestGetContract", TestGetContract)
	testframework.TFramework.RegTestCase("TestGetTransactionHeight", TestGetTransactionHeight)
}
