package rm

import (
	"github.com/transaction-mesh/starfish/pkg/base/meta"
	"github.com/transaction-mesh/starfish/pkg/base/model"
)

type ResourceManagerInbound interface {
	// Commit a branch transaction.
	BranchCommit(branchType meta.BranchType, xid string, branchID int64, resourceID string, applicationData []byte) (meta.BranchStatus, error)

	// Rollback a branch transaction.
	BranchRollback(branchType meta.BranchType, xid string, branchID int64, resourceID string, applicationData []byte) (meta.BranchStatus, error)
}

type ResourceManagerOutbound interface {
	// Branch register long.
	BranchRegister(branchType meta.BranchType, resourceID string, clientID string, xid string, applicationData []byte, lockKeys string) (int64, error)

	// Branch report.
	BranchReport(branchType meta.BranchType, xid string, branchID int64, status meta.BranchStatus, applicationData []byte) error

	// Lock query boolean.
	LockQuery(branchType meta.BranchType, resourceID string, xid string, lockKeys string) (bool, error)
}

type ResourceManager interface {
	ResourceManagerInbound
	ResourceManagerOutbound

	// Register a Resource to be managed by Resource Manager.
	RegisterResource(resource model.IResource)

	// Unregister a Resource from the Resource Manager.
	UnregisterResource(resource model.IResource)

	// Get all resources managed by this manager.
	GetManagedResources() map[string]model.IResource

	// Get the BranchType.
	GetBranchType() meta.BranchType
}
