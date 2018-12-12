/*
from boa.interop.System.Runtime import Notify
from boa.interop.Ontology.Header import GetVersion,GetMerkleRoot,GetConsensusData,GetNextConsensus
from boa.interop.Ontology.Runtime import GetCurrentBlockHash
from boa.interop.System.Blockchain import GetHeight,GetHeader

def Main(operation, args):
    if operation == 'getVersion':
        return getVersion()
    if operation == 'getMerkleRoot':
        return getMerkleRoot()
    if operation == 'getConsensusData':
        return getConsensusData()
    if operation == 'getNextConsensus':
        return getNextConsensus()
    return False

def getVersion():
    Hash = GetCurrentBlockHash()
    header = GetHeader(Hash)
    version = GetVersion(header)
    Notify(version)
    return True

def getMerkleRoot():
    Hash = GetCurrentBlockHash()
    header = GetHeader(Hash)
    merkleRoot = GetMerkleRoot(header)
    Notify(Hash)
    Notify(merkleRoot)
    return True

def getConsensusData():
    Hash = GetCurrentBlockHash()
    header = GetHeader(Hash)
    ConsensusData = GetConsensusData(header)
    Notify(Hash)
    Notify(ConsensusData)
    return True

def getNextConsensus():
    Hash = GetCurrentBlockHash()
    header = GetHeader(Hash)
    NextConsensus = GetNextConsensus(header)
    Notify(Hash)
    Notify(NextConsensus)
    return True
 */
package Header

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"strconv"
	"strings"
	"time"
)

var	codeAddress common.Address

func TestGetVersion(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "5dc56b6a00527ac46a51527ac46a00c30a67657456657273696f6e9c6409006590026c7566616a00c30d6765744d65726b6c65526f6f749c64090065bd016c7566616a00c310676574436f6e73656e737573446174619c64090065e4006c7566616a00c3106765744e657874436f6e73656e7375739c640900650b006c756661006c756657c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c368204f6e746f6c6f67792e4865616465722e4765744e657874436f6e73656e737573616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961006c756658c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c368204f6e746f6c6f67792e4865616465722e476574436f6e73656e73757344617461616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681d4f6e746f6c6f67792e4865616465722e4765744d65726b6c65526f6f74616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756657c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681a4f6e746f6c6f67792e4865616465722e47657456657273696f6e616a52527ac46a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
	codeAddress, _ = utils.GetContractAddress(code)

	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		true,
		code,
		"TestDomainSmartContract",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestDomainSmartContract DeploySmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getVersion", []interface{}{[]byte("getVersion")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}

	//GetEventOfContract
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	notify := events.Notify[0].States.(string)
	if notify != "00" {
		ctx.LogError("TestGetVersion error")
		return false
	}

	return true
}

func TestGetMerkleRoot(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getMerkleRoot", []interface{}{[]byte("getMerkleRoot")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}

	//GetEventOfContract
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	txRoot := block.Header.TransactionsRoot.ToHexString()

	merkleRoot := events.Notify[1].States.(string)
	count = strings.Count(merkleRoot, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, merkleRoot[i-2:i])
	}
	merkleRoot1 := strings.Join(s, "")

	if txRoot != merkleRoot1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}

func TestGetConsensusData(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getConsensusData", []interface{}{[]byte("getConsensusData")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}

	//GetEventOfContract
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	cd := block.Header.ConsensusData

	ConsensusData := events.Notify[1].States.(string)
	count = strings.Count(ConsensusData, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, ConsensusData[i-2:i])
	}
	ConsensusData1, err := strconv.ParseUint(strings.Join(s, ""), 16, 64)
	if err != nil {
		ctx.LogError("ConsensusDataToUint error:%s", err)
		return false
	}

	if cd != ConsensusData1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}

func TestGetNextConsensus(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getNextConsensus", []interface{}{[]byte("getNextConsensus")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
		return false
	}

	//GetEventOfContract
	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	Hash := events.Notify[0].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	Hash1 := strings.Join(s, "")
	block, err := ctx.Ont.GetBlockByHash(Hash1)
	if err != nil {
		ctx.LogError("GetBlockByHash error:%s", err)
		return false
	}
	nc := block.Header.NextBookkeeper.ToHexString()

	NextConsensus := events.Notify[1].States.(string)
	count = strings.Count(NextConsensus, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, NextConsensus[i-2:i])
	}
	NextConsensus1 := strings.Join(s, "")

	if nc != NextConsensus1 {
		ctx.LogError("TestGetMerkleRoot error")
		return false
	}

	return true
}
