package input

import "mime/multipart"

type CommonService interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	GetResource(resource string) (string, error)
}
