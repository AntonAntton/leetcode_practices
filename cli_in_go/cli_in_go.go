package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	port    string
	host    string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "mytool",
	Short: "A CLI HTTP server and API tester",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a local HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%s", host, port)
		fmt.Printf("Serving on http://%s\n", addr)
		http.ListenAndServe(addr, nil)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Make a GET request to a URL",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Printf("GET %s\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer resp.Body.Close()
		fmt.Printf("Status: %s\n", resp.Status)
		if verbose {
			fmt.Printf("Headers: %v\n", resp.Header)
		}
	},
}

func main() {
	// serve flags
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen on")
	serveCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to listen on")

	// get flags
	getCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show response headers")

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.Execute()
}
