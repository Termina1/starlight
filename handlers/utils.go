package handlers

import(
  "github.com/google/go-github/github"
  "encoding/base64"
  "encoding/json"
  "github.com/golang/glog"
)

func searchForFile(files []string, file string) bool {
  for _, b := range files {
    if b == file {
      return true
    }
  }
  return false
}

func getFileContents(client *github.Client, file, owner, repo string) ([]byte, error) {
  content, _, response, error := client.Repositories.GetContents(owner, repo, file, &github.RepositoryContentGetOptions{})
  checkResponse(response)
  if error == nil {
    result, error := base64.StdEncoding.DecodeString(*content.Content)
    if error == nil {
      return result, nil
    } else {
      glog.Errorln("Couldn't decode base64 of ", file, " for ", owner, "/", repo, ": ", error)
      return nil, error
    }
  } else {
    glog.Errorln("Couldn't get contents of file", file, " for ", owner, "/", repo, ": ", error)
    return nil, error
  }
}

func getJsonFileContents(client *github.Client, file, owner, repo string, i interface{}) error {
  contents, error := getFileContents(client, file, owner, repo)
  if error == nil {
    error = json.Unmarshal(contents, &i)
    if error != nil {
      glog.Errorln("Couldn't decode json of ", file, " for ", owner, "/", repo, ": ", error)
      return error
    }
  } else {
    return error
  }
  return nil
}

func checkResponse(resp *github.Response) {

}