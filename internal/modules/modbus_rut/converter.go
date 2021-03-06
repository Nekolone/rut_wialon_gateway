package modbus_rut

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

type DataFromBytes struct {
	dataType  int
	dataValue string
}

func modbusConvertService(byteArray []byte, registerMap []param) (result string) {
	var str string
	for _, p := range registerMap {
		str, byteArray = getData(p, byteArray)
		addToResult(&result, str)
	}
	return result
}

func addToResult(s *string, data string) {
	if len(*s) == 0 {
		*s = data
		return
	}
	*s = fmt.Sprintf("%s,%s", *s, data)
}

func getData(p param, array []byte) (string, []byte) {
	if uint16(len(array)) < p.ByteSize {
		log.Printf("out of byte array. Array>%b", array)
		return "", array
	}
	res := convertByteByMap(p, array[:p.ByteSize])
	return fmt.Sprintf("%s:%d:%s", p.ParamId, res.dataType, res.dataValue), array[p.ByteSize:]
}

func convertByteByMap(p param, array []byte) DataFromBytes {
	switch p.DataType {
	case "int8":
		return DataFromBytes{1, fmt.Sprintf("%v", GetSigInt8(array[0]))}
	case "uint8":
		return DataFromBytes{1, fmt.Sprintf("%v", GetUnsInt8(array[0]))}
	case "int16":
		return DataFromBytes{1, fmt.Sprintf("%v", GetSigInt16(array,
			p.ByteOrder[0], p.ByteOrder[1]))}
	case "uint16":
		return DataFromBytes{1, fmt.Sprintf("%v", GetUnsInt16(array,
			p.ByteOrder[0], p.ByteOrder[1]))}
	case "int32":
		return DataFromBytes{1, fmt.Sprintf("%v", GetSigInt32(array,
			p.ByteOrder[0], p.ByteOrder[1], p.ByteOrder[2], p.ByteOrder[3]))}
	case "uint32":
		return DataFromBytes{1, fmt.Sprintf("%v", GetUnsInt32(array,
			p.ByteOrder[0], p.ByteOrder[1], p.ByteOrder[2], p.ByteOrder[3]))}
	case "float32":
		//return DataFromBytes{2, fmt.Sprintf("%0.4f",GetFloat32(array,unpack4())))}
		return DataFromBytes{2, fmt.Sprintf("%.4f", GetFloat32(array,
			p.ByteOrder[0], p.ByteOrder[1], p.ByteOrder[2], p.ByteOrder[3]))}
	case "float64":
		//return DataFromBytes{2, fmt.Sprintf("%0.4f",GetFloat32(array,unpack4())))}
		return DataFromBytes{2, fmt.Sprintf("%.4f", GetFloat64(array,
			p.ByteOrder[0], p.ByteOrder[1], p.ByteOrder[2], p.ByteOrder[3], p.ByteOrder[4],
			p.ByteOrder[5], p.ByteOrder[6], p.ByteOrder[7]))}

	case "ASCII":
		return DataFromBytes{3, GetAscii(array)}

	default:
		log.Printf("dataType error>%s", p.DataType)
		return DataFromBytes{3, " "}
	}
}

func byteToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func byteToAscii(bytes []byte) string {
	return string(bytes)
}

func byteToInt16(bytes []byte) int16 {
	return int16(binary.BigEndian.Uint16(bytes))
}

func byteToUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

func byteToInt32(bytes []byte) int32 {
	return int32(binary.BigEndian.Uint32(bytes))
}
func byteToUint32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

func byteToFloat32(bytes []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(bytes))
}

func byteToFloat64(bytes []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(bytes))
}
