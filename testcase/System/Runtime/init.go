package Runtime

import "github.com/ontio/ontology-test/testframework"

func TestRuntime() {
	testframework.TFramework.RegTestCase("TestCheckWitness", TestCheckWitness)
	//testframework.TFramework.RegTestCase("TestLog", TestLog)
	testframework.TFramework.RegTestCase("TestNotify", TestNotify)
	testframework.TFramework.RegTestCase("TestGetTime", TestGetTime)
	testframework.TFramework.RegTestCase("TestSerialize", TestSerialize)
}
