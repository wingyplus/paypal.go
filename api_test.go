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
    Describe("Endpoint", func() {
        It("should be return https://api.paypal.com when set live mode", func() {
            var newApi *Api = NewApi(map[string]string { "mode": "live" })

            Expect(newApi.Endpoint()).To(Equal("https://api.paypal.com"))
        })

        It("should be return https://api.sandbox.paypal.com when set sandbox mode", func() {
            var newApi *Api = NewApi(map[string]string { "mode": "sandbox" })

            Expect(newApi.Endpoint()).To(Equal("https://api.sandbox.paypal.com"))
        })

        It("should be return https://custom-endpoint.paypal.com when set custum endpoint", func() {
            var newApi *Api = NewApi(map[string]string { "endpoint": "https://custom-endpoint.paypal.com" })

            Expect(newApi.Endpoint()).To(Equal("https://custom-endpoint.paypal.com"))
        })
    })
})
