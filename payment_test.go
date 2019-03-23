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

func TestPayment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Payment Suite")
}

var _ = Describe("Payment", func() {
	var (
		host string
	)

	const (
		paymentsPath   = "/v1/payments"
		paymentPath    = "/v1/payment"
		externalUri    = "localhost:8080"
		unknownPayment = "5c97320fa86e346013ee6489"
	)

	Context("initialisation", func() {
		host = startServer()
		payment.CleanCollection()
	})

	/*
		List payments
	*/
	Context("when posting 3 payments", func() {
		_, r1, err := postPayment(`{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		_, r2, err := postPayment(`{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		_, r3, err := postPayment(`{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		Context("listing all payments", func() {
			s, r, err := getPayments(host, paymentsPath)
			if err != nil {
				Fail("Payment error")
			}
			It("has a 201 response status code", func() {
				Expect(s).Should(Equal(http.StatusOK))
			})
			It("has 3 payments in total", func() {
				Expect(3).Should(Equal(len(r.Data)))
			})
			Context("when deleting these 3 payments", func() {
				_, err := deletePayment(r1.Data.Id, host, paymentPath)
				if err != nil {
					Fail("Payment error")
				}
				_, err = deletePayment(r2.Data.Id, host, paymentPath)
				if err != nil {
					Fail("Payment error")
				}
				_, err = deletePayment(r3.Data.Id, host, paymentPath)
				if err != nil {
					Fail("Payment error")
				}
				Context("listing all payments", func() {
					s, r, err := getPayments(host, paymentsPath)
					if err != nil {
						Fail("Payment error")
					}
					It("has a 201 response status code", func() {
						Expect(s).Should(Equal(http.StatusOK))
					})
					It("has 0 payment in total", func() {
						Expect(0).Should(Equal(len(r.Data)))
					})
				})
			})
		})
	})

	/*
		Nominal cases
	*/
	Context("when posting a payment", func() {
		s, r, err := postPayment(
			&swagger.PaymentCreation{
				Data: &swagger.Payment{
					Version:        1,
					Type_:          "payment",
					OrganisationId: "organisation",
					Attributes: &swagger.PaymentAttributes{
						Amount: "10.0",
					},
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
	Context("when posting a payment", func() {
		_, r, err := postPayment(&swagger.PaymentCreation{
			Data: &swagger.Payment{
				Version:        1,
				Type_:          "payment",
				OrganisationId: "organisation",
				Attributes: &swagger.PaymentAttributes{
					Amount: "10.0",
				},
			},
		}, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		Context("when updating the amount", func() {
			s, r, err := putPayment(&swagger.PaymentCreation{
				Data: &swagger.Payment{
					Version:        1,
					Type_:          "payment",
					OrganisationId: "organisation",
					Attributes: &swagger.PaymentAttributes{
						Amount: "11.0",
					},
				},
			}, r.Data.Id, host, paymentPath)
			if err != nil {
				Fail("Payment error")
			}
			It("has a 201 response status code", func() {
				Expect(s).Should(Equal(http.StatusCreated))
			})
			It("has the updated amount", func() {
				Expect("11.0").Should(Equal(r.Data.Attributes.Amount))
			})
			Context("when getting this updated payment", func() {
				s, r, err := getPayment(r.Data.Id, host, paymentPath)
				if err != nil {
					Fail("Payment error")
				}
				It("has a 200 response status code", func() {
					Expect(s).Should(Equal(http.StatusOK))
				})
				It("has the updated amount", func() {
					Expect("11.0").Should(Equal(r.Attributes.Amount))
				})
			})
		})
	})
	Context("when posting a payment", func() {
		_, r, err := postPayment(&swagger.PaymentCreation{
			Data: &swagger.Payment{
				Version:        1,
				Type_:          "payment",
				OrganisationId: "organisation",
				Attributes: &swagger.PaymentAttributes{
					Amount: "10.0",
				},
			},
		}, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		Context("when deleting this payment", func() {
			s, _ := deletePayment(r.Data.Id, host, paymentPath)
			if err != nil {
				Fail("Payment error")
			}
			It("has a 204 response status code", func() {
				Expect(s).Should(Equal(http.StatusNoContent))
			})
			Context("when deleting again this payment", func() {
				s, _ := deletePayment(r.Data.Id, host, paymentPath)
				if err != nil {
					Fail("Payment error")
				}
				It("has a 404 response status code", func() {
					Expect(s).Should(Equal(http.StatusNotFound))
				})
			})
		})
	})

	/*
		Invalid payment
	*/
	Context("when posting an invalid payment", func() {
		s, _, err := postPayment(`{{}`, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
		})
	})
	Context("when putting an invalid payment", func() {
		s, _, err := putPayment(`{{}`, unknownPayment, host, paymentPath)
		if err != nil {
			Fail("Payment error")
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
		})
	})

	/*
		Unknown payment
	*/
	Context("when getting an unknown payment", func() {
		s, _, err := getPayment(unknownPayment, host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 404 response status code", func() {
			Expect(s).Should(Equal(http.StatusNotFound))
		})
	})
	Context("when putting an unknown payment", func() {
		s, _, err := putPayment(`{}`, unknownPayment, host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 404 response status code", func() {
			Expect(s).Should(Equal(http.StatusNotFound))
		})
	})
	Context("when deleting an unknown payment", func() {
		s, err := deletePayment(unknownPayment, host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 404 response status code", func() {
			Expect(s).Should(Equal(http.StatusNotFound))
		})
	})

	/*
		Invalid identifier
	*/
	Context("when getting a payment with an invalid id", func() {
		s, _, err := getPayment("x", host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
		})
	})
	Context("when putting a payment with an invalid id", func() {
		s, _, err := putPayment(`{}`, "x", host, paymentPath)
		if err != nil {
			Fail("Server error:" + err.Error())
		}
		It("has a 400 response status code", func() {
			Expect(s).Should(Equal(http.StatusBadRequest))
		})
	})
	Context("when deleting a payment with an invalid id", func() {
		s, err := deletePayment("x", host, paymentPath)
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

func checkLink(payment swagger.PaymentWithId, externalUri, paymentPath string) {
	Expect(payment.Links.Self).Should(Equal(externalUri + paymentPath + "/" + payment.Id))
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

func getPayments(host, paymentsPath string) (int, *swagger.PaymentDetailsListResponse, error) {
	response := swagger.PaymentDetailsListResponse{}

	r, err := resty.R().
		SetResult(&response).
		Get(host + paymentsPath)

	if err != nil {
		return 0, nil, err
	}

	return r.RawResponse.StatusCode, &response, nil
}

func deletePayment(id, host, paymentPath string) (
	int, error) {
	response := swagger.PaymentUpdateResponse{}
	r, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&response).
		Delete(host + paymentPath + "/" + id)
	if err != nil {
		utils.Logger.Error("Server error", zap.Error(err))
		return 0, err
	}

	return r.RawResponse.StatusCode, nil
}

func putPayment(body interface{}, id, host, paymentPath string) (
	int, *swagger.PaymentUpdateResponse, error) {
	response := swagger.PaymentUpdateResponse{}
	r, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Put(host + paymentPath + "/" + id)
	if err != nil {
		utils.Logger.Error("Server error", zap.Error(err))
		return 0, nil, err
	}

	return r.RawResponse.StatusCode, &response, nil
}

func postPayment(body interface{}, host, paymentPath string) (
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
