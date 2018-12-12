package Native

import "github.com/ontio/ontology-test/testframework"

func TestNative() {
	testframework.TFramework.RegTestCase("TestInvoke", TestInvoke)
}
