package main

import (
	"fmt"
	"os"

	"github.com/guionardo/go-gstools/cmd/shared"
	"github.com/guionardo/go-gstools/correios"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "consulta-correios [RASTREAMENTO]",
		Short: "Consulta rastreamento dos correios",
		Run: func(cmd *cobra.Command, args []string) {
			rastreio, err := correios.GetRastreio(codRastreio)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = shared.SaveOutput(rastreio, "json", output)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("É necessário informar um código de rastreamento para consulta")
			}
			codRastreio = args[0]
			return nil
		},
	}
	
	codRastreio string
	output      string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&output, "output", "console", "Saída da consulta [console, caminho do arquivo]")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
