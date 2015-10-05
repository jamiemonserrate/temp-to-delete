package integration_test

import (
	"os/exec"
	"syscall"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var execPath string
var redisSession *gexec.Session

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	var err error
	execPath, err = gexec.Build("github.com/pivotal-cf/redis-cpu")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func execBin(configPath string) *gexec.Session {
	var cmd *exec.Cmd
	if configPath != "" {
		cmd = exec.Command(execPath, "--config", configPath)
	} else {
		cmd = exec.Command(execPath)
	}

	session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit())
	return session
}

func startRedis(configPath string) {
	cmd := exec.Command("redis-server", configPath)
	var err error
	redisSession, err = gexec.Start(cmd, nil, nil)
	Expect(err).ToNot(HaveOccurred())
	time.Sleep(time.Second * 5)
}

func stopRedis() {
	redisSession.Signal(syscall.SIGINT)
	redisSession.Wait(time.Second * 5)
}
