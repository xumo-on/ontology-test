package Datatype

import (
	"time"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestByteArray(ctx *testframework.TestFrameworkContext) bool {
	code := "55c56b6a00527ac46a51527ac46a00c36a51c39c6c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestByteArray GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		false,
		code,
		"TestByteArray",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestArray DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestArray WaitForGenerateBlock error:%s", err)
		return false
	}

	arg1 := []byte("Hello")
	arg2 := []byte("World")

	if !testByteArray(ctx, codeAddress, arg1, arg1, true) {
		return false
	}
	if !testByteArray(ctx, codeAddress, arg1, arg2, false) {
		return false
	}
	return true
}

func testByteArray(ctx *testframework.TestFrameworkContext, code common.Address, arg1, arg2 []byte, expect bool) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{arg1, arg2},
	)
	if err != nil {
		ctx.LogError("testByteArray InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToBool()
	if err != nil {
		ctx.LogError("testByteArray Result.ToBool error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(resValue, expect)
	if err != nil {
		ctx.LogError("testByteArray test failed %s", err)
		return false
	}
	return true
}