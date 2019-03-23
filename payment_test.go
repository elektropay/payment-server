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
		//statusCode int
		//payment    *swagger.PaymentWithId
		host string
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
		s, r, err := postPaymentWithBody(`{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		It("has a 201 response status code", func() {
			Expect(s).Should(Equal(http.StatusCreated))
		})
		It("can be retrieved with a get request", func() {
			checkPaymentDoesExist(r.Data.Id, host, paymentPath)
		})
		It("has link matching its id", func() {
			checkLink(*r.Data, externalUri, paymentPath)
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
		It("has 201 response status code", func() {
			Expect(s).Should(Equal(http.StatusCreated))
		})
		It("can be retrieved with a get request", func() {
			checkPaymentDoesExist(r.Data.Id, host, paymentPath)
		})
		It("has link matching its id", func() {
			checkLink(*r.Data, externalUri, paymentPath)
		})
		It("has matching version", func() {
			Expect(int32(1)).Should(Equal(r.Data.Version))
		})
		It("has matching type", func() {
			Expect("payment").Should(Equal(r.Data.Type_))
		})
		It("has matching organisation", func() {
			Expect("organisation").Should(Equal(r.Data.OrganisationId))
		})
		It("has matching amount", func() {
			Expect("10.0").Should(Equal(r.Data.Attributes.Amount))
		})
	})

	Context("when posting an invalid payment", func() {
		s, _, err := postPaymentWithBody(`{{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
		})
	})

	Context("when getting an unknown payment", func() {
		s, _, err := getPayment("5c97320fa86e346013ee6489", host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 404 response status code", func() {
			Expect(s).Should(Equal(http.StatusNotFound))
		})
	})

	Context("when getting a payment with an invalid id", func() {
		s, _, err := getPayment("x", host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
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

func checkPaymentDoesExist(id, host, paymentPath string) {
	statusCode, _, err := getPayment(id, host, paymentPath)
	if err != nil {
		Fail("Server error: " + err.Error())
		return
	}
	Expect(http.StatusOK).Should(Equal(statusCode))
}

func getPayment(id, host, paymentPath string) (int, *swagger.PaymentWithId, error) {
	response := swagger.PaymentWithId{}

	r, err := resty.R().
		SetResult(&response).
		Get(host + paymentPath + "/" + id)

	if err != nil {
		return 0, nil, err
	}

	return r.RawResponse.StatusCode, &response, nil
}

func checkLink(payment swagger.PaymentWithId, externalUri, paymentPath string) {
	Expect(payment.Links.Self).Should(Equal(externalUri + paymentPath + "/" + payment.Id))
}

func postPayment(payment *swagger.Payment, host, paymentPath string) (
	int, *swagger.PaymentCreationResponse, error) {
	return postPaymentWithBody(&swagger.PaymentCreation{
		Data: payment,
	}, host, paymentPath)
}

func postPaymentWithBody(body interface{}, host, paymentPath string) (
	int, *swagger.PaymentCreationResponse, error) {
	response := swagger.PaymentCreationResponse{}
	r, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(host + paymentPath)
	if err != nil {
		utils.Logger.Error("Server error", zap.Error(err))
		return 0, nil, err
	}

	return r.RawResponse.StatusCode, &response, nil
}
