package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/guionardo/go-gstools/cep"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
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
		err = cepOutput(cepResp, output_type, output)
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

var (
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
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbName = path.Join(home, ".consulta-cep.db")
}

func cepOutput(response *cep.CEP, output_type string, output_destiny string) error {
	var content []byte
	var err error
	if output_type == "json" {
		content, err = json.MarshalIndent(response, "", "  ")
	} else if output_type == "yaml" {
		content, err = yaml.Marshal(response)
	} else if output_type == "toml" {
		content, err = toml.Marshal(response)
	} else {
		err = fmt.Errorf("Formato de saída inválido [%s]", output_type)
	}
	if err != nil {
		return err
	}
	if output_destiny == "console" {
		fmt.Print(string(content))
		return nil
	}
	file, err := os.Create(output_destiny)
	if err == nil {
		defer file.Close()
		_, err = file.Write(content)
	}
	return err
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
