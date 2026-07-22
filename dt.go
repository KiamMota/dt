package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("uso: dt <arquivos> <pasta_destino>")
		os.Exit(1)
	}

	itens := os.Args[1 : len(os.Args)-1]
	destino := os.Args[len(os.Args)-1]

	err := os.MkdirAll(destino, 0755)
	if err != nil {
		fmt.Printf("erro criando destino: %v\n", err)
		os.Exit(1)
	}

	for _, item := range itens {
		nome := filepath.Base(item)

		if nome == destino {
			continue
		}

		destinoFinal := filepath.Join(destino, nome)

		err := os.Rename(item, destinoFinal)
		if err != nil {
			fmt.Printf("erro movendo %s: %v\n", item, err)
			continue
		}

		fmt.Printf("movido: %s -> %s\n", item, destinoFinal)
	}
}
