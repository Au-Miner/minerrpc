package utils

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	var err error
	conn, _, err = zk.Connect([]string{"localhost:2181"}, time.Second)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Zookeeper: %v", err))
	}

	children, _, err := conn.Children("/MinerRPC/Hello")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("children: ", children)
}

func TestB(t *testing.T) {

	a, b := GetAllInstances("hello")
	if b != nil {
		t.Error(b)
	}
	fmt.Println("a: ", a)
}

func TestC(t *testing.T) {
	var err error
	conn, _, err := zk.Connect([]string{"localhost:2181"}, time.Second)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Zookeeper: %v", err))
	}
	defer conn.Close()

	// 检查 /MinerRPC 节点是否存在
	exists, _, err := conn.Exists("/MinerRPC")
	if err != nil {
		t.Fatalf("Error checking for /MinerRPC existence: %v", err)
	}
	if !exists {
		// 如果不存在，创建 /MinerRPC
		_, err = conn.Create("/MinerRPC", []byte{}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			t.Fatalf("Failed to create /MinerRPC: %v", err)
		}
	}

	// 检查 /MinerRPC/hello 节点是否存在
	exists, _, err = conn.Exists("/MinerRPC/hello")
	if err != nil {
		t.Fatalf("Error checking for /MinerRPC/hello existence: %v", err)
	}
	if !exists {
		// 如果不存在，创建 /MinerRPC/hello
		_, err = conn.Create("/MinerRPC/hello", []byte{}, 0, zk.WorldACL(zk.PermAll))
		if err != nil {
			t.Fatalf("Failed to create /MinerRPC/hello: %v", err)
		}
	}

	// 现在可以安全地获取子节点了
	children, _, err := conn.Children("/MinerRPC/hello")
	if err != nil {
		t.Fatalf("Error getting children of /MinerRPC/hello: %v", err)
	}
	fmt.Println("children: ", children)
}
