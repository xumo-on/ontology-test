/*
from boa.interop.System.Runtime import Notify
from boa.interop.Ontology.Transaction import GetType,GetAttributes
from boa.interop.Ontology.Runtime import GetCurrentBlockHash
from boa.interop.System.Blockchain import GetBlock
from boa.interop.System.Block import GetTransactionByIndex,GetTransactionCount


def Main(operation, args):
    if operation == 'getType':
        return getType()
    if operation == 'getAttributes':
        return getAttributes()
    return False

def getType():
    Hash = GetCurrentBlockHash()
    block = GetBlock(Hash)
    tx = GetTransactionByIndex(block ,0)
    Type = GetType(tx)
    Notify(Hash)
    Notify(Type)
    return True

def getAttributes():
    Hash = GetCurrentBlockHash()
    block = GetBlock(Hash)
    count = GetTransactionCount(block)
    Notify(count)
    if count != None:
        tx = GetTransactionByIndex(block ,0)
        attr = GetAttributes(tx)
        Notify(tx)
        return True
    return False
 */
package Transaction

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"strconv"
	"strings"
	"time"
)

var codeAddress common.Address

func TestGetType(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "57c56b6a00527ac46a51527ac46a00c307676574547970659c640900650b006c756661006c756659c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a52527ac46a52c3681c4f6e746f6c6f67792e5472616e73616374696f6e2e47657454797065616a53527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"getType", []interface{}{[]byte("getType")}})
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
	txType := uint64(block.Transactions[0].TxType)

	Type := events.Notify[1].States.(string)
	count = strings.Count(Type, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Type[i-2:i])
	}
	Type = strings.Join(s, "")
	Type1, err := strconv.ParseUint(Type, 16, 64)

	if txType != Type1 {
		ctx.LogError("TestGetType error")
		return false
	}

	return true
}

//func TestGetAttributes(ctx *testframework.TestFrameworkContext) bool {
//	//DeployContract
//	code := "59c56b6a00527ac46a51527ac46a00c307676574547970659c64090065ef006c7566616a00c30d676574417474726962757465739c640900650b006c756661006c756658c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a52527ac46a52c368224f6e746f6c6f67792e5472616e73616374696f6e2e47657441747472696275746573616a53527ac46a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c756659c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a52527ac46a52c3681c4f6e746f6c6f67792e5472616e73616374696f6e2e47657454797065616a53527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
//	codeAddress, _ = utils.GetContractAddress(code)
//
//	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
//	signer, err := ctx.GetDefaultAccount()
//	if err != nil {
//		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
//		return false
//	}
//
//	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
//		signer,
//		true,
//		code,
//		"TestDomainSmartContract",
//		"1.0",
//		"",
//		"",
//		"",
//	)
//
//	if err != nil {
//		ctx.LogError("TestDomainSmartContract DeploySmartContract error: %s", err)
//	}
//
//	//WaitForGenerateBlock
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
//		return false
//	}
//
//	//InvokeContract
//	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
//		signer,
//		codeAddress,
//		[]interface{}{"getType", []interface{}{[]byte("getType")}})
//	if err != nil {
//		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
//	}
//
//	//WaitForGenerateBlock
//	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
//	if err != nil {
//		ctx.LogError("TestInvokeSmartContract WaitForGenerateBlock error:%s", err)
//		return false
//	}
//
//	//GetEventOfContract
//	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
//	if err != nil {
//		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
//		return false
//	}
//	if events.State == 0 {
//		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
//		return false
//	}
//
//	Hash := events.Notify[0].States.(string)
//	count := strings.Count(Hash, "") - 1
//	s := []string{}
//	for i := count; i > 0; i -= 2 {
//		s = append(s, Hash[i-2:i])
//	}
//	Hash1 := strings.Join(s, "")
//	block, err := ctx.Ont.GetBlockByHash(Hash1)
//	if err != nil {
//		ctx.LogError("GetBlockByHash error:%s", err)
//		return false
//	}
//	txType := uint64(block.Transactions[0].TxType)
//
//	Type := events.Notify[1].States.(string)
//	count = strings.Count(Type, "") - 1
//	s = []string{}
//	for i := count; i > 0; i -= 2 {
//		s = append(s, Type[i-2:i])
//	}
//	Type = strings.Join(s, "")
//	Type1, err := strconv.ParseUint(Type, 16, 64)
//
//	if txType != Type1 {
//		ctx.LogError("TestGetType error")
//		return false
//	}
//
//	return true
//}
