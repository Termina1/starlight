package handlers

import(
  "github.com/google/go-github/github"
  "encoding/json"
  "github.com/golang/glog"
  "fmt"
  "errors"
  "io/ioutil"
  "net/http"
)

const GITHUB__RAW_URL = "https://raw.githubusercontent.com/"

func searchForFile(files []string, file string) bool {
  for _, b := range files {
    if b == file {
      return true
    }
  }
  return false
}

func getFileContents(client *github.Client, file, owner, repo string) ([]byte, error) {

  repoUrl := fmt.Sprintf("%v%v/%v/master/%v", GITHUB__RAW_URL, owner, repo, file)

  resp, error := http.Get(repoUrl)

  if resp.StatusCode != 200 {
    return nil, errors.New("Couldn't read file " + repoUrl)
  }

  content, error := ioutil.ReadAll(resp.Body)

  if error == nil {
    return content, nil
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
