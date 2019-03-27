package cmd

import (
  "fmt"
  "io/ioutil"

  "github.com/xeipuuv/gojsonschema"
  "github.com/spf13/cobra"
  "github.com/lrner/lrner-cli/helpers"

)

var File string
var Directory string

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func init() {
  rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
  Use:   "validate",
  Short: "Print the validate number of Hugo",
  Long:  `All software has validates. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    file_or_directory := helpers.Path{Path: args[0]}

    fmt.Println(file_or_directory.GetJSON())

    contents, err := ioutil.ReadFile("./cmd/skill_schema.json")
    check(err)
    skillSchema := string(contents)
    schemaLoader := gojsonschema.NewStringLoader(skillSchema)

    var documentLoader gojsonschema.JSONLoader

    documentLoader = gojsonschema.NewStringLoader(file_or_directory.GetJSON())

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
  },
}
