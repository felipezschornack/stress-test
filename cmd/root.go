package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type RunEFunc func(cmd *cobra.Command, args []string)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "Sistema para realizar testes de carga em um serviço web",
	Long: `Sistema CLI para realizar testes de carga em um serviço web. 
Forneça a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas, e
o sistema gerará um relatório com informações específicas após a execução dos testes.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
