/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package logging

import (
	"context"
	"fmt"

	logging "github.com/wso2/apk/adapter/pkg/logging"
)

var logContext context.Context

type logContextKey string

// SetValueToContext used to set the value in the context
func SetValueToContext(ctx context.Context, key logContextKey, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetValueFromContext used to retrieve the value from the context
func GetValueFromContext(ctx context.Context, key logContextKey) interface{} {
	return ctx.Value(key)
}

// InitializeContext used to initialize logContext
func InitializeContext() {
	logContext = context.Background()
}

// SetValueToLogContext used to set values to logContext
func SetValueToLogContext(key logContextKey, value interface{}) {
	if logContext == nil {
		InitializeContext()
	}
	logContext = SetValueToContext(logContext, key, value)
}

// GetValueFromLogContext used to retrieve values from logContext
func GetValueFromLogContext(key logContextKey) interface{} {
	if logContext == nil {
		return nil
	}
	return GetValueFromContext(logContext, key)
}

// RemoveValueFromLogContext used to remove values from logContext
func RemoveValueFromLogContext(key logContextKey) {
	if logContext == nil {
		return
	}
	logContext = SetValueToContext(logContext, key, nil)
}

// GetErrorMessageByCode retrieve the error message corresponds to the provided error code
func GetErrorMessageByCode(code int) string {
	errorLog, ok := Mapper[code]
	if !ok {
		return fmt.Sprintf("No error message found for error code: %v", code)
	}
	return errorLog.Message
}

// PrintError prints the error details
func PrintError(code int, severity string, message string, args ...interface{}) logging.ErrorDetails {
	errorLog := logging.ErrorDetails{
		ErrorCode: code,
		Message:   fmt.Sprintf(message, args...),
		Severity:  severity,
	}
	return errorLog
}

// PrintErrorWithDefaultMessage prints the error details with default message
func PrintErrorWithDefaultMessage(code int, severity string) logging.ErrorDetails {
	errorLog := logging.ErrorDetails{
		ErrorCode: code,
		Message:   GetErrorMessageByCode(code),
		Severity:  severity,
	}
	return errorLog
}
