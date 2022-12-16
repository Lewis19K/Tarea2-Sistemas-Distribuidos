package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea2-Grupo08/Proto"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("dist029:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewMessageServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("****************************************************")
		fmt.Println("*                 MENU DE REBELDES                 *")
		fmt.Println("****************************************************")
		fmt.Println("*  1- Obtener LOGISTICA                            *")
		fmt.Println("*  2- Obtener FINANCIERA                           *")
		fmt.Println("*  3- Obtener MILITAR                              *")
		fmt.Println("*  0- Cerrar programa                              *")
		fmt.Println("****************************************************")
		menuOption, _ := reader.ReadString('\n')
		menuOption = strings.TrimSpace(menuOption)
		var tipo string
		if menuOption == "1" {
			tipo = "LOGISTICA"
		} else if menuOption == "2" {
			tipo = "FINANCIERA"
		} else if menuOption == "3" {
			tipo = "MILITAR"
		} else if menuOption == "0" {
			res, err := serviceClient.Close(context.Background(), &pb.ReqClose{})
			if err != nil {
				panic("cannot connect with server " + err.Error())
			}
			fmt.Println(res.Data)
		} else {
			panic("No se reconocio una opcion valida")
		}
		res, err := serviceClient.Download(context.Background(), &pb.ReqDownload{
			Tipo: tipo,
		})
		if err != nil {
			panic("cannot connect with server " + err.Error())
		}
		for i := 0; i < len(res.Descarga); i++ {
			fmt.Println(res.Descarga[i])
		}
	}
}
