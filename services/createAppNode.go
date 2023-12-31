package services

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

		var cmd *exec.Cmd

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell", "-Command", command)
		} else {
			cmd = exec.Command("bash", "-c", command)
		}

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
		"init.md",
		".env.example",
	}

	for _, createFile := range files {
		_, err := os.Create(createFile)

		if err != nil {
			panic("erro")
		}

		ignore := `\node_modules 
.env
`
		err = os.WriteFile(".gitignore", []byte(ignore), os.ModePerm)

		if err != nil {
			panic("erro bugado")
		}
		err = os.WriteFile("init.md", []byte(readmeFile()), os.ModePerm)

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

func readmeFile() string {

	file := "- Para iniciar o server.ts basta utilizar o comando `npx tsx server.ts` "
	return file

}
