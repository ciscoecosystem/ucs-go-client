package models

import (

	"fmt"


	"github.com/beevik/etree"

)

const LsServerClassName = "lsServer"

type LsServer struct {
    BaseAttributes
    LsServerAttributes
}

type LsServerAttributes struct{
    Agent_policy_name string `xml:",omitempty"`
    Assign_state string `xml:",omitempty"`
    Assoc_state string `xml:",omitempty"`
    Bios_profile_name string `xml:",omitempty"`
    Boot_policy_name string `xml:",omitempty"`
    Child_action string `xml:",omitempty"`
    Config_qualifier string `xml:",omitempty"`
    Config_state string `xml:",omitempty"`
    Dynamic_con_policy_name string `xml:",omitempty"`
    Ext_ip_pool_name string `xml:",omitempty"`
    Ext_ip_state string `xml:",omitempty"`
    Flt_aggr string `xml:",omitempty"`
    Fsm_descr string `xml:",omitempty"`
    Fsm_flags string `xml:",omitempty"`
    Fsm_prev string `xml:",omitempty"`
    Fsm_progr string `xml:",omitempty"`
    Fsm_rmt_inv_err_code string `xml:",omitempty"`
    Fsm_rmt_inv_err_descr string `xml:",omitempty"`
    Fsm_rmt_inv_rslt string `xml:",omitempty"`
    Fsm_stage_descr string `xml:",omitempty"`
    Fsm_stamp string `xml:",omitempty"`
    Fsm_status string `xml:",omitempty"`
    Fsm_try string `xml:",omitempty"`
    Graphics_card_policy_name string `xml:",omitempty"`
    Host_fw_policy_name string `xml:",omitempty"`
    Ident_pool_name string `xml:",omitempty"`
    Int_id string `xml:",omitempty"`
    Kvm_mgmt_policy_name string `xml:",omitempty"`
    Local_disk_policy_name string `xml:",omitempty"`
    Maint_policy_name string `xml:",omitempty"`
    Mgmt_access_policy_name string `xml:",omitempty"`
    Mgmt_fw_policy_name string `xml:",omitempty"`
    Oper_bios_profile_name string `xml:",omitempty"`
    Oper_boot_policy_name string `xml:",omitempty"`
    Oper_dynamic_con_policy_name string `xml:",omitempty"`
    Oper_ext_ip_pool_name string `xml:",omitempty"`
    Oper_graphics_card_policy_name string `xml:",omitempty"`
    Oper_host_fw_policy_name string `xml:",omitempty"`
    Oper_ident_pool_name string `xml:",omitempty"`
    Oper_kvm_mgmt_policy_name string `xml:",omitempty"`
    Oper_local_disk_policy_name string `xml:",omitempty"`
    Oper_maint_policy_name string `xml:",omitempty"`
    Oper_mgmt_access_policy_name string `xml:",omitempty"`
    Oper_mgmt_fw_policy_name string `xml:",omitempty"`
    Oper_power_policy_name string `xml:",omitempty"`
    Oper_power_sync_policy_name string `xml:",omitempty"`
    Oper_scrub_policy_name string `xml:",omitempty"`
    Oper_sol_policy_name string `xml:",omitempty"`
    Oper_src_templ_name string `xml:",omitempty"`
    Oper_state string `xml:",omitempty"`
    Oper_stats_policy_name string `xml:",omitempty"`
    Oper_vcon_profile_name string `xml:",omitempty"`
    Oper_vmedia_policy_name string `xml:",omitempty"`
    Owner string `xml:",omitempty"`
    Pn_dn string `xml:",omitempty"`
    Policy_level string `xml:",omitempty"`
    Policy_owner string `xml:",omitempty"`
    Power_policy_name string `xml:",omitempty"`
    Power_sync_policy_name string `xml:",omitempty"`
    Prop_acl string `xml:",omitempty"`
    Resolve_remote string `xml:",omitempty"`
    Sacl string `xml:",omitempty"`
    Scrub_policy_name string `xml:",omitempty"`
    Sol_policy_name string `xml:",omitempty"`
    Src_templ_name string `xml:",omitempty"`
    Stats_policy_name string `xml:",omitempty"`
    Svnic_config string `xml:",omitempty"`
    Type string `xml:",omitempty"`
    Usr_lbl string `xml:",omitempty"`
    Uuid string `xml:",omitempty"`
    Uuid_suffix string `xml:",omitempty"`
    Vcon_profile_name string `xml:",omitempty"`
    Vmedia_policy_name string `xml:",omitempty"`
    

}

