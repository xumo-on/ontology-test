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
	"encoding/hex"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/xumo-on/ontology-test/testframework"
	"time"
)

var codeAddress common.Address
//Notice:Every time this program is executed, the migrateCode have to be replaced.Any contract undeployed is ok.
var migrateCode = "59c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c640e006a51c300c365af006c7566616a00c3096765745363726970749c640900650b006c756661006c756656c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c36c756657c56b6a00527ac46a00c301320131013101310131013156795179587275517275557952795772755272755479537956727553727568194f6e746f6c6f67792e436f6e74726163742e4d696772617465616a51527ac46a51c3641b00144d696772617465207375636365737366756c6c796c756661006c7566006c7566"

func TestGetScript(ctx *testframework.TestFrameworkContext) bool {

	//DeployContract
	code := "59c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c640e006a51c300c365af006c7566616a00c3096765745363726970749c640900650b006c756661006c756656c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c36c756657c56b6a00527ac46a00c301310131013101310131013156795179587275517275557952795772755272755479537956727553727568194f6e746f6c6f67792e436f6e74726163742e4d696772617465616a51527ac46a51c3641b00144d696772617465207375636365737366756c6c796c756661006c7566006c7566"
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
	txHash, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"getScript", []interface{}{[]byte("getScript")}})
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

	if hex.EncodeToString(result) != "59c56b6a00527ac46a51527ac46a00c30f4d696772617465436f6e74726163749c640e006a51c300c365af006c7566616a00c3096765745363726970749c640900650b006c756661006c756656c56b682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e6753637269707448617368616a00527ac46a00c3681d53797374656d2e426c6f636b636861696e2e476574436f6e7472616374616a51527ac46a51c3681b4f6e746f6c6f67792e436f6e74726163742e476574536372697074616a52527ac46a52c36c756657c56b6a00527ac46a00c301310131013101310131013156795179587275517275557952795772755272755479537956727553727568194f6e746f6c6f67792e436f6e74726163742e4d696772617465616a51527ac46a51c3641b00144d696772617465207375636365737366756c6c796c756661006c7566006c7566" {
		ctx.LogError("TestGetScript error")
		return false
	}

	return true
}

func TestMigrate(ctx *testframework.TestFrameworkContext) bool {

	//PreExecInvokeContract
	txHash, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(codeAddress, []interface{}{"MigrateContract", []interface{}{migrateCode}})
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

	if hex.EncodeToString(result) != "4d696772617465207375636365737366756c6c79" {
		ctx.LogError("TestGetContext error")
		return false
	}

	return true
}
