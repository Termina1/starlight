package handlers

import(
  "github.com/google/go-github/github"
  "regexp"
)

func ExtractGemspec(client *github.Client, owner string, repo string, files []string, out chan string) {
  file := repo + ".gemspec"
  if searchForFile(files, file) {
    content, error := getFileContents(client, file, owner, repo)
    contentS := string(content)

    if error == nil {
      patterns := []*regexp.Regexp{
        regexp.MustCompile(`\.description\s*=\s*("|'|%q\{|%Q\{)(.*?)("|'|\})`),
        regexp.MustCompile(`\.name\s*=\s*"(|')(.*?)("|')`),
        regexp.MustCompile(`\.summary\s*=\s*("|'|%q\{|%Q\{)(.*?)("|'|\})`),
      }

      var result []string
      for _, regex := range patterns {
        result = regex.FindStringSubmatch(contentS)
        if len(result) > 1 {
          out <- result[2]
        }
      }
    }
  }
  close(out)
}