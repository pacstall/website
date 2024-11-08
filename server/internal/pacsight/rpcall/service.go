package rpcall

import (
	"net/rpc"

	"github.com/joomcode/errorx"
	"pacstall.dev/webserver/internal/pacsight/repology"
	"pacstall.dev/webserver/pkg/common/pacsight"
)

func RegisterService() {
	rpc.Register(new(Service))
	rpc.HandleHTTP()
}

type Service int

func (s *Service) GetRepologyProject(args *pacsight.GetRepologyProjectArgs, reply *pacsight.GetRepologyProjectReply) error {
	project, err := repology.FindRepologyProject(repology.Packages, args.Filters)
	if err != nil {
		return errorx.ExternalError.Wrap(err, "pacsight rpc call failed")
	}

	reply.Project = project
	return nil
}
