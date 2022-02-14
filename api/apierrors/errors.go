package apierrors

import "errors"

var (
	NO_ENTITY_DELETED_ERROR = errors.New("No entity was deleted")
	NO_ENTITY_UPDATED_ERROR = errors.New("No entity was updated")
)
