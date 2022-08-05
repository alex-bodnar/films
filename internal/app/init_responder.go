package app

import "films-api/pkg/http/responder"

// initResponder init lib responder in app struct.
func (a *App) initResponder() {
	a.responder = responder.New()
}
