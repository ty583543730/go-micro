package logic

import (
	//"../common"
	proto "../proto"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

type CommentReadAoLogic struct {
	engine *xorm.Engine
	engine_master *xorm.Engine
}


func NewCommentReadAoLogic(cfg common.GlobalCfg) (*CommentReadAoLogic, error) {
	handle := new(CommentReadAoLogic)

	engine, err := xorm.NewEngine("mysql", cfg.Mysql["xproject"])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	fdMapper := core.NewPrefixMapper(core.SnakeMapper{}, "F")

	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(fdMapper)

	engineMaster, err := xorm.NewEngine("mysql", cfg.Mysql["xproject"])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tbMasterMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	fbMasterMapper := core.NewPrefixMapper(core.SnakeMapper{}, "F")

	engineMaster.SetTableMapper(tbMasterMapper)
	engineMaster.SetColumnMapper(fbMasterMapper)

	handle.engine = engine
	handle.engine_master = engineMaster

	return handle, nil
}

func (logic *CommentReadAoLogic) GetCommentDealInfo( ctx context.Context,req *proto.GetCommentRequest, rsp *proto.GetCommentResponse) error {
	if req.RealTime {
		return logic.getCommentDealInfo(logic.engine_master, ctx, req, rsp)
	}else {
		return logic.getCommentDealInfo(logic.engine, ctx, req, rsp)
	}
}

func (logic *CommentReadAoLogic)getCommentDealInfo(engine *xorm.Engine, ctx context.Context, req *proto.GetCommentRequest, rsp *proto.GetCommentResponse) error {
	header := new(proto.ResponseHeader)
	header.ErrNo = 0
	header.ErrMsg = "success"

	renewDealInfo := common.CommentDealInfo{}
	renewDealInfo.Deal_id = req.DealId

	// 读取数据
	_, err := engine.Table("t_Comment_v2_deal").Get(&renewDealInfo)
	if err != nil {
		header.ErrNo = 10001
		header.ErrMsg = err.Error()
	}

	// 填充protobuf结构体
	CommentDealInfo := new(proto.CommentDealInfo)
	CommentDealInfo.DealId = renewDealInfo.Deal_id
	CommentDealInfo.Uid = renewDealInfo.Uid
	CommentDealInfo.InsuranceBenificiary = renewDealInfo.Insurance_benificiary
	CommentDealInfo.InsuranceBenificiaryIdentity = renewDealInfo.Insurance_benificiary_identity
	CommentDealInfo.InsuranceBenificiaryBirth = renewDealInfo.Insurance_benificiary_birth
	CommentDealInfo.InsuranceBenificiarySex = renewDealInfo.Insurance_benificiary_sex
	CommentDealInfo.RenewPayState = renewDealInfo.Renew_pay_state

	// 返回请求
	rsp.Header = header
	rsp.DealInfo = CommentDealInfo

	fmt.Println(rsp.String())
	return nil
}


