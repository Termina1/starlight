package main

import (
  "github.com/google/go-github/github"
  "strings"
  "sync"
  "github.com/Termina1/starlight/handlers"
  "code.google.com/p/goauth2/oauth"
  "github.com/golang/glog"
)

type StarReaperWrap func(*github.Client, string, string, []string) <-chan string
type StarReaper func(*github.Client, string, string, []string, chan string)

func ExtractorWrapper(ex StarReaper) StarReaperWrap {
  return func(client *github.Client, owner string, repo string, files []string) <-chan string {
    out := make(chan string)
    go func() {
      ex(client, owner, repo, files, out)
    }()
    return out
  }
}

func StarExtractor(mconf MongoConf, token string) func(string) {
  transp := &oauth.Transport{
    Token: &oauth.Token{AccessToken: token},
  }
  whandlers := []StarReaperWrap{ExtractorWrapper(handlers.ExtractReadme), ExtractorWrapper(handlers.ExtractPackageJson)}
  client := github.NewClient(transp.Client())
  glog.Info("Acquired client with token")
  mongoSession := CreateMongoClient(mconf)
  glog.Info("Create MongoDB connection for extractor")

  return func(repo string) {
    splitted := strings.Split(repo, "/")
    owner, repository := splitted[0], splitted[1]
    fileNames, error := handlers.ExtractFileNames(client, owner, repository)
    info, error := handlers.ExtractRepoInfo(client, owner, repository)
    if error != nil {
      return
    }
    repositoryInfo := StarRepo{info.FullName, true, info.StargazersCount, info.ForksCount, info.Description, nil}

    allChannels := make([]<-chan string, len(whandlers))
    for i, extractor := range whandlers {
      allChannels[i] = extractor(client, owner, repository, fileNames)
    }
    out := StarExtractorCompose(allChannels)
    var searchField string
    for value := range out {
      searchField += " " + value
    }

    repositoryInfo.SearchField = &searchField
    StarRepoUpdate(mongoSession, repo, &repositoryInfo)
  }
}

func StarExtractorCompose(chans []<-chan string) <-chan string {
  out := make(chan string)
  process := func(channel <-chan string, wg *sync.WaitGroup) {
    for result := range channel {
      out <- result
    }
    wg.Done()
  }
  go func() {
    var wg sync.WaitGroup
    wg.Add(len(chans))
    for  _, channel := range chans {
      go process(channel, &wg)
    }
    wg.Wait()
    close(out)
  }()
  return out
}