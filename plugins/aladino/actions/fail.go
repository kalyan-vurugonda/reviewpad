// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package plugins_aladino_actions

import (
	"github.com/reviewpad/reviewpad/v3/handler"
	"github.com/reviewpad/reviewpad/v3/lang/aladino"
)

func Fail() *aladino.BuiltInAction {
	return &aladino.BuiltInAction{
		Type:           aladino.BuildFunctionType([]aladino.Type{aladino.BuildStringType()}, nil),
		Code:           failCode,
		SupportedKinds: []handler.TargetEntityKind{handler.PullRequest, handler.Issue},
	}
}

func failCode(e aladino.Env, args []aladino.Value) error {
	failMessage := args[0].(*aladino.StringValue).Val

	e.GetBuiltInsReportedMessages()[aladino.SEVERITY_FATAL] = append(e.GetBuiltInsReportedMessages()[aladino.SEVERITY_FATAL], failMessage)

	return nil
}
