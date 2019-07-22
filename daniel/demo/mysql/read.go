package main

import (
	"./common"
	"./logic"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/astaxie/beedb"
)

func main()  {
	db, err := logic.NewMysql()
	logic.CheckErr(err)
	defer db.Close()

	orm := beedb.New(db)

	var AllRenewalDealInfos []common.RenewalDealInfo
	err = orm.SetTable("t_renewal_v2_deal").Limit(1).OrderBy("Faddtime desc").FindAll(&AllRenewalDealInfos)
	logic.CheckErr(err)

	for _, renewalDealInfo := range AllRenewalDealInfos {
		fmt.Println(renewalDealInfo.Fdeal_id)
		fmt.Println(renewalDealInfo.Finsurance_benificiary)
	}

	dataMap, err := orm.SetTable("t_renewal_v2_deal").SetPK("Fdeal_id").Limit(1).OrderBy("Faddtime desc").FindMap()
	logic.CheckErr(err)

	for _, value := range dataMap {
		for key, value1 := range value {
			//fmt.Println(key)
			if key == "Fdeal_id" {
				var dealId int
				buf := bytes.NewBuffer(value1)
				binary.Read(buf, binary.LittleEndian, &dealId)
				fmt.Println(dealId)
			}
		}
	}
}
