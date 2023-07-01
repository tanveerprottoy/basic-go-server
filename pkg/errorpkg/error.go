package errorpkg

import (
	"errors"
	"net/http"

	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/pkg/response"
)

func NewError(m string) error {
	return errors.New(m)
}

func HandleDBError(err error, w http.ResponseWriter) {
	if err.Error() == "sql: no rows in result set" {
		response.RespondError(http.StatusNotFound, errors.New(constant.NotFound), w)
		return
	}
	response.RespondError(http.StatusBadRequest, err, w)
}
