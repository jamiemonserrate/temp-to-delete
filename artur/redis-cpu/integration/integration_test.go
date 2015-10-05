package integration_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

type metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

var _ = Describe("redis-cpu", func() {
	var (
		session *gexec.Session
		params  []string
	)

	BeforeEach(func() {
		params = []string{"--host", redisHost, "--port", redisPort}
	})

	JustBeforeEach(func() {
		session = runBinary(params...)
	})

	It("exits with a zero exit code", func() {
		Expect(session).To(gexec.Exit(0))
	})

	It("outputs proper JSON", func() {
		msg := session.Out.Contents()
		var jsondata interface{}
		err := json.Unmarshal(msg, &jsondata)
		Expect(err).ToNot(HaveOccurred())
	})

	It("outputs a metric", func() {
		msg := session.Out.Contents()
		var m metric
		err := json.Unmarshal(msg, &m)
		Expect(err).ToNot(HaveOccurred())

		Expect(m.Name).ToNot(Equal(""))
		Expect(m.Unit).ToNot(Equal(""))
	})

	Context("when a password param is passed", func() {
		BeforeEach(func() {
			params = append(params, "--password", "bar")
		})

		It("exits with a zero error code", func() {
			Expect(session).To(gexec.Exit(0))
		})
	})

	Context("when no host param is passed", func() {
		BeforeEach(func() {
			params = []string{}
		})

		It("exits with a non-zero error code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("provides a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Host not defined"))
		})
	})

	Context("when no port param is passed", func() {
		BeforeEach(func() {
			params = []string{"--host", "localhost"}
		})

		It("exits with a non-zero error code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("provides a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Port not defined"))
		})
	})

	Context("when we cannot connect to redis", func() {
		BeforeEach(func() {
			params = []string{"--host", "dummynet", "--port", "1234"}
		})

		It("exits with a non-zero exit code", func() {
			Expect(session.ExitCode()).ToNot(Equal(0))
		})

		It("provides a meaningful error message", func() {
			Expect(session.Err).To(gbytes.Say("Cannot connect to redis"))
		})
	})
})
