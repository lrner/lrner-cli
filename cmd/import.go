package cmd

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os/user"
  "github.com/spf13/viper"
  "github.com/spf13/cobra"

  "github.com/lrner/lrner-cli/helpers"

)

func init() {
  rootCmd.AddCommand(importCmd)
}

func getToken() string {
  usr, err := user.Current()
  check(err)

  contents, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", usr.HomeDir, viper.Get("token_file")))
  check(err)

  return  string(contents)
}

func postSkill(postData string) {
  token := getToken()

  updateSkillURL := fmt.Sprintf("%s/%s", "http://localhost:8080/api/1", "skills")

  log.Println("Here is the data being posted")
  log.Println(postData)

  client := &http.Client{}

  req, err := http.NewRequest("POST", updateSkillURL, bytes.NewReader([]byte(postData)))
  check(err)

  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
  resp, err := client.Do(req)
  log.Println("Response Body: ")
  log.Println(resp.Body)
  check(err)
  defer resp.Body.Close()

}

var importCmd = &cobra.Command{
  Use:   "import",
  Short: "Print the import number of Hugo",
  Long:  `All software has imports. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
    file_or_directory := helpers.Path{Path: args[0]}

    file_or_directory.Summary()
    if file_or_directory.SkillYAMLExists(file_or_directory.Path) {

      var postData = file_or_directory.GetJSON()
      log.Printf("Importing skill from %s\n", args[0])
      postSkill(postData)
    } else {
      log.Printf("Importing skills from %s\n", args[0])
      var postData string
      var index int
      for index, postData = range file_or_directory.EachJSON() {

        log.Printf("----------------skill %s-------------------", index)

        log.Println("------------------------------------------")
        postSkill(postData)      }
        log.Println("------------------------------------------")

    }
  },
}
