package static

import (
	"log"
	"net/http"
	"os"
)

var home, workingDir string


func init() {
	home, _ = os.UserHomeDir()
	workingDir = home + "/workspace/dropship"
}

func Run() error {
	log.Println("static server serve, http://localhost:3000")
	fs := http.FileServer(http.Dir(workingDir+"/static"))
	http.Handle("/", fs)
	err := http.ListenAndServe(":3000", nil)
	log.Println("static server terminate")
	return err
}