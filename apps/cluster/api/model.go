package api

import (
	"github.com/infraboard/mcube/types/tree"
	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/deploy"
)

func ClusterSetToTreeSet(set *cluster.ClusterSet) (*tree.ArcoDesignTreeSet, error) {
	trees := tree.NewArcoDesignTreeSet()
	set.ForEatch(func(c *cluster.Cluster) {
		svc := c.Service
		svcNode := trees.GetOrCreateTreeByRootKey(svc.Meta.Id, svc.Spec.Name)
		clusterNode := svcNode.GetOrCreateChildrenByKey(c.Meta.Id, c.Spec.Name, 1)
		c.Deployments.ForEatch(func(item *deploy.Deployment) {
			clusterNode.GetOrCreateChildrenByKey(item.Meta.Id, item.Spec.Name, 1)
		})
	})

	return nil, nil
}
