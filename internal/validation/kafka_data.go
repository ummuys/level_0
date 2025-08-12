package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ummuys/level_0/internal/models"
)

func DecodeOrder(b []byte) (models.OrderData, error) {
	var od models.OrderData
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&od); err != nil {
		return models.OrderData{}, fmt.Errorf("bad json: %w", err)
	}
	return od, nil
}

func Validate(od models.OrderData) error {
	var sErr []string

	if od.OrderUID == "" {
		sErr = append(sErr, "order_uid is required")
	}
	if od.TrackNumber == "" {
		sErr = append(sErr, "track_number is required")
	}
	if od.CustomerID == "" {
		sErr = append(sErr, "customer_id is required")
	}
	if od.Payment.Currency == "" {
		sErr = append(sErr, "payment.currency is required")
	}
	if od.DateCreated.IsZero() {
		sErr = append(sErr, "date_created is required")
	}
	if len(od.Items) == 0 {
		sErr = append(sErr, "items must be non-empty")
	}

	if od.Payment.Amount < 0 || od.Payment.GoodsTotal < 0 || od.Payment.DeliveryCost < 0 || od.Payment.CustomFee < 0 {
		sErr = append(sErr, "amounts must be >= 0")
	}
	if od.SmID < 0 {
		sErr = append(sErr, "sm_id must be >= 0")
	}

	now := time.Now().Add(5 * time.Minute)
	if od.DateCreated.After(now) {
		sErr = append(sErr, "date_created is incorrect")
	}
	if od.Payment.PaymentDT > now.Unix() {
		sErr = append(sErr, "payment_dt is incorrect")
	}

	var goodsTotal int
	for i, it := range od.Items {
		if it.Price < 0 || it.TotalPrice < 0 || it.Sale < 0 {
			sErr = append(sErr, fmt.Sprintf("items[%d] amounts must be >= 0", i))
		}
		if it.TrackNumber != "" && it.TrackNumber != od.TrackNumber {
			sErr = append(sErr, fmt.Sprintf("items[%d].track_number != order.track_number", i))
		}
		goodsTotal += it.TotalPrice
	}
	if goodsTotal != od.Payment.GoodsTotal {
		sErr = append(sErr, fmt.Sprintf("goods_total mismatch: items=%d payment=%d", goodsTotal, od.Payment.GoodsTotal))
	}
	if od.Payment.GoodsTotal+od.Payment.DeliveryCost+od.Payment.CustomFee != od.Payment.Amount {
		sErr = append(sErr, "payment.amount != goods_total + delivery_cost + custom_fee")
	}

	if len(sErr) > 0 {
		return fmt.Errorf(strings.Join(sErr, ", "))
	}
	return nil
}
