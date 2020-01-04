package qcloud

import (
	"GoDDNS/util"

	cns "github.com/athurg/go-qcloud-cns-sdk"
)

type QcloudCommon struct {
	Record struct {
		Name  string
		Type  string
		Value string
	}
	Create struct {
		Id         int
		Ttl        int
		Value      string
		Enabled    int
		Status     string
		UpdatedOn  string `json:"updated_on"`
		QProjectId int    `json:"q_project_id"`
		Name       string
		Line       string
		LineId     string `json:"line_id"`
		Type       string
		Remark     string
		Mx         int
		Hold       string
	}
}

func common() (cli *cns.Client) {
	setting := util.Setting()
	cli = cns.New(setting.Qcloud.SecretId, setting.Qcloud.SecretKey)
	return cli
}

func (QcloudCommon *QcloudCommon) DDNS() {
	setting := util.Setting()
	for _, Item := range setting.Qcloud.List {
		RecordList, err := common().RecordList(Item.Domain)
		util.CheckErr(err)
		for _, Record := range RecordList {
			if Record.Name == Item.SubDomain {
				var Value string
				if Item.Type == "A" {
					Value = util.IPv4()
				} else if Item.Type == "AAAA" {
					Value = util.IPv6()
				}
				if Record.Value == Value {
					util.CheckErrCustom("IP一致，无需变更。")
				} else if Record.Type == Item.Type && Record.Value != Value {
					QcloudCommon.RecordModify(Record.Id, Item.Domain, Item.SubDomain, Item.Type)
				}
			}
		}
	}
	return
}

func (QcloudCommon *QcloudCommon) RecordCreate(Domain, SubDomain, Type string) {
	var Value string
	if Type == "A" {
		Value = util.IPv4()
	} else if Type == "AAAA" {
		Value = util.IPv6()
	}
	QcloudCommon.Create.Name = SubDomain
	QcloudCommon.Create.Type = Type
	QcloudCommon.Create.Line = "默认"
	QcloudCommon.Create.Value = Value
	RecordID, err := common().RecordCreate(Domain, QcloudCommon.Create)
	util.CheckErr(err)
	if RecordID != 0 {
		util.CheckErrCustom("\n当前IP：" + Value + "\n记录IP：" + Value)
	}
}

func (QcloudCommon *QcloudCommon) RecordModify(ID int, Domain, SubDomain, Type string) {
	var Value string
	if Type == "A" {
		Value = util.IPv4()
	} else if Type == "AAAA" {
		Value = util.IPv6()
	}
	QcloudCommon.Create.Id = ID
	QcloudCommon.Create.Name = SubDomain
	QcloudCommon.Create.Type = Type
	QcloudCommon.Create.Line = "默认"
	QcloudCommon.Create.Value = Value
	err := common().RecordModify(Domain, QcloudCommon.Create)
	util.CheckErr(err)
	if err != nil {
		util.CheckErrCustom("\n当前IP：" + Value + "记录IP：" + Value)
	}
}
