package main

import (
	"proxy/module"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var env string
	rootCmd := &cobra.Command{
		Use:   "proxy",
		Short: "simple proxy serve",
		Long:  "just simple proxy",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		},
	}
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", ".env", "environment variable\n")

	rootCmd.AddCommand(run())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("failure to launch...")
	}

}

func run() *cobra.Command {
	var port string
	var reverse string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "proxy serve",
		Long:  "simple proxy serve to forwarding requests",
		Run: func(cmd *cobra.Command, args []string) {
			remote, err := url.Parse(reverse)
			if err != nil {
				panic(err)
			}

			proxy := module.GoReverseProxy(&module.RProxy{
				Remote: remote,
			})

			log.Println("代理地址： " + reverse + " 本地监听： http://127.0.0.1:" + port)

			serveErr := http.ListenAndServe(":"+port, proxy)

			if serveErr != nil {
				panic(serveErr)
			}
		},
	}
	cmd.PersistentFlags().StringVar(&port, "port", "1874", "proxy local port")
	cmd.PersistentFlags().StringVar(&reverse, "remote", "", "remote url")

	return cmd
}
