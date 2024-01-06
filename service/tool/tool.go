package tool

import (
	"context"
	"encoding/json"
	"ginson/pkg/conf"
	pkgutil "ginson/pkg/util"
	"github.com/easonchen147/foundation/log"
	"path/filepath"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// FaceLocation 本地调试使用 docker run -v %cd%:/data/images/ face_location:latest /data/images/%s
// FaceLocation 服务器使用  python3 face_location.py %s
func (c *Service) FaceLocation(ctx context.Context, filePath string) (map[string]interface{}, error) {
	//output, code, err := pkgutil.ExecCmd("python3", []string{"face_location.py", filePath})
	output, code, err := pkgutil.ExecCmd("docker", []string{"run", "--name", filepath.Base(filePath),
		"-v", conf.ExtConf().UploadImagePath + ":/data/images/", "face_location:latest", "/data/images/" + filepath.Base(filePath)})
	if err != nil {
		log.Error(ctx, "get face location failed, exitcode: %d error: %v", code, err)
		return nil, err
	}
	defer pkgutil.ExecCmd("docker", []string{"rm", filepath.Base(filePath)})

	var result map[string]interface{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		log.Error(ctx, "face location result json unmarshal failed, error: %v", err)
		return nil, err
	}
	return result, nil
}
