// Code generated by "stringer -type=ChunkType"; DO NOT EDIT.

package main

import "fmt"

const _ChunkType_name = "TextChunkTagChunkObjChunk"

var _ChunkType_index = [...]uint8{0, 9, 17, 25}

func (i ChunkType) String() string {
	if i < 0 || i >= ChunkType(len(_ChunkType_index)-1) {
		return fmt.Sprintf("ChunkType(%d)", i)
	}
	return _ChunkType_name[_ChunkType_index[i]:_ChunkType_index[i+1]]
}
