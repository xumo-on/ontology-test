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

func TestGetIndex(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "5dc56b6a00527ac46a51527ac46a00c308676574496e6465789c6409006551026c7566616a00c307676574486173689c6409006595016c7566616a00c30b67657450726576486173689c64090065d1006c7566616a00c30c67657454696d657374616d709c640900650b006c756661006c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681a53797374656d2e4865616465722e47657454696d657374616d70616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681953797374656d2e4865616465722e4765745072657648617368616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681553797374656d2e4865616465722e47657448617368616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681653797374656d2e4865616465722e476574496e646578616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"getIndex", []interface{}{[]byte("getIndex")}})
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

func TestGetHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
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

	Hash := events.Notify[1].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	s1 := strings.Join(s, "")

	height := events.Notify[0].States.(string)
	count = strings.Count(height, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, height[i-2:i])
	}
	s2 := strings.Join(s, "")
	uheight, err := strconv.ParseInt(s2, 16, 64)
	if err != nil {
		ctx.LogError("TransferStringToInt64 error:%s", err)
		return false
	}
	bhash, err := ctx.Ont.GetBlockHash(uint32(uheight))
	if err != nil {
		ctx.LogError("GetBlockHash error:%s", err)
		return false
	}

	if bhash.ToHexString() != s1 {
		ctx.LogError("TestGetHash error:%s", err)
		return false
	}

	ctx.LogInfo("hash:", bhash.ToHexString())
	ctx.LogInfo("hash1:", s1)

	return true
}

func TestGetPrevHash(ctx *testframework.TestFrameworkContext) bool {

	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"getPrevHash", []interface{}{[]byte("getPrevHash")}})
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

	Hash := events.Notify[1].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	s1 := strings.Join(s, "")

	height := events.Notify[0].States.(string)
	count = strings.Count(height, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, height[i-2:i])
	}
	s2 := strings.Join(s, "")
	uheight, err := strconv.ParseInt(s2, 16, 64)
	if err != nil {
		ctx.LogError("TransferStringToInt64 error:%s", err)
		return false
	}
	bhash, err := ctx.Ont.GetBlockHash(uint32(uheight - 1))
	if err != nil {
		ctx.LogError("GetBlockHash error:%s", err)
		return false
	}

	if bhash.ToHexString() != s1 {
		ctx.LogError("TestGetHash error:%s", err)
		return false
	}

	ctx.LogInfo("hash:", bhash.ToHexString())
	ctx.LogInfo("hash1:", s1)

	return true
}

func TestGetTimestamp(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "5dc56b6a00527ac46a51527ac46a00c308676574496e6465789c6409006551026c7566616a00c307676574486173689c6409006595016c7566616a00c30b67657450726576486173689c64090065d1006c7566616a00c30c67657454696d657374616d709c640900650b006c756661006c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681a53797374656d2e4865616465722e47657454696d657374616d70616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681953797374656d2e4865616465722e4765745072657648617368616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681553797374656d2e4865616465722e47657448617368616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b681b53797374656d2e426c6f636b636861696e2e476574486569676874616a00527ac46a00c3681b53797374656d2e426c6f636b636861696e2e476574486561646572616a51527ac46a51c3681653797374656d2e4865616465722e476574496e646578616a52527ac46a00c3681553797374656d2e52756e74696d652e4e6f74696679616a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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
		[]interface{}{"getTimestamp", []interface{}{[]byte("getTimestamp")}})
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

	Hash := events.Notify[1].States.(string)
	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	s1 := strings.Join(s, "")
	utime, err := strconv.ParseInt(s1, 16, 64)

	height := events.Notify[0].States.(string)
	count = strings.Count(height, "") - 1
	s = []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, height[i-2:i])
	}
	s2 := strings.Join(s, "")
	uheight, err := strconv.ParseInt(s2, 16, 64)
	if err != nil {
		ctx.LogError("TransferStringToInt64 error:%s", err)
		return false
	}
	block, err := ctx.Ont.GetBlockByHeight(uint32(uheight))
	if err != nil {
		ctx.LogError("GetBlockHash error:%s", err)
		return false
	}

	timestamp := block.Header.Timestamp

	ctx.LogInfo("timestamp:", timestamp)
	ctx.LogInfo("timestamp1:", utime)

	if timestamp != uint32(utime) {
		ctx.LogError("TestGetTimestamp error:%s", err)
		return false
	}

	return true
}

