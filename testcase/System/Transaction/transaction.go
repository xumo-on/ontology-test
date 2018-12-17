/*
from boa.interop.System.Runtime import Notify
from boa.interop.System.Transaction	import GetTransactionHash
from boa.interop.System.Blockchain import GetHeight,GetBlock,GetTransactionHeight
from boa.interop.System.Block import GetTransactionByIndex
from boa.interop.Ontology.Runtime import GetCurrentBlockHash

def Main(operation, args):
    if operation == 'getHash':
        return getHash()
    return False

def getHash():
    Hash = GetCurrentBlockHash()
    Block = GetBlock(Hash)
    tx = GetTransactionByIndex(Block, 0)
    Hash = GetTransactionHash(tx)
    Notify(Hash)
    return True
 */
package Transaction

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology-test/testframework"
	"strings"
	"time"
)

func TestGetHash(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "57c56b6a00527ac46a51527ac46a00c307676574486173689c640900650b006c756661006c756658c56b68244f6e746f6c6f67792e52756e74696d652e47657443757272656e74426c6f636b48617368616a00527ac46a00c3681a53797374656d2e426c6f636b636861696e2e476574426c6f636b616a51527ac46a51c3007c681b53797374656d2e426c6f636b2e4765745472616e73616374696f6e616a52527ac46a52c3681a53797374656d2e5472616e73616374696f6e2e47657448617368616a00527ac46a00c3681553797374656d2e52756e74696d652e4e6f7469667961516c7566"
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

	Hash := events.Notify[0].States.(string)

	count := strings.Count(Hash, "") - 1
	s := []string{}
	for i := count; i > 0; i -= 2 {
		s = append(s, Hash[i-2:i])
	}
	s1 := strings.Join(s, "")

	if txHash.ToHexString() != s1 {
		ctx.LogError("TestGetHash error")
		return false
	}

	return true
}
