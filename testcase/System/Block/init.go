package Block

import "github.com/ontio/ontology-test/testframework"

func TestBlock() {
	testframework.TFramework.RegTestCase("TestGetTransactionCount", TestGetTransactionCount)
	testframework.TFramework.RegTestCase("TestGetTransactions & TestGetTransactionByIndex", TestGetTransactions)
}