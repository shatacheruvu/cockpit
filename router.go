package cockpit

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

// EntityResource is a child route of a parent entity
type EntityResource struct {
	Description string `json:"description,omitempty"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	Handler     string `json:"handler"`
}

// Entity represents structure of a single entity with several resources
type Entity struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Prefix      string           `json:"prefix"`
	Resources   []EntityResource `json:"resources"`
}

// AppRoutes is a structure for the main configuration file
type AppRoutes struct {
	ServiceName string   `json:"service_name"`
	Entity      []Entity `json:"entity"`
}

func (c *Cockpit) RegisterRoutes(path string) *AppRoutes {
	var err error
	var configFile []byte
	var appRoutes *AppRoutes

	if configFile, err = ioutil.ReadFile(path); err != nil {
		log.Fatalln(fmt.Errorf("failed reading routes config file %s because %v", path, err.Error()))
	}

	if err = json.Unmarshal(configFile, &appRoutes); err != nil {
		log.Fatalln(fmt.Errorf("failed to unmarshall route config: %s", err.Error()))
	}

	registerEntityResources(appRoutes, c)
	return appRoutes
}

func registerEntityResources(a *AppRoutes, c *Cockpit) {
	for _, e := range a.Entity {
		var prefixRouter *mux.Router = c.AppRouter.PathPrefix(e.Prefix).Subrouter()
		for _, r := range e.Resources {
			prefixRouter.Methods([]string{r.Method}...).Path(r.Path).Handler(buildHandler(r.Handler, c))
		}
	}
}

func buildHandler(name string, c *Cockpit) http.Handler {
	return http.HandlerFunc(reflect.ValueOf(&c.AppHandler).MethodByName(name).Interface().(func(w http.ResponseWriter, r *http.Request)))
}

