package main

import (
  "github.com/lrner/lrner-cli/cmd"
  "github.com/spf13/viper"
)

func main() {
  viper.SetDefault("token_file", ".lrner_io_token")

  cmd.Execute()
}
