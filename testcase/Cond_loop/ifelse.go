package Cond_loop

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/xumo-on/ontology-test/testframework"
	"time"
)

func TestIfElse(ctx *testframework.TestFrameworkContext) bool {
	code := "59c56b6a00527ac46a51527ac46a00c3066966656c73659c6424006a51c300c36a52527ac46a51c351c36a53527ac46a52c36a53c37c650b006c756661006c756659c56b6a00527ac46a51527ac46a00c36a51c3a0640700516c7566616a00c36a51c39f6407004f6c756661006c7566006c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestIfElse GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,

		false,
		code,
		"TestIfElse",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestIfElse DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestIfElse WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testIfElse(ctx, codeAddress, 23, 2) {
		return false
	}

	if !testIfElse(ctx, codeAddress, 2, 23) {
		return false
	}

	if !testIfElse(ctx, codeAddress, 0, 0) {
		return false
	}

	return true
}

func testIfElse(ctx *testframework.TestFrameworkContext, code common.Address, a, b int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{"ifelse",[]interface{}{a,b}},
	)
	if err != nil {
		ctx.LogError("TestIfElse InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestIfElse Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, condIfElse(a, b))
	if err != nil {
		ctx.LogError("TestIfElse test %d ifelse %d failed %s", a, b, err)
		return false
	}
	return true
}

func condIfElse(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}