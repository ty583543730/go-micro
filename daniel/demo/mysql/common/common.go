package common

type DBConfig struct {
	Host      string // 地址IP
	Port      int    // 数据库端口
	Database  string // 数据库名称
	Username  string // 账号
	Password  string // 密码
}


type RenewalDealInfo struct{
	Fdeal_id uint64
	Fbdeal_id uint64
	Fdisdeal_id string

	Finsurance_id uint64
	Fgroup_id uint64
	Fproduct_id uint64
	Fproduct_title string
	Fuid uint64

	Finsurance_sn string
	Finsurance_benificiary string
	Finsurance_benificiary_identity string
	Finsurance_benificiary_identity_type uint64
	Finsurance_benificiary_sex uint64
	Finsurance_benificiary_birth string
	Fchn string

	Frenew_deal_id uint64
	Frenew_bdeal_id uint64
	Frenew_disdeal_id string
	Frenew_uid uint64
	Frenew_pay_state uint64
	Fhas_buy uint64
	Fbuy_deal_id uint64
}