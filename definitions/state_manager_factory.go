package definitions

import "github.com/google/uuid"

type StateManagerFactory interface {
	CreateStateManager(id uuid.UUID) StateManager
}