func NewLsServer(lsServerRn,parentDn,description string, lsServerAttributes LsServerAttributes) *LsServer { 
    dn := fmt.Sprintf("%s/%s", parentDn, lsServerRn)
	return &LsServer{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			ClassName:         LsServerClassName,
			Status:            "created, modified",
			Description:       description,
		},

		LsServerAttributes: lsServerAttributes,
	}
}

func (lsServer *LsServer) ToMap() (map[string]string, error) {
	lsServerMap, err := lsServer.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}
    lsServerMap["agentPolicyName"] = lsServer.Agent_policy_name
    lsServerMap["assignState"] = lsServer.Assign_state
    lsServerMap["assocState"] = lsServer.Assoc_state
    lsServerMap["biosProfileName"] = lsServer.Bios_profile_name
    lsServerMap["bootPolicyName"] = lsServer.Boot_policy_name
    lsServerMap["childAction"] = lsServer.Child_action
    lsServerMap["configQualifier"] = lsServer.Config_qualifier
    lsServerMap["configState"] = lsServer.Config_state
    lsServerMap["dynamicConPolicyName"] = lsServer.Dynamic_con_policy_name
    lsServerMap["extIPPoolName"] = lsServer.Ext_ip_pool_name
    lsServerMap["extIPState"] = lsServer.Ext_ip_state
    lsServerMap["fltAggr"] = lsServer.Flt_aggr
    lsServerMap["fsmDescr"] = lsServer.Fsm_descr
    lsServerMap["fsmFlags"] = lsServer.Fsm_flags
    lsServerMap["fsmPrev"] = lsServer.Fsm_prev
    lsServerMap["fsmProgr"] = lsServer.Fsm_progr
    lsServerMap["fsmRmtInvErrCode"] = lsServer.Fsm_rmt_inv_err_code
    lsServerMap["fsmRmtInvErrDescr"] = lsServer.Fsm_rmt_inv_err_descr
    lsServerMap["fsmRmtInvRslt"] = lsServer.Fsm_rmt_inv_rslt
    lsServerMap["fsmStageDescr"] = lsServer.Fsm_stage_descr
    lsServerMap["fsmStamp"] = lsServer.Fsm_stamp
    lsServerMap["fsmStatus"] = lsServer.Fsm_status
    lsServerMap["fsmTry"] = lsServer.Fsm_try
    lsServerMap["graphicsCardPolicyName"] = lsServer.Graphics_card_policy_name
    lsServerMap["hostFwPolicyName"] = lsServer.Host_fw_policy_name
    lsServerMap["identPoolName"] = lsServer.Ident_pool_name
    lsServerMap["intId"] = lsServer.Int_id
    lsServerMap["kvmMgmtPolicyName"] = lsServer.Kvm_mgmt_policy_name
    lsServerMap["localDiskPolicyName"] = lsServer.Local_disk_policy_name
    lsServerMap["maintPolicyName"] = lsServer.Maint_policy_name
    lsServerMap["mgmtAccessPolicyName"] = lsServer.Mgmt_access_policy_name
    lsServerMap["mgmtFwPolicyName"] = lsServer.Mgmt_fw_policy_name
    lsServerMap["operBiosProfileName"] = lsServer.Oper_bios_profile_name
    lsServerMap["operBootPolicyName"] = lsServer.Oper_boot_policy_name
    lsServerMap["operDynamicConPolicyName"] = lsServer.Oper_dynamic_con_policy_name
    lsServerMap["operExtIPPoolName"] = lsServer.Oper_ext_ip_pool_name
    lsServerMap["operGraphicsCardPolicyName"] = lsServer.Oper_graphics_card_policy_name
    lsServerMap["operHostFwPolicyName"] = lsServer.Oper_host_fw_policy_name
    lsServerMap["operIdentPoolName"] = lsServer.Oper_ident_pool_name
    lsServerMap["operKvmMgmtPolicyName"] = lsServer.Oper_kvm_mgmt_policy_name
    lsServerMap["operLocalDiskPolicyName"] = lsServer.Oper_local_disk_policy_name
    lsServerMap["operMaintPolicyName"] = lsServer.Oper_maint_policy_name
    lsServerMap["operMgmtAccessPolicyName"] = lsServer.Oper_mgmt_access_policy_name
    lsServerMap["operMgmtFwPolicyName"] = lsServer.Oper_mgmt_fw_policy_name
    lsServerMap["operPowerPolicyName"] = lsServer.Oper_power_policy_name
    lsServerMap["operPowerSyncPolicyName"] = lsServer.Oper_power_sync_policy_name
    lsServerMap["operScrubPolicyName"] = lsServer.Oper_scrub_policy_name
    lsServerMap["operSolPolicyName"] = lsServer.Oper_sol_policy_name
    lsServerMap["operSrcTemplName"] = lsServer.Oper_src_templ_name
    lsServerMap["operState"] = lsServer.Oper_state
    lsServerMap["operStatsPolicyName"] = lsServer.Oper_stats_policy_name
    lsServerMap["operVconProfileName"] = lsServer.Oper_vcon_profile_name
    lsServerMap["operVmediaPolicyName"] = lsServer.Oper_vmedia_policy_name
    lsServerMap["owner"] = lsServer.Owner
    lsServerMap["pnDn"] = lsServer.Pn_dn
    lsServerMap["policyLevel"] = lsServer.Policy_level
    lsServerMap["policyOwner"] = lsServer.Policy_owner
    lsServerMap["powerPolicyName"] = lsServer.Power_policy_name
    lsServerMap["powerSyncPolicyName"] = lsServer.Power_sync_policy_name
    lsServerMap["propAcl"] = lsServer.Prop_acl
    lsServerMap["resolveRemote"] = lsServer.Resolve_remote
    lsServerMap["sacl"] = lsServer.Sacl
    lsServerMap["scrubPolicyName"] = lsServer.Scrub_policy_name
    lsServerMap["solPolicyName"] = lsServer.Sol_policy_name
    lsServerMap["srcTemplName"] = lsServer.Src_templ_name
    lsServerMap["statsPolicyName"] = lsServer.Stats_policy_name
    lsServerMap["svnicConfig"] = lsServer.Svnic_config
    lsServerMap["type"] = lsServer.Type
    lsServerMap["usrLbl"] = lsServer.Usr_lbl
    lsServerMap["uuid"] = lsServer.Uuid
    lsServerMap["uuidSuffix"] = lsServer.Uuid_suffix
    lsServerMap["vconProfileName"] = lsServer.Vcon_profile_name
    lsServerMap["vmediaPolicyName"] = lsServer.Vmedia_policy_name
    
	return lsServerMap, nil

}

