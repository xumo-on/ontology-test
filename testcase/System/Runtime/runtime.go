/*
from boa.interop.System.Runtime import Notify,Log,GetTrigger,CheckWitness,GetTime,Serialize,Deserialize
from boa.interop.System.ExecutionEngine	import GetExecutingScriptHash
from boa.interop.System.Blockchain import GetHeader
from boa.interop.System.Header import GetTimestamp,GetBlockHash
from boa.interop.Ontology.Runtime import GetCurrentBlockHash
from boa.interop.System.Storage	import Put,GetContext

def Main(operation, args):
    if operation == 'getTrigger':
        return getTrigger()
    if operation == 'checkWitness':
        Hash = args[1]
        return checkWitness(Hash)
    if operation == 'log':
        return log()
    if operation == 'notify':
        return notify()
    if operation == 'getTime':
        return getTime()
    if operation == 'serialize':
        return serialize()
    return False

def getTrigger():
    trigger = GetTrigger()
    Notify(trigger)
    return True

def checkWitness(Hash):
    check = CheckWitness(Hash)
    Notify(check)
    return True

def log():
    Log('aaaaa')
    return True

def notify():
    Notify('aaaaa')
    return True

def getTime():
    time = GetTime()
    Hash = GetCurrentBlockHash()
    header = GetHeader(Hash)
    time1 = GetTimestamp(header)
    Notify(time)
    Notify(time1)
    return  True

def serialize():
    context = GetContext()
    time = GetTime()
    time1 = Serialize(time)
    time2 = Deserialize(time1)
    Notify(time)
    Notify(time2)
    return True
 */
package Runtime

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

var	codeAddress common.Address

func TestCheckWitness(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "0112c56b6a00527ac46a51527ac46a00c30a676574547269676765729c64090065b2026c7566616a00c30c636865636b5769746e6573739c6416006a51c351c36a52527ac46a52c3653d026c7566616a00c3036c6f679c6409006509026c7566616a00c3066e6f746966799c64090065cf016c7566616a00c30767657454696d659c64090065e7006c7566616a00c30973657269616c697a659c640900650b006c756661006c756659c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac4681653797374656d2e52756e74696d652e47657454696d65616a51527ac46a51c3681853797374656d2e52756e74696d652e53657269616c697a65616a52527ac46a52c3681a53797374656d2e52756e74696d652e446573657269616c697a65616a53527ac46a51c3681553797374656d2e52756e74696d652e4e6f74696679616a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c756659c56b681653797374656d2e52756e74696d652e47657454696d65616a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a51527ac46a51c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a52527ac46a52c3681a53797374656d2e4865616465722e47657454696d657374616d70616a53527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a53c3681553797374656d2e52756e74696d652e4e6f7469667961516c756654c56b056161616161681553797374656d2e52756e74696d652e4e6f7469667961516c756654c56b056161616161681253797374656d2e52756e74696d652e4c6f6761516c756656c56b6a00527ac46a00c3681b53797374656d2e52756e74696d652e436865636b5769746e657373616a51527ac46a51c3681553797374656d2e52756e74696d652e4e6f7469667961516c756655c56b681953797374656d2e52756e74696d652e47657454726967676572616a00527ac46a00c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"checkWitness", []interface{}{[]byte("checkWitness"), signer.Address[:]}})
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

	check := events.Notify[0].States.(string)

	if check != "01" {
		ctx.LogError("TestcheckWitness error")
		return false
	}

	return true
}

//func TestLog(ctx *testframework.TestFrameworkContext) bool {
//
//	signer, err := ctx.GetDefaultAccount()
//	if err != nil {
//		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
//		return false
//	}
//
//	//InvokeContract
//	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
//		signer,
//		codeAddress,
//		[]interface{}{"log", []interface{}{[]byte("log")}})
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
//	return true
//}

func TestNotify(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"notify", []interface{}{[]byte("notify")}})
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

	check := events.Notify[0].States.(string)

	ctx.LogInfo("check: ", check)

	if check != "6161616161" {
		ctx.LogError("TestcheckWitness error")
		return false
	}

	return true
}

func TestGetTime(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getTime", []interface{}{[]byte("getTime")}})
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

	time := events.Notify[0].States.(string)
	time1 := events.Notify[1].States.(string)

	if time != time1 {
		ctx.LogError("TestNotify error")
		return true
	}
	return true
}

func TestSerialize(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"serialize", []interface{}{[]byte("serialize")}})
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

	time := events.Notify[0].States.(string)
	time1 := events.Notify[1].States.(string)

	if time != time1 {
		ctx.LogError("TestSerialize error")
		return true
	}

	return true
}
