package main

import (
  "github.com/joho/godotenv"
  "log"
  "fmt"
  "os"
  "strings"
  "bufio"
  "flag"
)

func replace_df(file_name string, output_file *string, envs map[string]string) int {
  fin, err := os.Open(file_name)
  if err != nil {
    log.Fatal("Cannot open Dockerfile")
  }

  defer fin.Close()

  scanner := bufio.NewScanner(fin)

  fout, err := os.Create(*output_file)
  if err != nil {
    log.Fatal("Cannot create output file")
  }

  defer fout.Close()

  output := bufio.NewWriter(fout)
  for scanner.Scan() {
    if strings.Contains(scanner.Text(), "DOTENV") {
      for k, v := range envs {
        output.WriteString(fmt.Sprintf("ENV %s %s\n", k, v))
      }
    }else{
      output.WriteString(scanner.Text() + "\n")
    }
  }

  output.Flush()

  if err != nil {
    log.Fatal("Cannot find DOTENV pattern")
  }

  return -1
}

func main() {
  dotenvPtr := flag.String("dotenv-file", ".env", "path to .env file or it different name")
  dockerfilePtr := flag.String("dockerfile-file", "Dockerfile.env", "path to Dockerfile before rendering")
  outputPtr := flag.String("output", "Dockerfile", "path to output Dockerfile")

  flag.Parse()

  fmt.Println("Args: dotenv-file: " + *dotenvPtr)
  fmt.Println("Args: dockerfile-file: " + *dockerfilePtr)
  fmt.Println("Args: output: " + *outputPtr)

  myEnvs, err := godotenv.Read(*dotenvPtr)
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  replace_df(*dockerfilePtr, outputPtr, myEnvs)

}
