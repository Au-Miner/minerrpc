package transportServer

import (
	"io"
	"jrpc/src/rpc-core/handler"
	"jrpc/src/rpc-core/provider"
	"jrpc/src/rpc-core/serializer"
	servicesRegister "jrpc/src/rpc-core/services/register"
	transportUtils "jrpc/src/rpc-core/transport/utils"
	"log"
	"net"
)

type SocketServer struct {
	ServicesProvider provider.ServiceProvider
	ServiceRegister  servicesRegister.ServiceRegister
	Serializer       serializer.CommonSerializer
	RequestHandler   handler.RequestHandler
	Addr             *net.TCPAddr
}

func NewDefaultSocketServer(addrStr string) (*SocketServer, error) {
	return NewSocketServer(addrStr, DEFAULT_SERIALIZER)
}

func NewSocketServer(addrStr string, serializerId int) (*SocketServer, error) {
	addr, err := net.ResolveTCPAddr("tcp", addrStr)
	if err != nil {
		return nil, err
	}
	return &SocketServer{
		ServicesProvider: provider.NewServiceProvider(),
		ServiceRegister:  servicesRegister.NewZkServiceRegister(),
		Serializer:       serializer.GetByCode(serializerId),
		RequestHandler:   handler.NewRequestHandlerImpl(),
		Addr:             addr,
	}, nil
}

func (ss *SocketServer) Register(iClass interface{}) {
	ss.ServicesProvider.AddServiceProvider(iClass)
}

func (ss *SocketServer) Start() {
	l, err := net.Listen("tcp", ss.Addr.String())
	if err != nil {
		log.Printf("listen on %s err: %v\n", ss.Addr, err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
			continue
		}
		// 每接收到一个rpc请求，就开启一个goroutine处理
		go func() {
			nowObjReader := transportUtils.NewObjectReader(conn)
			nowObjWriter := transportUtils.NewObjectWriter(conn)
			transportUtils.NewObjectWriter(conn)
			for {
				decReq, err := nowObjReader.ReadObject()
				if err != nil {
					if err != io.EOF {
						log.Printf("read err: %v\n", err)
						return
					}
				}
				f, err := ss.ServicesProvider.GetFunc(decReq.Name)
				if err != nil {
					log.Printf("service provider err: %v\n", err)
					return
				}
				resp := ss.RequestHandler.Execute(decReq, f)
				err = nowObjWriter.WriteObject(resp, ss.Serializer)
				if err != nil {
					log.Printf("write err: %v\n", err)
					return
				}
			}
		}()
	}
}
