package micro_grade

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	s := http.Server{
		Addr:    ":15401",
		Handler: nil,
	}

	log.Info("Listening to localhost" + s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Unexpected http server error: " + err.Error())
	}
}
