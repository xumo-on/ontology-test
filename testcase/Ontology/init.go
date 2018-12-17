package Ontology

import (
	"github.com/ontio/ontology-test/testcase/Ontology/Contract"
	"github.com/ontio/ontology-test/testcase/Ontology/Header"
	"github.com/ontio/ontology-test/testcase/Ontology/Native"
	"github.com/ontio/ontology-test/testcase/Ontology/Runtime"
	"github.com/ontio/ontology-test/testcase/Ontology/Transaction"
)

func TestOntology() {
	//Attribute.TestAttribute()
	Contract.TestContract()
	Header.TestHeader()
	Native.TestNative()
	Runtime.TestRuntime()
	Transaction.TestTransaction()
}
