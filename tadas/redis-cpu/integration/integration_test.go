package integration_test

import (
	"bytes"
	"encoding/json"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/redis-cpu/metrics"
)

var _ = Describe("redis-cpu", func() {
	var session *gexec.Session

	Context("when redis is running", func() {
		BeforeEach(func() {
			path := filepath.Join(".", "assets", "redis.conf")
			startRedis(path)
			session = execBin(path)
		})

		AfterEach(func() {
			stopRedis()
		})

		It("returns a zero exit code", func() {
			Expect(session.ExitCode()).To(Equal(0))
		})

		It("outputs proper metrics JSON", func() {
			decoder := json.NewDecoder(bytes.NewBuffer(session.Out.Contents()))
			m := new(metrics.Metrics)
			err := decoder.Decode(m)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when the --config flag is not provided", func() {
		BeforeEach(func() {
			session = execBin("")
		})

		It("returns with non-zero exit code", func() {
			Expect(session.ExitCode()).NotTo(Equal(0))
		})

		It("provides a meaningful error mesage", func() {
			Expect(session.Err).To(gbytes.Say("Missing flag"))
		})
	})

	Context("when the redis config file does not exist", func() {
		BeforeEach(func() {
			session = execBin("/path/to/nowhere")
		})

		It("returns with non-zero exit code", func() {
			Expect(session.ExitCode()).NotTo(Equal(0))
		})

		It("provides a meaningful error mesage", func() {
			Expect(session.Err).To(gbytes.Say("Redis config does not exist"))
		})
	})

	Context("when the redis config file is invalid", func() {
		BeforeEach(func() {
			session = execBin(filepath.Join(".", "assets", "invalid.conf"))
		})

		It("returns with non-zero exit code", func() {
			Expect(session.ExitCode()).NotTo(Equal(0))
		})

		It("provides a meaningful error mesage", func() {
			Expect(session.Err).To(gbytes.Say("Invalid redis config"))
		})
	})

	Context("when redis is not reachable", func() {
		BeforeEach(func() {
			session = execBin(filepath.Join(".", "assets", "redis.conf"))
		})

		It("returns with non-zero exit code", func() {
			Expect(session.ExitCode()).NotTo(Equal(0))
		})

		It("provides a meaningful error mesage", func() {
			Expect(session.Err).To(gbytes.Say("Can not connect to Redis"))
		})
	})
})
