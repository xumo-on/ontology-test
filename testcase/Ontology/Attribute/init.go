package Attribute

import "github.com/ontio/ontology-test/testframework"

func TestAttribute() {
	testframework.TFramework.RegTestCase("TestGetUsage", TestGetUsage)
	testframework.TFramework.RegTestCase("TestGetData", TestGetData)
}
