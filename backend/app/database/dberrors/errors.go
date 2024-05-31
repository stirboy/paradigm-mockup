package dberrors

import "errors"

var (
	// OptimisticLockingError occurs when no rows are updated.
	OptimisticLockingError = errors.New("optimistic locking error")
	// NotFoundError occurs when an entity is not found.
	NotFoundError = errors.New("entity not found")
)
