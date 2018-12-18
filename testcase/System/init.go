package System

import (
	"github.com/xumo-on/ontology-test/testcase/System/Action"
	"github.com/xumo-on/ontology-test/testcase/System/App"
	"github.com/xumo-on/ontology-test/testcase/System/Block"
	"github.com/xumo-on/ontology-test/testcase/System/Blockchain"
	"github.com/xumo-on/ontology-test/testcase/System/Contract"
	"github.com/xumo-on/ontology-test/testcase/System/ExecutionEngine"
	"github.com/xumo-on/ontology-test/testcase/System/Header"
	"github.com/xumo-on/ontology-test/testcase/System/Runtime"
	"github.com/xumo-on/ontology-test/testcase/System/Storage"
	"github.com/xumo-on/ontology-test/testcase/System/StorageContext"
	"github.com/xumo-on/ontology-test/testcase/System/Transaction"
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
