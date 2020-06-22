package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	// Llamamos el paquete de gRPC
	"google.golang.org/grpc"

	// Llamamos el compilado que nos generó protoc
	//pb "github.com/zetogk//notificador"

	"net/http"

	pb "github.com/ajme-prog/proyecto2sopes/notificador"
	//"os"
	// Llamamos el paquete de gRPC
	// Llamamos el compilado que nos generó protoc
)

const (
	address = "127.0.0.1:50051" // Definimos por que host y puerto nos comunicamos
)

func reply(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==============================")
	fmt.Println("ENTRE A REPLY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if err := r.ParseForm(); err != nil {
		fmt.Println("el error es", err)
	}

	//this is my first impulse.  It makes the most sense to me.
	//fmt.Println(r.PostForm);          //out -> `map[]`  would be `map[string]string` I think
	//fmt.Println(r.PostForm["hat"]);   //out -> `[]`  would be `fez` or `["fez"]`
	// fmt.Println(r.Body);              //out -> `&{0xc82000e780 <nil> <nil> false true {0 0} false false false}`

	type Persona struct {
		Nombre       string
		Departamento string
		Edad         int32
		Contagio     string
		Estado       string
	}

	//this is the way the linked SO post above said should work.  I don't see how the r.Body could be decoded.
	decoder := json.NewDecoder(r.Body)
	var t Persona
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Println(err)
	}
	if t.Nombre != "" {
		fmt.Println("ESTOY RECIBIENDO A LA PERSONA CON NOMBRE id: ", t.Nombre)

		//----------------CONFIGURAR CLIENTE GRPC PARA ENVIAR LOS DATOS A PYTHON------------------------
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Error en la conexión: %v", err)
		}
		defer conn.Close()
		c := pb.NewNotificadorClient(conn)

		// Enviamos la petición al servidor
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Enviardatos(ctx, &pb.DatosRequest{Nombre: t.Nombre, Departamento: t.Departamento, Edad: t.Edad, Formacontagio: t.Contagio, Estado: t.Estado})

		//	time.Sleep(2 * time.Second)
		if err != nil {
			log.Fatalf("Error en la petición: %v", err)
		}
		// Imprimimos la respuesta
		log.Println("¿la persona fue enviada?: ", r.Enviado)

	}
}

func main() {

	http.HandleFunc("/post", reply)
	log.Fatal(http.ListenAndServe(":3030", nil))
	// configurar la conexión al servidor
	/*conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error en la conexión: %v", err)
	}
	defer conn.Close()
	c := pb.NewNotificadorClient(conn)

	// Enviamos la petición al servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Enviardatos(ctx, &pb.DatosRequest{Nombre: "Alan", Departamento: "Chiquimula", Edad: 55, Formacontagio: "beso", Estado: "grave"})
	if err != nil {
		log.Fatalf("Error en la petición: %v", err)
	}
	// Imprimimos la respuesta
	log.Println("¿El correo fue enviado?: ", r.Enviado)*/
}
