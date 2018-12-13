/*
contract code:
from boa.interop.System.Runtime import Notify
from boa.interop.System.Blockchain import GetContract
from boa.interop.System.ExecutionEngine	import GetExecutingScriptHash
from boa.interop.System.Contract import GetStorageContext,Destroy
from boa.interop.System.Storage	import GetContext
from boa.interop.System.Storage import Put
from boa.interop.Ontology.Contract import GetScript

context = GetContext()

def Main(operation, args):
    if operation == 'getStorageContext':
        return getStorageContext()
    if operation == 'destroy':
        return destroy()
    return False

def getStorageContext():
    Hash = GetExecutingScriptHash()
    contract = GetContract(Hash)
    context1 = GetStorageContext(contract)
    #Put(context, "get", context)  # Not support interface to bytearray
    #if context1 == context: # Not support interface to biginteger
        #return True
    Notify(context1)
    Notify(context)
    return True

def destroy():
    Destroy()
    return True
 */

package Contract

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

var	codeAddress common.Address

func TestGetStorageContext(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "59c56b6a00527ac46a51527ac46a00c31167657453746f72616765436f6e746578749c6409006542006c7566616a00c30764657374726f799c640900650b006c756661006c756654c56b681753797374656d2e436f6e74726163742e44657374726f7961516c756659c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a51527ac46a51c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a52527ac46a52c3682153797374656d2e436f6e74726163742e47657453746f72616765436f6e74657874616a53527ac46a53c3681553797374656d2e52756e74696d652e4e6f74696679616a00c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"getStorageContext", []interface{}{[]byte("getStorageContext")}})
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

	events, err := ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}
	notify := events.Notify[0]
	notify1 := events.Notify[1]
	ctx.LogInfo("notify: ", notify.States)
	ctx.LogInfo("notify1: ", notify1.States)

	if notify.States != notify1.States {
		ctx.LogError("TestGetStorageContext error")
		return false
	}
	return true
}

func TestContractDestroy(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestGetContract - GetDefaultAccount error:%s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"destroy", []interface{}{[]byte("destroy")}})
	if err != nil {
		ctx.LogError("TestContractDestroy InvokeSmartContract error: %s", err)
		return false
	}

	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestContractDestroy WaitForGenerateBlock error: %s", err)
		return false
	}

	_, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"destroy", []interface{}{[]byte("destroy")}})
	if err == nil {
		return false
	}

	return true
}
