package Datatype

import (
	"time"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestArray(ctx *testframework.TestFrameworkContext) bool {
	code := "57c56b6a00527ac46a51527ac46a00c30548656c6c6f9c640c006a51c3650b006c756661006c756655c56b6a00527ac46a00c3c06a51527ac46a51c36c7566"
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
		ctx.LogError("TestArray DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestArray WaitForGenerateBlock error:%s", err)
		return false
	}
	params := []interface{}{[]byte("Hello"), []byte("world")}
	if !testArray(ctx, codeAddress, params) {
		return false
	}

	params = []interface{}{[]byte("Hello"), []byte("world"), "123456", 8}
	if !testArray(ctx, codeAddress, params) {
		return false
	}
	return true
}

func testArray(ctx *testframework.TestFrameworkContext, code common.Address, params []interface{}) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{"Hello", params},
	)
	if err != nil {
		ctx.LogError("TestArray InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestArray Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, len(params))
	if err != nil {
		ctx.LogError("TestArray test failed %s", err)
		return false
	}
	return true
}
