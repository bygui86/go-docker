package main

import (
	"fmt"
	"time"

	"github.com/bygui86/go-docker/containers"
	"github.com/bygui86/go-docker/images"
)

func main() {

	timeout := time.Duration(int64(5))

	fmt.Println()
	fmt.Println("Go-Docker sample project")
	fmt.Println()

	fmt.Println("Pull 'alpine' image")
	images.Pull("docker.io/library/alpine")
	fmt.Println("Pull 'bitnami/minideb' image")
	images.Pull("docker.io/bitnami/minideb")
	fmt.Println("Pull 'registry' image")
	images.Pull("docker.io/library/registry")
	fmt.Println()

	fmt.Println("Tag 'bitnami/minideb' image as 'bitnami/minideb'")
	images.Tag("docker.io/bitnami/minideb", "localhost:5000/bitnami/minideb")
	fmt.Println()

	fmt.Println("List images")
	images.List(true)
	fmt.Println()

	fmt.Println("Run 'docker.io/library/registry'")
	contId := containers.Run("docker.io/library/registry", []string{"5000"}, "registry")
	fmt.Printf("Container ID: %s", contId)
	fmt.Println()
	fmt.Println()

	time.Sleep(timeout * time.Second)

	fmt.Println("Logs from 'docker.io/library/registry'")
	containers.Logs(contId, false)
	fmt.Println()

	fmt.Println("List containers")
	containers.List(true, 10)
	fmt.Println()

	fmt.Println("Push 'localhost:5000/bitnami/minideb'")
	images.Push("localhost:5000/bitnami/minideb")
	fmt.Println()

	fmt.Println("Stop 'docker.io/library/registry'")
	containers.Stop(contId, timeout)
	fmt.Println()

	time.Sleep(timeout * time.Second)

	// NOT WORKING :(
	// error message: Conflict, cannot remove the default name of the container
	// fmt.Println("Remove 'docker.io/library/registry'")
	// containers.Remove(contId, true)
	// fmt.Println()

	fmt.Println("Prune")
	containers.Prune()
	fmt.Println()
}
