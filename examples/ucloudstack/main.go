package main

import (
	"fmt"

	"github.com/ucloud/ucloud-sdk-go/services/ucloudstack"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

func loadUcloudStackConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.BaseUrl = "http://console.dev.ucloudstack.com/api"

	credential := auth.NewCredential()
	credential.PrivateKey = "gqwtkueCBF6fCPxstu5AvuFgZ-Eid92InCh7cIBkLiFHp7RsJsTIxTXUjb10krxb"
	credential.PublicKey = "1c64Ps3wN53H8MwjWub5euxIOGRIwKkVnm899FAw"

	return &cfg, &credential
}

func main() {

	// for i := 0; i < 60; i++ {
	// 	createVM("my-vm-" + strconv.Itoa(i+1))
	// }

	// stopVM("vm-uGvOT3TZg")
	// stopAllVM(10)

	// deleteVM("vm-lxK-oqTZg")
	// deleteAllVM(100)

	// terminateResource("dddd")
	terminateAllResource(100)

	// describeVM()

	// describeMetric("vm-hu74T3oZR")

}

func createVM(name string) {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// create request
	createReq := ucloudstackClient.NewCreateVMInstanceRequest()

	// 地域
	createReq.Region = ucloud.String("cn")
	createReq.Zone = ucloud.String("zone-01")
	// 配置
	createReq.ImageID = ucloud.String("cn-image-centos-74")
	createReq.CPU = ucloud.Int(1)
	createReq.Memory = ucloud.Int(2048)
	createReq.BootDiskType = ucloud.String("Normal")
	createReq.DataDiskType = ucloud.String("Normal")
	createReq.VMType = ucloud.String("Normal")
	// 网络
	createReq.VPCID = ucloud.String("vpc-Ci9vkUVpm")
	createReq.SubnetID = ucloud.String("subnet-Ci9vkUVpm")
	createReq.WANSGID = ucloud.String("sg-Ci9vkUVpm")
	// 认证方式
	createReq.Name = ucloud.String(name)
	createReq.Password = ucloud.String("ucloud.cn132")
	// 计费方式
	createReq.ChargeType = ucloud.String("Month")

	// send request
	newVMInstance, err := ucloudstackClient.CreateVMInstance(createReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if newVMInstance != nil {
		fmt.Printf("resource id of the VM: %s\n", newVMInstance.VMID)
	}

}

func stopVM(vmID string) {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// delete request
	stopReq := ucloudstackClient.NewStopVMInstanceRequest()
	stopReq.Region = ucloud.String("cn")
	stopReq.Zone = ucloud.String("zone-01")
	stopReq.VMID = ucloud.String(vmID)

	// send request
	stopResp, err := ucloudstackClient.StopVMInstance(stopReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if stopResp != nil {
		fmt.Printf("RetCode: %d\n", stopResp.RetCode)
	}

}

func deleteVM(vmID string) {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// delete request
	deleteReq := ucloudstackClient.NewDeleteVMInstanceRequest()
	deleteReq.Region = ucloud.String("cn")
	deleteReq.Zone = ucloud.String("zone-01")
	deleteReq.VMID = ucloud.String(vmID)

	// send request
	delResp, err := ucloudstackClient.DeleteVMInstance(deleteReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if delResp != nil {
		fmt.Printf("RetCode: %d\n", delResp.RetCode)
	}

}

func terminateResource(resourceID string) {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// delete request
	terminateReq := ucloudstackClient.NewTerminateResourceRequest()
	terminateReq.Region = ucloud.String("cn")
	terminateReq.Zone = ucloud.String("zone-01")
	terminateReq.ResourceID = ucloud.String(resourceID)

	// send request
	delResp, err := ucloudstackClient.TerminateResource(terminateReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if delResp != nil {
		fmt.Printf("RetCode: %d\n", delResp.RetCode)
	}

}

func describeVM() {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeVMInstanceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	describeReq.Limit = ucloud.Int(10)

	// send request
	descResp, err := ucloudstackClient.DescribeVMInstance(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if descResp.TotalCount > 0 {
		fmt.Printf("fisrt of VMs: %s\n", descResp.Infos[0].VMID)
	}

}

func stopAllVM(n int) {
	if n == 0 {
		return
	}

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeVMInstanceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	if n > 0 { // 负数表示all
		describeReq.Limit = ucloud.Int(n)
	}

	// send request
	descResp, err := ucloudstackClient.DescribeVMInstance(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if descResp.TotalCount > 0 {
		for _, info := range descResp.Infos {
			if info.State == "Running" {
				stopVM(info.VMID)
				fmt.Printf("stop vm: %s\n", info.VMID)
			}
		}
	}

}

func deleteAllVM(n int) {
	if n == 0 {
		return
	}

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeVMInstanceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	if n > 0 { // 负数表示all
		describeReq.Limit = ucloud.Int(n)
	}

	// send request
	descResp, err := ucloudstackClient.DescribeVMInstance(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if descResp.TotalCount > 0 {
		for _, info := range descResp.Infos {
			if info.State == "Stopped" {
				deleteVM(info.VMID)
				fmt.Printf("delete vm: %s\n", info.VMID)
			}
		}
	}

}

func terminateAllResource(n int) {
	if n == 0 {
		return
	}

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// describe Request
	describeReq := ucloudstackClient.NewDescribeRecycledResourceRequest()
	describeReq.Region = ucloud.String("cn")
	describeReq.Zone = ucloud.String("zone-01")
	if n > 0 { // 负数表示all
		describeReq.Limit = ucloud.Int(n)
	}

	// send request
	descResp, err := ucloudstackClient.DescribeRecycledResource(describeReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if descResp.TotalCount > 0 {
		for _, info := range descResp.Infos {
			terminateResource(info.ResourceID)
			fmt.Printf("terminate resource(status=%d): %s\n", info.Status, info.ResourceID)
		}
	}

}

func describeMetric(vmID string) {

	// 认证
	cfg, credential := loadUcloudStackConfig()
	ucloudstackClient := ucloudstack.NewClient(cfg, credential)

	// metric Request
	metricReq := ucloudstackClient.NewDescribeMetricRequest()
	metricReq.Region = ucloud.String("cn")
	metricReq.Zone = ucloud.String("zone-01")
	metricReq.ResourceID = ucloud.String(vmID)
	metricReq.MetricName = []string{"CPUUtilization"}
	metricReq.BeginTime = ucloud.String("1571819416")
	metricReq.EndTime = ucloud.String("1571823016")

	// send request
	metricResp, err := ucloudstackClient.DescribeMetric(metricReq)
	if err != nil {
		fmt.Printf("something bad happened: %s\n", err)
	}

	if metricResp.TotalCount > 0 && len(metricResp.Infos[0].Infos) > 0 {
		fmt.Printf("value of %s at %d: %f\n",
			metricResp.Infos[0].MetricName,
			metricResp.Infos[0].Infos[0].Timestamp,
			metricResp.Infos[0].Infos[0].Value)
	}

}
