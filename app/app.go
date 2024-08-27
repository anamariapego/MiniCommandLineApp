package app

import (
	"fmt"
	"log"
	"net"
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com.br",
		},
	}

	//Criando os comandos
	app.Commands = []cli.Command{
		// Criando o primeiro comando
		{
			Name:   "ips",
			Usage:  "Buscar Ips de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome do servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

// Função para buscar e printar os IPs no terminal
func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("Erro ao buscar IPs para o host %s: %v", host, err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Função para buscar os nomes dos servidores
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host) // name server
	if err != nil {
		log.Fatalf("Erro ao buscar nomes dos servidores para o host %s: %v", host, err)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}