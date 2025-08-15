package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

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

func ValidateOrder(od models.OrderData) (string, error) {
	var errs []string
	add := func(path, format string, args ...any) {
		errs = append(errs, fmt.Sprintf("%s: %s", path, fmt.Sprintf(format, args...)))
	}

	// INFO
	if od.OrderUID == "" {
		add("order_uid", "required field")
	} else if l := len(od.OrderUID); l > 32 {
		add("order_uid", "must be <= 32 characters (got %d)", l)
	}

	if od.TrackNumber == "" {
		add("track_number", "required field")
	} else if l := len(od.TrackNumber); l > 32 {
		add("track_number", "must be <= 32 characters (got %d)", l)
	}

	if od.Entry == "" {
		add("entry", "required field")
	} else if l := len(od.Entry); l > 16 {
		add("entry", "must be <= 16 characters (got %d)", l)
	}

	if od.Locale == "" {
		add("locale", "required field")
	} else if l := len(od.Locale); l > 10 {
		add("locale", "must be <= 10 characters (got %d)", l)
	}

	if l := len(od.InternalSignature); l > 512 {
		add("internal_signature", "must be <= 512 characters (got %d)", l)
	}

	if od.CustomerID == "" {
		add("customer_id", "required field")
	} else if l := len(od.CustomerID); l > 64 {
		add("customer_id", "must be <= 64 characters (got %d)", l)
	}

	if od.DeliveryService == "" {
		add("delivery_service", "required field")
	} else if l := len(od.DeliveryService); l > 32 {
		add("delivery_service", "must be <= 32 characters (got %d)", l)
	}

	if od.ShardKey == "" {
		add("shardkey", "required field")
	} else if l := len(od.ShardKey); l > 8 {
		add("shardkey", "must be <= 8 characters (got %d)", l)
	}

	if od.OofShard == "" {
		add("oof_shard", "required field")
	} else if l := len(od.OofShard); l > 8 {
		add("oof_shard", "must be <= 8 characters (got %d)", l)
	}

	if l := len(od.InternalSignature); l > 64 {
		add("internal_signature", "must be <= 64 characters (got %d)", l)
	}

	if od.SmID < 0 {
		add("sm_id", "must be >= 0")
	}

	if od.DateCreated.IsZero() {
		add("date_created", "required field")
	}

	// DELIVERY
	if strings.TrimSpace(od.Delivery.Name) == "" {
		add("delivery.name", "required field")
	} else if l := len(od.Delivery.Name); l > 200 {
		add("delivery.name", "must be <= 200 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.Phone) == "" {
		add("delivery.phone", "required field")
	} else if l := len(od.Delivery.Phone); l > 32 {
		add("delivery.phone", "must be <= 32 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.Zip) == "" {
		add("delivery.zip", "required field")
	} else if l := len(od.Delivery.Zip); l > 12 {
		add("delivery.zip", "must be <= 12 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.City) == "" {
		add("delivery.city", "required field")
	} else if l := len(od.Delivery.City); l > 120 {
		add("delivery.city", "must be <= 120 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.Address) == "" {
		add("delivery.address", "required field")
	} else if l := len(od.Delivery.Address); l > 255 {
		add("delivery.address", "must be <= 255 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.Region) == "" {
		add("delivery.region", "required field")
	} else if l := len(od.Delivery.Region); l > 120 {
		add("delivery.region", "must be <= 120 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Delivery.Email) == "" {
		add("delivery.email", "required field")
	} else if l := len(od.Delivery.Email); l > 254 {
		add("delivery.email", "must be <= 254 characters (got %d)", l)
	}

	//  PAYMENT
	if strings.TrimSpace(od.Payment.Transaction) == "" {
		add("payment.transaction", "required field")
	} else if l := len(od.Payment.Transaction); l > 128 {
		add("payment.transaction", "must be <= 128 characters (got %d)", l)
	}

	if l := len(od.Payment.RequestID); l > 128 {
		add("payment.request_id", "must be <= 128 characters (got %d)", l)
	}

	if od.Payment.Currency == "" {
		add("payment.currency", "required field")
	} else if l := len(od.Payment.Currency); l != 3 {
		add("payment.currency", "must be exactly 3 characters (got %d)", l)
	}

	if strings.TrimSpace(od.Payment.Provider) == "" {
		add("payment.provider", "required field")
	} else if l := len(od.Payment.Provider); l > 32 {
		add("payment.provider", "must be <= 32 characters (got %d)", l)
	}

	if l := len(od.Payment.Bank); l > 64 {
		add("payment.bank", "must be <= 64 characters (got %d)", l)
	}

	if od.Payment.Amount < 0 || od.Payment.GoodsTotal < 0 || od.Payment.DeliveryCost < 0 || od.Payment.CustomFee < 0 {
		add("payment", "amount fields must be >= 0")
	}

	if od.Payment.PaymentDT == 0 {
		add("payment.payment_dt", "invalid payment_dt")
	}

	// ITEMS
	if len(od.Items) == 0 {
		add("items", "must contain at least one item")
	}

	for i, it := range od.Items {
		if it.ChrtID <= 0 {
			add(fmt.Sprintf("items[%d].chrt_id", i), "must be > 0")
		}
		if it.NmID <= 0 {
			add(fmt.Sprintf("items[%d].nm_id", i), "must be > 0")
		}
		if it.Status < 0 {
			add(fmt.Sprintf("items[%d].status", i), "must be >= 0")
		}

		if it.Price < 0 {
			add(fmt.Sprintf("items[%d].price", i), "must be >= 0")
		}
		if it.TotalPrice < 0 {
			add(fmt.Sprintf("items[%d].total_price", i), "must be >= 0")
		}
		if it.Sale < 0 || it.Sale > 100 {
			add(fmt.Sprintf("items[%d].sale", i), "must be in range 0..100")
		}

		if it.TrackNumber != "" && len(it.TrackNumber) > 32 {
			add(fmt.Sprintf("items[%d].track_number", i), "must be <= 32 characters")
		}
		if it.TrackNumber != "" && od.TrackNumber != "" && it.TrackNumber != od.TrackNumber {
			add(fmt.Sprintf("items[%d].track_number", i), "must match order.track_number (%s)", od.TrackNumber)
		}

		if it.Name == "" {
			add(fmt.Sprintf("items[%d].name", i), "required field")
		} else if len(it.Name) > 255 {
			add(fmt.Sprintf("items[%d].name", i), "must be <= 255 characters")
		}

		if it.Size == "" {
			add(fmt.Sprintf("items[%d].size", i), "required field")
		} else if len(it.Size) > 32 {
			add(fmt.Sprintf("items[%d].size", i), "must be <= 32 characters")
		}

		if strings.TrimSpace(it.RID) == "" {
			add(fmt.Sprintf("items[%d].rid", i), "required field")
		} else if len(it.RID) > 128 {
			add(fmt.Sprintf("items[%d].rid", i), "must be <= 128 characters")
		}

		if len(it.Brand) > 128 {
			add(fmt.Sprintf("items[%d].brand", i), "must be <= 128 characters")
		}

	}

	if len(errs) > 0 {
		return "", fmt.Errorf(strings.Join(errs, "; "))
	}
	return od.OrderUID, nil
}

func Convert(order models.OrderDataDb) models.OrderData {
	out := models.OrderData{
		OrderUID:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.ShardKey,
		SmID:              order.SmID,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,

		Delivery: order.DeliveryData,
		Payment:  order.PaymentData,
	}

	if order.Items != nil {
		items := make([]models.ItemData, len(order.Items))
		copy(items, order.Items)
		out.Items = items
	} else {
		out.Items = []models.ItemData{}
	}

	return out

}

func ConvertSlice(orders []models.OrderDataDb) []models.OrderData {
	out := make([]models.OrderData, len(orders))
	for i, o := range orders {
		out[i] = Convert(o)
	}
	return out
}
