package payment_server

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/phayes/freeport"
	"github.com/teivah/payment-server/payment"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
	"net/http"
	"testing"
)

func TestCart(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment Suite")
}

var _ = Describe("Payment", func() {
	var (
		statusCode int
		payment    *swagger.PaymentWithId
		host       string
	)

	const (
		paymentsPath = "/v1/payments"
		paymentPath  = "/v1/payment"
		externalUri  = "localhost:8080"
	)

	Context("initialisation", func() {
		host = startServer()
	})

	Context("when posting an empty payment", func() {
		s, r, err := postPayment(&swagger.Payment{}, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		statusCode = s
		payment = r.Data
		It("has 201 status code", func() {
			Expect(statusCode).Should(Equal(http.StatusCreated))
		})
		It("has link matching its id", func() {
			checkLink(*payment, externalUri, paymentPath)
		})
	})

	Context("when posting a payment", func() {
		s, r, err := postPayment(&swagger.Payment{
			Version:        1,
			Type_:          "payment",
			OrganisationId: "organisation",
			Attributes: &swagger.PaymentAttributes{
				Amount: "10.0",
			},
		}, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		statusCode = s
		payment = r.Data
		fmt.Printf("%v\n", payment.Version)
		It("has 201 status code", func() {
			Expect(statusCode).Should(Equal(http.StatusCreated))
		})
		It("has link matching its id", func() {
			checkLink(*payment, externalUri, paymentPath)
		})
		It("has matching version", func() {
			Expect(int32(1)).Should(Equal(payment.Version))
		})
		It("has matching type", func() {
			Expect("payment").Should(Equal(payment.Type_))
		})
		It("has matching organisation", func() {
			Expect("organisation").Should(Equal(payment.OrganisationId))
		})
		It("has matching amount", func() {
			Expect("10.0").Should(Equal(payment.Attributes.Amount))
		})
	})
})

func startServer() string {
	p, err := freeport.GetFreePort()
	if err != nil {
		panic("Unable to get a free port")
	}
	host := fmt.Sprintf("http://localhost:%d", p)

	go func() {
		mux := payment.NewRouter()
		err = http.ListenAndServe(fmt.Sprintf(":%d", p), mux)
		panic(err)
	}()

	return host
}

func checkLink(payment swagger.PaymentWithId, externalUri, paymentPath string) {
	Expect(payment.Links.Self).Should(Equal(externalUri + paymentPath + "/" + payment.Id))
}

func postPayment(payment *swagger.Payment, host, paymentPath string) (int, *swagger.PaymentCreationResponse, error) {
	response := swagger.PaymentCreationResponse{}
	r, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&swagger.PaymentCreation{
			Data: payment,
		}).
		SetResult(&response).
		Post(host + paymentPath)
	if err != nil {
		utils.Logger.Error("Server error", zap.Error(err))
		return 0, nil, err
	}

	return r.StatusCode(), &response, nil
}
