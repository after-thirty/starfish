/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package meta

import (
	"errors"
)

type TransactionExceptionCode byte

const (
	// Unknown transaction exception code.
	TransactionExceptionCodeUnknown TransactionExceptionCode = iota

	// BeginFailed
	TransactionExceptionCodeBeginFailed

	// Lock key conflict transaction exception code.
	TransactionExceptionCodeLockKeyConflict

	// Io transaction exception code.
	IO

	// Branch rollback failed retriable transaction exception code.
	TransactionExceptionCodeBranchRollbackFailedRetriable

	// Branch rollback failed unretriable transaction exception code.
	TransactionExceptionCodeBranchRollbackFailedUnretriable

	// Branch register failed transaction exception code.
	TransactionExceptionCodeBranchRegisterFailed

	// Branch report failed transaction exception code.
	TransactionExceptionCodeBranchReportFailed

	// Lockable check failed transaction exception code.
	TransactionExceptionCodeLockableCheckFailed

	// Branch transaction not exist transaction exception code.
	TransactionExceptionCodeBranchTransactionNotExist

	// Global transaction not exist transaction exception code.
	TransactionExceptionCodeGlobalTransactionNotExist

	// Global transaction not active transaction exception code.
	TransactionExceptionCodeGlobalTransactionNotActive

	// Global transaction status invalid transaction exception code.
	TransactionExceptionCodeGlobalTransactionStatusInvalid

	// Failed to send branch commit request transaction exception code.
	TransactionExceptionCodeFailedToSendBranchCommitRequest

	// Failed to send branch rollback request transaction exception code.
	TransactionExceptionCodeFailedToSendBranchRollbackRequest

	// Failed to add branch transaction exception code.
	TransactionExceptionCodeFailedToAddBranch

	// Failed to lock global transaction exception code.
	TransactionExceptionCodeFailedLockGlobalTranscation

	// FailedWriteSession
	TransactionExceptionCodeFailedWriteSession

	// Failed to holder exception code
	FailedStore
)

// TransactionException
type TransactionException struct {
	Code    TransactionExceptionCode
	Message string
	Err     error
}

// Error
func (e *TransactionException) Error() string {
	return "TransactionException: " + e.Message
}

// Unwrap
func (e *TransactionException) Unwrap() error { return e.Err }

// TransactionExceptionOption used to construct TransactionException
type TransactionExceptionOption func(exception *TransactionException)

// WithTransactionExceptionCode
func WithTransactionExceptionCode(code TransactionExceptionCode) TransactionExceptionOption {
	return func(exception *TransactionException) {
		exception.Code = code
	}
}

// WithMessage
func WithMessage(message string) TransactionExceptionOption {
	return func(exception *TransactionException) {
		exception.Message = message
	}
}

// NewTransactionException
func NewTransactionException(err error, opts ...TransactionExceptionOption) *TransactionException {
	var ex *TransactionException
	if errors.As(err, &ex) {
		return ex
	}
	ex = &TransactionException{
		Code:    TransactionExceptionCodeUnknown,
		Message: err.Error(),
		Err:     err,
	}
	for _, o := range opts {
		o(ex)
	}
	return ex
}
