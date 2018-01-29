package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/api/admin/health", handleAPIHealth)
	http.HandleFunc("/", handler)
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
