package handlers

import(
  "github.com/google/go-github/github"
  "encoding/base64"
  "encoding/json"
)

func searchForFile(files []string, file string) bool {
  for _, b := range files {
    if b == file {
      return true
    }
  }
  return false
}

func getFileContents(client *github.Client, file, owner, repo string) []byte {
  content, _, _, error := client.Repositories.GetContents(owner, repo, file, &github.RepositoryContentGetOptions{})
  if error == nil {
    result, _ := base64.StdEncoding.DecodeString(*content.Content)
    return result
  }
  return nil
}

func getJsonFileContents(client *github.Client, file, owner, repo string, i interface{}) {
  contents := getFileContents(client, file, owner, repo)
  json.Unmarshal(contents, &i)
}