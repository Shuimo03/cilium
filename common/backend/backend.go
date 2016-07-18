package backend

import (
	"net"

	"github.com/noironetworks/cilium-net/common/ipam"
	"github.com/noironetworks/cilium-net/common/types"

	"github.com/gorilla/websocket"
)

type bpfBackend interface {
	EndpointJoin(ep types.Endpoint) error
	EndpointLeave(epID uint16) error
	EndpointLeaveByDockerEPID(dockerEPID string) error
	EndpointGet(epID uint16) (*types.Endpoint, error)
	EndpointGetByDockerEPID(dockerEPID string) (*types.Endpoint, error)
	EndpointsGet() ([]types.Endpoint, error)
	EndpointUpdate(epID uint16, opts types.OptionMap) error
	EndpointSave(ep types.Endpoint) error
	EndpointLabelsGet(epID uint16) (*types.OpLabels, error)
	EndpointLabelsUpdate(epID uint16, op types.LabelOP, labels types.Labels) error
}

type ipamBackend interface {
	GetIPAMConf(ipamType ipam.IPAMType, options ipam.IPAMReq) (*ipam.IPAMConfigRep, error)
	AllocateIP(ipamType ipam.IPAMType, options ipam.IPAMReq) (*ipam.IPAMRep, error)
	ReleaseIP(ipamType ipam.IPAMType, options ipam.IPAMReq) error
}

type labelBackend interface {
	PutLabels(labels types.Labels, contĨD string) (*types.SecCtxLabel, bool, error)
	GetLabels(uuid uint32) (*types.SecCtxLabel, error)
	GetLabelsBySHA256(sha256sum string) (*types.SecCtxLabel, error)
	DeleteLabelsByUUID(uuid uint32, contĨD string) error
	DeleteLabelsBySHA256(sha256sum, contID string) error
	GetMaxID() (uint32, error)
}

type policyBackend interface {
	PolicyAdd(path string, node *types.PolicyNode) error
	PolicyDelete(path string) error
	PolicyGet(path string) (*types.PolicyNode, error)
	PolicyCanConsume(ctx *types.SearchContext) (*types.SearchContextReply, error)
}

type control interface {
	Ping() (*types.PingResponse, error)
	Update(opts types.OptionMap) error
	SyncState(path string, clean bool) error
}

type ui interface {
	GetUIIP() (*net.TCPAddr, error)
	RegisterUIListener(conn *websocket.Conn) (chan types.UIUpdateMsg, error)
}

// CiliumBackend is the interface for both client and daemon.
type CiliumBackend interface {
	bpfBackend
	control
	ipamBackend
	labelBackend
	policyBackend
}

// CiliumDaemonBackend is the interface for daemon only.
type CiliumDaemonBackend interface {
	CiliumBackend
	ui
}
