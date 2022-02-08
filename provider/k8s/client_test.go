package k8s_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mpaas/provider/k8s"
)

var (
	kubeConfig = ""
)

func TestGetter(t *testing.T) {
	should := assert.New(t)
	client, err := k8s.NewClient(kubeConfig)
	should.NoError(err)
	v, err := client.ServerVersion()
	should.NoError(err)
	fmt.Println(v)
	fmt.Println(client.CurrentContext())
	fmt.Println(client.CurrentCluster())
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	kc, err := ioutil.ReadFile(filepath.Join(wd, "kube_config.yml"))
	if err != nil {
		panic(err)
	}
	kubeConfig = string(kc)
}
