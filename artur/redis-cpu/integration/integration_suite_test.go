package integration_test

import (
	"io/ioutil"
	"os/exec"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/cf-redis-broker/integration"

	"testing"
)

var (
	redisHost     = "127.0.0.1"
	redisPort     = strconv.Itoa(integration.RedisPort)
	redisPassword = "foobar"

	pathToBinary string
	redisRunner  *integration.RedisRunner
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	var err error
	pathToBinary, err = gexec.Build("github.com/pivotal-cf/redis-cpu")
	Expect(err).ToNot(HaveOccurred())

	redisRunner = new(integration.RedisRunner)

	redisRunner.Dir, err = ioutil.TempDir("", "redis")
	Expect(err).ToNot(HaveOccurred())

	redisRunner.Start([]string{
		"--bind", redisHost,
		"--port", redisPort,
		"--requirepass", redisPassword,
	})
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	redisRunner.Stop()
})

func runBinary(params ...string) *gexec.Session {
	command := exec.Command(pathToBinary, params...)
	session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	session.Wait()
	return session
}
