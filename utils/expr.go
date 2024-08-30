package utils

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/google/uuid"
	"math/rand/v2"
	"os"
	"strings"
)

var (
	ErrNestedExpressionsNotSupported = fmt.Errorf("nested expressions are not supported, use the `+` operator instead")
	ErrFailedToCompileExpression     = fmt.Errorf("failed to compile expression")
	ErrFailedToRunExpression         = fmt.Errorf("failed to run expression")
	ErrorUnfinishedExpression        = fmt.Errorf("unfinished expression found")
)

var EvaluateExpression = evaluateExpression

// uuidFunc is a custom expression function that generates a new UUID string.
// Usage: ${uuid()}
var uuidFunc = expr.Function("uuid", func(params ...any) (any, error) {
	return uuid.NewString(), nil
})

// randomFunc is a custom expression function that returns a random integer up to a specified maximum value.
// Usage: ${random(100)}  // returns a random number between 0 and 99
var randomFunc = expr.Function("random", func(params ...any) (any, error) {
	if len(params) != 1 {
		return nil, fmt.Errorf("random function expects exactly 1 argument")
	}

	maxNumber, ok := params[0].(int)
	if !ok {
		return nil, fmt.Errorf("random function expects an integer argument")
	}

	return rand.IntN(maxNumber), nil
})

// envVarFunc is a custom expression function that retrieves the value of an environment variable.
// Usage: ${getEnv("HOME")}  // returns the value of the HOME environment variable
var envVarFunc = expr.Function("getEnv", func(params ...any) (any, error) {
	if len(params) != 1 {
		return nil, fmt.Errorf("env function expects exactly 1 argument")
	}

	key, ok := params[0].(string)
	if !ok {
		return nil, fmt.Errorf("env function expects a string argument")
	}

	return os.Getenv(key), nil
})

// evaluateExpression parses and evaluates expressions embedded within a string.
// The input string may contain expressions enclosed in `${}` that are evaluated based on the provided data map.
// Supported functions include uuid(), random(), and getEnv().
// Returns the input string with all expressions evaluated, or an error if the evaluation fails.
func evaluateExpression(input string, data map[string]interface{}) (string, error) {
	var result strings.Builder
	var expression strings.Builder
	inExpression := false
	escapeNext := false

	for i := 0; i < len(input); i++ {
		char := input[i]

		if escapeNext {
			if inExpression {
				expression.WriteByte(char)
			} else {
				result.WriteByte(char)
			}
			escapeNext = false
			continue
		}

		if char == '\\' {
			escapeNext = true
			continue
		}

		if char == '$' && i+1 < len(input) && input[i+1] == '{' {
			if inExpression {
				return "", ErrNestedExpressionsNotSupported
			}
			inExpression = true
			i++ // skip '{'
			continue
		}

		if char == '}' && inExpression {
			exprStr := expression.String()
			program, err := expr.Compile(exprStr, expr.Env(data), uuidFunc, envVarFunc, randomFunc)
			if err != nil {
				return "", fmt.Errorf("%w: %v", ErrFailedToCompileExpression, err)
			}

			exprResult, err := expr.Run(program, data)
			if err != nil {
				return "", fmt.Errorf("%w: %v", ErrFailedToRunExpression, err)
			}

			result.WriteString(fmt.Sprintf("%v", exprResult))
			expression.Reset()
			inExpression = false
			continue
		}

		if inExpression {
			expression.WriteByte(char)
		} else {
			result.WriteByte(char)
		}
	}

	// If we end with an unfinished expression, return an error
	if inExpression {
		return "", ErrorUnfinishedExpression
	}

	return result.String(), nil
}
