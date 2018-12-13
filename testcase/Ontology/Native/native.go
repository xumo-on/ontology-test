/*
from boa.interop.Ontology.Native import Invoke
from boa.builtins import ToScriptHash, state
from boa.interop.System.Runtime import Notify
from boa.interop.System.ExecutionEngine import GetExecutingScriptHash


# ONG Big endian Script Hash: 0x0200000000000000000000000000000000000000
OngContract = ToScriptHash("AFmseVrdL9f9oyCzZefL9tG6UbvhfRZMHJ")
fromAccount = ToScriptHash("ASwaf8mj2E3X18MHvcJtXoDsMqUjJswRWS")


selfContractAddress = GetExecutingScriptHash()

def Main(operation, args):
    if operation == "transferOngToContract":
        return transferOngToContract()
    if operation == "checkSelfContractONGAmount":
        return checkSelfContractONGAmount()
    return False

def transferOngToContract():
    Notify(["111_transferOngToContract", selfContractAddress])
    param = state(fromAccount, selfContractAddress, 1)
    res = Invoke(0, OngContract, 'transfer', [param])
    if res and res == b'\x01':
        Notify('transfer Ong succeed')
        return True
    else:
        Notify('transfer Ong failed')
        return False

def checkSelfContractONGAmount():
    param = state(selfContractAddress)
    # do not use [param]
    res = Invoke(0, OngContract, 'balanceOf', param)
    res1 = res
    Notify(res)
    return res
 */
package Native

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"strconv"
	"strings"
	"time"
)

func TestInvoke(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "59c56b6a00527ac46a51527ac46a00c3157472616e736665724f6e67546f436f6e74726163749c6409006527016c7566616a00c31a636865636b53656c66436f6e74726163744f4e47416d6f756e749c640900650b006c756661006c756659c56b2241466d73655672644c3966396f79437a5a65664c397447365562766866525a4d484a751400000000000000000000000000000000000000026a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a51527ac46a51c351c66b6a00527ac46c6a52527ac4006a00c30962616c616e63654f666a52c3537951795572755172755279527954727552727568164f6e746f6c6f67792e4e61746976652e496e766f6b65616a53527ac46a53c36a54527ac46a53c3681553797374656d2e52756e74696d652e4e6f74696679616a53c36c75665dc56b2241466d73655672644c3966396f79437a5a65664c397447365562766866525a4d484a751400000000000000000000000000000000000000026a00527ac4224153776166386d6a3245335831384d4876634a74586f44734d71556a4a737752575375147a7f5c8c364ef70e52904daf1a99a49450ccef0f6a51527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a52527ac4193131315f7472616e736665724f6e67546f436f6e74726163746a52c352c176c9681553797374656d2e52756e74696d652e4e6f74696679616a51c36a52c35153c66b6a52527ac46a51527ac46a00527ac46c6a53527ac4006a00c3087472616e736665726a53c351c176c9537951795572755172755279527954727552727568164f6e746f6c6f67792e4e61746976652e496e766f6b65616a54527ac46a54c3643d006a54c301019c643400147472616e73666572204f6e672073756363656564681553797374656d2e52756e74696d652e4e6f7469667961516c756661137472616e73666572204f6e67206661696c6564681553797374656d2e52756e74696d652e4e6f7469667961006c7566006c75665ec56b6a00527ac46a51527ac46a51c36a00c3946a52527ac46a52c3c56a53527ac4006a54527ac46a00c36a55527ac461616a00c36a51c39f6433006a54c36a55c3936a56527ac46a56c36a53c36a54c37bc46a54c351936a54527ac46a55c36a54c3936a00527ac462c8ff6161616a53c36c7566"
	codeAddress, _ := utils.GetContractAddress(code)

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
		[]interface{}{"checkSelfContractONGAmount", []interface{}{[]byte("checkSelfContractONGAmount")}})
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

	hexbalance := events.Notify[0].States.(string)
	count := strings.Count(hexbalance, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, hexbalance[i-2:i])
	}
	s1 := strings.Join(s, "")
	balance, err := strconv.ParseUint(s1, 16, 32)

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"transferOngToContract", []interface{}{[]byte("transferOngToContract")}})
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

	//InvokeContract
	txHash, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"checkSelfContractONGAmount", []interface{}{[]byte("checkSelfContractONGAmount")}})
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
	events, err = ctx.Ont.GetSmartContractEvent(txHash.ToHexString())
	if err != nil {
		ctx.LogError("TestInvokeSmartContract GetSmartContractEvent error:%s", err)
		return false
	}
	if events.State == 0 {
		ctx.LogError("TestInvokeSmartContract failed invoked exec state return 0")
		return false
	}

	hexbalance = events.Notify[0].States.(string)
	count = strings.Count(hexbalance, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, hexbalance[i-2:i])
	}
	s1 = strings.Join(s, "")
	balance1, err := strconv.ParseUint(s1, 16, 32)

	if balance+1 != balance1 {
		ctx.LogError("TestInvoke error")
		return false
	}

	return true
}
