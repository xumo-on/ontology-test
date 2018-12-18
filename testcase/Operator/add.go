package Operator

import (
	"time"

	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/xumo-on/ontology-test/testframework"
	"github.com/ontio/ontology/common"
)

func TestOperationAdd(ctx *testframework.TestFrameworkContext) bool {
	code := "59c56b6a00527ac46a51527ac46a00c3036164649c6424006a51c300c36a52527ac46a51c351c36a53527ac46a52c36a53c37c650b006c756661006c756655c56b6a00527ac46a51527ac46a00c36a51c3936c7566"
	codeAddress, err := utils.GetContractAddress(code)
	if err != nil {
		ctx.LogError("TestOperationAdd GetContractAddress error:%s", err)
		return false
	}
	ctx.LogInfo("TestOperationAdd contact address:%s", codeAddress.ToHexString())
	signer, err := ctx.GetDefaultAccount()
	if err != nil {
		ctx.LogError("TestOperationAdd GetDefaultAccount error:%s", err)
		return false
	}
	tx, err := ctx.Ont.NeoVM.DeployNeoVMSmartContract(ctx.GetGasPrice(), ctx.GetGasLimit(),
		signer,
		false,
		code,
		"TestOperationAdd",
		"1.0",
		"",
		"",
		"",
	)
	if err != nil {
		ctx.LogError("TestOperationAdd DeploySmartContract error:%s", err)
		return false
	}
	ctx.LogInfo("DeployContract TxHash:%s", tx.ToHexString())
	//等待出块
	_, err = ctx.Ont.WaitForGenerateBlock(30*time.Second, 1)
	if err != nil {
		ctx.LogError("TestOperationAdd WaitForGenerateBlock error:%s", err)
		return false
	}

	if !testOperationAdd(ctx, codeAddress, 1, 2) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, -1, 1) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, -1, -2) {
		return false
	}

	if !testOperationAdd(ctx, codeAddress, 0, 0) {
		return false
	}

	return true
}

func testOperationAdd(ctx *testframework.TestFrameworkContext, codeAddress common.Address, a, b int) bool {
	res, err := ctx.Ont.NeoVM.PreExecInvokeNeoVMContract(
		codeAddress,
		[]interface{}{"add", []interface{}{a, b}},
	)
	if err != nil {
		ctx.LogError("TestOperationAdd InvokeSmartContract error:%s", err)
		return false
	}
	resValue,err := res.Result.ToInteger()
	if err != nil {
		ctx.LogError("TestOperationAdd Result.ToInteger error:%s", err)
		return false
	}
	err = ctx.AssertToInt(resValue, a+b)
	if err != nil {
		ctx.LogError("TestOperationAdd test failed %s , %d, %d", err, a, b)
		return false
	}
	return true
}
