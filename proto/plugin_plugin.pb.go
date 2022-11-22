//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.9
// source: proto/plugin.proto

package imgPlugin

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const ActuatorPluginAPIVersion = 1

//export actuator_api_version
func _actuator_api_version() uint64 {
	return ActuatorPluginAPIVersion
}

var actuator Actuator

func RegisterActuator(p Actuator) {
	actuator = p
}

//export actuator_modify
func _actuator_modify(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req PluginRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := actuator.Modify(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
