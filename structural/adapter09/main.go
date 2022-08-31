package adapter09

import "fmt"

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct {
}

func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, men: %f", cpu, mem)
	return nil
}

type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

type AliyunClient struct {
}

func (a *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
