// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package main

import "fmt"

const _TokenType_name = "IdentifierTypeKeywordTypeRelationTypeValueType"

var _TokenType_index = [...]uint8{0, 14, 25, 37, 46}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}