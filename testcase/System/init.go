package System

import (
	"github.com/ontio/ontology-test/testcase/System/Action"
	"github.com/ontio/ontology-test/testcase/System/App"
	"github.com/ontio/ontology-test/testcase/System/Block"
	"github.com/ontio/ontology-test/testcase/System/Blockchain"
	"github.com/ontio/ontology-test/testcase/System/Contract"
	"github.com/ontio/ontology-test/testcase/System/ExecutionEngine"
	"github.com/ontio/ontology-test/testcase/System/Header"
	"github.com/ontio/ontology-test/testcase/System/Runtime"
	"github.com/ontio/ontology-test/testcase/System/Storage"
	"github.com/ontio/ontology-test/testcase/System/StorageContext"
	"github.com/ontio/ontology-test/testcase/System/Transaction"
)

func TestSystem() {
	Action.TestAction()
	App.TestApp()
	Blockchain.TestBlockchain()
	Block.TestBlock()
	Contract.TestContract()
	ExecutionEngine.TestExecutionEngine()
	Header.TestHeader()
	Runtime.TestRuntime()
	StorageContext.TestStorageContext()
	Storage.TestStorage()
	Transaction.TestTransaction()
}
