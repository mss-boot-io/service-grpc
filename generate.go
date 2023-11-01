package main

/*
 * @Author: lwnmengjing
 * @Date: 2023/10/31 23:07:00
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/10/31 23:07:00
 */

//go:generate protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative proto/*.proto
