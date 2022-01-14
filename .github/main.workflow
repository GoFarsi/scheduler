workflow "Unit Tests" {
  resolves = ["cedrickring/golang-action@1.3.0"]
  on = "push"
}

action "cedrickring/golang-action@1.3.0" {
  uses = "cedrickring/golang-action@1.3.0"
}