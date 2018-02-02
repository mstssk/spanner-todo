//go:generate jwg -output model_json.go -transcripttag swagger .

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/favclip/ucon"
	"github.com/favclip/ucon/swagger"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {

	ucon.Middleware(UseAppengineContext)
	ucon.Orthodox()
	ucon.Middleware(swagger.RequestValidator())

	swPlugin := swagger.NewPlugin(&swagger.Options{
		Object: &swagger.Object{
			Info: &swagger.Info{Title: "spanner-todo", Version: "v1"},
		},
		DefinitionNameModifier: func(refT reflect.Type, defName string) string {
			// クライアント側とのやり取り用に自動生成したstruct名の末尾から「JSON」を除去
			if strings.HasSuffix(defName, "JSON") {
				return defName[:len(defName)-4]
			}
			return defName
		},
	})
	ucon.Plugin(swPlugin)

	ucon.HandleFunc("GET", "/swagger-ui/", func(w http.ResponseWriter, r *http.Request, c context.Context) {
		localPath := "./api/swagger-ui/" + r.URL.Path[len("/swagger-ui/"):]
		http.ServeFile(w, r, localPath)
	})

	ucon.HandleFunc("GET", "/api/admin/test", func(c context.Context) error {
		todo := Todo{Title: "test"}
		s := TodoStore{}
		newTodo, err := s.Insert(c, &todo)
		if err != nil {
			return err
		}
		log.Debugf(c, "%#v", todo)
		log.Debugf(c, "%#v", newTodo)
		return nil
	})
	ucon.HandleFunc("GET", "/api/admin/health", handleAPIHealth)
	ucon.HandleFunc("GET", "/", handler)

	todoSetup(swPlugin)

	ucon.DefaultMux.Prepare()
	http.Handle("/", ucon.DefaultMux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is dummy.")
}

func handleAPIHealth(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	log.Debugf(c, "%#v", c)

	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{
		"version":                  runtime.Version(),
		"app_id":                   appengine.AppID(c),
		"module_name":              appengine.ModuleName(c),
		"version_id":               appengine.VersionID(c),
		"datacenter":               appengine.Datacenter(c),
		"default_version_hostname": appengine.DefaultVersionHostname(c),
		"instance_id":              appengine.InstanceID(),
		"server_software":          appengine.ServerSoftware(),
	}
	json, _ := json.Marshal(res)
	fmt.Fprint(w, string(json))
}
