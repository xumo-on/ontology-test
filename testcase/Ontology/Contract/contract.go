/*
from boa.interop.Ontology.Contract import Migrate
from boa.interop.System.Runtime import Notify
from boa.interop.System.ExecutionEngine	import GetExecutingScriptHash
from boa.interop.System.Blockchain import GetContract


def Main(operation, args):
    if operation == "MigrateContract":
        return MigrateContract(args[1])
    if operation == 'getScript':
        return getScript()
    return False

def MigrateContract(code):
    res = Migrate(code, "1", "1", "1", "1", "1", "1")
    res1 = res
    if res:
        Notify("Migrate successfully")
        return True
    else:
        return False

def getScript():
    script = GetExecutingScriptHash()
    contract = GetContract(script)
    sc = GetScript(contract)
    Notify(sc)
    return True
 */
package Contract

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
	"time"
)

var codeAddress common.Address
//Notice:Every time this program is executed, the migrateCode have to be replaced.Any contract undeployed is ok.
var migrateCode = "57c56b6a00527ac46a51527ac46a00c3036c6f679c640900650b006c756661006c756654c56b0761616141414141681253797374656d2e52756e74696d652e4c6f6761516c7566"

func TestGetScript(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "59c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c640e006a51c351c365c8006c7566616a00c3096765745363726970749c640900650b006c756661006c756657c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b6a00527ac46a00c301310131013101310131013156795179587275517275557952795772755272755479537956727553727568194f6e746f6c6f67792e436f6e74726163742e4d696772617465616a51527ac46a51c3643400144d696772617465207375636365737366756c6c79681553797374656d2e52756e74696d652e4e6f7469667961516c756661006c7566006c7566"
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
		[]interface{}{"getScript", []interface{}{[]byte("getScript")}})
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

	notify := events.Notify[0].States.(string)

	if notify != "59c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c640e006a51c351c365c8006c7566616a00c3096765745363726970749c640900650b006c756661006c756657c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c3681553797374656d2e52756e74696d652e4e6f7469667961516c756658c56b6a00527ac46a00c301310131013101310131013156795179587275517275557952795772755272755479537956727553727568194f6e746f6c6f67792e436f6e74726163742e4d696772617465616a51527ac46a51c3643400144d696772617465207375636365737366756c6c79681553797374656d2e52756e74696d652e4e6f7469667961516c756661006c7566006c7566" {
		ctx.LogError("TestGetScript error")
		return false
	}

	return true
}

func TestMigrate(ctx *testframework.TestFrameworkContext) bool {
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestDomainSmartContract GetDefaultAccount error:%s", err)
		return false
	}

	//InvokeContract
	txHash, err := ctx.Ont.NeoVM.InvokeNeoVMContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		codeAddress,
		[]interface{}{"MigrateContract", []interface{}{[]byte("MigrateContract"),migrateCode}})
	if err != nil {
		ctx.LogError("Please change the migrateCode.")
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

	notify := events.Notify[0].States.(string)

	if notify != "4d696772617465207375636365737366756c6c79" {
		ctx.LogError("TestGetContext error")
		return false
	}

	return true
}
