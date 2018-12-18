package Cond_loop

import (
	"time"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestSwitch(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b6c766b00527ac4616c766b00c36c766b51527ac46c766b51c3519c630600620e00516c766b52527ac4620e00006c766b52527ac46203006c766b52c3616c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestArray GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,

		false,
		code,
		"TestArray",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestSwitch DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestSwitch WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testSwitch(ctx, codeAddress, 23) {
		return false
	}

	if !testSwitch(ctx, codeAddress, 1) {
		return false
	}

	if !testSwitch(ctx, codeAddress, 0) {
		return false
	}

	return true
}

func testSwitch(ctx *testframework.TestFrameworkContext, code common.Address, a int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{a},
	)
	if err != nil {
		ctx.LogError("TestSwitch InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestSwitch Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, tswitch(a))
	if err != nil {
		ctx.LogError("TestSwitch test switch %d failed %s", a, err)
		return false
	}
	return true
}

func tswitch(a int) int {
	switch a {
	case 1:
		return 1
	default:
		return 0
	}
}