package logic

import (
	"../common"
	proto "../proto"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

type RenewalReadAoLogic struct {
	engine *xorm.Engine
	engine_master *xorm.Engine
}


func NewRenewalReadAoLogic(cfg common.GlobalCfg) (*RenewalReadAoLogic, error) {
	handle := new(RenewalReadAoLogic)

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

func (logic *RenewalReadAoLogic) GetRenewalDealInfo( ctx context.Context,req *proto.GetRenewalRequest, rsp *proto.GetRenewalResponse) error {
	if req.RealTime {
		return logic.getRenewalDealInfo(logic.engine_master, ctx, req, rsp)
	}else {
		return logic.getRenewalDealInfo(logic.engine, ctx, req, rsp)
	}
}

func (logic *RenewalReadAoLogic)getRenewalDealInfo(engine *xorm.Engine, ctx context.Context, req *proto.GetRenewalRequest, rsp *proto.GetRenewalResponse) error {
	header := new(proto.ResponseHeader)
	header.ErrNo = 0
	header.ErrMsg = "success"

	renewDealInfo := common.RenewalDealInfo{}
	renewDealInfo.Deal_id = req.DealId

	// 读取数据
	_, err := engine.Table("t_renewal_v2_deal").Get(&renewDealInfo)
	if err != nil {
		header.ErrNo = 10001
		header.ErrMsg = err.Error()
	}

	// 填充protobuf结构体
	renewalDealInfo := new(proto.RenewalDealInfo)
	renewalDealInfo.DealId = renewDealInfo.Deal_id
	renewalDealInfo.Uid = renewDealInfo.Uid
	renewalDealInfo.InsuranceBenificiary = renewDealInfo.Insurance_benificiary
	renewalDealInfo.InsuranceBenificiaryIdentity = renewDealInfo.Insurance_benificiary_identity
	renewalDealInfo.InsuranceBenificiaryBirth = renewDealInfo.Insurance_benificiary_birth
	renewalDealInfo.InsuranceBenificiarySex = renewDealInfo.Insurance_benificiary_sex
	renewalDealInfo.RenewPayState = renewDealInfo.Renew_pay_state

	// 返回请求
	rsp.Header = header
	rsp.DealInfo = renewalDealInfo

	fmt.Println(rsp.String())
	return nil
}


