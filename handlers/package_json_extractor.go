package handlers

import(
  "github.com/google/go-github/github"
  "strings"
)

type pjson struct {
  Name string
  Description string
  Keywords []string
}



func ExtractPackageJson(client *github.Client, owner string, repo string, files []string, out chan string) {
  file := "package.json"
  if searchForFile(files, file) {
    var pack pjson
    error := getJsonFileContents(client, file, owner, repo, &pack)
    if error == nil {
      out <- pack.Name
      out <- pack.Description
      out <- strings.Join(pack.Keywords, " ")
    }
  }
  close(out)
}