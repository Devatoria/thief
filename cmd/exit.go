package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:     "exit",
	Short:   "re-joins the PID 1 cgroups",
	Example: "thief exit",
	Run: func(cmd *cobra.Command, args []string) {
		if err := driver.Exit(); err != nil {
			fmt.Printf("Error exiting current cgroup: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Exited successfully!")
	},
}
