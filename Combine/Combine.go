package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea2-Grupo08/Proto"
	"google.golang.org/grpc"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func main() {
	var IDs []string
	fmt.Println("Subir data:")
	conn, err := grpc.Dial("dist029:50051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewMessageServiceClient(conn)

	// ############### Recibir mensajes de la terminal ###############
	fmt.Print("Ingresar data con el siguiente formato: ")
	fmt.Println("<TIPO> : <ID> : <DATA>")
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		data = strings.TrimSpace(data)
		var aux []string = strings.Split(data, " : ") // separa la data recibida (<TIPO> : <ID> : <DATA>), para poder verificar si el ID es valido o no esta disponible
		var valido = true
		for _, a := range IDs {
			if a == aux[1] {
				valido = false
			}
		}
		if !valido {
			fmt.Println("el ID ya se encuentra utilzado")
			continue
		}
		_, err := serviceClient.Upload(context.Background(), &pb.ReqUpload{
			Data: data,
		})
		if err != nil {
			panic("error:  " + err.Error())
		}
		fmt.Println("Mensaje enviado Exitosamente")
		IDs = append(IDs, aux[1])
	}
}
