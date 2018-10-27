package tools

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/Casxt/GoExcel/config"
)

func GetUserGroup(w http.ResponseWriter, r *http.Request) (group []string) {

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		http.Error(w, "Not authorized", 401)
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return nil
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		http.Error(w, "Not authorized", 401)
		return nil
	}

	if config.Users[pair[0]].Pass != pair[1] {
		http.Error(w, "Not authorized", 401)
		return nil
	}

	return config.Users[pair[0]].Group
}
