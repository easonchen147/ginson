package tool

import (
	"context"
	"encoding/json"

	pkgutil "ginson/pkg/util"

	"github.com/easonchen147/foundation/log"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (c *Service) FaceLocation(ctx context.Context, filePath string) (map[string]interface{}, error) {
	output, code, err := pkgutil.ExecCmd("python3", []string{"face_location.py", filePath})
	if err != nil {
		log.Error(ctx, "get face location failed, exitcode: %d error: %v", code, err)
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		log.Error(ctx, "face location result json unmarshal failed, error: %v", err)
		return nil, err
	}
	return result, nil
}
