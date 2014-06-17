package main


func main() {
  // worker := func(repo string) {
  //   fmt.Println(repo)
  // }
  // beam := CreateStarBeam(worker, 10)
  // beam.launch()
  StarExtractor()()
  select {}
}