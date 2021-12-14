package main

import (
	"github.com/testground/sdk-go/run"
)

var testcases = map[string]interface{}{
	"evaluate": run.InitializedTestCaseFn(RunSimulation),
}

func main() {
	run.InvokeMap(testcases)
}
