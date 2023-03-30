package network

import (
	"context"
	"fmt"
	"net"

	"gitlab.com/meta-node/meta-node/pkg/bls"
	"gitlab.com/meta-node/meta-node/pkg/common"
	p_common "gitlab.com/meta-node/meta-node/pkg/common"
	"gitlab.com/meta-node/meta-node/pkg/config"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
)

type ISocketServer interface {
	Listen(string) error
	Stop()

	OnConnect(IConnection)
	OnDisconnect(IConnection)

	HandleConnection(IConnection) error
	GetHandler() chan interface{}
}

type SocketServer struct {
	connectionsManager IConnectionsManager
	listener           net.Listener
	handler            IHandler
	config             config.IConfig
	keyPair            *bls.KeyPair
	ctx                context.Context
	cancelFunc         context.CancelFunc
}

func NewSockerServer(
	config config.IConfig,
	keyPair *bls.KeyPair,
	connectionsManager IConnectionsManager,
	handler IHandler,
) ISocketServer {
	s := &SocketServer{
		config:             config,
		keyPair:            keyPair,
		connectionsManager: connectionsManager,
		handler:            handler,
	}
	s.ctx, s.cancelFunc = context.WithCancel(context.Background())
	return s
}
func (s *SocketServer) GetHandler() chan interface{}{
	return s.handler.GetChData()
		
	
}

func (s *SocketServer) Listen(listenAddress string) error {
	var err error
	s.listener, err = net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}
	defer func() {
		s.listener.Close()
		s.listener = nil
	}()
	logger.Info(fmt.Sprintf("Listening at %v", listenAddress))
	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			tcpConn, err := s.listener.Accept()
			if err != nil {
				logger.Warn(fmt.Sprintf("Error when accept connection %v\n", err))
				continue
			}
			conn, err := ConnectionFromTcpConnection(tcpConn)
			if err != nil {
				logger.Warn(fmt.Sprintf("error when create connection from tcp connection: %v", err))
				continue
			}
			s.OnConnect(conn)
			go s.HandleConnection(conn)
			SendMessage(conn, s.keyPair, common.InitConnection, &pb.InitConnection{
				Address: s.keyPair.GetAddress().Bytes(),
				Type:    s.config.GetNodeType(),
			}, p_common.Sign{}, s.config.GetVersion())
		}
	}
}

func (s *SocketServer) Stop() {
	s.cancelFunc()
}

func (s *SocketServer) OnConnect(conn IConnection) {
	logger.Info(fmt.Sprintf("On Connect with %v:%v type %v", conn.GetIp(), conn.GetPort(), conn.GetType()))
}

func (s *SocketServer) OnDisconnect(conn IConnection) {
	logger.Warn(
		fmt.Sprintf(
			"On Disconnect with %v:%v - address %v - type %v",
			conn.GetIp(),
			conn.GetPort(),
			conn.GetAddress(),
			conn.GetType(),
		),
	)
	s.connectionsManager.RemoveConnection(conn)
}

func (s *SocketServer) HandleConnection(conn IConnection) error {
	logger.Debug(fmt.Sprintf("handle connection %v", conn.GetAddress()))
	defer func() {
		conn.Disconnect()
		s.OnDisconnect(conn)
	}()
	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			request, err := conn.ReadRequest()
			if err != nil {
				if err != ErrDisconnected {
					logger.Warn(fmt.Sprintf("error when read request %v", err))
				}
				return err
			}
			err = s.handler.HandleRequest(request)
			if err != nil {
				logger.Warn(fmt.Sprintf("error when process request %v", err))
				continue
			}
		}
	}
}
