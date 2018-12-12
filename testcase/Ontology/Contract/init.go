package Contract

import "github.com/ontio/ontology-test/testframework"

func TestContract() {
	testframework.TFramework.RegTestCase("TestGetScript", TestGetScript)
	testframework.TFramework.RegTestCase("TestMigrate", TestMigrate)
}
