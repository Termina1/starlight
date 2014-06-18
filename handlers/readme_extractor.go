package handlers

import(
  "encoding/base64"
  "github.com/google/go-github/github"
)

func ExtractReadme(client *github.Client, owner string, repo string, files []string, out chan string) {
  readme, _, _ := client.Repositories.GetReadme(owner, repo, &github.RepositoryContentGetOptions{})
  result, _ := base64.StdEncoding.DecodeString(*readme.Content)
  out <- string(result)
  close(out)
}