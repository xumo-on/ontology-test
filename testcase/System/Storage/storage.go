/*
from boa.interop.System.Runtime import Notify
from boa.interop.System.Storage	import GetContext,GetReadOnlyContext,Get,Put,Delete
from boa.interop.System.Contract import GetStorageContext
from boa.interop.System.Blockchain import GetContract
from boa.interop.System.ExecutionEngine	import GetExecutingScriptHash
from boa.interop.System.StorageContext import AsReadOnly

def Main(operation, args):
    if operation == 'getContext':
        return getContext()
    if operation == 'put':
        return put()
    if operation == 'get':
        return get()
    if operation == 'delete':
        return delete()
    return False

context = GetContext()

def getContext():
    script = GetExecutingScriptHash()
    contract = GetContract(script)
    context1 = GetStorageContext(contract)
    Notify(context)
    Notify(context1)
    return True

def put():
    Put(context, 'get', 'aaaaa')
    return True

def get():
    Put(context, 'get', 'aaaaa')
    value = Get(context, 'get')
    Notify(value)
    return True

def delete():
    Put(context, 'get', 'aaaaa')
    Delete(context, 'get')
    value = Get(context, 'get')
    Notify(value)
    return True
 */
package Storage

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

var	codeAddress common.Address

func TestGetContext(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "5dc56b6a00527ac46a51527ac46a00c30a676574436f6e746578749c64090065c2016c7566616a00c3037075749c6409006564016c7566616a00c3036765749c64090065c9006c7566616a00c30664656c6574659c640900650b006c756661006c756658c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac46a00c3036765740561616161615272681253797374656d2e53746f726167652e507574616a00c3036765747c681553797374656d2e53746f726167652e44656c657465616a00c3036765747c681253797374656d2e53746f726167652e476574616a51527ac46a51c3681553797374656d2e52756e74696d652e4e6f7469667961516c756657c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac46a00c3036765740561616161615272681253797374656d2e53746f726167652e507574616a00c3036765747c681253797374656d2e53746f726167652e476574616a51527ac46a51c3681553797374656d2e52756e74696d652e4e6f7469667961516c756655c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac46a00c3036765740561616161615272681253797374656d2e53746f726167652e50757461516c756659c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a51527ac46a51c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a52527ac46a52c3682153797374656d2e436f6e74726163742e47657453746f72616765436f6e74657874616a53527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"getContext", []interface{}{[]byte("getContext")}})
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

	context := events.Notify[0].States.(string)
	context1 := events.Notify[1].States.(string)

	if context != context1 {
		ctx.LogError("TestGetContext error")
		return false
	}

	return true
}

func TestPut(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}
	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"put", []interface{}{[]byte("put")}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//WaitForGenerateBlock
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestDomainSmartContract WaitForGenerateBlock error: %s", err)
		return false
	}

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value := hex.EncodeToString(svalue)

	if value != "6161616161" {
		ctx.LogError("TestPut error")
		return false
	}
	return true
}

func TestGet(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"get", []interface{}{[]byte("get")}})
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

	value := events.Notify[0].States.(string)

	if value != "6161616161" {
		ctx.LogError("TestPut error")
		return false
	}
	return true
}

func TestDelete(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"delete", []interface{}{[]byte("delete")}})
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

	value := events.Notify[0].States.(string)
	if value != "" {
		ctx.LogError("TestDelete error")
		return false
	}

	return true
}
