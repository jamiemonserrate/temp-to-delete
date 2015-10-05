package metrics_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/redis-cpu/metrics"
	"github.com/pivotal-cf/redis-cpu/metrics/fakes"
)

var _ = Describe("CPULoad", func() {
	var client *fakes.FakeClient

	BeforeEach(func() {
		client = new(fakes.FakeClient)
		client.InfoReturns(map[string]string{"used_cpu_sys": "123.45"}, nil)
	})

	It("calls client.Info()", func() {
		metrics.CPULoad(client)
		Expect(client.InfoCallCount()).To(Equal(1))
	})

	It("does not return an error", func() {
		_, err := metrics.CPULoad(client)
		Expect(err).ToNot(HaveOccurred())
	})

	It("reports the right CPU load metric", func() {
		m, _ := metrics.CPULoad(client)
		Expect(m.Key).To(Equal("cpu_load"))
		Expect(m.Value).To(Equal(123.45))
		Expect(m.Unit).To(Equal("load"))
	})

	Context("client.Info() errors", func() {
		var expectedErr = errors.New("some-error")

		BeforeEach(func() {
			client.InfoReturns(map[string]string{}, expectedErr)
		})

		It("returns the error", func() {
			_, err := metrics.CPULoad(client)
			Expect(err).To(Equal(expectedErr))
		})
	})

	Context("client info has invalid value", func() {
		BeforeEach(func() {
			client.InfoReturns(map[string]string{"used_cpu_sys": "nofloat"}, nil)
		})

		It("returns the error", func() {
			_, err := metrics.CPULoad(client)
			Expect(err).To(HaveOccurred())
		})
	})

})
