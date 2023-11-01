package services

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func EslintConfig() {
	eslintOptions()
}

func eslintOptions() {
	var types = []string{"Airbnb", "Standard", "Prettier", "Manual"}

	prompt := &survey.Select{
		Message: "Selecione o tipo de configuração do ESLint",
		Options: types,
	}

	var selectedType string

	survey.AskOne(prompt, &selectedType, nil)

	var eslintConfig string

	switch selectedType {
	case "Airbnb":
		eslintConfig = "eslint-config-airbnb-base"
	case "Standard":
		eslintConfig = "eslint-config-standard"
	case "Prettier":
		eslintConfig = "eslint-config-prettier"
	case "Custom":
		fmt.Println("Você escolheu configurar o ESLint manualmente.")
		return
	}

	fmt.Printf("Instalando configuração %s...\n", eslintConfig)

	cmd := exec.Command("npm", "install", "--save-dev", eslintConfig)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}
