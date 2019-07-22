package daniel_xproject_handle_renewal

import (
	renewalDao "../../dao/renewal"
	proto "../../proto/renewal"
	"context"
)


type RenewalDealReadAoHandle struct {
	RenewalDealDao renewalDao.RenewalDealDao
}

func(renewal *RenewalDealReadAoHandle) GetRenewalDealByDealId(ctx context.Context ,req *proto.RenewalDealId, rsp *proto.SingleRenewalDeal) error {
	header := new(proto.ResponseHeader)
	header.ErrNo = 0
	header.ErrMsg = "success"

	realTime := req.GetRealTime()
	dealId := req.GetDealId()

	renewalDealInfo, err := renewal.RenewalDealDao.GetRenewalDealByDealId(dealId, realTime)
	if err != nil {
		// todo 写日志文件
		header.ErrNo = 10001
		header.ErrMsg = err.Error()

		rsp.Header = header
		rsp.RenewalDealInfo = nil
		return nil
	}

	rsp.Header = header
	rsp.RenewalDealInfo = renewalDealInfo
	return nil
}


