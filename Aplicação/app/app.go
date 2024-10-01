package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			Action: searchIp,
		},
		{
			Name: "server",
			Usage: "Busca Servidores na Internet",
			Flags: flags,
			Action: searchServer,
		},
	}
	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Resultados")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Resultado")
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}