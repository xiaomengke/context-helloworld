package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type meshInfo struct {
	Did        string     `json:"did"`
	GatewayId  string     `json:"gateway_id"`
	Rssi       float64    `json:"rssi"`
	Hops       int        `json:"hops"`
	CreateTime int        `json:"createtime"`
	BackList   []meshInfo `json:"back_list"`
	Parent     string     `json:"parent"`
}

var (
	BleCacheFile          = "ble_cache_redis_cluster.json"
	BleDeviceOnlinePrefix = "device_cache/ble_device_online/%s"
	BleRelationPrefix     = "bleRelation/%s"
	BleDeviceParentPrefix = "device_cache/ble_device_parent/%s"
	MeshPrefix            = "device_cache/blemesh_gateway_info/%s"
)

func main34() {
	re, err := Command([]string{"gdh", "mesh"}, 2, map[string]string{"did": "123456"})
	if err != nil {
		println(err)
	}
	fmt.Println(re)
}

func Command(commandArg []string, currentIndex int, flags map[string]string) (string, error) {
	// check did
	did := flags["did"]
	if did == "" {
		return "did为必传参数,eg:getdeviceheartbeat [ble|mesh] -did=xxx", nil
	}
	// check and request
	isBLEcmd := strings.EqualFold(commandArg[1], "ble")
	//conditional judge
	var ret string
	var err error
	if isBLEcmd {
		ret, err = execBle(did)
	} else {
		ret, err = execMesh(did)
	}
	if err != nil {
		return ret, err
	}
	return ret, nil
}
func execBle(did string) (string, error) {
	//获取1,子设备时间戳,信号强度 2,此子设备网关,信号强度时间戳 3,此子设备直连网关did
	bleDeviceOnlineKey := fmt.Sprintf(BleDeviceOnlinePrefix, did)
	bleRelationKey := fmt.Sprintf(BleRelationPrefix, did)
	bleDeviceParentKey := fmt.Sprintf(BleDeviceParentPrefix, did)
	sliceBleKey := []string{bleDeviceOnlineKey, bleRelationKey, bleDeviceParentKey}
	sliceBleValue := []interface{}{"1721807924|1034523337|-76", `{"463609756":"-63|1724828096"}`, "463609756"}
	mapBleKeyValue := make(map[string]string, 3)
	for i, v := range sliceBleValue {
		if v != nil {
			mapBleKeyValue[sliceBleKey[i]] = v.(string)
		}
	}
	resp := fmt.Sprintf("蓝牙设备did:%s\n", did)
	//解析在线信息
	if v, ok := mapBleKeyValue[bleDeviceOnlineKey]; ok {
		valueSplitArr := strings.Split(v, "|")
		if len(valueSplitArr) != 3 {
			return fmt.Sprintf("根据key(%v)查询的value(%v)格式错误", v, valueSplitArr), nil
		}
		lastHeartBeatTimestamp, err := strconv.ParseInt(valueSplitArr[0], 10, 64)
		if err != nil {
			return fmt.Sprintf("Error converting string to int64:%v", valueSplitArr[0]), err
		}
		lastHeartBeatDate := time.Unix(lastHeartBeatTimestamp, 0)
		timeDiff2Now := getTimeDiffToNow(lastHeartBeatDate)
		gatewayId := valueSplitArr[1]
		resp += fmt.Sprintf("主网关did:%v\n", gatewayId)
		resp += fmt.Sprintf("最后一次心跳:%v\n", lastHeartBeatDate)
		resp += fmt.Sprintf("--在%s之前\n", timeDiff2Now)
		signalStrength, err := strconv.ParseFloat(valueSplitArr[2], 64)
		resp += fmt.Sprintf("信号强度为%v\n", signalStrength)
		if err != nil {
			return fmt.Sprintf("信号强度类型转换错误:(%v)", valueSplitArr[2]), err
		}
	}
	//解析蓝牙子设备的网关信息
	var relationDeviceMap map[string]string
	resp += fmt.Sprintf("此蓝牙设备网关信息列表:\n")
	if v, ok := mapBleKeyValue[bleRelationKey]; ok {
		if len(v) > 0 {
			err := json.Unmarshal([]byte(v), &relationDeviceMap)
			if err != nil {
				return fmt.Sprintf("Json.Unmarshal错误:%v", v), err
			}
			for k, v := range relationDeviceMap {
				vSplitArr := strings.Split(v, "|")
				if len(vSplitArr) < 2 {
					continue
				}
				heartBeat, err := strconv.ParseInt(vSplitArr[1], 10, 64)
				if err != nil {
					return fmt.Sprintf("Error converting string to int64:%v", vSplitArr[1]), err
				}
				heartBeatDate := time.Unix(heartBeat, 0)
				s := fmt.Sprintf("信号强度%v,上次心跳时间%v\n", vSplitArr[0], heartBeatDate)
				relationDeviceMap[k] = s
			}
		} else {
			resp += fmt.Sprintf("--为空\n")
		}
	}
	for k, v := range relationDeviceMap {
		resp += fmt.Sprintf("--%s:,%s", k, v)
	}
	if v, ok := mapBleKeyValue[bleDeviceParentKey]; ok {
		resp += fmt.Sprintf("直连网关为%v", v)
	} else {
		resp += fmt.Sprintf("直连网关为：空\n")
	}
	return resp, nil
}

