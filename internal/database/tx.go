// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package database

import (
	"context"
	"fmt"

	"github.com/apex/log"
	"github.com/lrstanley/liam.sh/internal/database/ent"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %w", err, rerr)
	}

	return err
}

type TxFn func(ctx context.Context, logger log.Interface, db *ent.Tx) error

func RunWithTx(ctx context.Context, logger log.Interface, db *ent.Client, fn TxFn) error {
	tx, err := db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	err = fn(ctx, logger, tx)

	if err != nil {
		return Rollback(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
