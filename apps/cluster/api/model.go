package api

import (
	"fmt"

	"github.com/infraboard/mcube/v2/types"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

// clusterA/v1
func ClusterSetToTreeSet(set *cluster.ClusterSet) *types.ArcoDesignTree {
	tree := types.NewArcoDesignTree()
	set.ForEatch(func(c *cluster.Cluster) {
		clusterNode := tree.GetOrCreateTreeByRootKey(c.Meta.Id, c.Spec.Name, "cluster")
		clusterNode.Labels = c.Spec.Labels
		c.Deployments.ForEatch(func(item *deploy.Deployment) {
			// 服务
			serviceNode := clusterNode.GetOrCreateChildrenByKey(item.Spec.ServiceId, item.Spec.ServiceName, "service")
			serviceNode.Extra["type"] = c.Spec.Kind.String()

			// 部署
			deployNode := serviceNode.GetOrCreateChildrenByKey(item.Meta.Id, item.Spec.ServiceVersion, "deploy")
			for k, podStr := range item.Spec.K8STypeConfig.Pods {
				// Pod
				podNode := deployNode.GetOrCreateChildrenByKey(k, k, "pod")
				podObj := &v1.Pod{}
				if err := yaml.Unmarshal([]byte(podStr), podObj); err != nil {
					continue
				}
				podNode.Labels = podObj.Labels
				podNode.Extra["deploy_group"] = item.Spec.Group
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
