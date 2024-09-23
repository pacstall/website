package pacsight

import (
	"net/rpc"
	"strconv"

	"pacstall.dev/webserver/internal/pacnexus/config"
	"pacstall.dev/webserver/pkg/common/types"
)

type GetRepologyProjectArgs struct {
	Filters []string
}

type GetRepologyProjectReply struct {
	Project types.RepologyApiProject
}

type PacsightRpcService struct {
	client *rpc.Client
}

func NewPacsightRpcService(addr string, port int) (*PacsightRpcService, error) {
	if !config.Repology.Enabled {
		return nil, nil
	}

	client, err := rpc.DialHTTP("tcp", addr+":"+strconv.FormatInt(int64(port), 10))
	return &PacsightRpcService{client}, err
}

func (s *PacsightRpcService) GetRepologyProject(filters []string) (types.RepologyApiProject, error) {
	args := &GetRepologyProjectArgs{Filters: filters}
	reply := &GetRepologyProjectReply{}

	err := s.client.Call("Service.GetRepologyProject", args, reply)
	if err != nil {
		return types.RepologyApiProject{}, err
	}

	return reply.Project, nil
}