func execMesh(did string) (string, error) {
	//获取mesh设备的网关信息
	meshKey := fmt.Sprintf(MeshPrefix, did)
	//meshValue, err := r.Get(meshKey)
	meshValue := `{"did":"1116841508","gateway_id":"1066480572","rssi":-81,"hops":0,"createtime":1721808117,"back_list":[{"did":"1116841508","gateway_id":"1077142098","rssi":-81,"hops":0,"createtime":1721808117,"json":0}],"parent":"1077142098"}`
	if len(meshValue) <= 0 {
		return fmt.Sprintf("redis key(%v) get 空字符串", meshKey), nil
	}
	//解析mesh设备的网关信息
	var meshJsonValue meshInfo
	err := json.Unmarshal([]byte(meshValue), &meshJsonValue)
	if err != nil {
		return fmt.Sprintf("根据key查询的value(%v)格式错误,Json.Unmarshal失败", meshValue), err
	}
	//calculate time diff
	lastHeartBeatDate := time.Unix(int64(meshJsonValue.CreateTime), 0) //最新心跳时间
	timeDiff2Now := getTimeDiffToNow(lastHeartBeatDate)
	// tidy resp and return
	resp := fmt.Sprintf("Mesh设备did: %s\n", meshJsonValue.Did)
	resp += fmt.Sprintf("主网关did:%v\n", meshJsonValue.GatewayId)
	resp += fmt.Sprintf("最后一次心跳:%v\n", lastHeartBeatDate)
	resp += fmt.Sprintf("--在%s之前\n", timeDiff2Now)
	resp += fmt.Sprintf("信号强度:%v\n", meshJsonValue.Rssi)
	resp += fmt.Sprintf("hops数:%v\n", meshJsonValue.Hops)
	resp += fmt.Sprintf("此设备back_list:\n")
	if len(meshJsonValue.BackList) == 0 {
		resp += fmt.Sprintf("为空\n")
	} else {
		for _, v := range meshJsonValue.BackList {
			resp += fmt.Sprintf("--did:%v,网关id:%v,信号强度:%v,hops:%v,创建时间:%v\n", v.Did, v.GatewayId, v.Rssi, v.Hops, time.Unix(int64(v.CreateTime), 0))
		}
	}
	if len(meshJsonValue.Parent) <= 0 {
		resp += fmt.Sprintf("直连网关为：空\n")
	} else {
		resp += fmt.Sprintf("直连网关为:%v\n", meshJsonValue.Parent)
	}
	return resp, nil

}

func getTimeDiffToNow(lastHeartBeatDate time.Time) string {
	diff := time.Now().Sub(lastHeartBeatDate)
	days := diff.Truncate(24*time.Hour).Hours() / 24
	hours := diff.Truncate(time.Hour).Minutes()/60 - days*24
	minutes := diff.Truncate(time.Minute).Seconds()/60 - hours*60 - days*24*60
	seconds := diff.Seconds() - minutes*60 - hours*3600 - days*86400
	return fmt.Sprintf("%v天%v时%v分%v秒", days, hours, minutes, int(seconds))
}
