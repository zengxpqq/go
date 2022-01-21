package tool

import (
	"fmt"
	"sync"
	"time"
)

// EpochTimestamp 初始时间
const EpochTimestamp int64 = 550281600000

// Generator 生成器
type Generator struct {
	dc               int64 // 手动分配， 不同部署环境的 分配不同的值
	worker           int64 // 多节点部署的情况下， 需要不同的值
	nodeID           int64 // dc 与 worker 生成唯一值
	lastTimestamp    int64 // 最后的时间
	sequence         int64 // 相同时间生成的个数
	sequenceOverload int64 // 相同时间生成的个数查过 4095 后增加1
	errors           int64 // 发生错误的个数
	generatedIds     int64 // 生成唯一标识的总个数
	sync.Mutex
}

// InitGenerator 初始化
func InitGenerator(dc, worker int64) Generator {
	nodeID := ((dc & 0xf) << 6) | (worker & 0x3f)
	return Generator{
		dc:               dc,
		worker:           worker,
		nodeID:           nodeID,
		lastTimestamp:    EpochTimestamp,
		sequence:         0,
		sequenceOverload: 0,
		errors:           0,
		generatedIds:     0,
	}
}

// GetNextID 获取唯一标识
func (g *Generator) GetNextID() (generatedID int64, err error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	currTime := time.Now().UnixNano() / int64(time.Millisecond)
	if currTime < g.lastTimestamp {
		g.errors++
		err = fmt.Errorf("clock went backwards! %d < %d", currTime, g.lastTimestamp)
		return
	}

	if currTime > g.lastTimestamp {
		g.sequence = 0
		g.lastTimestamp = currTime
	}
	g.sequence++
	if g.sequence > 4095 {
		g.sequenceOverload++
		time.Sleep(time.Millisecond)
		g.sequence = 0
		currTime = time.Now().UnixNano() / int64(time.Millisecond)
		g.lastTimestamp = currTime
	}
	generatedID = ((currTime - EpochTimestamp) << 22) | (g.nodeID << 12) | g.sequence
	g.generatedIds++
	return
}
