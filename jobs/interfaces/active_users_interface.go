// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package interfaces

import "github.com/spline-fu/mattermost-server/v5/model"

type ActiveUsersJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
