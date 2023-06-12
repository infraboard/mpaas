package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mpaas/apps/approval"
	"github.com/infraboard/mpaas/apps/job"
	"github.com/infraboard/mpaas/apps/k8s"
	"github.com/infraboard/mpaas/apps/pipeline"
	"github.com/infraboard/mpaas/apps/task"
	"github.com/infraboard/mpaas/apps/trigger"
	"github.com/infraboard/mpaas/conf"

	// 加载并初始化Runner
	"github.com/infraboard/mpaas/apps/task/runner"
	_ "github.com/infraboard/mpaas/apps/task/runner/k8s"
	"github.com/infraboard/mpaas/apps/task/webhook"
)

func init() {
	ioc.RegistryController(&impl{})
}

type impl struct {
	jcol *mongo.Collection
	pcol *mongo.Collection
	log  logger.Logger
	task.UnimplementedJobRPCServer
	task.UnimplementedPipelineRPCServer
	ioc.IocObjectImpl

	job      job.Service
	pipeline pipeline.Service
	cluster  k8s.Service
	approval approval.Service
	hook     *webhook.WebHook
	trigger  trigger.Service

	mcenter *rpc.ClientSet
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.jcol = db.Collection("job_tasks")
	i.pcol = db.Collection("pipeline_tasks")
	i.log = zap.L().Named(i.Name())
	i.job = ioc.GetController(job.AppName).(job.Service)
	i.pipeline = ioc.GetController(pipeline.AppName).(pipeline.Service)
	i.cluster = ioc.GetController(k8s.AppName).(k8s.Service)
	i.approval = ioc.GetController(approval.AppName).(approval.Service)
	i.trigger = ioc.GetController(trigger.AppName).(trigger.Service)
	i.mcenter = rpc.C()
	if err := runner.Init(); err != nil {
		return err
	}

	i.hook = webhook.NewWebHook()
	return nil
}

func (i *impl) Name() string {
	return task.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	task.RegisterJobRPCServer(server, i)
	task.RegisterPipelineRPCServer(server, i)
}
