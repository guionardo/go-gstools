package main

import (
	"fmt"
	"os"
	"path"

	"github.com/guionardo/go-gstools/cep"
	"github.com/guionardo/go-gstools/cmd/shared"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "consulta-cep [CEP]",
		Short: "Consulta o CEP",
		Run: func(cmd *cobra.Command, args []string) {
			service, err := cep.NewSQliteCEPService(dbName)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			cepResp, err := service.GetCEP(codCep)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = shared.SaveOutput(cepResp, output_type, output)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("É necessário informar um CEP para consulta")
			}
			codCep = args[0]
			return nil
		},
	}

	dbName      string
	codCep      string
	output      string
	output_type string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&output, "output", "console", "Saída da consulta [console, caminho do arquivo]")
	rootCmd.PersistentFlags().StringVar(&output_type, "type", "json", "Formato de saída [json, yaml, toml]")
}

func initConfig() {	
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbName = path.Join(home, ".consulta-cep.db")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
