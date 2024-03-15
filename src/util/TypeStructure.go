package util

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"
)


type Resolution struct {
	Width  int
	Height int
}

type PipelineConfig struct {
	General    GeneralConfig
	Encoder    Encoder
	Decoder    Decoder
}

type GeneralConfig struct {
	BufferSize         int
	LogPath            string
	GPU_NAME           string
}

type Encoder struct {
	H264 string
	H265 string
}

type Decoder struct {
	H264 string
	H265 string
}

type PipelineInfo struct {
	RtspInfo     RtspInfo
}

type RtspInfo struct {
	ID              int
	NAME            string
	RTSP            string
	CODEC           string
	MODEL           string
	FPS             float64
	IN_WIDTH        int
	IN_HEIGHT       int
	GPU             int
	ENCODER         string
	DECODER         string
	OrgRtspAddr     string
	ResizeRtspAddr  string

	BufferSize int
	LogPath    string
	GPU_NAME     string
	DataTypeSize int
}


type Ticker struct {
	Period time.Duration
	Ticker time.Ticker
}

func CreateTicker(period time.Duration) *Ticker {
	return &Ticker{period, *time.NewTicker(period)}
}
func (t *Ticker) ResetTicker() {
	t.Ticker = *time.NewTicker(t.Period)
}
func (t *Ticker) FlushTicker() {
	for len(t.Ticker.C) > 0 {
		<-t.Ticker.C
	}
	t.Ticker.Stop()
}

type CustomError struct {
	Code    string
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %s, message: %s", e.Code, e.Message)
}

func Float32toBytes(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}
func Float32fromBytes(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
