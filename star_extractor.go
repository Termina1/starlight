package main

import (
  "github.com/google/go-github/github"
  "fmt"
  // "encoding/base64"
)

func StarExtractor() func() {
  client := github.NewClient(nil)
  return func() {
    fileNames := getFileNames(client)
    fmt.Println(fileNames)
  }
}

func getFileNames(client *github.Client) []string {
  _, dir, _, _ := client.Repositories.GetContents("google", "go-github", "/", &github.RepositoryContentGetOptions{})
  fileNames := make([]string, len(dir))
  for i, file := range dir {
    // result, _ := base64.StdEncoding.DecodeString(*file.Content)
    fileNames[i] = *file.Name
  }
  return fileNames
}