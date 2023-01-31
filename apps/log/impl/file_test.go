package impl_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/infraboard/mpaas/apps/log"
)

func TestUploadFile(t *testing.T) {
	buf := io.NopCloser(strings.NewReader("test log 1222222sdfsdfsdsfdf"))
	defer buf.Close()

	req := log.NewUploadFileRequest("task_log", "test.log", buf)
	err := impl.UploadFile(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestD(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	req := log.NewDownloadFileRequest("task_log", "test.log", buf)
	err := impl.Download(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.String())
}
