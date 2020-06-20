package main

import (
	"context"
	"log"
	"net"

	// Llamamos el paquete de gRPC
	"google.golang.org/grpc"

	// Llamamos el compilado que nos generó protoc
	//pb "github.com/zetogk//notificador"

	pb "github.com/ajme-prog/proyecto2sopes"
)

const (
	// Puerto por donde expondremos el servicio
	port = ":50051"
)

// server es usado para implementar Notificador.
type server struct{}

// EnviarCorreo implementar Notificador.EnviarCorreo
// El parámetro in, corresponde al mensaje CorreoRequest
func (s *server) Enviardatos(ctx context.Context, in *pb.DatosRequest) (*pb.Responsevalor, error) {
	log.Printf("Enviando datos al server: %v a %v", in.Nombre, in.Departamento)
	// Generamos y retornamos la respuesta de tipo CorreoResponse
	return &pb.Responsevalor{Enviado: true}, nil
}

func main() {
	// Exponemos nuestro servidor para que reciba llamados
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificadorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
