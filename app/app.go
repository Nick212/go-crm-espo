package app

import "enube.visualstudio.com/enube-bs-cloudsolv-ssa/db"

// App contains configurations for application.
type App struct {
	Config   *Config
	Database *db.Database
	block    *Block
}

// NewContext create new context for application.
func (a *App) NewContext(version, languageID, requestID, IP string) *Context {
	return &Context{
		Logger:     NewLogger(),
		LanguageID: languageID,
		RequestID:  requestID,
		Version:    version,
		IP:         IP,
		HTTPPrefix: a.Config.HTTPPrefix,
	}
}

// New create new application instance.
func New() (app *App, err error) {
	app = &App{}

	app.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := db.InitConfig()
	if err != nil {
		return nil, err
	}

	app.Database, err = db.New(dbConfig)
	if err != nil {
		return nil, err
	}

	app.block = &Block{}

	return app, err
}

// Close finalize objects of application instance.
func (a *App) Close() error {
	// return a.Database.Close()
	return nil
}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
