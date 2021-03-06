package peer

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ops/internal/model"
	"go-ops/internal/service"

	"github.com/luxingwen/pnet"
	"github.com/luxingwen/pnet/config"
	"github.com/luxingwen/pnet/log"
)

type OspPeer struct {
	*pnet.PNet
}

func (self *OspPeer) HandlerFunc(msg interface{}, msgID []byte, srcID, rpath string, pn *pnet.PNet) {
	switch v := msg.(type) {
	case *model.ResponseResCmd:
		fmt.Println("res cmd:", v.Jobid)
		err := service.Task().UpdateSubScriptTask(context.Background(), v)
		if err != nil {
			fmt.Println("update err:", err)
		}
	case *model.DownloadFileJobRes:
		err := service.Task().UpdateDownloadFileTask(context.Background(), v)
		if err != nil {
			fmt.Println("update err:", err)
		}
	case *model.PeerInfo:
		err := service.VM().Check(context.Background(), v)
		if err != nil {
			fmt.Println("check vm err:", err)
		}
	default:
		b, _ := json.Marshal(v)
		log.Error("msg handler not found,msg: %s", string(b))
	}
}

var ospPeer *OspPeer

func InitOspPeer(id string, conf *config.Config) (err error) {

	ospPeer = &OspPeer{}

	pn, err := NewPnet(id, conf, ospPeer.HandlerFunc)
	if err != nil {
		return
	}
	ospPeer.PNet = pn
	err = pn.Start()
	if err != nil {
		return
	}
	return
}

func GetOspPeer() *OspPeer {
	if ospPeer == nil {
		panic("osp peer is nil")
	}
	return ospPeer
}
