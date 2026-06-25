package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/siperal/hYZmet-localai/core/config"
	"github.com/siperal/hYZmet-localai/core/trace"
	"github.com/siperal/hYZmet-localai/pkg/grpc/proto"
	"github.com/siperal/hYZmet-localai/pkg/model"
)

func FaceAnalyze(
	ctx context.Context,
	img string,
	actions []string,
	antiSpoofing bool,
	loader *model.ModelLoader,
	appConfig *config.ApplicationConfig,
	modelConfig config.ModelConfig,
) (*proto.FaceAnalyzeResponse, error) {
	opts := ModelOptions(modelConfig, appConfig)
	faceModel, err := loader.Load(opts...)
	if err != nil {
		recordModelLoadFailure(appConfig, modelConfig.Name, modelConfig.Backend, err, nil)
		return nil, err
	}
	if faceModel == nil {
		return nil, fmt.Errorf("could not load face recognition model")
	}

	var startTime time.Time
	if appConfig.EnableTracing {
		trace.InitBackendTracingIfEnabled(appConfig.TracingMaxItems, appConfig.TracingMaxBodyBytes)
		startTime = time.Now()
	}

	res, err := faceModel.FaceAnalyze(ctx, &proto.FaceAnalyzeRequest{
		Img:          img,
		Actions:      actions,
		AntiSpoofing: antiSpoofing,
	})

	if appConfig.EnableTracing {
		errStr := ""
		if err != nil {
			errStr = err.Error()
		}
		trace.RecordBackendTrace(trace.BackendTrace{
			Timestamp: startTime,
			Duration:  time.Since(startTime),
			Type:      trace.BackendTraceFaceAnalyze,
			ModelName: modelConfig.Name,
			Backend:   modelConfig.Backend,
			Error:     errStr,
		})
	}

	return res, err
}
