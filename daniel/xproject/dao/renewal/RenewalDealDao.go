package daniel_xproject_dao_renewal

import (
	proto "../../proto/renewal"
	"../../common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/micro/go-micro/config"
	"github.com/pkg/errors"
	"log"
	"xorm.io/core"
)

type IRenewalDealDao interface {
	GetRenewalDealByDealId(dealId uint64, realTime bool) (proto.SingleRenewalDeal, error)
}

type RenewalDealDao struct {
	engine *xorm.Engine
	engine_master *xorm.Engine
}

func NewRenewalDealDao() (*RenewalDealDao, error) {
	var cfg common.GlobalCfg
	config.LoadFile("../../config/Mysql.json")
	err := config.Scan(&cfg)
	if err != nil {
		// todo 等下写日志到文件
		return nil, err
	}

	handle := new(RenewalDealDao)

	engine, err := xorm.NewEngine("mysql", cfg.Mysql["xproject"])
	if err != nil {
		return nil, err
	}

	// 结构体和数据库表字段使用驼峰命名的映射规则
	engine.SetMapper(core.SameMapper{})

	// 为表名基于SnakeMapper统一加前缀
	tbMapper := core.NewPrefixMapper(core.SameMapper{}, "t_")
	// 为字段名基于SnakeMapper统一加前缀
	fdMapper := core.NewPrefixMapper(core.SameMapper{}, "F")

	engine.SetTableMapper(tbMapper)
	engine.SetColumnMapper(fdMapper)

	// 主库默认用从库的配置
	engineMaster, err := xorm.NewEngine("mysql", cfg.Mysql["xproject"])
	if err != nil {
		// todo 等下写日志到文件
		return nil, err
	}

	engineMaster.SetMapper(core.SameMapper{})

	tbMasterMapper := core.NewPrefixMapper(core.SameMapper{}, "t_")
	fbMasterMapper := core.NewPrefixMapper(core.SameMapper{}, "F")

	engineMaster.SetTableMapper(tbMasterMapper)
	engineMaster.SetColumnMapper(fbMasterMapper)

	handle.engine = engine
	handle.engine_master = engineMaster
	return handle, nil
}

func (dao *RenewalDealDao) GetRenewalDealByDealId(dealId uint64, realTime bool) (*proto.RenewalDealInfo, error) {
	if realTime {
		return getRenewalDealInfoByDealId(dao.engine_master, dealId)
	}else {
		return getRenewalDealInfoByDealId(dao.engine, dealId)
	}
}

func getRenewalDealInfoByDealId(engine *xorm.Engine, dealId uint64) (*proto.RenewalDealInfo, error){
	renewDealInfo := new(common.RenewalDealInfo)
	renewDealInfo.Deal_id = dealId

	has, err := engine.Table("t_renewal_v2_deal").Get(renewDealInfo)
	if err != nil {
		log.Println(err) // todo 等下写日志到文件
		return nil, err
	}

	if !has {
		log.Println("不存在该续保订单")
		err := errors.New("不存在该续保订单")
		return nil, err
	}

	renewalDealInfo := common.ConvertRenewalDealFromOrg(*renewDealInfo)
	return renewalDealInfo, nil
}

