package apiserver

import (
	"github.com/LilitMilante/restapi/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ApiServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewApiServer(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ApiServer) Start() error {
	err := s.configureLogger()
	if err != nil {
		return err
	}

	s.configureRouter()

	err = s.configureStore()
	if err != nil {
		return err
	}

	s.logger.Infof("starting api server")

	return http.ListenAndServe(":"+s.config.Port, s.router)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello()).Methods(http.MethodGet)
}

func (s *ApiServer) configureStore() error {
	st := store.NewStore(s.config.Store)

	err := st.Open()
	if err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *ApiServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
