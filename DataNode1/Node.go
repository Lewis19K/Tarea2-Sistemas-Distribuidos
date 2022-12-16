package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea2-Grupo08/protoDataNode"
	"google.golang.org/grpc"
)

// ############### Server Interface ###############
type server struct {
	pb.UnimplementedDataNodeServiceServer
}

func (s *server) Node1(ctx context.Context, msg *pb.ReqSaveData) (*pb.ResSaveData, error) {

	// Ingresa al DATA.txt de NameNode
	f2, err := os.OpenFile("DataNode1/DATA.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	_, err = f2.WriteString(msg.Data + "\n")
	if err != nil {
		panic(err)
	}

	return &pb.ResSaveData{}, nil
}

func (s *server) Getid(ctx context.Context, msg *pb.ReqData) (*pb.ResData, error) {
	file, err := os.Open("DataNode1/DATA.txt")
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
		var aux []string = strings.Split(each_ln, " : ") // separa la data recibida (<TIPO> : <ID> : <DATA>)
		if aux[1] == msg.ID {
			fmt.Printf("Solicitud de NameNode recibida, mensaje enviado: %s\n", aux[2])
			return &pb.ResData{Data: aux[2]}, nil
		}
	}

	return &pb.ResData{Data: "ERROR: no se encontro el dato"}, nil
}

func shutDownNode() {
	time.Sleep(3 * time.Second)
	defer os.Exit(0)
	return
}

func (s *server) ShutDown(ctx context.Context, msg *pb.ReqShutDown) (*pb.ResShutDown, error) {
	go shutDownNode()
	return &pb.ResShutDown{Data: "Cerrando DataNode1"}, nil
}

// ############### Main ###############

func main() {
	if _, err := os.Stat("DataNode1/DATA.txt"); err == nil {
		e := os.Remove("DataNode1/DATA.txt")
		if e != nil {
			log.Fatal(e)
		}
	}
	fmt.Println("Nodo funcionando ...")
	listner, err := net.Listen("tcp", ":50052")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterDataNodeServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

	// Se crea DATA.txt
	f, err := os.Create("DataNode1/DATA.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
