package paypalrestsdk

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
)

func TestApi(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Paypal Rest API - API Suite")
}

var _ = Describe("API", func() {
    Describe("Configure", func() {
        It("should be has field mode", func() {
            var configure = Configure{ Mode: "live" }

            Expect(configure.Mode).To(Equal("live"))
        })

        It("should be has field endpoint", func() {
            var configure = Configure{ Endpoint: "https://api.sandbox.paypal.com" }

            Expect(configure.Endpoint).To(Equal("https://api.sandbox.paypal.com"))
        })
    })

    Describe("Endpoint", func() {
        It("should be return https://api.paypal.com when set live mode", func() {
            var api *Api = NewApi(Configure { Mode: "live" })

            Expect(api.Endpoint()).To(Equal("https://api.paypal.com"))
        })

        It("should be return https://api.sandbox.paypal.com when set sandbox mode", func() {
            var api *Api = NewApi(Configure { Mode: "sandbox" })

            Expect(api.Endpoint()).To(Equal("https://api.sandbox.paypal.com"))
        })

        It("should be return https://custom-endpoint.paypal.com when set custum endpoint", func() {
            var api *Api = NewApi(Configure { Endpoint: "https://custom-endpoint.paypal.com" })

            Expect(api.Endpoint()).To(Equal("https://custom-endpoint.paypal.com"))
        })
    })
})
