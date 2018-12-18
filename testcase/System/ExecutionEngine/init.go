package ExecutionEngine

import "github.com/xumo-on/ontology-test/testframework"

func TestExecutionEngine() {
	testframework.TFramework.RegTestCase("TestGetScriptContainer", TestGetScriptContainer)
	testframework.TFramework.RegTestCase("TestGetExecutingScriptHash", TestGetExecutingScriptHash)
	testframework.TFramework.RegTestCase("TestGetCallingScriptHash&GetEntryScriptHash", TestGetCallingScriptHash)
}
