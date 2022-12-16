package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	//"strconv"
	"strings"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea2-Grupo08/Proto"
	pbDataNode "github.com/Sistemas-Distribuidos-2022-2/Tarea2-Grupo08/protoDataNode"
	"google.golang.org/grpc"
)

// ############### Constantes ###############
const node1 = "Grunt"
const node2 = "Synth"
const node3 = "Cremator"
const port1 = "50052" // aqui no use el 50051 porque lo usaste en el de combine
const port2 = "50053"
const port3 = "50054"

// ############### Server Interface ###############
type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Upload(ctx context.Context, msg *pb.ReqUpload) (*pb.ResUpload, error) {
	var aux []string = strings.Split(msg.Data, " : ") // separa la data recibida (<TIPO> : <ID> : <DATA>)
	nodeName := ""
	nodePort := ""

	// Elige al azar a cual DataNode enviar la data

	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	var i = r2.Intn(3)
	var hostS string
	switch i {
	case 0:
		nodeName = node1
		nodePort = port1
		hostS = "dist030"
	case 1:
		nodeName = node2
		nodePort = port2
		hostS = "dist031"
	case 2:
		nodeName = node3
		nodePort = port3
		hostS = "dist032"
	}

	// Ingresa al DATA.txt de NameNode
	f2, err := os.OpenFile("NameNode/DATA.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	_, err = f2.WriteString(aux[0] + " : " + aux[1] + " : " + nodeName + "\n")
	if err != nil {
		panic(err)
	}

	// Enviar la data al DataNode
	conn, err := grpc.Dial(hostS+":"+nodePort, grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}
	serviceClient := pbDataNode.NewDataNodeServiceClient(conn)
	switch i {
	case 0:
		_, err := serviceClient.Node1(context.Background(), &pbDataNode.ReqSaveData{
			Data: msg.Data,
		})
		if err != nil {
			panic("error:  " + err.Error())
		}
	case 1:
		_, err := serviceClient.Node2(context.Background(), &pbDataNode.ReqSaveData{
			Data: msg.Data,
		})
		if err != nil {
			panic("error:  " + err.Error())
		}
	case 2:
		_, err := serviceClient.Node3(context.Background(), &pbDataNode.ReqSaveData{
			Data: msg.Data,
		})
		if err != nil {
			panic("error:  " + err.Error())
		}
	}

	fmt.Printf("Mensaje subido: %s\n", msg.Data)
	return &pb.ResUpload{}, nil
}

func (s *server) Download(ctx context.Context, msg *pb.ReqDownload) (*pb.ResDownload, error) {
	var downloadList []string
	//var index = 0
	if _, err := os.Stat("NameNode/DATA.txt"); err == nil { // chequea que exista DATA.txt, en caso de no existir no descarga nada (evitando errores)
		file, err := os.Open("NameNode/DATA.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		var text []string

		for scanner.Scan() {
			text = append(text, scanner.Text())
		}

		for _, each_ln := range text {
			//fmt.Println(each_ln)
			var aux []string = strings.Split(each_ln, " : ") // separa la data recibida (<TIPO> : <ID> : <NODO>)
			if aux[0] == msg.Tipo {
				id := aux[1]
				var nodoNombre = aux[2]
				switch nodoNombre {
				case node1:
					// Solicitar dato al DataNode
					conn, err := grpc.Dial("dist030:"+port1, grpc.WithInsecure())
					if err != nil {
						panic("cannot connect with server " + err.Error())
					}
					serviceClient := pbDataNode.NewDataNodeServiceClient(conn)
					res, err := serviceClient.Getid(context.Background(), &pbDataNode.ReqData{
						ID: id,
					})
					if err != nil {
						panic("error:  " + err.Error())
					}
					downloadList = append(downloadList, aux[1]+" : "+res.Data)
				case node2:
					// Solicitar dato al DataNode
					conn, err := grpc.Dial("dist031:"+port2, grpc.WithInsecure())
					if err != nil {
						panic("cannot connect with server " + err.Error())
					}
					serviceClient := pbDataNode.NewDataNodeServiceClient(conn)
					res, err := serviceClient.Getid2(context.Background(), &pbDataNode.ReqData{
						ID: id,
					})
					if err != nil {
						panic("error:  " + err.Error())
					}
					downloadList = append(downloadList, aux[1]+" : "+res.Data)
				case node3:
					// Solicitar dato al DataNode
					//fmt.Println("CASE NODE 3")
					conn, err := grpc.Dial("dist032:"+port3, grpc.WithInsecure())
					if err != nil {
						panic("cannot connect with server " + err.Error())
					}
					//fmt.Println("CREO LA CONEXION NODE 3")
					serviceClient := pbDataNode.NewDataNodeServiceClient(conn)
					//fmt.Println("CREO EL SERVICECLIENT NODE 3")
					res, err := serviceClient.Getid3(context.Background(), &pbDataNode.ReqData{
						ID: id,
					})
					if err != nil {
						panic("error:  " + err.Error())
					}
					//fmt.Println("PASO EL GETPORID 3")
					downloadList = append(downloadList, aux[1]+" : "+res.Data)
				}
			}
		}
	} else {
		return &pb.ResDownload{Descarga: downloadList}, nil
	}
	return &pb.ResDownload{Descarga: downloadList}, nil
}

func shutDownNode() {
	time.Sleep(3 * time.Second)
	defer os.Exit(0)
	return
}

func (s *server) Close(ctx context.Context, msg *pb.ReqClose) (*pb.ResClose, error) {
	var port string
	var host string
	for i := 0; i < 3; i++ {
		if i == 0 {
			port = port1
			host = "dist030"
		} else if i == 1 {
			port = port2
			host = "dist031"
		} else if i == 2 {
			port = port3
			host = "dist032"
		}

		conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
		if err != nil {
			panic("cannot connect with server " + err.Error())
		}
		serviceClient := pbDataNode.NewDataNodeServiceClient(conn)
		res, err := serviceClient.ShutDown(context.Background(), &pbDataNode.ReqShutDown{})
		if err != nil {
			panic("error:  " + err.Error())
		}

		fmt.Println(res.Data)
	}

	go shutDownNode()

	return &pb.ResClose{Data: "Cerrando NameNode"}, nil
}

// ############### Main ###############

func main() {
	if _, err := os.Stat("NameNode/DATA.txt"); err == nil {
		e := os.Remove("NameNode/DATA.txt")
		if e != nil {
			log.Fatal(e)
		}
	}
	fmt.Println("Recibiendo solicitudes ...")
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMessageServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

	// Se crea DATA.txt
	f, err := os.Create("NameNode/DATA.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
