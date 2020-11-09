// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package main

import (
	"net/http"

	"github.com/spline-fu/mattermost-server/v5/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn, brw, err := hj.Hijack()
	if conn == nil || brw == nil || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.Write([]byte("HTTP/1.1 200\n\nOK"))
	conn.Close()
}

func main() {
	plugin.ClientMain(&Plugin{})
}
