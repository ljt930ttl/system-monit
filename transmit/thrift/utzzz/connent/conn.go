package connent

import (
	"runtime/debug"
	"sync"
	"system-monit/transmit/thrift/utzzz"

	"github.com/apache/thrift/lib/go/thrift"
	"gorm.io/gorm/logger"
)

type ZzzPool struct {
	// 注册了的节点
	pool map[*PMANode]bool
	//	pool *HashTable
	// 在线用户  domain+name
	poolNode map[string]map[*PMANode]bool
	//	poolUser *HashTable
	// 从连接器中注册请求
	register chan *PMANode
	// 从连接器中注销请求
	unregister chan *PMANode
	// 锁
	rwLock *sync.RWMutex
}

func (this *ZzzPool) AddAuthUser(pNode *PMANode) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("AddAuthUser,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	if _, ok := this.pool[pNode]; ok {
		// name, _ := GetLoginName(pNode.UserTid)
		if len(this.poolNode[name]) == 0 {
			this.poolNode[name] = make(map[*PMANode]bool)
		}
		this.poolNode[name][pNode] = true
	}
}

type PMANode struct {
	Ts   thrift.TTransport
	Msg  *utzzz.PMAMsg
	Name string
	Id   string
	// Register	bool
}
