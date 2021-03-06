package Operator

import (
	"time"

	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestOperationNotEqual(ctx *testframework.TestFrameworkContext) bool {
	code := "52C56B6C766B00527AC46C766B51527AC46C766B00C36C766B51C39C009C616C7566"
	codeAddress, _ := utils.GetContractAddress(code)
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestOperationNotEqual GetDefaultAccount error:%s", err)
		return false
	}
	_, err = ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,

		false,
		code,
		"TestOperationNotEqual",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestOperationNotEqual DeploySmartContract error:%s", err)
		return false
	}
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationNotEqual WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationNotEqual(ctx, codeAddress, -1, 1) {
		return false
	}

	if !testOperationNotEqual(ctx, codeAddress, -1, -1) {
		return false
	}

	if !testOperationNotEqual(ctx, codeAddress, 1, 1) {
		return false
	}

	if !testOperationNotEqual(ctx, codeAddress, 0, 0) {
		return false
	}

	return true
}

func testOperationNotEqual(ctx *testframework.TestFrameworkContext, code common.Address, a, b int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		code,
		[]interface{}{a, b},
	)
	if err != nil {
		ctx.LogError("TestOperationNotEqual InvokeSmartContract error:%s", err)
		return false
	}
	resValue, err := res.Result.ToBool()
	if err != nil {
		ctx.LogError("TestOperationNotEqual Result.ToBool error:%s", err)
		return false
	}
	err = ctx.AssertToBoolean(resValue, a != b)
	if err != nil {
		ctx.LogError("TestOperationNotEqual test failed %s", err)
		return false
	}
	return true
}

/*
using Neo.SmartContract.Framework;
using Neo.SmartContract.Framework.Services.Neo;

class A : SmartContract
{
    public static bool Main(int a, int b)
    {
        return a != b;
    }
}
*/
