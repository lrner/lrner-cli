package cmd

import (
  "bufio"
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "os/user"
  "strings"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

func init() {
  rootCmd.AddCommand(loginCmd)
  loginCmd.PersistentFlags().Bool("local", true, "Run against localhost");
}

var loginCmd = &cobra.Command{
  Use:   "login",
  Short: "Print the login number of Hugo",
  Long:  `All software has logins. This is Hugo's`,
  Run: func(cmd *cobra.Command, args []string) {
  loginURL := fmt.Sprintf("%s/%s", "http://localhost:8080/api/1", "login")


  fmt.Println("Email: ")
  emailReader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
  var emailStorageString string
  emailStorageString, _ = emailReader.ReadString('\n')

  fmt.Println("Password: ")
  passwordReader := bufio.NewReader(os.Stdin) //create new reader, assuming bufio imported
  var passwordStorageString string
  passwordStorageString, _ = passwordReader.ReadString('\n')


  message := map[string]interface{}{
    "email": strings.Trim(emailStorageString, "\n"),
    "password": strings.Trim(passwordStorageString, "\n"),
  }

  bytesRepresentation, err := json.Marshal(message)
  if err != nil {
    log.Fatalln(err)
  }


  resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(bytesRepresentation))
  if err != nil {
    log.Fatalln(err)
  }

  var result map[string]interface{}

  json.NewDecoder(resp.Body).Decode(&result)

  var token string
  token = result["token"].(string)

  tokenBytes := []byte(string(token))

  usr, err := user.Current()
  check(err)

  err = ioutil.WriteFile(fmt.Sprintf("%s/%s", usr.HomeDir, viper.Get("token_file")), tokenBytes, 0644)
  check(err)

  log.Println(fmt.Sprintf("Wrote login token to %s/%s", usr.HomeDir, viper.Get("token_file")))

  },
}
