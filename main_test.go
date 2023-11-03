package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/kvii/show-new-features-of-protoc-gen-go/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var mockTime = time.Date(2006, 1, 2, 15, 4, 5, 0, time.Local)

var marshalOptions = protojson.MarshalOptions{
	EmitUnpopulated: true,
}

func TestMarshal(t *testing.T) {
	testCases := []struct {
		name string
		m    *pb.UserInfo
		b    []byte
	}{
		{
			name: "normal",
			m: &pb.UserInfo{
				UserId:    1,
				CreatedAt: timestamppb.New(mockTime),
				Score: &wrapperspb.Int32Value{
					Value: 100,
				},
			},
			b: []byte(`{"userId":1, "createdAt":"2006-01-02T07:04:05Z", "score":100}`),
		},
		{
			name: "zero value",
			m: &pb.UserInfo{
				UserId:    0,
				CreatedAt: nil,
				Score: &wrapperspb.Int32Value{
					Value: 0,
				},
			},
			b: []byte(`{"userId":0, "createdAt":null, "score":0}`),
		},
		{
			name: "nil value",
			m: &pb.UserInfo{
				UserId:    1,
				CreatedAt: nil,
				Score:     nil,
			},
			b: []byte(`{"userId":1, "createdAt":null, "score":null}`),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b1, err := marshalOptions.Marshal(tc.m)
			if err != nil {
				t.Errorf("proto json marshal %v", err)
			}
			if !bytes.Equal(b1, tc.b) {
				t.Errorf("proto json marshal want %v, got %v", string(tc.b), string(b1))
			}

			b2, err := json.Marshal(tc.m)
			if err != nil {
				t.Errorf("gin json marshal %v", err)
			}
			if !bytes.Equal(b2, tc.b) {
				t.Errorf("gin json marshal want %v, got %v", string(tc.b), string(b2))
			}
		})
	}
}
