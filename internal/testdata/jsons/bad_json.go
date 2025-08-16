package jsons

// Набор намеренно невалидных JSON-пэйлоадов
var BadJSONs = []string{
	//NIL
	``,

	//NO ORDERS
	`{
   "order_uid": "b563feb7b2b84b6testaaa",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`,

	// NO DELIVERY
	`{
   "order_uid": "b563feb7b2b84b6testbbb",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btest",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`,

	// NO DELIVERY + ITEMS
	`{
   "order_uid": "b563feb7b2b84b6testccc",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`,

	// NO PAYMENTS
	`{
   "order_uid": "b563feb7b2b84b6testddd",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btest",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`,

	// ORDER_UID REPEATS
	`{
   "order_uid": "b563feb7b2b84b6test",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "payment": {
      "transaction": "b563feb7b2b84b6test",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btest",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}`,
	// CURRENCY INVALID
	`{
  "order_uid": "ORD-BAD-0001",
  "track_number": "WBUS-NYC-BAD1",
  "entry": "WBIL",
  "delivery": {
    "name": "John Doe",
    "phone": "+12025550100",
    "zip": "10007",
    "city": "New York",
    "address": "200 Broadway",
    "region": "NY",
    "email": "john.doe@example.com"
  },
  "payment": {
    "transaction": "tx-bad-001",
    "request_id": "",
    "currency": "USDAAA",
    "provider": "wbpay",
    "amount": 1700,
    "payment_dt": 1715325331,
    "bank": "chase",
    "delivery_cost": 500,
    "goods_total": 1200,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500001,
      "track_number": "WBUS-NYC-BAD1",
      "price": 1200,
      "rid": "rid-bad-001",
      "name": "Desk Lamp",
      "sale": 0,
      "size": "std",
      "total_price": 1200,
      "nm_id": 910999,
      "brand": "BrightNest",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0001",
  "delivery_service": "ups",
  "shardkey": "1",
  "sm_id": 1,
  "date_created": "2024-05-10T09:15:31Z",
  "oof_shard": "1"
}`,

	// NO ORDER_UID
	`{
  "track_number": "WBUS-BOS-BAD2",
  "entry": "WBIL",
  "delivery": {
    "name": "Emma Clark",
    "phone": "+16175550123",
    "zip": "02108",
    "city": "Boston",
    "address": "24 Beacon St",
    "region": "MA",
    "email": "emma.clark@example.com"
  },
  "payment": {
    "transaction": "tx-bad-002",
    "request_id": "",
    "currency": "USD",
    "provider": "stripe",
    "amount": 900,
    "payment_dt": 1717000000,
    "bank": "boa",
    "delivery_cost": 200,
    "goods_total": 700,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500002,
      "track_number": "WBUS-BOS-BAD2",
      "price": 700,
      "rid": "rid-bad-002",
      "name": "Bluetooth Speaker",
      "sale": 0,
      "size": "std",
      "total_price": 700,
      "nm_id": 911000,
      "brand": "HomeWave",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0002",
  "delivery_service": "fedex",
  "shardkey": "2",
  "sm_id": 2,
  "date_created": "2024-06-01T12:00:00Z",
  "oof_shard": "2"
}`,

	// NO ITEMS
	`{
  "order_uid": "ORD-BAD-0003",
  "track_number": "WBUS-SEA-BAD3",
  "entry": "WBIL",
  "delivery": {
    "name": "Liam White",
    "phone": "+12065550123",
    "zip": "98101",
    "city": "Seattle",
    "address": "600 4th Ave",
    "region": "WA",
    "email": "liam.white@example.com"
  },
  "payment": {
    "transaction": "tx-bad-003",
    "request_id": "",
    "currency": "USD",
    "provider": "paypal",
    "amount": 400,
    "payment_dt": 1718000000,
    "bank": "wellsfargo",
    "delivery_cost": 50,
    "goods_total": 350,
    "custom_fee": 0
  },
  "items": [],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0003",
  "delivery_service": "usps",
  "shardkey": "3",
  "sm_id": 3,
  "date_created": "2024-06-10T08:00:00Z",
  "oof_shard": "3"
}`,

	// INVALID ITEMS
	`{
  "order_uid": "ORD-BAD-0004",
  "track_number": "WBGB-LON-BAD4",
  "entry": "WBIL",
  "delivery": {
    "name": "Oliver Smith",
    "phone": "+442079460111",
    "zip": "SW1A 1AA",
    "city": "London",
    "address": "The Mall",
    "region": "England",
    "email": "oliver.smith@example.co.uk"
  },
  "payment": {
    "transaction": "tx-bad-004",
    "request_id": "req-bad-004",
    "currency": "GBP",
    "provider": "wbpay",
    "amount": 810,
    "payment_dt": 1719000000,
    "bank": "hsbc",
    "delivery_cost": 110,
    "goods_total": 700,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500004,
      "track_number": "WBGB-LON-BAD4",
      "price": 700,
      "rid": "rid-bad-004",
      "name": "Cordless Vacuum",
      "sale": 150,
      "size": "",
      "total_price": 700,
      "nm_id": 911002,
      "brand": "CleanJet",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0004",
  "delivery_service": "royalmail",
  "shardkey": "4",
  "sm_id": 4,
  "date_created": "2024-06-20T10:00:00Z",
  "oof_shard": "4"
}`,

	// INVALID ITEMS
	`{
  "order_uid": "ORD-BAD-0005",
  "track_number": "WBUS-SF-0005",
  "entry": "WBIL",
  "delivery": {
    "name": "Sophia Lee",
    "phone": "+14155550123",
    "zip": "94102",
    "city": "San Francisco",
    "address": "1 Dr Carlton B Goodlett Pl",
    "region": "CA",
    "email": "sophia.lee@example.com"
  },
  "payment": {
    "transaction": "tx-bad-005",
    "request_id": "",
    "currency": "USD",
    "provider": "stripe",
    "amount": 1500,
    "payment_dt": 1720000000,
    "bank": "chase",
    "delivery_cost": 200,
    "goods_total": 1300,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500005,
      "track_number": "MISMATCH-TRACK",
      "price": 1300,
      "rid": "rid-bad-005",
      "name": "Coffee Grinder",
      "sale": 0,
      "size": "std",
      "total_price": 1300,
      "nm_id": 911003,
      "brand": "CaffèMax",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0005",
  "delivery_service": "fedex",
  "shardkey": "5",
  "sm_id": 5,
  "date_created": "2024-06-30T09:00:00Z",
  "oof_shard": "5"
}`,

	//	INVALID PAYMENT
	`{
  "order_uid": "ORD-BAD-0006",
  "track_number": "WBUS-ATX-BAD6",
  "entry": "WBIL",
  "delivery": {
    "name": "Charlotte Nguyen",
    "phone": "+15125550123",
    "zip": "78701",
    "city": "Austin",
    "address": "1100 Congress Ave",
    "region": "TX",
    "email": "charlotte.nguyen@example.com"
  },
  "payment": {
    "transaction": "tx-bad-006",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": -10,
    "payment_dt": 1721000000,
    "bank": "wellsfargo",
    "delivery_cost": -5,
    "goods_total": -1,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500006,
      "track_number": "WBUS-ATX-BAD6",
      "price": 999,
      "rid": "rid-bad-006",
      "name": "Webcam",
      "sale": 0,
      "size": "std",
      "total_price": 999,
      "nm_id": 911004,
      "brand": "VisionPro",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0006",
  "delivery_service": "usps",
  "shardkey": "6",
  "sm_id": 6,
  "date_created": "2024-07-05T08:00:00Z",
  "oof_shard": "6"
}`,

	// INVALID PAYMENT_DT
	`{
  "order_uid": "ORD-BAD-0007",
  "track_number": "WBUS-CHI-BAD7",
  "entry": "WBIL",
  "delivery": {
    "name": "Ava Robinson",
    "phone": "+13125550123",
    "zip": "60601",
    "city": "Chicago",
    "address": "233 S Wacker Dr",
    "region": "IL",
    "email": "ava.robinson@example.com"
  },
  "payment": {
    "transaction": "tx-bad-007",
    "request_id": "",
    "currency": "USD",
    "provider": "stripe",
    "amount": 1200,
    "payment_dt": 0,
    "bank": "chase",
    "delivery_cost": 200,
    "goods_total": 1000,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500007,
      "track_number": "WBUS-CHI-BAD7",
      "price": 1000,
      "rid": "rid-bad-007",
      "name": "Mechanical Keyboard",
      "sale": 0,
      "size": "TKL",
      "total_price": 1000,
      "nm_id": 911005,
      "brand": "KeyForge",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0007",
  "delivery_service": "ups",
  "shardkey": "7",
  "sm_id": -5,
  "date_created": "2024-07-10T10:00:00Z",
  "oof_shard": "7"
}`,

	// INVALID track_number (>32) + internal_signature (>64)
	`{
  "order_uid": "ORD-BAD-0008",
  "track_number": "THIS-IS-A-VERY-LONG-TRACK-NUMBER-OVER-32-CHARS",
  "entry": "WBIL",
  "delivery": {
    "name": "Mia Thompson",
    "phone": "+12125550123",
    "zip": "11201",
    "city": "Brooklyn",
    "address": "209 Joralemon St",
    "region": "NY",
    "email": "mia.thompson@example.com"
  },
  "payment": {
    "transaction": "tx-bad-008",
    "request_id": "",
    "currency": "USD",
    "provider": "paypal",
    "amount": 1300,
    "payment_dt": 1722000000,
    "bank": "boa",
    "delivery_cost": 200,
    "goods_total": 1100,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500008,
      "track_number": "THIS-IS-A-VERY-LONG-TRACK-NUMBER-OVER-32-CHARS",
      "price": 1100,
      "rid": "rid-bad-008",
      "name": "Air Fryer",
      "sale": 0,
      "size": "5L",
      "total_price": 1100,
      "nm_id": 911006,
      "brand": "CookLite",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
  "customer_id": "cust-bad-0008",
  "delivery_service": "fedex",
  "shardkey": "8",
  "sm_id": 8,
  "date_created": "2024-07-15T12:00:00Z",
  "oof_shard": "8"
}`,

	// INVALID DecodeOrder
	`{
  "order_uid": "ORD-BAD-0009",
  "track_number": "WBGB-MAN-BAD9",
  "entry": "WBIL",
  "delivery": {
    "name": "Ethan Harris",
    "phone": "+441612345678",
    "zip": "M2 5DB",
    "city": "Manchester",
    "address": "1 King Street",
    "region": "England",
    "email": "ethan.harris@example.co.uk"
  },
  "payment": {
    "transaction": "tx-bad-009",
    "request_id": "",
    "currency": "GBP",
    "provider": "wbpay",
    "amount": 600,
    "payment_dt": 1672531201,
    "bank": "barclays",
    "delivery_cost": 60,
    "goods_total": 540,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500009,
      "track_number": "WBGB-MAN-BAD9",
      "price": 540,
      "rid": "rid-bad-009",
      "name": "Electric Kettle",
      "sale": 0,
      "size": "1.7L",
      "total_price": 540,
      "nm_id": 911007,
      "brand": "BoilQuick",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0009",
  "delivery_service": "royalmail",
  "shardkey": "9",
  "sm_id": 9,
  "date_created": "2023-01-01T00:00:00Z",
  "oof_shard": "9",
  "unknown_field": "should trigger decoder error"
}`,

	// INVALID DecodeOrder
	`{
  "order_uid": "ORD-BAD-0010",
  "track_number": "WBUS-MIA-BAD10",
  "entry": "WBIL",
  "delivery": {
    "name": "Noah Miller",
    "phone": "+13055550123",
    "zip": "33130",
    "city": "Miami",
    "address": "350 SW 1st Ave",
    "region": "FL",
    "email": "noah.miller@example.com"
  },
  "payment": {
    "transaction": "tx-bad-010",
    "request_id": "req-bad-010",
    "currency": "USD",
    "provider": "stripe",
    "amount": 950,
    "payment_dt": 1723000000,
    "bank": "chase",
    "delivery_cost": 150,
    "goods_total": 800,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 500010,
      "track_number": "WBUS-MIA-BAD10",
      "price": 800,
      "rid": "rid-bad-010",
      "name": "Toaster Oven",
      "sale": 0,
      "size": "std",
      "total_price": 800,
      "nm_id": 911008,
      "brand": "HeatBox",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-bad-0010",
  "delivery_service": "usps",
  "shardkey": "10",
  "sm_id": "110",
  "date_created": "2024-07-20T07:30:00Z",
  "oof_shard": "10"
}`,
}
