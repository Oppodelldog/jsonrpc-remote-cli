package main

import (
	"fmt"
	"github.com/Oppodelldog/dockertest"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

const waitingTimeout = time.Minute
const goVersion = "golang:1.17.3"

func main() {
	var projectDir string
	var goPath = os.Getenv("GOPATH")
	if val, ok := os.LookupEnv("TRAVIS_BUILD_DIR"); ok {
		projectDir = val
	} else {
		val, err := os.Getwd()
		panicOnErr(err)
		parentOfWd := filepath.Dir(val)
		projectDir = parentOfWd
	}

	fmt.Printf("START BUILD IN '%s'\n", projectDir)

	fmt.Println("NEW DOCKERTEST SESSION")
	test, err := dockertest.NewSession()
	panicOnErr(err)

	go cancelSessionOnSigTerm(test)

	fmt.Println("CLEANUP REMAINS")
	test.CleanupRemains()

	var testResult = TestResult{ExitCode: -1}
	defer cleanup(test, &testResult)

	test.SetLogDir(path.Join(projectDir, "test-logs"))

	net, err := test.CreateBasicNetwork("test-network").Create()
	panicOnErr(err)

	basicConfiguration := test.NewContainerBuilder().
		Image(goVersion).
		Connect(net).
		WorkingDir("/app").
		Env("GOPATH", goPath).
		Mount(goPath, goPath).
		Mount(projectDir, "/app")

	api, err := basicConfiguration.NewContainerBuilder().
		Name("api").
		Cmd("go run cmd/server/main.go").
		Build()
	panicOnErr(err)

	tests, err := basicConfiguration.NewContainerBuilder().
		Name("client").
		CmdArgs("bash", "-c", `go run cmd/generator/main.go --client-folder=cmd/client --endpoint-uri=http://api:8080/rpc && \
		go run cmd/client/main.go add -a=9990 -b=7`).
		Link(api, "api", net).
		Build()
	panicOnErr(err)

	fmt.Println("START API CONTAINER")
	err = api.Start()
	panicOnErr(err)

	fmt.Println("WAIT FOR API READINESS")
	err = <-test.NotifyContainerLogContains(api, waitingTimeout, "starting JSON-RPC Example")
	panicOnErr(err)

	fmt.Println("START CLIENT CONTAINER")
	err = tests.Start()
	panicOnErr(err)

	err = <-test.NotifyContainerLogContains(tests, waitingTimeout, "9997")
	panicOnErr(err)
	fmt.Println("WAIT FOR CLIENT RPC CALL TO SUCCEED")

	<-test.NotifyContainerExit(tests, waitingTimeout)

	testResult.ExitCode, err = tests.ExitCode()
	panicOnErr(err)

	test.DumpContainerLogsToDir(api, tests)
}

func cancelSessionOnSigTerm(session *dockertest.Session) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	session.Cancel()
}

func cleanup(test *dockertest.Session, testResult *TestResult) {
	fmt.Println("CLEANUP-START")
	test.Cleanup()
	fmt.Println("CLEANUP-DONE")

	if r := recover(); r != nil {
		fmt.Printf("ERROR: %v\n", r)
	}

	os.Exit(testResult.ExitCode)
}

type TestResult struct {
	ExitCode int
}

func panicOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
