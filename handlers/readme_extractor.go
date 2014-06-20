package handlers

import(
  "encoding/base64"
  "github.com/google/go-github/github"
  "github.com/golang/glog"
)

func ExtractReadme(client *github.Client, owner string, repo string, files []string, out chan string) {
  readme, response, error := client.Repositories.GetReadme(owner, repo, &github.RepositoryContentGetOptions{})
  checkResponse(response)
  if error != nil {
    glog.Errorln("Couldn't get readme for ", owner, "/", repo, ": ", error)
    close(out)
    return
  }
  var result []byte
  if readme.Content != nil {
    result, error = base64.StdEncoding.DecodeString(*readme.Content)
    close(out)
    return
  }
  if error != nil {
    glog.Errorln("Couldn't decode base64 sequence of readme for ", owner, "/", repo, ": ", error)
    close(out)
    return
  }
  out <- string(result)
  close(out)
}