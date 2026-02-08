package usecase

import (
	"RuoYi-Go/config"
	"RuoYi-Go/internal/ports/input"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
)

type CommonService struct {
	logger *zap.Logger
	config config.AppConfig
}

func NewCommonService(logger *zap.Logger, config config.AppConfig) input.CommonService {
	return &CommonService{logger: logger, config: config}
}

func (s *CommonService) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// Profile path, usually configured in config
	basePath := s.config.RuoYi.Profile
	if basePath == "" {
		basePath = "./upload"
	}

	// Generate file path: /upload/2021/01/01/uuid.ext
	fileName := file.Filename
	ext := filepath.Ext(fileName)

	now := time.Now()
	datePath := now.Format("2006/01/02")

	// Use uuid or similar for filename to avoid collision? RuoYi uses original name or uuid.
	// Let's use uuid or timestamp for simplicity + original name?
	// Actually RuoYi-Vue uses `FileUploadUtils.upload`.
	// For simplicity, let's just append random string or use timestamp.
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	relativePath := fmt.Sprintf("/upload/%s/%s", datePath, newFileName)
	absolutePath := filepath.Join(basePath, "upload", datePath)

	if err := os.MkdirAll(absolutePath, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("create directory failed: %v", err)
	}

	dst := filepath.Join(absolutePath, newFileName)
	src, err := file.Open()
	if err != nil {
		return "", "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", "", err
	}
	defer out.Close()

	if _, err = io.Copy(out, src); err != nil {
		return "", "", err
	}

	// Returns absolute path (for debugging?) and relative url
	// Usually returns url, fileName, newFileName, originalFilename
	url := "/profile" + relativePath
	return url, newFileName, nil
}

func (s *CommonService) GetResource(resource string) (string, error) {
	// Resource usually starts with /profile
	// Map /profile to local path
	basePath := s.config.RuoYi.Profile
	if basePath == "" {
		basePath = "./upload" // default
	}

	if strings.HasPrefix(resource, "/profile") {
		relPath := strings.TrimPrefix(resource, "/profile")
		return filepath.Join(basePath, relPath), nil
	}
	return "", fmt.Errorf("invalid resource path")
}
