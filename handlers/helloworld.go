/*
 * @Author: lwnmengjing
 * @Date: 2023/10/31 23:28:06
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2023/10/31 23:28:06
 */

package handlers

import (
	"context"

	"github.com/mss-boot-io/mss-boot/pkg/server/handler"
	pb "service-grpc/proto"
)

type Handler struct {
	handler.Handler
	pb.UnimplementedHelloworldServer
}

func NewHandler() *Handler {
	return &Handler{}
}

// Call is a single request handlers called via client.Call or the generated client code
func (e *Handler) Call(c context.Context, req *pb.CallRequest) (*pb.CallResponse, error) {
	h := handler.Make(c)
	h.Log.Info("Received Call request")
	rsp := &pb.CallResponse{}
	rsp.Message = "Hello " + req.Name
	return rsp, nil
}
