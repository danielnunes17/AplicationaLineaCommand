package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Recebe o nome"
	app.Usage = "Recebe o usuário"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips de endereço na iternet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nome do servidor na iternet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}
func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
