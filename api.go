package xjob

import (
	"fmt"
	"net/http"

	"github.com/twiglab/xjob/pkg/xe"
)

func Registry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p xe.RegistryParam
		if err := bind(r.Body, &p); err != nil {
			jsonTo(http.StatusInternalServerError, xe.Failure(err.Error()), w)
			return
		}

		fmt.Println(p)
		jsonTo(http.StatusOK, xe.Success("OK"), w)
	}
}
