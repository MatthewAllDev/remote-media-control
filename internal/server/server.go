package server

import (
	"RemoteMediaControl/internal/keys"
	"RemoteMediaControl/internal/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"text/template"
)

type App struct {
	Mux             http.Handler
	key             *keys.Key
	cachedTemplates *template.Template
	debugMode       bool
	staticPath      string
}

func NewApp(staticPath string, debug bool) (*App, error) {
	app := &App{}
	app.Mux = utils.ChainMiddleware(app.setupRouter(), utils.CspMiddleware, utils.LoggingMiddleware)
	app.debugMode = debug
	app.staticPath = staticPath
	var err error
	app.key, err = keys.NewKey()
	if err != nil {
		return nil, err
	}
	err = app.initTemplates()
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homeHandler)
	mux.HandleFunc("/api", app.apiHandler)
	mux.HandleFunc("/static/", app.staticFileHandler)
	return mux
}

func (app *App) initTemplates() error {
	var templates []string
	re := regexp.MustCompile(fmt.Sprintf(`^[a-z0-9]+(_%v)*.html$`, runtime.GOOS))
	err := filepath.WalkDir(filepath.Join(app.staticPath, "template"), func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && re.MatchString(filepath.Base(path)) {
			templates = append(templates, path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(templates) == 0 {
		return fmt.Errorf("no valid templates found in directory")
	}
	app.cachedTemplates, err = template.ParseFiles(templates...)
	return err
}
