package cmd

import (
	"fmt"
	"stress-test/internal/model"
	"stress-test/internal/service"

	"github.com/spf13/cobra"
)

var url string
var numberOfRequests int
var concurrencyLevel int
var output model.StressTestOutput

var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Comando para executar o teste de stress",

	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("url: %s\n", url)
		fmt.Printf("requests: %d\n", numberOfRequests)
		fmt.Printf("concurrency: %d\n", concurrencyLevel)
	},

	Run: func(cmd *cobra.Command, args []string) {

		input := model.StressTestInput{
			Url:              url,
			NumberOfRequests: numberOfRequests,
			ConcurrencyLevel: concurrencyLevel,
		}

		output = service.Run(input)
	},

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Tempo total gasto na execução: %s\n", output.ExecutionTime)
		fmt.Printf("Quantidade total de requests realizados: %d\n", output.RequestsCount)
		for status, quantidade := range output.ResponseStatusCounter {
			fmt.Printf("Quantidade de requests com status HTTP %d: %d\n", status, quantidade)
		}
	},
}

func init() {
	rootCmd.AddCommand(stressCmd)
	createUrlFlag()
	createRequestsFlag()
	createConcurrencyFlag()
}

func createUrlFlag() {
	stressCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	stressCmd.MarkFlagRequired("url")
}

func createRequestsFlag() {
	stressCmd.Flags().IntVarP(&numberOfRequests, "requests", "r", 0, "Número total de requests")
	stressCmd.MarkFlagRequired("requests")
}

func createConcurrencyFlag() {
	stressCmd.Flags().IntVarP(&concurrencyLevel, "concurrency", "c", 0, "Número de chamadas simultâneas")
	stressCmd.MarkFlagRequired("concurrency")
}
