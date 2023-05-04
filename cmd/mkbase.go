package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var mkBase = &cobra.Command{
	Use:   "base [string to print]",
	Short: "Genarates html/template base in /base folder based on the template name",
	Long:  `This enables you to generate the html template base in the /base folder using the template name provide. e.g Base.html, footer.html`,
	Run: func(cmd *cobra.Command, args []string) {
		tempName := strings.Join(args, " ")
		fmt.Println("generating package: " + tempName)
		err := os.MkdirAll("./templates/base", os.ModePerm)
		if err != nil {
			fmt.Println("Error in generating package", err)
		}
		tempNameext := "./templates/base/" + tempName + ".html"
		file, err2 := os.Create(tempNameext)

		if err2 != nil {
			fmt.Println("Error in generating package", err)
		} else {
			fmt.Println("Generated go file:", file.Name())

		}
		defer file.Close()

		file.WriteString("{{ define \"base/" + tempName + "\" }} \n<p>first paragraph</p>\n{{end}}")

	},
}
