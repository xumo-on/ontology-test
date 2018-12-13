package StorageContext

import "github.com/xumo-on/ontology-test/testframework"

func TestStorageContext() {
	testframework.TFramework.RegTestCase("TestAsReadOnly", TestAsReadOnly)
}
