package flow

import (
	"fmt"
	"strconv"
	"time"

	pb "github.com/openzee/point-flow/proto"
	xlsx "github.com/openzee/xlsx-loader"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 离散的一级数采的一级点位，在cache中交换的目标
type Point struct {
	Original        *xlsx.Point //原始的点位信息
	Value           interface{} //数据值
	ChangeTimestamp time.Time   //数据的变化时间
	pbPoint         *pb.Point   //协议点位
}

func (obj *Point) PointPrimaryKey() string {
	return strconv.FormatUint(obj.Original.PointPrimaryKey, 10)
}

func (obj *Point) String() string {
	return fmt.Sprintf("changeTime:%v ID:%v ioAddr:%v value:%v",
		obj.ChangeTimestamp.Format("15:04:05.000"),
		obj.Original.PointPrimaryKey,
		obj.Original.IOAddr,
		obj.Value)
}

func (obj *Point) Marshal() ([]byte, error) {

	pt, err := obj.toPbPoint()
	if err != nil {
		return nil, err
	}

	return proto.Marshal(pt)
}

// 该函数负责将各种不通协议的数据类型，统一对齐到采集后的协议类型
func (obj *Point) toPbPoint() (*pb.Point, error) {

	pt := &pb.Point{
		CreatedAt: timestamppb.New(obj.ChangeTimestamp),
		VId:       obj.Original.PointPrimaryKey,
	}

	switch x := obj.Value.(type) {
	case bool:
		pt.VType = pb.DataType_Boolean
		pt.VData = &pb.Point_BoolValue{BoolValue: x}
	case float32:
		pt.VType = pb.DataType_Float
		pt.VData = &pb.Point_FloatValue{FloatValue: x}
	case float64:
		pt.VType = pb.DataType_Double
		pt.VData = &pb.Point_DoubleValue{DoubleValue: x}
	case uint8:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case int8:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case int16:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case uint16:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case int32:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case uint32:
		pt.VType = pb.DataType_Integer
		pt.VData = &pb.Point_IntValue{IntValue: int32(x)}
	case int64:
		pt.VType = pb.DataType_Integer64
		pt.VData = &pb.Point_Int64Value{Int64Value: int64(x)}
	case uint64:
		pt.VType = pb.DataType_Integer64
		pt.VData = &pb.Point_Int64Value{Int64Value: int64(x)}
	case string:
		pt.VType = pb.DataType_String
		pt.VData = &pb.Point_StrValue{StrValue: x}

	}

	return pt, nil
}
