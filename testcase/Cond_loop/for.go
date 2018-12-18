package Cond_loop

import (
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/xumo-on/ontology-test/testframework"
	"time"
)

func TestFor(ctx *testframework.TestFrameworkContext) bool {
	code := "58c56b6a00527ac46a51527ac46a00c3055768696c659c6416006a51c300c36a52527ac46a52c3650b006c756661006c756659c56b6a00527ac4006a51527ac4006a52527ac461616a52c36a00c39f641c006a51c36a52c3936a51527ac46a52c351936a52527ac462dfff6161616a51c36c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestFor GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,

		false,
		code,
		"TestFor",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestFor DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestFor WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testFor(ctx, codeAddress, 23) {
		return false
	}

	if !testFor(ctx, codeAddress, -23) {
		return false
	}

	if !testFor(ctx, codeAddress, 0) {
		return false
	}

	return true
}

func testFor(ctx *testframework.TestFrameworkContext, code common.Address, a int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{"While",[]interface{}{a}},
	)
	if err != nil {
		ctx.LogError("TestFor InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestFor Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, forloop(a))
	if err != nil {
		ctx.LogError("TestFor test for %d failed %s", a, err)
		return false
	}
	return true
}

func forloop(a int) int {
	b := 0
	for i := 0; i < a; i++ {
		b = b + i
	}
	return b
}
