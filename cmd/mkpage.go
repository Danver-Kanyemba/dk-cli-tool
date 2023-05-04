package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var mkPage = &cobra.Command{
	Use:   "page [string to print]",
	Short: "Genarates html/template page in the /page folder based on the template name",
	Long:  `This enables you to generate the html template page in the template folder using the template name provide. e.g Home.html, index.html`,
	Run: func(cmd *cobra.Command, args []string) {
		pageName := strings.Join(args, " ")
		fmt.Println("generating page: " + pageName)
		err := os.MkdirAll("./templates/pages", os.ModePerm)
		if err != nil {
			fmt.Println("Error in generating page:", err)
		}
		pageNameext := "./templates/pages/" + pageName + ".html"
		file, err2 := os.Create(pageNameext)

		if err2 != nil {
			fmt.Println("Error in generating page", err)
		} else {
			fmt.Println("Generated go file:", file.Name())

		}
		defer file.Close()

		file.WriteString("{{ define \"pages/" + pageName + "\" }} \n<p>first paragraph</p>\n{{end}}")

	},
}
