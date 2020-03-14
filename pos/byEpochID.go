package pos

import (
	"github.com/gorilla/websocket"
	"github.com/iwantogo/common"
)

type posEpochIDExecutor interface {
	getActivity(key string, c *websocket.Conn)
	getEpochIncentivePayDetail(key string, c *websocket.Conn)
	getEpochLeadersByEpochID(key string, c *websocket.Conn)
	getEpochIncentiveBlockNumber(key string, c *websocket.Conn)
	getRandomProposersByEpochID(key string, c *websocket.Conn)
	getLeaderGroupByEpochID(key string, c *websocket.Conn)
	getValidatorActivity(key string, c *websocket.Conn)
	getMaxBlockNumber(key string, c *websocket.Conn)
	getEpochStakeOut(key string, c *websocket.Conn)
	getSlotActivity(key string, c *websocket.Conn)
	getTimeByEpochID(key string, c *websocket.Conn)
}

type posEpochIDParamsPreSign struct {
	EpochID   int    `json:"epochID"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
}

type posEpochIDMessagePreSign struct {
	JSONRPC string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  posEpochIDParamsPreSign `json:"params"`
	ID      int                     `json:"id"`
}

type posEpochIDParamsPostSign struct {
	EpochID   int    `json:"epochID"`
	ChainType string `json:"chainType"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type posEpochIDMessagePostSign struct {
	JSONRPC string                   `json:"jsonrpc"`
	Method  string                   `json:"method"`
	Params  posEpochIDParamsPostSign `json:"params"`
	ID      int                      `json:"id"`
}

// NewReqByEpochID instantiates a new RPC-JSON call
func NewReqByEpochID(id int) *posEpochIDMessagePreSign {
	timeStamp := common.GetTimeStamp()
	msg := &posEpochIDMessagePreSign{
		JSONRPC: "2.0",
		Params: posEpochIDParamsPreSign{
			EpochID:   id,
			ChainType: "WAN",
			Timestamp: timeStamp,
		},
		ID: 1,
	}
	return msg
}

func (m *posEpochIDMessagePreSign) getActivity(key string, c *websocket.Conn) {
	m.Method = "getActivity"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getActivity",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetActivity returns activity info for the epoch ID provided
func GetActivity(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getActivity(k, c)
}

func (m *posEpochIDMessagePreSign) getEpochIncentivePayDetail(key string, c *websocket.Conn) {
	m.Method = "getEpochIncentivePayDetail"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getEpochIncentivePayDetail",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetEpochIncentivePayDetail returns reward info of all the
// verification nodes and clients for the epoch ID provided
// e.g. RNP reward, EL reward and chunk reward
func GetEpochIncentivePayDetail(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getEpochIncentivePayDetail(k, c)
}

func (m *posEpochIDMessagePreSign) getEpochLeadersByEpochID(key string, c *websocket.Conn) {
	m.Method = "getEpochLeadersByEpochID"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getEpochLeadersByEpochID",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetEpochLeadersByEpochID returns the public key list of
// epoch leaders for the epoch ID provided
func GetEpochLeadersByEpochID(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getEpochLeadersByEpochID(k, c)
}

func (m *posEpochIDMessagePreSign) getEpochIncentiveBlockNumber(key string, c *websocket.Conn) {
	m.Method = "getEpochIncentiveBlockNumber"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getEpochIncentiveBlockNumber",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetEpochIncentiveBlockNumber returns the block number in
// which the epoch incentive transaction is included for the
// epoch ID provided
func GetEpochIncentiveBlockNumber(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getEpochIncentiveBlockNumber(k, c)
}

func (m *posEpochIDMessagePreSign) getRandomProposersByEpochID(key string, c *websocket.Conn) {
	m.Method = "getRandomProposersByEpochID"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getRandomProposersByEpochID",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetRandomProposersByEpochID returns the RNP's public keys
// for the epoch ID provided
func GetRandomProposersByEpochID(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getRandomProposersByEpochID(k, c)
}

func (m *posEpochIDMessagePreSign) getLeaderGroupByEpochID(key string, c *websocket.Conn) {
	m.Method = "getLeaderGroupByEpochID"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getLeaderGroupByEpochID",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetLeaderGroupByEpochID returns the ELs and RNPs' public keys
// plus addresses for the epoch ID provided
func GetLeaderGroupByEpochID(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getLeaderGroupByEpochID(k, c)
}

func (m *posEpochIDMessagePreSign) getValidatorActivity(key string, c *websocket.Conn) {
	m.Method = "getValidatorActivity"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getValidatorActivity",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetValidatorActivity returns activity info of the ELd and RNPs
// for the epoch ID provided
func GetValidatorActivity(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getValidatorActivity(k, c)
}

func (m *posEpochIDMessagePreSign) getMaxBlockNumber(key string, c *websocket.Conn) {
	m.Method = "getMaxBlockNumber"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getMaxBlockNumber",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetMaxBlockNumber returns the highest block number for the
// epoch ID provided
func GetMaxBlockNumber(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getMaxBlockNumber(k, c)
}

func (m *posEpochIDMessagePreSign) getEpochStakeOut(key string, c *websocket.Conn) {
	m.Method = "getEpochStakeOut"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getEpochStakeOut",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetEpochStakeOut returns the stake-out transactions
// for the epoch ID provided
func GetEpochStakeOut(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getEpochStakeOut(k, c)
}

func (m *posEpochIDMessagePreSign) getSlotActivity(key string, c *websocket.Conn) {
	m.Method = "getSlotActivity"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getSlotActivity",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetSlotActivity returns activity info of the slot leader
// for the epoch ID provided
func GetSlotActivity(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getSlotActivity(k, c)
}

func (m *posEpochIDMessagePreSign) getTimeByEpochID(key string, c *websocket.Conn) {
	m.Method = "getTimeByEpochID"
	sig := common.GenSig(m, key)
	msgSend := &posEpochIDMessagePostSign{
		JSONRPC: "2.0",
		Method:  "getTimeByEpochID",
		Params: posEpochIDParamsPostSign{
			EpochID:   m.Params.EpochID,
			ChainType: "WAN",
			Timestamp: m.Params.Timestamp,
			Signature: sig,
		},
		ID: 1,
	}

	common.SendMessage(msgSend, c)
}

// GetTimeByEpochID returns the start time in UTC time of
// the epoch ID provided
func GetTimeByEpochID(pe posEpochIDExecutor, k string, c *websocket.Conn) {
	pe.getTimeByEpochID(k, c)
}
