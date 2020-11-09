// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package slashcommands

import (
	goi18n "github.com/mattermost/go-i18n/i18n"
	"github.com/spline-fu/mattermost-server/v5/app"
	"github.com/spline-fu/mattermost-server/v5/model"
)

type ShortcutsProvider struct {
}

const (
	CMD_SHORTCUTS = "shortcuts"
)

func init() {
	app.RegisterCommandProvider(&ShortcutsProvider{})
}

func (me *ShortcutsProvider) GetTrigger() string {
	return CMD_SHORTCUTS
}

func (me *ShortcutsProvider) GetCommand(a *app.App, T goi18n.TranslateFunc) *model.Command {
	return &model.Command{
		Trigger:          CMD_SHORTCUTS,
		AutoComplete:     true,
		AutoCompleteDesc: T("api.command_shortcuts.desc"),
		AutoCompleteHint: "",
		DisplayName:      T("api.command_shortcuts.name"),
	}
}

func (me *ShortcutsProvider) DoCommand(a *app.App, args *model.CommandArgs, message string) *model.CommandResponse {
	// This command is handled client-side and shouldn't hit the server.
	return &model.CommandResponse{
		Text:         args.T("api.command_shortcuts.unsupported.app_error"),
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
	}
}
