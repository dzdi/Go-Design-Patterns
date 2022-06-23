package adapter09

import "testing"

func TestAliyunClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AliyunClientAdapter{
		AliyunClient{},
	}

	a.CreateServer(1.0, 2.0)
}
func TestAwsClientAdapter_CreateServer(t *testing.T) {
	var a ICreateServer = &AwsClientAdapter{
		Client: AWSClient{},
	}
	a.CreateServer(1.0, 2.0)
}
