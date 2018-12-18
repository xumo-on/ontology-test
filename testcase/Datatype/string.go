package Datatype

import (
	"time"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
)

func TestString(ctx *testframework.TestFrameworkContext) bool {
	code := "53c56b0b48656c6c6f20576f726c646c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestString GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,

		false,
		code,
		"TestString",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestString DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestString WaitForGenerateBlock error:%s", err)
		return false
	}
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		codeAddress,
		[]interface{}{},
	)
	if err != nil {
		ctx.LogError("TestString InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToString()
	if err != nil {
		ctx.LogError("TestString Result.ToString error:%s", err)
		return false
	}
	err = ctx.AssertToString(resValue, "Hello World")
	if err != nil {
		ctx.LogError("TestString test failed %s", err)
		return false
	}
	return true
}
