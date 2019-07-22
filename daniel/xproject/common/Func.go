package common

import proto "../proto/renewal"

func ConvertRenewalDealFromOrg(renewDealInfo RenewalDealInfo) *proto.RenewalDealInfo {
	renewalDealInfo := new(proto.RenewalDealInfo)

	renewalDealInfo.DealId = renewDealInfo.Deal_id
	renewalDealInfo.BdealId = renewDealInfo.Bdeal_id
	renewalDealInfo.DisdealId = renewDealInfo.Disdeal_id
	renewalDealInfo.InsuranceId = renewDealInfo.Insurance_id
	renewalDealInfo.GroupId = renewDealInfo.Group_id
	renewalDealInfo.ProductId = renewDealInfo.Product_id
	renewalDealInfo.ProductTitle = renewDealInfo.Product_title
	renewalDealInfo.Uid = renewDealInfo.Uid

	renewalDealInfo.InsuranceSn = renewDealInfo.Insurance_sn
	renewalDealInfo.InsuranceBenificiary = renewDealInfo.Insurance_benificiary
	renewalDealInfo.InsuranceBenificiaryIdentity = renewDealInfo.Insurance_benificiary_identity
	renewalDealInfo.InsuranceBenificiaryIdentityType = renewDealInfo.Insurance_benificiary_identity_type
	renewalDealInfo.InsuranceBenificiarySex = renewDealInfo.Insurance_benificiary_sex
	renewalDealInfo.InsuranceBenificiaryBirth = renewDealInfo.Insurance_benificiary_birth
	renewalDealInfo.Chn = renewDealInfo.Chn

	renewalDealInfo.RenewDealId = renewDealInfo.Renew_deal_id
	renewalDealInfo.RenewBdealId = renewDealInfo.Renew_bdeal_id
	renewalDealInfo.RenewDisdealId = renewDealInfo.Renew_disdeal_id
	renewalDealInfo.RenewUid = renewDealInfo.Renew_uid
	renewalDealInfo.RenewPayState = renewDealInfo.Renew_pay_state
	renewalDealInfo.HasBuy = renewDealInfo.Has_buy
	renewalDealInfo.BuyDealId = renewDealInfo.Buy_deal_id

	return renewalDealInfo
}

func ConvertRenewalDealFromProto(renewDealInfo *proto.RenewalDealInfo) *RenewalDealInfo {
	renewalDealInfo := new(RenewalDealInfo)

	renewalDealInfo.Deal_id = renewDealInfo.DealId
	renewalDealInfo.Bdeal_id = renewDealInfo.BdealId
	renewalDealInfo.Disdeal_id = renewDealInfo.DisdealId
	renewalDealInfo.Insurance_id = renewDealInfo.InsuranceId
	renewalDealInfo.Group_id = renewDealInfo.GroupId
	renewalDealInfo.Product_id = renewDealInfo.ProductId
	renewalDealInfo.Product_title = renewDealInfo.ProductTitle
	renewalDealInfo.Uid = renewDealInfo.Uid

	renewalDealInfo.Insurance_sn = renewDealInfo.InsuranceSn
	renewalDealInfo.Insurance_benificiary = renewDealInfo.InsuranceBenificiary
	renewalDealInfo.Insurance_benificiary_identity = renewDealInfo.InsuranceBenificiaryIdentity
	renewalDealInfo.Insurance_benificiary_identity_type = renewDealInfo.InsuranceBenificiaryIdentityType
	renewalDealInfo.Insurance_benificiary_sex = renewDealInfo.InsuranceBenificiarySex
	renewalDealInfo.Insurance_benificiary_birth = renewDealInfo.InsuranceBenificiaryBirth
	renewalDealInfo.Chn = renewDealInfo.Chn

	renewalDealInfo.Renew_deal_id = renewDealInfo.RenewDealId
	renewalDealInfo.Renew_bdeal_id = renewDealInfo.RenewBdealId
	renewalDealInfo.Renew_disdeal_id = renewDealInfo.RenewDisdealId
	renewalDealInfo.Renew_uid = renewDealInfo.RenewUid
	renewalDealInfo.Renew_pay_state = renewDealInfo.RenewPayState
	renewalDealInfo.Has_buy = renewDealInfo.HasBuy
	renewalDealInfo.Buy_deal_id = renewDealInfo.BuyDealId

	return renewalDealInfo
}
