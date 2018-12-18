package Datatype

import (
	"time"

	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestReturnType(ctx *testframework.TestFrameworkContext) bool {
	code := "5ac56b6a00527ac46a51527ac46a52527ac400c176c96a53527ac46a53c36a00c3c86a53c36a51c3c86a53c36a52c3c86a53c36c7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestReturnType GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		false,
		code,
		"TestReturnType",
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
		ctx.LogError("TestReturnType WaitForGenerateBlock error:%s", err)
		return false
	}
	if !testReturnType(ctx, codeAddress, []int{100343, 2433554}, []byte("Hello world")) {
		return false
	}
	return true
}

func testReturnType(ctx *testframework.TestFrameworkContext, code common.Address, args []int, arg3 []byte) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{args[0], args[1], arg3},
	)
	if err != nil {
		ctx.LogError("TestReturnType InvokeSmartContract error:%s", err)
		return false
	}

	rt, err := res.Result.ToArray()
	if err != nil {
		ctx.LogError("TestReturnType Result.ToArray error:%s", err)
		return false
	}
	a1, err := rt[0].ToInteger()
	if err != nil {
		ctx.LogError("TestReturnType Result.ToByteArray error:%s", err)
		return false
	}
	err = ctx.AssertToInt(a1, args[0])
	if err != nil {
		ctx.LogError("TestReturnType AssertToInt error:%s", err)
		return false
	}
	a2, err := rt[1].ToInteger()
	if err != nil {
		ctx.LogError("TestReturnType Result.ToByteArray error:%s", err)
		return false
	}
	err = ctx.AssertToInt(a2, args[1])
	if err != nil {
		ctx.LogError("TestReturnType AssertToInt error:%s", err)
		return false
	}
	a3, err := rt[2].ToByteArray()
	if err != nil {
		ctx.LogError("TestReturnType ToByteArray error:%s", err)
		return false
	}
	err = ctx.AssertToByteArray(a3, arg3)
	if err != nil {
		ctx.LogError("AssertToByteArray error:%s", err)
		return false
	}

	return true
}
