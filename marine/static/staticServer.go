package static

import (
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
	"os"
)

var home, workingDir string
var authenticator auth.BasicAuth

func init() {
	home, _ = os.UserHomeDir()
	workingDir = home + "/workspace/dropship"
	secrets := auth.HtpasswdFileProvider(home+"/.dropship/security.htpasswd")
	authenticator = auth.BasicAuth{Realm: "localhost", Secrets: secrets}
}

func Secret(user, realm string) string {

	if user == "john" {
		// password is "hello"
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

func handleFileServer() http.HandlerFunc {
	fs := http.FileServer(http.Dir(workingDir+"/static"))
	handler := fs.ServeHTTP
	return func(w http.ResponseWriter, req *http.Request) {
		handler(w, req)
	}
}

func Run() error {
	log.Println("static server serve, http://localhost:3000")
	http.HandleFunc("/", auth.JustCheck(&authenticator, handleFileServer()))

	err := http.ListenAndServe(":3000", nil)
	log.Println("static server terminate")
	return err
}