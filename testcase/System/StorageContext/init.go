package StorageContext

import "github.com/ontio/ontology-test/testframework"

func TestStorageContext() {
	testframework.TFramework.RegTestCase("TestAsReadOnly", TestAsReadOnly)
}
