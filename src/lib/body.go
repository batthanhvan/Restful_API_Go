package lib

import (
	"fmt"
	"reflect"
	"runtime"

	"strconv"

	pb "github.com/batthanhvan/src/pb"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func ParseInt32Val(a string) int32 {
	if a == "" {
		return 0
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	i32 := int32(v)
	return i32
}

func ParseInt32Ptr(a string) *int32 {
	if a == "" {
		return nil
	}
	v, err := strconv.Atoi(a)
	if err != nil {
		return nil
	}
	i32 := int32(v)
	return &i32
}

func Pagination(offset, limit int32, total *int64) *pb.Pagination {
	if offset == 0 && limit == 0 && total == nil {
		return nil
	}
	return &pb.Pagination{
		Offset: offset,
		Limit:  limit,
		Total:  Int64Val(total),
	}
}

func GetFunctionName(funcname interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(funcname).Pointer()).Name()
}

func RecordError(span trace.Span, err error) {
	span.RecordError(fmt.Errorf("%+v", err))
	span.SetStatus(codes.Error, err.Error())
}
