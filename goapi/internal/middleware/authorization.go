package middleware

import(
	"errors"
	"net/http"

	"github.com/sparky0520/goapi/api"
	"github.com/sparky0520/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

