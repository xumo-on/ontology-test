package Native

import "github.com/xumo-on/ontology-test/testframework"

func TestNative() {
	testframework.TFramework.RegTestCase("TestInvoke", TestInvoke)
}
