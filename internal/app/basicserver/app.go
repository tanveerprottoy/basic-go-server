package app

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/basic-go-server/internal/app/basicserver/module/resource"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/internal/pkg/router"
	"github.com/tanveerprottoy/basic-go-server/pkg/data/sqlxpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/validatorpkg"
)

// App struct
type App struct {
	DBClient      *sqlxpkg.Client
	router        *router.Router
	ResourceModule    *resource.Module
	Validate      *validator.Validate
}

func NewApp() *App {
	a := new(App)
	a.initComponents()
	return a
}

func (a *App) initDB() {
	a.DBClient = sqlxpkg.GetInstance()
}

func (a *App) initModules() {
	a.ResourceModule = resource.NewModule(a.DBClient.DB, a.Validate)
}

func (a *App) initModuleRouters() {
	router.RegisterUserRoutes(a.router, constant.V1, a.ResourceModule)
}

func (a *App) initValidators() {
	a.Validate = validator.New()
	_ = a.Validate.RegisterValidation("notempty", validatorpkg.NotEmpty)
}

/* func (a *App) initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"proxy.log",
	}
	cfg.Build()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	task := "taskName"
	logger.Info("failed to do task",
		// Structured context as strongly typed Field values.
		zap.String("url", task),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
} */

// Init app
func (a *App) initComponents() {
	a.initDB()
	a.router = router.NewRouter()
	a.initModules()
	a.initModuleRouters()
	a.initValidators()
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) RunDisableHTTP2() {
	srv := &http.Server{
		Handler:      a.router.Mux,
		Addr:         "127.0.0.1:8080",
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	log.Fatal(srv.ListenAndServe())
}
