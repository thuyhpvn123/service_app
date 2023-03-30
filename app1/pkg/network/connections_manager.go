package network

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type IConnectionsManager interface {
	AddConnection(IConnection)
	RemoveConnection(IConnection)
	GetConnectionsByAddress(common.Address) []IConnection
	GetConnectionByAddress(common.Address) IConnection
	GetConnectionsByTypeAtIdx(cType string, idx int) []IConnection
	GetConnectionsByType(cType string) map[common.Address][]IConnection
	GetTotalAddressByType(cType string) int
	AddParentConnection(IConnection)
	GetParentConnection() IConnection
}

type ConnectionsManager struct {
	mu                    sync.RWMutex
	connections           []IConnection
	parentConnection      IConnection
	mapAddressConnections map[common.Address][]IConnection
	mapTypeAddress        map[string]map[common.Address]struct{}
}

func NewConnectionsManager(
	connectionTypes []string,
) IConnectionsManager {
	cm := &ConnectionsManager{}
	cm.mapTypeAddress = make(map[string]map[common.Address]struct{})
	cm.mapAddressConnections = make(map[common.Address][]IConnection)
	for _, v := range connectionTypes {
		cm.mapTypeAddress[v] = make(map[common.Address]struct{})
	}
	return cm
}

func (cm *ConnectionsManager) AddParentConnection(conn IConnection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections = append(cm.connections, conn)
	cm.mapAddressConnections[conn.GetAddress()] = append(cm.mapAddressConnections[conn.GetAddress()], conn)
	cm.parentConnection = conn
}

func (cm *ConnectionsManager) GetParentConnection() IConnection {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	return cm.parentConnection
}

func (cm *ConnectionsManager) AddConnection(conn IConnection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.connections = append(cm.connections, conn)
	address := conn.GetAddress()
	if (address != common.Address{}) {
		cm.mapAddressConnections[conn.GetAddress()] = append(cm.mapAddressConnections[conn.GetAddress()], conn)
		cType := conn.GetType()
		if cType != "" && cm.mapTypeAddress[cType] != nil {
			cm.mapTypeAddress[cType][address] = struct{}{}
		}
	}
}

func removeConnectionAtIndex(s []IConnection, index int) []IConnection {
	if len(s) == 1 {
		return []IConnection{}
	}
	return append(s[:index], s[index+1:]...)
}

func (cm *ConnectionsManager) RemoveConnection(conn IConnection) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for i, v := range cm.connections {
		if v == conn {
			cm.connections = removeConnectionAtIndex(cm.connections, i)
		}
	}
	address := conn.GetAddress()
	if (address != common.Address{}) {
		connections := cm.mapAddressConnections[address]
		for i, v := range connections {
			if v == conn {
				cm.mapAddressConnections[address] = removeConnectionAtIndex(cm.mapAddressConnections[address], i)
			}
		}
		if len(cm.mapAddressConnections[address]) == 0 {
			delete(cm.mapAddressConnections, address)
			cType := conn.GetType()
			if cType != "" && cm.mapTypeAddress[cType] != nil && len(cm.mapAddressConnections[address]) == 0 {
				delete(cm.mapTypeAddress[cType], address)
			}
		}

	}
}

func (cm *ConnectionsManager) GetConnectionsByAddress(address common.Address) []IConnection {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return cm.mapAddressConnections[address]
}

func (cm *ConnectionsManager) GetConnectionByAddress(address common.Address) IConnection {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	if len(cm.mapAddressConnections[address]) > 0 {
		return cm.mapAddressConnections[address][0]
	} else {
		return nil
	}
}

func (cm *ConnectionsManager) GetConnectionsByType(cType string) map[common.Address][]IConnection {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	addresses := cm.mapTypeAddress[cType]
	rs := make(map[common.Address][]IConnection, len(addresses))
	for i, _ := range addresses {
		rs[i] = cm.mapAddressConnections[i]
	}
	return rs
}

func (cm *ConnectionsManager) GetConnectionsByTypeAtIdx(cType string, idx int) []IConnection {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	count := 0
	for _, v := range cm.GetConnectionsByType(cType) {
		if count == idx {
			return v
		}
		count++
	}
	return nil
}

func (cm *ConnectionsManager) GetTotalAddressByType(cType string) int {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	return len(cm.mapTypeAddress[cType])
}

func MapAddressConnectionsToInterface(data map[common.Address][]IConnection) map[common.Address]interface{} {
	rs := make(map[common.Address]interface{})
	for i, v := range data {
		rs[i] = v
	}
	return rs
}
