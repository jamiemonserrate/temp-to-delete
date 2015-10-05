package integration_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/cf-redis-broker/integration"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var (
	binPath     string
	redisRunner *integration.RedisRunner
)

var _ = BeforeSuite(func() {
	var err error
	binPath, err = gexec.Build("github.com/pivotal-cf/redis-cpu")
	Expect(err).ToNot(HaveOccurred())
	redisRunner = &integration.RedisRunner{}
	startRedis()
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
	stopRedis()
})

func stopRedis() {
	redisRunner.Stop()
}

func startRedis() {
	redisRunner.Start([]string{"--port", "6480"})
}

func runBin(configPath string) *gexec.Session {
	var cmd *exec.Cmd
	if configPath != "" {
		cmd = exec.Command(binPath, "--config", configPath)
	} else {
		cmd = exec.Command(binPath)
	}

	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	Eventually(session).Should(gexec.Exit())

	return session
}
