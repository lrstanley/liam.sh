package graphql

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/apex/log"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/liam.sh/internal/ent"
	"github.com/lrstanley/liam.sh/internal/ent/privacy"
	"github.com/lrstanley/liam.sh/internal/graphql/resolver"
	"github.com/lrstanley/liam.sh/internal/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:embed schema/*.gql
var rawSchema embed.FS

var Schema string

func init() {
	files, err := rawSchema.ReadDir("schema")
	if err != nil {
		panic(err)
	}

	buf := bytes.Buffer{}

	for i, file := range files {
		if file.IsDir() {
			continue
		}

		path := filepath.Join("schema", file.Name())

		f, err := rawSchema.Open(path)
		if err != nil {
			panic(err)
		}

		if i != 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(fmt.Sprintf("# source: %s\n", path))
		_, _ = buf.ReadFrom(f)
	}

	Schema = buf.String()
}

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

func New(db *ent.Client, cli *clix.CLI[models.Flags]) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(db, cli))
	srv.AroundOperations(requestLogger)
	srv.SetRecoverFunc(recoverFunc)
	srv.SetErrorPresenter(errorPresenter)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	if !cli.Debug {
		srv.SetQueryCache(lru.New(512))
		srv.Use(extension.FixedComplexityLimit(100))
	}

	srv.Use(extension.Introspection{})
	srv.Use(entgql.Transactioner{
		TxOpener: db,
	})

	return srv
}
