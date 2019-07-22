package common

type GlobalCfg struct {
	Consul map[string]string `json:"consul"`
	Mysql map[string]string `json:"mysql"`
}

type RenewalDealInfo struct{
	Deal_id uint64
	Bdeal_id uint64
	Disdeal_id string

	Insurance_id uint64
	Group_id uint64
	Product_id uint64
	Product_title string
	Uid uint64

	Insurance_sn string
	Insurance_benificiary string
	Insurance_benificiary_identity string
	Insurance_benificiary_identity_type uint64
	Insurance_benificiary_sex uint32
	Insurance_benificiary_birth string
	Chn string

	Renew_deal_id uint64
	Renew_bdeal_id uint64
	Renew_disdeal_id string
	Renew_uid uint64
	Renew_pay_state uint32
	Has_buy uint64
	Buy_deal_id uint64
}