package handlers

import(
  "github.com/google/go-github/github"
  "github.com/golang/glog"
)

func ExtractReadme(client *github.Client, owner string, repo string, files []string, out chan string) {
  readme, error := getFileContents(client, "README.md", owner, repo)
  if error != nil {
    glog.Errorln("Couldn't get readme for ", owner, "/", repo, ": ", error)
    close(out)
    return
  }
  
  if readme == nil {
    glog.Errorln("Content of readme is nil ", owner, "/", repo)
    close(out)
    return
  }

  out <- string(readme)
  close(out)
}
