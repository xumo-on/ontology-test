package Blockchain

import (
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"strconv"
	"strings"
	"time"
)

var codeAddress common.Address
var code string

func TestGetHeight(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code = "0113c56b6a00527ac46a51527ac46a00c3096765744865696768749c64090065b0046c7566616a00c3096765744865616465729c640900651d046c7566616a00c308676574426c6f636b9c6409006581036c7566616a00c30e6765745472616e73616374696f6e9c6409006503026c7566616a00c30b676574436f6e74726163749c640900655d016c7566616a00c3146765745472616e73616374696f6e4865696768749c6409006562006c7566616a00c31367657443757272656e74426c6f636b486173689c640900650b006c756661006c756654c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c36c756658c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a52527ac46a52c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a53527ac46a53c3682653797374656d2e426c6f636b636861696e2e4765745472616e73616374696f6e486569676874616a54527ac46a54c36c756656c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c36c75665cc56b681953797374656d2e53746f726167652e476574436f6e74657874616a00527ac468244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a51527ac46a51c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a52527ac46a52c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a53527ac46a53c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a54527ac46a54c3682053797374656d2e426c6f636b636861696e2e4765745472616e73616374696f6e616a55527ac46a55c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a56527ac46a00c307676574486173686a54c35272681253797374656d2e53746f726167652e507574616a00c30867657448617368316a56c35272681253797374656d2e53746f726167652e50757461516c756656c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3682053797374656d2e426c6f636b2e4765745472616e73616374696f6e436f756e74616a52527ac46a52c36c756656c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681553797374656d2e4865616465722e47657448617368616a52527ac46a52c36c756654c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c36c7566"
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
	value, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getHeight", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err := value.Result.ToByteArray()

	//SdkGetBlockHeight
	SdkGetBlockHeight, err := ctx.Ont.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.Ont.GetCurrentBlockHeight error:%s", err)
		return false
	}
	ctx.LogInfo("SdkGetBlockHeight:", SdkGetBlockHeight)

	//TransferValueToUint
	Svalue := hex.EncodeToString(bValue)

	count := strings.Count(Svalue, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Svalue[i-2:i])
	}
	s1 := strings.Join(s, "")

	height, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("ContractGetBlockHeight:", height)

	if  uint32(height) != SdkGetBlockHeight {
		ctx.LogError("TestGetHeight error")
		return false
	}

	return true
}

func TestGetHeader(ctx *testframework.TestFrameworkContext) bool {

	//PreExecInvokeContract
	value, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getHeader", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err := value.Result.ToByteArray()

	shash := hex.EncodeToString(bValue)
	count := strings.Count(shash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, shash[i-2:i])
	}
	s1 := strings.Join(s, "")
	uhash, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("hash:", uhash)

	//PreExecInvokeContract
	value, err = ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getCurrentBlockHash", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err = value.Result.ToByteArray()

	sheader := hex.EncodeToString(bValue)
	count1 := strings.Count(sheader, "") - 1
	s2 := []string{}
	for i := count1; i > 0; i -= 2 {
		s2 = append(s2, sheader[i-2:i])
	}
	s3 := strings.Join(s2, "")
	uheader, err := strconv.ParseUint(s3, 16, 32)
	ctx.LogInfo("header:", uheader)

	if uheader != uhash {
		ctx.LogError("TestHeader error")
		return false
	}
	return true
}

func TestGetBlock(ctx *testframework.TestFrameworkContext) bool {
	//PreExecInvokeContract
	value, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getBlock", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err := value.Result.ToByteArray()

	str := hex.EncodeToString(bValue)
	count := strings.Count(str, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, str[i-2:i])
	}
	s1 := strings.Join(s, "")
	Count, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("Count:", Count)

	if Count != 1 {
		ctx.LogError("TestGetBlock error", err)
		return false
	}
	return true
}

func TestGetTransaction(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	_, err = ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getTransaction", []interface{}{[]byte("getTransaction")}})
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
	hash, err := ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("getHash"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	txHash := hex.EncodeToString(hash)
	ctx.LogInfo("	TxHash:", txHash)

	hash, err = ctx.Ont.GetStorage(codeAddress.ToHexString(), []byte("getHash1"))
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetStorageItem key:hello error: %s", err)
		return false
	}
	txHash1 := hex.EncodeToString(hash)
	ctx.LogInfo("	TxHash1:", txHash1)

	if txHash != txHash1 {
		ctx.LogError("TestGetTransaction error", err)
		return false
	}
	return true
}

func TestGetContract(ctx *testframework.TestFrameworkContext) bool {

	//PreExecInvokeContract
	value, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getContract", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err := value.Result.ToByteArray()
	script := hex.EncodeToString(bValue)
	ctx.LogInfo("	script:", script)

	if script != code {
		ctx.LogError("TestGetTransaction error")
		return false
	}
	return true
}

func TestGetTransactionHeight(ctx *testframework.TestFrameworkContext) bool {

	//PreExecInvokeContract
	value, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getTransactionHeight", []interface{}{}})
	if err != nil {
		ctx.LogError("TestDomainSmartContract PreExecInvokeSmartContract error: %s", err)
		return false
	}
	bValue, err := value.Result.ToByteArray()

	//SdkGetTransactionHeight
	SdkGetTransactionHeight, err := ctx.Ont.GetCurrentBlockHeight()
	if err != nil {
		ctx.LogError("ctx.Ont.GetCurrentTransactionHeighHeight error:%s", err)
		return false
	}
	ctx.LogInfo("SdkGetTransactionHeight:", SdkGetTransactionHeight)

	value1 := hex.EncodeToString(bValue)
	count := strings.Count(value1, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, value1[i-2:i])
	}
	s1 := strings.Join(s, "")
	height, err := strconv.ParseUint(s1, 16, 32)
	ctx.LogInfo("ContractGetTransactionHeight:", height)

	if uint32(height) - SdkGetTransactionHeight > 2 {
		ctx.LogError("TestGetTransactionHeight error: %s")
		return false
	}

	return true
}
