package domain

import "context"

type FileRepo interface {
	AddFile(ctx context.Context, path string, user_id int64) error
}
