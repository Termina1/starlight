package handlers

import(
  "github.com/google/go-github/github"
  "github.com/golang/glog"
)

func ExtractRepoInfo(client *github.Client, owner string, repo string) (*github.Repository, error) {
  info, response, error := client.Repositories.Get(owner, repo)
  checkResponse(response)
  if error != nil {
    glog.Errorln("Coulnd't get repository info ", owner, "/", repo, ": ", error)
    return info, error
  } else {
    return info, nil
  }
}