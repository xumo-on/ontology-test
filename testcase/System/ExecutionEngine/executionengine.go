package ExecutionEngine

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"strings"
	"time"
)

var	codeAddress common.Address

func TestGetScriptContainer(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "5bc56b6a00527ac46a51527ac46a00c312676574536372697074436f6e7461696e65729c640900658d016c7566616a00c316676574457865637574696e67536372697074486173689c64090065cf006c7566616a00c31467657443616c6c696e67536372697074486173689c640900650b006c756661006c756657c56b682b53797374656d2e457865637574696f6e456e67696e652e47657443616c6c696e6753637269707448617368616a00527ac4682953797374656d2e457865637574696f6e456e67696e652e476574456e74727953637269707448617368616a51527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a51c3681553797374656d2e52756e74696d652e4e6f7469667961516c756657c56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac4682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a51527ac46a00c3036765746a51c35272681253797374656d2e53746f726167652e507574616a51c3681553797374656d2e52756e74696d652e4e6f7469667961516c75665cc56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac4682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e6572616a51527ac46a51c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a52527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a53527ac46a53c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a54527ac46a54c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a55527ac46a55c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a56527ac46a00c3036765746a52c35272681253797374656d2e53746f726167652e507574616a00c304676574316a56c35272681253797374656d2e53746f726167652e50757461516c7566"
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
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getScriptContainer", []interface{}{[]byte("getScriptContainer")}})
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

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value := hex.EncodeToString(svalue)

	svalue1, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get1"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value1 := hex.EncodeToString(svalue1)

	ctx.LogInfo("hash:",value)
	ctx.LogInfo("hash1:",value1)

	if value != value1 {
		ctx.LogError("TestGetScriptContainer error")
		return false
	}
	return true
}

func TestGetExecutingScriptHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getExecutingScriptHash", []interface{}{[]byte("getExecutingScriptHash")}})
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

	//GetSvalueInStorage
	svalue, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("get"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	value := hex.EncodeToString(svalue)

	count := strings.Count(value, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, value[i-2:i])
	}
	s1 := strings.Join(s, "")
	ctx.LogInfo("hash:",s1)

	if s1 != codeAddress.ToHexString() {
		ctx.LogError("TestGetExecutingScriptHash error")
		return false
	}

	return true
}

func TestGetCallingScriptHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getCallingScriptHash", []interface{}{[]byte("getCallingScriptHash")}})
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
	notify := events.Notify[0]
	notify1 := events.Notify[1]
	ctx.LogInfo("notify: ", notify.States)
	ctx.LogInfo("notify1: ", notify1.States)

	if notify.States != notify1.States {
		ctx.LogError("getCallingScriptHash error")
		return false
	}

	return true
}
