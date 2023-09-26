package api

import (
	"fmt"

	"github.com/infraboard/mcube/types/tree"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

// clusterA/v1
func ClusterSetToTreeSet(set *cluster.ClusterSet) *tree.ArcoDesignTree {
	tree := tree.NewArcoDesignTree()
	set.ForEatch(func(c *cluster.Cluster) {
		svc := c.Service
		clusterNode := tree.GetOrCreateTreeByRootKey(c.Spec.Name, c.Spec.Describe)
		c.Deployments.ForEatch(func(item *deploy.Deployment) {
			deployNode := clusterNode.GetOrCreateChildrenByKey(item.Meta.Id,
				fmt.Sprintf("%s_%s", svc.Spec.Name, item.Spec.Name))
			for k, podStr := range item.Spec.K8STypeConfig.Pods {
				podNode := deployNode.GetOrCreateChildrenByKey(k, k)
				podObj := &v1.Pod{}
				if err := yaml.Unmarshal([]byte(podStr), podObj); err != nil {
					continue
				}
				podNode.Extra["cluster_id"] = item.Spec.K8STypeConfig.ClusterId
				podNode.Extra["namespace"] = podObj.Namespace
				podNode.Extra["pod_name"] = podObj.Name
				podNode.Extra["status"] = string(podObj.Status.Phase)
				podNode.Extra["message"] = fmt.Sprintf("%s,%s", podObj.Status.Reason, podObj.Status.Message)
				podNode.SetTitle(podObj.Status.PodIP)
				podNode.IsLeaf = true
			}
		})
	})

	return tree
}
