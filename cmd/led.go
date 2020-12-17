/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	rpi "github.com/nathan-osman/go-rpigpio"
	"time"

	"github.com/spf13/cobra"
)

// ledCmd represents the led command
var ledCmd = &cobra.Command{
	Use:   "led",
	Short: "led",
	Long:  `led control`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("led called")
		led()
	},
}

var ioid int

func init() {
	ledCmd.Flags().IntVarP(&ioid, "ioid", "i", 2, "merchantId")
	ledCmd.MarkFlagRequired("ioid")
	rootCmd.AddCommand(ledCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ledCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ledCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func led() {
	p, err := rpi.OpenPin(ioid, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	//set high
	p.Write(rpi.HIGH)
	fmt.Println("start gpio test")

	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			p.Write(rpi.HIGH)
			fmt.Println("on")
			time.Sleep(time.Millisecond * 100)
			p.Write(rpi.LOW)
			fmt.Println("off")
		}
	}()

	time.Sleep(time.Hour * 2)
}
