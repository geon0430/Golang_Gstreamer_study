package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	apipkg "go_vms/src/api/util"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func InfoConverter(
	rtspConfig apipkg.APIstruct,
	apiConfig apipkg.RTSPstruct,
	globalConfig PipelineConfig) PipelineInfo {

	var pipelineInfo PipelineInfo
	pipelineInfo.RtspInfo.ID = apiConfig.ID
	pipelineInfo.RtspInfo.NAME = apiConfig.NAME
	pipelineInfo.RtspInfo.RTSP = apiConfig.RTSP
	pipelineInfo.RtspInfo.CODEC = apiConfig.CODEC
	pipelineInfo.RtspInfo.MODEL = returnModelName(apiConfig)
	pipelineInfo.RtspInfo.FPS = apiConfig.FPS
	pipelineInfo.RtspInfo.IN_WIDTH = apiConfig.IN_WIDTH
	pipelineInfo.RtspInfo.IN_HEIGHT = apiConfig.IN_HEIGHT
	pipelineInfo.RtspInfo.GPU = rtspConfig.GPU
	pipelineInfo.RtspInfo.ENCODER = returnEncoder(apiConfig.CODEC, globalConfig)
	pipelineInfo.RtspInfo.DECODER = returnDecoder(apiConfig.CODEC, globalConfig)
	pipelineInfo.RtspInfo.BufferSize = globalConfig.General.BufferSize
	pipelineInfo.RtspInfo.LogPath = globalConfig.General.LogPath
	pipelineInfo.RtspInfo.GPU_NAME = returnGPUName(globalConfig)

	return pipelineInfo
}

func returnEncoder(CODEC string, globalConfig PipelineConfig) string {
	codec := strings.ToUpper(CODEC) // codec = H264

	r := reflect.ValueOf(globalConfig.Encoder)
	encoder := reflect.Indirect(r).FieldByName(codec)
	return encoder.String()
}

func returnDecoder(CODEC string, globalConfig PipelineConfig) string {
	codec := strings.ToUpper(CODEC) // codec = H264

	r := reflect.ValueOf(globalConfig.Decoder)
	decoder := reflect.Indirect(r).FieldByName(codec)
	return decoder.String()
}

func returnGPUName(globalConfig PipelineConfig) string {
	return globalConfig.General.GPU_NAME
}

func returnModelName(apiConfig apipkg.RTSPstruct) string {

	modelName := apiConfig.MODEL

	gpuNames, err := apipkg.GetGPUNames()
	if err != nil {
		return modelName
	}

	re := regexp.MustCompile(`NVIDIA GeForce RTX 4090`)

	for _, gpuName := range gpuNames {
		if re.MatchString(gpuName) {
			modelName += "_4090"
			break
		}
	}

	return modelName
}
