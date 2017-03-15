// Code generated by protoc-gen-grpc-gateway
// source: mailinglist.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_MailingList_SendEmail_0(ctx context.Context, marshaler runtime.Marshaler, client MailingListClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SendEmailRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.SendEmail(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_MailingList_Subscribe_0(ctx context.Context, marshaler runtime.Marshaler, client MailingListClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SubscribeRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Subscribe(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterMailingListHandlerFromEndpoint is same as RegisterMailingListHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMailingListHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMailingListHandler(ctx, mux, conn)
}

// RegisterMailingListHandler registers the http handlers for service MailingList to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMailingListHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewMailingListClient(conn)

	mux.Handle("POST", pattern_MailingList_SendEmail_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_MailingList_SendEmail_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_MailingList_SendEmail_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_MailingList_Subscribe_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_MailingList_Subscribe_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_MailingList_Subscribe_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MailingList_SendEmail_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"_appscode", "api", "mail", "v1beta1", "json"}, ""))

	pattern_MailingList_Subscribe_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"_appscode", "api", "mailinglist", "v1beta1", "subscribe", "json"}, ""))
)

var (
	forward_MailingList_SendEmail_0 = runtime.ForwardResponseMessage

	forward_MailingList_Subscribe_0 = runtime.ForwardResponseMessage
)
