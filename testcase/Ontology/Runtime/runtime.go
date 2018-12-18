/*
from boa.interop.Ontology.Runtime import Base58ToAddress,AddressToBase58,GetRandomHash
from boa.interop.System.Runtime import Notify
from boa.interop.System.Storage	import Put,GetContext

context = GetContext()

def Main(operation, args):
    if operation == 'BTA':
        return BTA()
    if operation == 'ATB':
        return ATB()
    if operation == 'getHash':
        return getHash()
    return False

def BTA():
    bta = Base58ToAddress('ASwaf8mj2E3X18MHvcJtXoDsMqUjJswRWS')
    Put(context, 'get', bta)
    return bta

def ATB():
    bta = BTA()
    atb = AddressToBase58(bta)
    Put(context, 'get', atb)
    return True

def getHash():
    blockhash = GetCurrentBlockHash()
    Put(context, 'get', blockhash)
    Notify(blockhash)
    return True
 */
package Runtime

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"strings"
	"time"
)

var codeAddress common.Address

func TestBase58ToAddress(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "5bc56b6a00527ac46a51527ac46a00c3034254419c64090065e3006c7566616a00c3034154429c6409006595006c7566616a00c307676574486173689c640900650b006c756661006c756656c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a51527ac46a00c3036765746a51c35272681253797374656d2e53746f726167652e507574616c756655c56b6539006a00527ac46a00c368204f6e746f6c6f67792e52756e74696d652e41646472657373546f426173653538616a51527ac46a51c36c756654c56b224153776166386d6a3245335831384d4876634a74586f44734d71556a4a737752575368204f6e746f6c6f67792e52756e74696d652e426173653538546f41646472657373616a00527ac46a00c36c7566"
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

	//PreExecInvokeContract
	txHash ,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"BTA", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//GetResult
	result, err := txHash.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetResult error: %s", err)
		return false
	}

	if hex.EncodeToString(result) != "7a7f5c8c364ef70e52904daf1a99a49450ccef0f" {
		ctx.LogError("TestBTA error")
		return false
	}

	return true
}

func TestAddressToBase58(ctx *testframework.TestFrameworkContext) bool {
	//PreExecInvokeContract
	txHash ,err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"ATB", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract InvokeNeoVMSmartContract error: %s", err)
		return false
	}

	//GetResult
	result, err := txHash.Result.ToByteArray()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetResult error: %s", err)
		return false
	}

	if hex.EncodeToString(result) != "4153776166386d6a3245335831384d4876634a74586f44734d71556a4a7377525753" {
		ctx.LogError("TestATB error")
		return false
	}

	return true
}

func TestGetRandomHash(ctx *testframework.TestFrameworkContext) bool {
	//SmartontractGetHash
	ctx.LogInfo("=====CodeAddress===%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getHash", []interface{}{[]byte("getHash")}})
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
	ctx.LogInfo("SmartContractGetHash:", value)

	//SdkGetBlockHash
	SdkGetBlockHash, err := ctx.Ont.GetCurrentBlockHash()
	if err != nil {
		ctx.LogError("ctx.Ont.GetCurrentBlockHash error:%s", err)
		return false
	}

	SdkHash := SdkGetBlockHash.ToHexString()
	count := strings.Count(SdkHash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, SdkHash[i-2:i])
	}
	s1 := strings.Join(s, "")
	ctx.LogInfo("SdkGetBlockHash:", s1)

	if s1 != value {
		ctx.LogError("Test error.")
		return false
	}

	return true
}
