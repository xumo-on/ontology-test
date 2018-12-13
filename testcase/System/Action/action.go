/*
from boa.interop.System.Action import RegisterAction
from boa.interop.System.Runtime import Notify

Transfer = RegisterAction('test','a')


def Main(operation, args):
    if operation == "test":
        return test()
    return False

def test():
    a = 1
    Transfer(a)
    return True
 */

package Action

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"time"
)

func TestRegisterAction(ctx *testframework.TestFrameworkContext) bool {
	//DeployContract
	code := "57c56b6a00527ac46a51527ac46a00c304746573749c640900650b006c756661006c756655c56b516a00527ac46a00c3047465737452c1681553797374656d2e52756e74696d652e4e6f74696679516c7566"
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
		[]interface{}{"test", []interface{}{[]byte("test")}})
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

	test := events.Notify[0].States

	ctx.LogInfo("test:", test)

	return true
}
