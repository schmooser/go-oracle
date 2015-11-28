package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-oci8"
)

type Api struct {
}

func NewApi() (api *Api) {
	nlsLang := os.Getenv("NLS_LANG")
	if !strings.HasSuffix(nlsLang, "UTF8") {
		i := strings.LastIndex(nlsLang, ".")
		if i < 0 {
			os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
		} else {
			nlsLang = nlsLang[:i+1] + "AL32UTF8"
			Log.Error(fmt.Sprintf("NLS_LANG error: should be %s, not %s!\n",
				nlsLang, os.Getenv("NLS_LANG")))
		}
	}

}

// Issue function invokes pkg_api.issue procedure and updates provided issue
// with actual database values.
func (api *Api) Issue(i *Issue) error {

}
