package shot

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	"time"
)

type Service struct {
	cfg       *Config
	ch        chan string
	sriptPath string

	blockMu sync.Mutex
	block   map[string]struct{}
}

var wg sync.WaitGroup

func New(filepath string) (*Service, error) {
	cfg, e := NewConfig(filepath)
	if e != nil {
		return nil, e
	}
	return NewFromConfig(cfg)
}

func NewFromConfig(cfg *Config) (*Service, error) {
	s := new(Service)
	s.cfg = cfg
	s.ch = make(chan string)
	s.block = map[string]struct{}{}
	return s, nil
}

func (s *Service) Run() error {
	bts, e := Asset("r.js")
	if e != nil {
		return fmt.Errorf("script not found")
	}
	f, e := ioutil.TempFile("", "shot")
	if e != nil {
		return fmt.Errorf("script not found")
	}
	_, e = f.Write(bts)
	if e != nil {
		return e
	}
	s.sriptPath = f.Name()
	f.Close()
	fmt.Printf("Run %d workers\n", s.cfg.App.Workers)
	for i := 0; i < s.cfg.App.Workers; i++ {
		go s.Work(s.ch)
	}
	return nil
}

func (s *Service) Work(ch chan string) {
	for {
		select {
		case u := <-ch:
			s.Shot(u)
		}
	}
}

func (s *Service) Send(u string) {
	s.blockMu.Lock()
	defer s.blockMu.Unlock()
	if _, ok := s.block[u]; ok {
		return
	}
	s.ch <- u
	s.block[u] = struct{}{}
}

func (s *Service) Wait() {
	wg.Wait()
}

func (s *Service) Shot(u string) {
	wg.Add(1)
	cmd := exec.Command(s.cfg.App.PhantomJsPath,
		//"--web-security=false",
		//"--disk-cache=true",
		s.sriptPath,
		u,
		u,
		"20000",
		s.cfg.App.CachePath)
	var out bytes.Buffer
	cmd.Stdout = &out

	start := time.Now()
	e := cmd.Run()
	if e != nil {
		log.Fatal(e)
	}

	fmt.Printf("in all caps: %q\n", out.String())

	fmt.Println(time.Since(start))
	wg.Done()
	s.blockMu.Lock()
	delete(s.block, u)
	s.blockMu.Unlock()
}