func LsServerFromDoc(doc *etree.Document, rootClass string) *LsServer {
	element, err := GetMoElement(doc, rootClass, LsServerClassName)
	if err != nil {
		return nil
	}
	return &LsServer{
		BaseAttributes: BaseAttributes{
			DistinguishedName: element.SelectAttrValue("dn", ""),
			ClassName:         LsServerClassName,
			Description:       element.SelectAttrValue("descr", ""),
			Status:            element.SelectAttrValue("status", ""),
		},
    

		LsServerAttributes: LsServerAttributes{

        
			Agent_policy_name:  element.SelectAttrValue("agentPolicyName", ""),
			Assign_state:  element.SelectAttrValue("assignState", ""),
			Assoc_state:  element.SelectAttrValue("assocState", ""),
			Bios_profile_name:  element.SelectAttrValue("biosProfileName", ""),
			Boot_policy_name:  element.SelectAttrValue("bootPolicyName", ""),
			Child_action:  element.SelectAttrValue("childAction", ""),
			Config_qualifier:  element.SelectAttrValue("configQualifier", ""),
			Config_state:  element.SelectAttrValue("configState", ""),
			Dynamic_con_policy_name:  element.SelectAttrValue("dynamicConPolicyName", ""),
			Ext_ip_pool_name:  element.SelectAttrValue("extIPPoolName", ""),
			Ext_ip_state:  element.SelectAttrValue("extIPState", ""),
			Flt_aggr:  element.SelectAttrValue("fltAggr", ""),
			Fsm_descr:  element.SelectAttrValue("fsmDescr", ""),
			Fsm_flags:  element.SelectAttrValue("fsmFlags", ""),
			Fsm_prev:  element.SelectAttrValue("fsmPrev", ""),
			Fsm_progr:  element.SelectAttrValue("fsmProgr", ""),
			Fsm_rmt_inv_err_code:  element.SelectAttrValue("fsmRmtInvErrCode", ""),
			Fsm_rmt_inv_err_descr:  element.SelectAttrValue("fsmRmtInvErrDescr", ""),
			Fsm_rmt_inv_rslt:  element.SelectAttrValue("fsmRmtInvRslt", ""),
			Fsm_stage_descr:  element.SelectAttrValue("fsmStageDescr", ""),
			Fsm_stamp:  element.SelectAttrValue("fsmStamp", ""),
			Fsm_status:  element.SelectAttrValue("fsmStatus", ""),
			Fsm_try:  element.SelectAttrValue("fsmTry", ""),
			Graphics_card_policy_name:  element.SelectAttrValue("graphicsCardPolicyName", ""),
			Host_fw_policy_name:  element.SelectAttrValue("hostFwPolicyName", ""),
			Ident_pool_name:  element.SelectAttrValue("identPoolName", ""),
			Int_id:  element.SelectAttrValue("intId", ""),
			Kvm_mgmt_policy_name:  element.SelectAttrValue("kvmMgmtPolicyName", ""),
			Local_disk_policy_name:  element.SelectAttrValue("localDiskPolicyName", ""),
			Maint_policy_name:  element.SelectAttrValue("maintPolicyName", ""),
			Mgmt_access_policy_name:  element.SelectAttrValue("mgmtAccessPolicyName", ""),
			Mgmt_fw_policy_name:  element.SelectAttrValue("mgmtFwPolicyName", ""),
			Oper_bios_profile_name:  element.SelectAttrValue("operBiosProfileName", ""),
			Oper_boot_policy_name:  element.SelectAttrValue("operBootPolicyName", ""),
			Oper_dynamic_con_policy_name:  element.SelectAttrValue("operDynamicConPolicyName", ""),
			Oper_ext_ip_pool_name:  element.SelectAttrValue("operExtIPPoolName", ""),
			Oper_graphics_card_policy_name:  element.SelectAttrValue("operGraphicsCardPolicyName", ""),
			Oper_host_fw_policy_name:  element.SelectAttrValue("operHostFwPolicyName", ""),
			Oper_ident_pool_name:  element.SelectAttrValue("operIdentPoolName", ""),
			Oper_kvm_mgmt_policy_name:  element.SelectAttrValue("operKvmMgmtPolicyName", ""),
			Oper_local_disk_policy_name:  element.SelectAttrValue("operLocalDiskPolicyName", ""),
			Oper_maint_policy_name:  element.SelectAttrValue("operMaintPolicyName", ""),
			Oper_mgmt_access_policy_name:  element.SelectAttrValue("operMgmtAccessPolicyName", ""),
			Oper_mgmt_fw_policy_name:  element.SelectAttrValue("operMgmtFwPolicyName", ""),
			Oper_power_policy_name:  element.SelectAttrValue("operPowerPolicyName", ""),
			Oper_power_sync_policy_name:  element.SelectAttrValue("operPowerSyncPolicyName", ""),
			Oper_scrub_policy_name:  element.SelectAttrValue("operScrubPolicyName", ""),
			Oper_sol_policy_name:  element.SelectAttrValue("operSolPolicyName", ""),
			Oper_src_templ_name:  element.SelectAttrValue("operSrcTemplName", ""),
			Oper_state:  element.SelectAttrValue("operState", ""),
			Oper_stats_policy_name:  element.SelectAttrValue("operStatsPolicyName", ""),
			Oper_vcon_profile_name:  element.SelectAttrValue("operVconProfileName", ""),
			Oper_vmedia_policy_name:  element.SelectAttrValue("operVmediaPolicyName", ""),
			Owner:  element.SelectAttrValue("owner", ""),
			Pn_dn:  element.SelectAttrValue("pnDn", ""),
			Policy_level:  element.SelectAttrValue("policyLevel", ""),
			Policy_owner:  element.SelectAttrValue("policyOwner", ""),
			Power_policy_name:  element.SelectAttrValue("powerPolicyName", ""),
			Power_sync_policy_name:  element.SelectAttrValue("powerSyncPolicyName", ""),
			Prop_acl:  element.SelectAttrValue("propAcl", ""),
			Resolve_remote:  element.SelectAttrValue("resolveRemote", ""),
			Sacl:  element.SelectAttrValue("sacl", ""),
			Scrub_policy_name:  element.SelectAttrValue("scrubPolicyName", ""),
			Sol_policy_name:  element.SelectAttrValue("solPolicyName", ""),
			Src_templ_name:  element.SelectAttrValue("srcTemplName", ""),
			Stats_policy_name:  element.SelectAttrValue("statsPolicyName", ""),
			Svnic_config:  element.SelectAttrValue("svnicConfig", ""),
			Type:  element.SelectAttrValue("type", ""),
			Usr_lbl:  element.SelectAttrValue("usrLbl", ""),
			Uuid:  element.SelectAttrValue("uuid", ""),
			Uuid_suffix:  element.SelectAttrValue("uuidSuffix", ""),
			Vcon_profile_name:  element.SelectAttrValue("vconProfileName", ""),
			Vmedia_policy_name:  element.SelectAttrValue("vmediaPolicyName", ""),
		},
    
	}
}