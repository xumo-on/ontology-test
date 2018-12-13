package Header

import "github.com/xumo-on/ontology-test/testframework"

func TestHeader() {
	testframework.TFramework.RegTestCase("TestGetVersion", TestGetVersion)
	testframework.TFramework.RegTestCase("TestGetMerkleRoot", TestGetMerkleRoot)
	testframework.TFramework.RegTestCase("TestGetConsensusData", TestGetConsensusData)
	testframework.TFramework.RegTestCase("TestGetNextConsensus", TestGetNextConsensus)
}
