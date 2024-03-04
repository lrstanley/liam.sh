// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package graphql

import (
	"context"
	"errors"
	"strings"

	"ariga.io/entcache"
	"github.com/99designs/gqlgen/graphql"
	"github.com/apex/log"
	"github.com/lrstanley/chix"
	"github.com/lrstanley/liam.sh/internal/database/ent"
	"github.com/lrstanley/liam.sh/internal/database/ent/privacy"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func errorPresenter(ctx context.Context, err error) (gerr *gqlerror.Error) {
	defer func() {
		if errors.Is(err, privacy.Deny) {
			gerr.Message = "Permission denied"
		}
	}()

	if errors.As(err, &gerr) {
		if gerr.Path == nil {
			gerr.Path = graphql.GetPath(ctx)
		}

		return gerr
	}

	path := graphql.GetPath(ctx)
	logger := log.FromContext(ctx)

	logger.WithError(err).WithField("path", gerr.Path).Error("graphql internal failure")

	return gqlerror.ErrorPathf(path, "internal error")
}

func recoverFunc(ctx context.Context, err interface{}) error {
	logger := log.FromContext(ctx)

	switch e := err.(type) {
	case error:
		logger.WithError(e).Error("graphql panic")
	case string:
		logger.WithError(errors.New(e)).Error("graphql panic")
	default:
		logger.WithField("error", e).Error("graphql panic")
	}

	return gqlerror.Errorf("internal error")
}

func requestLogger(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)

	log.FromContext(ctx).WithFields(log.Fields{
		"op":    oc.OperationName,
		"query": strings.ReplaceAll(strings.ReplaceAll(oc.RawQuery, "  ", ""), "\n", " "),
	}).Debug("handling request")

	return next(ctx)
}

func injectClient(client *ent.Client) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = ent.NewContext(ctx, client)
		return next(ctx)
	}
}

func cacheEvictAdmin(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if chix.RolesFromContext(ctx).Has("admin") {
		ctx = entcache.Evict(ctx)
	}

	return next(ctx)
}
