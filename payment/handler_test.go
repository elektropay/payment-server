package payment

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teivah/payment-server/swagger"
	"gopkg.in/mgo.v2/bson"
)

func TestDecodeRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	body := ioutil.NopCloser(bytes.NewReader([]byte(`
	{
        "attributes": {
            "amount": "26.00"
		}
	}
	`)))
	request := http.Request{
		Body: body,
	}

	payment := swagger.Payment{}
	err := decodeRequest(&payment, recorder, &request)
	assert.Nil(t, err)
	assert.Equal(t, "26.00", payment.Attributes.Amount)
}

func TestDecodeInvalidRequest(t *testing.T) {
	recorder := httptest.NewRecorder()
	body := ioutil.NopCloser(bytes.NewReader([]byte(`
	{
        "attributes": {
            "amount": "26.00
		}
	}
	`)))
	request := http.Request{
		Body: body,
	}

	payment := swagger.Payment{}
	err := decodeRequest(&payment, recorder, &request)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.NotEqual(t, 0, recorder.Body.Len())
}

func TestCreatePaymentResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	createPaymentResponse(recorder, http.StatusTeapot, "id", "uri", &swagger.Payment{
		Attributes: &swagger.PaymentAttributes{
			Amount: "26.00",
		},
	}, func(id *swagger.PaymentWithId) interface{} {
		return swagger.PaymentCreationResponse{
			Data: id,
		}
	})
	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.NotEqual(t, 0, recorder.Body.Len())

	response := swagger.PaymentCreationResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, "26.00", response.Data.Attributes.Amount)
	assert.Equal(t, "id", response.Data.Id)
	assert.Equal(t, "uri/v1/payment/id", response.Data.Links.Self)
}

func TestCreatePaymentResponseWithNilResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	createPaymentResponse(recorder, http.StatusTeapot, "id", "uri", nil, nil)
	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, 0, recorder.Body.Len())
}

func TestCreatePaymentsResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	id1 := bson.NewObjectId()
	id2 := bson.NewObjectId()
	createPaymentsResponse(recorder, http.StatusTeapot, "uri", []Envelope{
		{
			Id: id1,
			Payment: &swagger.Payment{
				Attributes: &swagger.PaymentAttributes{
					Amount: "1",
				},
			},
		},
		{
			Id: id2,
			Payment: &swagger.Payment{
				Attributes: &swagger.PaymentAttributes{
					Amount: "2",
				},
			},
		},
	})

	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.NotEqual(t, 0, recorder.Body.Len())

	response := swagger.PaymentDetailsListResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(response.Data))
	assert.Equal(t, "uri/v1/payments", response.Links.Self)
	assert.Equal(t, "uri/v1/payment/"+id1.Hex(), response.Data[0].Links.Self)
	assert.Equal(t, "1", response.Data[0].Attributes.Amount)
	assert.Equal(t, "uri/v1/payment/"+id2.Hex(), response.Data[1].Links.Self)
	assert.Equal(t, "2", response.Data[1].Attributes.Amount)
}

func TestCreateErrorResponse(t *testing.T) {
	recorder := httptest.NewRecorder()
	createErrorResponse(recorder, http.StatusTeapot, &swagger.ApiError{
		ErrorMessage: "message",
		ErrorCode:    "code",
	})
	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, `{"error_message":"message","error_code":"code"}`, recorder.Body.String())
}

func TestCreateErrorResponseWithNilError(t *testing.T) {
	recorder := httptest.NewRecorder()
	createErrorResponse(recorder, http.StatusTeapot, nil)
	assert.Equal(t, http.StatusTeapot, recorder.Code)
	assert.Equal(t, 0, recorder.Body.Len())
}

func TestToBsonObjectId(t *testing.T) {
	hex, err := toBsonObjectId("5c97320fa86e346013ee6489")
	assert.Nil(t, err)
	assert.NotNil(t, hex)
}

func TestToBsonObjectIdWithInvalidId(t *testing.T) {
	_, err := toBsonObjectId("x")
	assert.NotNil(t, err)
}
