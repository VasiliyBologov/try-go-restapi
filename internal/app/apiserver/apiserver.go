package apiserver

import (
	"io"
	"net/http"

	"github.com/VasiliyBologov/try-go-restapi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API server ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Start api server!")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIserver) configureStore() error {

	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info("Hello world endpoint is runing")
		io.WriteString(w, "Hello world from Go lang")
	}
}
