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
	var (
		session    *gexec.Session
		configPath string
	)

	BeforeEach(func() {
		configPath = filepath.Join(".", "assets", "redis.conf")
	})

	JustBeforeEach(func() {
		session = runBin(configPath)
	})

	var decodeJson = func(session *gexec.Session) metrics.Metric {
		var m metrics.Metric
		err := json.NewDecoder(bytes.NewReader(session.Out.Contents())).Decode(&m)
		Expect(err).ToNot(HaveOccurred())
		return m
	}

	It("outputs a single cpu_load metric", func() {
		m := decodeJson(session)
		Expect(m.Key).To(Equal("cpu_load"))
		Expect(m.Unit).To(Equal("load"))
	})

	Context("when --config flag is not provided", func() {
		BeforeEach(func() {
			configPath = ""
		})

		It("exits with a non-zero exit code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("outputs a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Missing flag --config"))
		})
	})

	Context("when redis config file does not exist", func() {
		BeforeEach(func() {
			configPath = "/path/to/nowhere"
		})

		It("exits with a non-zero exit code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("outputs a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Error loading Redis config file"))
		})
	})

	Context("when we cannot connect to redis", func() {
		BeforeEach(func() {
			stopRedis()
		})

		AfterEach(func() {
			startRedis()
		})

		It("exits with a non-zero exit code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("outputs a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Error connecting to Redis"))
		})
	})
})
