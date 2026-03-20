package main

import "Phantom-Operator/cmd"
import "Phantom-Operator/pkg/logger"

func main() {
	err := cmd.Phantom-OperatorCli.Execute()
	if err != nil {
		logger.Error("Failed to execute havoc")
		return
	}
}
