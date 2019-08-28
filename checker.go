package healthchecks

import (
	"net/http"
)

const okMessage = "OK"
const notOkMessage = "NOT OK"

// Checker exposes an interface for k8s monitoring/health checks
type Checker interface {
	HealthHandlerFunc(w http.ResponseWriter, r *http.Request)
	ReadyHandlerFunc(w http.ResponseWriter, r *http.Request)
	SetReady(ready bool) bool
	SetHealthy(health bool) bool
}

type checker struct {
	healthy bool
	ready   bool
}

//New returns a concrete implementation of the checker
func New() Checker {
	return &checker{false, false}
}

func (c *checker) SetHealthy(healthy bool) bool {
	c.healthy = healthy
	return healthy
}

func (c *checker) HealthHandlerFunc(w http.ResponseWriter, r *http.Request) {
	c.handlerFunc(c.healthy, w, r)
}

func (c *checker) SetReady(ready bool) bool {
	c.ready = ready
	return ready
}

func (c *checker) ReadyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	c.handlerFunc(c.ready, w, r)
}

func (c *checker) handlerFunc(state bool, w http.ResponseWriter, r *http.Request) {
	message := okMessage
	if !state {
		message = notOkMessage
		w.WriteHeader(http.StatusInternalServerError)
	}

	bytes := []byte(message)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(bytes)
}
