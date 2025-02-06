package middleware

import(
	"errors"
	"net/http"

	"github.com/gevorgalaverdyan/go-playground/api"
	"github.com/gevorgalaverdyan/go-playground/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthError = errors.New("Invalid username/tokn")

func Auth(next http.Handler) http.Header{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if username=="" || token=="" {
			log.Error(UnauthError)
			api.RequestErrorHandler(w, UnauthError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDB()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthError)
			api.RequestErrorHandler(w, UnauthError)
			return
		}

		next.ServeHTTP(w, r)
	})
}