package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var mkpkgCmd = &cobra.Command{
	Use:   "pack ",
	Short: "Generates a package folder with the package_name.go in the folder based on the package name",
	Long:  `This enables you to generate packages in go containing the folder name and the go file based on the package name`,
	Run: func(cmd *cobra.Command, args []string) {
		pckname := strings.Join(args, " ")
		fmt.Println("generating package: " + pckname)
		err := os.MkdirAll("./"+pckname, os.ModePerm)
		if err != nil {
			fmt.Println("Error in generating package", err)
		}
		pcknameext := "./" + pckname + "/" + pckname + ".go"
		file, err2 := os.Create(pcknameext)

		if err2 != nil {
			fmt.Println("Error in generating package", err)
		} else {
			fmt.Println("Generated go file:", file.Name())

		}
		defer file.Close()

		file.WriteString("package " + pckname + "\n\nimport \"fmt\" \n\nfunc init(){\n\nfmt.Println(\"created package\")\n\n}")

	},
}
