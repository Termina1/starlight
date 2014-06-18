package handlers

import(
  "github.com/google/go-github/github"
  "github.com/golang/glog"
)

func ExtractFileNames(client *github.Client, owner string, repo string) ([]string, error) {
  _, dir, response, error := client.Repositories.GetContents(owner, repo, "/", &github.RepositoryContentGetOptions{})
  checkResponse(response)
  if error != nil {
    glog.Errorln("Couldn't get list of files for ", owner, "/", repo, ": ", error)
    return nil, error
  }
  fileNames := make([]string, len(dir))
  for i, file := range dir {
    fileNames[i] = *file.Name
  }
  return fileNames, nil
}