// Copyright (c) 2012-2014 The upper.io/db authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package db

import (
	"errors"
)

// Shared error messages.
var (
	ErrExpectingPointer        = errors.New(`Argument must be an address.`)
	ErrExpectingSlicePointer   = errors.New(`Argument must be a slice address.`)
	ErrExpectingSliceMapStruct = errors.New(`Argument must be a slice address of maps or structs.`)
	ErrExpectingMapOrStruct    = errors.New(`Argument must be either a map or a struct.`)
	ErrNoMoreRows              = errors.New(`There are no more rows in this result set.`)
	ErrNotConnected            = errors.New(`You're currently not connected.`)
	ErrMissingDatabaseName     = errors.New(`Missing database name.`)
	ErrMissingCollectionName   = errors.New(`Missing collection name.`)
	ErrCollectionDoesNotExist  = errors.New(`Collection does not exist.`)
	ErrSockerOrHost            = errors.New(`You may connect either to a unix socket or a tcp address, but not both.`)
	ErrQueryLimitParam         = errors.New(`A query can accept only one db.Limit() parameter.`)
	ErrQuerySortParam          = errors.New(`A query can accept only one db.Sort{} parameter.`)
	ErrQueryOffsetParam        = errors.New(`A query can accept only one db.Offset() parameter.`)
	ErrMissingConditions       = errors.New(`Missing selector conditions.`)
	ErrUnsupported             = errors.New(`This action is currently unsupported on this database.`)
	ErrUndefined               = errors.New(`This value is undefined.`)
	ErrQueryIsPending          = errors.New(`Can't execute this instruction while the result set is still open.`)
	ErrUnsupportedDestination  = errors.New(`Unsupported destination type.`)
	ErrUnsupportedType         = errors.New(`This type does not support marshaling.`)
	ErrUnsupportedValue        = errors.New(`This value does not support unmarshaling.`)
	ErrUnknownConditionType    = errors.New(`Arguments of type %s can't be used as constraints.`)
)

// Deprecated but kept for backwards compatibility. See: https://github.com/upper/db/issues/18
var ErrCollectionDoesNotExists = ErrCollectionDoesNotExist
