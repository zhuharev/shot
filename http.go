package shot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	strUrl := r.FormValue("url")
	fname := filepath.Join(s.cfg.App.CachePath, strUrl+".png")
	if _, err := os.Stat(fname); err == nil {
		fmt.Println("serve cached file,", fname)
		http.ServeFile(w, r, fname)
		return
	}

	if strUrl != "" {
		fmt.Println("file not exists, get", fname)
		s.Send(strUrl)
		w.Header().Set("Content-Type", "image/png")
		s.Draw(strUrl, w)
	}
}

func (s *Service) Serve() error {
	wg.Add(1)
	defer wg.Done()
	http.Handle("/", s)
	port := s.cfg.App.Port
	log.Println("Listen:", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
