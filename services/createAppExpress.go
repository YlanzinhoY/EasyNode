package services

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func CreateExpressApp() {

	comands := []string{
		"npm i express",
		"npm i @types/express",
		"npm i dotenv",
	}

	for _, command := range comands {

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
			panic("erro ao rodar o comando")
		}
	}

	createServerTs()
	createEnv()
}

func createServerTs() {

	file, err := os.Create("server.ts")

	if err != nil {
		panic("falha ao criar o arquivo")
	}

	defer file.Close()

	code := `import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';

dotenv.config();
	
const app: Express = express();
const port = process.env.PORT;
	
app.get('/', (req: Request, res: Response) => {
	res.send('Express + TypeScript Server');
});
	
app.listen(port, () => {
	console.log('Server is running at http://localhost:' + port);
});

`

	_, err = file.WriteString(code)

	if err != nil {
		fmt.Println("falha ao escrever no arquvio server.ts")
		return
	}

}

func createEnv() {
	file, err := os.Create(`.env`)

	if err != nil {
		panic("erro ao criar o arquivo .env")
	}

	_, err = file.WriteString(`PORT=3333`)

	if err != nil {
		panic("erro ao criar um env")
	}

}
