package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mpaas/apps/cluster"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/conf"

	// 加载并初始化Runner
	"github.com/infraboard/mpaas/apps/task/runner"
	_ "github.com/infraboard/mpaas/apps/task/runner/k8s"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	jcol *mongo.Collection
	pcol *mongo.Collection
	log  logger.Logger
	task.UnimplementedJobRPCServer
	task.UnimplementedPipelineRPCServer

	job      job.Service
	pipeline pipeline.Service
	cluster  cluster.Service
}

func (i *impl) Config() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.jcol = db.Collection("job_tasks")
	i.pcol = db.Collection("pipeline_tasks")
	i.log = zap.L().Named(i.Name())
	i.job = app.GetInternalApp(job.AppName).(job.Service)
	i.pipeline = app.GetInternalApp(pipeline.AppName).(pipeline.Service)
	i.cluster = app.GetInternalApp(cluster.AppName).(cluster.Service)
	if err := runner.Init(); err != nil {
		return err
	}

	i.log.Debug("init task impl ok")
	return nil
}

func (i *impl) Name() string {
	return task.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	task.RegisterJobRPCServer(server, svr)
	task.RegisterPipelineRPCServer(server, svr)
}

func init() {
	app.RegistryInternalApp(svr)
	app.RegistryGrpcApp(svr)
}
