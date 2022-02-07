package k8s_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mpaas/provider/k8s"
)

var (
	kubeConfig = ""
)

func TestGetter(t *testing.T) {
	should := assert.New(t)
	k8s.NewClient()
}

func init() {
	kubeConfig = os.Getenv("KUBE_CONFIG")
}
