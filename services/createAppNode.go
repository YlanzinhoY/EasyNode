package services

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateNodeApp() {

	npmRun()
	createFiles()

}

func npmRun() {
	commands := []string{
		"npm init -y",
		"npm install -D typescript tsx",
		"npx tsc --init",
	}

	for _, command := range commands {
		cmd := exec.Command("powershell", "-Command", command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Erro ao executar o comando: %s\n", command)
			panic(err)
		}

		fmt.Printf("Comando executado com sucesso: %s\n", command)
	}
}

func createFiles() {
	files := []string{
		"dockerfile",
		"docker-compose.yml",
		".gitignore",
		"Prettier",
	}

	for _, createFile := range files {
		_, err := os.Create(createFile)

		if err != nil {
			panic("erro")
		}

		err = os.WriteFile(".gitignore", []byte("/node_modules"), os.ModePerm)

		if err != nil {
			panic("erro bugado")
		}

		if err != nil {
			fmt.Printf("Erro ao criar o arquivo %s", createFile)
		} else {
			fmt.Printf("Arquivo %s criado com scuesso!\n", createFile)
		}

	}
}
