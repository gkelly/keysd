package service

import (
	"context"
	"fmt"

	"github.com/urfave/cli"
)

func pullCommands(client *Client) []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:  "pull",
			Usage: "Get public keys and sigchains from the key server",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "all", Usage: "pull everything"},
				cli.StringFlag{Name: "kid, k", Usage: "kid"},
				cli.StringFlag{Name: "user, u", Usage: "user, eg. gabriel@github"},
			},
			Action: func(c *cli.Context) error {
				req := &PullRequest{
					KID:  c.String("kid"),
					User: c.String("user"),
					All:  c.Bool("all"),
				}
				resp, err := client.ProtoClient().Pull(context.TODO(), req)
				if err != nil {
					return err
				}
				for _, kid := range resp.KIDs {
					fmt.Printf("%s\n", kid)
				}
				return nil
			},
		},
		cli.Command{
			Name:    "push",
			Usage:   "Publish public key and sigchain to a key server",
			Aliases: []string{"publish"},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "kid, k", Usage: "kid"},
			},
			Action: func(c *cli.Context) error {
				req := &PushRequest{
					KID: c.String("kid"),
				}
				resp, err := client.ProtoClient().Push(context.TODO(), req)
				if err != nil {
					return err
				}
				for _, url := range resp.URLs {
					fmt.Println(url)
				}
				return nil
			},
		},
	}
}