package Block

import "github.com/xumo-on/ontology-test/testframework"

func TestBlock() {
	testframework.TFramework.RegTestCase("TestGetTransactionCount", TestGetTransactionCount)
	testframework.TFramework.RegTestCase("TestGetTransactions & TestGetTransactionByIndex", TestGetTransactions)
}
