package data_generator

import (
	"fmt"
	"math/rand"
)

const (
	CountOfMerchants = 1000
	MinCreatedAt = 1578765017
	MaxCreatedAt = 1591897817
)

var (
	MerchantIds []string
)

func GetRowData() map[string]interface{} {
	var methods []string
	var status []string

	methods = append(methods, "card")
	methods = append(methods, "netbanking")
	methods = append(methods, "wallet")
	methods = append(methods, "upi")

	status = append(status, "created")
	status = append(status, "captured")
	status = append(status, "authorized")
	status = append(status, "failed")

	return map[string]interface{}{
		"id":          fmt.Sprintf("pay_%s", GenerateID()),
		"merchant_id": MerchantIds[rand.Intn(CountOfMerchants-1)],
		"created_at":  rand.Intn(MaxCreatedAt - MinCreatedAt) + MinCreatedAt,
		"payment_data": map[string]interface{}{
			"amount":               rand.Intn(100000),
			"currency":             "INR",
			"base_amount":          rand.Intn(100000),
			"method":               methods[rand.Intn(len(methods)-1)],
			"status":               status[rand.Intn(len(status)-1)],
			"two_factor_auth":      "passed",
			"order_id":             fmt.Sprintf("order_%s", GenerateID()),
			"invoice_id":           nil,
			"transfer_id":          nil,
			"payment_link_id":      nil,
			"international":        false,
			"amount_authorized":    rand.Intn(100000),
			"amount_refunded":      0,
			"base_amount_refunded": 0,
			"amount_paidout":       0,
			"amount_transferred":   0,
			"refund_status":        nil,
			"description":          nil,
			"card_id":              fmt.Sprintf("card_%s", GenerateID()),
			"bank":                 nil,
			"wallet":               nil,
			"vpa":                  nil,
			"on_hold":              false,
			"on_hold_until":        nil,
			"emi_plan_id":          nil,
			"error_code":           nil,
			"internal_error_code":  nil,
			"error_description":    nil,
			"cancellation_reason":  nil,
			"global_customer_id":   nil,
			"receiver_id":          nil,
			"app_token":            nil,
			"emi_subvention":       nil,
			"auth_type":            "3ds",
			"acknowledged_at":      nil,
			"verify_at":            rand.Intn(MaxCreatedAt - MinCreatedAt) + MinCreatedAt,
			"refund_at":            nil,
			"reference13":          nil,
			"settled_by":           "hdfc",
			"reference16":          nil,
			"global_token_id":      nil,
			"email":                "test@gmail.com",
			"contact":              "+324332233334",
			"transaction_id":         "transaction_id",
			"authorized_at":          rand.Intn(MaxCreatedAt - MinCreatedAt) + MinCreatedAt,
			"auto_captured":          true,
			"gateway_captured":       true,
			"captured_at":            rand.Intn(MaxCreatedAt - MinCreatedAt) + MinCreatedAt,
			"gateway":                "hdfc",
			"terminal_id":            "terminal_id",
			"authentication_gateway": nil,
			"reference1":             nil,
			"reference2":             "011679",
			"cps_route":              2,
			"batch_id":               nil,
			"receiver_type":          nil,
			"signed":                 false,
			"verified":               nil,
			"verify_bucket":          0,
			"callback_url":           "https://api.test.in/v2/pay/response/com.test/callback",
			"fee":                    0,
			"mdr":                    0,
			"tax":                    0,
			"otp_attempts":           nil,
			"otp_count":              nil,
			"recurring":              false,
			"recurring_type":         nil,
			"save":                   false,
			"late_authorized":        false,
			"convert_currency":       nil,
			"disputed":               false,
			"updated_at":             rand.Intn(MaxCreatedAt - MinCreatedAt) + MinCreatedAt,
			"public_id":              fmt.Sprintf("pay_%s", GenerateID()),
			"captured":               true,
			"entity":       "payment",
			"fee_bearer":   "platform",
			"error_source": nil,
			"error_step":   nil,
			"error_reason": nil,
			"admin":        true,
			"mode":         "live",
			"notes": map[string]interface{}{
				"txn_uuid":       "txn_uuid",
				"transaction_id": "transaction_id",
			},
			"acquirer_data": map[string]interface{}{
				"auth_code": "000233",
			},
		},
	}
}

func GenerateMerchantIds() {
	for i := 0; i < CountOfMerchants; i++ {
		MerchantIds = append(MerchantIds, GenerateID())
	}
}
