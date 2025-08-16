package jsons

var GoodJSONs = []string{
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
	`{
  "order_uid": "ORD-0002-LA77CA",
  "track_number": "WBUS-LA-0002",
  "entry": "WBIL",
  "delivery": {
    "name": "Ben Turner",
    "phone": "+12135550123",
    "zip": "90012",
    "city": "Los Angeles",
    "address": "200 N Spring St",
    "region": "CA",
    "email": "ben.turner@example.com"
  },
  "payment": {
    "transaction": "tx-002-us",
    "request_id": "req-002",
    "currency": "USD",
    "provider": "stripe",
    "amount": 3000,
    "payment_dt": 1723756800,
    "bank": "boa",
    "delivery_cost": 800,
    "goods_total": 2200,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 6022201,
      "track_number": "WBUS-LA-0002",
      "price": 1200,
      "rid": "rid-002-a",
      "name": "Gaming Mouse",
      "sale": 0,
      "size": "std",
      "total_price": 1200,
      "nm_id": 920020,
      "brand": "ClickPro",
      "status": 202
    },
    {
      "chrt_id": 6022202,
      "track_number": "WBUS-LA-0002",
      "price": 1000,
      "rid": "rid-002-b",
      "name": "Laptop Stand",
      "sale": 0,
      "size": "std",
      "total_price": 1000,
      "nm_id": 920021,
      "brand": "DeskLift",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "sig-us-2",
  "customer_id": "cust-2002",
  "delivery_service": "fedex",
  "shardkey": "2",
  "sm_id": 102,
  "date_created": "2024-08-16T10:00:00Z",
  "oof_shard": "2"
}`,

	`{
  "order_uid": "ORD-0003-LN55UK",
  "track_number": "WBGB-LON-0003",
  "entry": "WBIL",
  "delivery": {
    "name": "James Walker",
    "phone": "+442079460123",
    "zip": "EC2A 2FA",
    "city": "London",
    "address": "25 City Road",
    "region": "England",
    "email": "james.walker@example.co.uk"
  },
  "payment": {
    "transaction": "tx-003-gb",
    "request_id": "",
    "currency": "GBP",
    "provider": "paypal",
    "amount": 750,
    "payment_dt": 1711920000,
    "bank": "natwest",
    "delivery_cost": 100,
    "goods_total": 650,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 7033401,
      "track_number": "WBGB-LON-0003",
      "price": 650,
      "rid": "rid-003-a",
      "name": "LED Desk Lamp",
      "sale": 0,
      "size": "std",
      "total_price": 650,
      "nm_id": 930040,
      "brand": "BrightNest",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-3003",
  "delivery_service": "royalmail",
  "shardkey": "3",
  "sm_id": 103,
  "date_created": "2024-04-01T09:30:00Z",
  "oof_shard": "3"
}`,

	`{
  "order_uid": "ORD-0004-CH21IL",
  "track_number": "WBUS-CHI-0004",
  "entry": "WBIL",
  "delivery": {
    "name": "Sophia Martinez",
    "phone": "+13125550111",
    "zip": "60601",
    "city": "Chicago",
    "address": "233 S Wacker Dr",
    "region": "IL",
    "email": "sophia.martinez@example.com"
  },
  "payment": {
    "transaction": "tx-004-us",
    "request_id": "req-004",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 2000,
    "payment_dt": 1706918400,
    "bank": "chase",
    "delivery_cost": 200,
    "goods_total": 1800,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 8044501,
      "track_number": "WBUS-CHI-0004",
      "price": 900,
      "rid": "rid-004-a",
      "name": "Noise Cancelling Earbuds",
      "sale": 0,
      "size": "std",
      "total_price": 900,
      "nm_id": 940050,
      "brand": "SoundPeak",
      "status": 202
    },
    {
      "chrt_id": 8044502,
      "track_number": "WBUS-CHI-0004",
      "price": 900,
      "rid": "rid-004-b",
      "name": "Portable SSD 1TB",
      "sale": 0,
      "size": "std",
      "total_price": 900,
      "nm_id": 940051,
      "brand": "FlashBox",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "sig-us-4",
  "customer_id": "cust-4004",
  "delivery_service": "ups",
  "shardkey": "4",
  "sm_id": 104,
  "date_created": "2024-02-03T14:00:00Z",
  "oof_shard": "4"
}`,
	`{
  "order_uid": "ORD-0005-TO88ON",
  "track_number": "WBCAN-TOR-0005",
  "entry": "WBIL",
  "delivery": {
    "name": "Liam Johnson",
    "phone": "+14165550123",
    "zip": "M5V 2T6",
    "city": "Toronto",
    "address": "100 Queen St W",
    "region": "ON",
    "email": "liam.johnson@example.ca"
  },
  "payment": {
    "transaction": "tx-005-ca",
    "request_id": "",
    "currency": "CAD",
    "provider": "stripe",
    "amount": 7000,
    "payment_dt": 1704326400,
    "bank": "rbc",
    "delivery_cost": 1200,
    "goods_total": 5800,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9055601,
      "track_number": "WBCAN-TOR-0005",
      "price": 3200,
      "rid": "rid-005-a",
      "name": "4K Monitor 27\"",
      "sale": 0,
      "size": "std",
      "total_price": 3200,
      "nm_id": 950060,
      "brand": "ViewPro",
      "status": 202
    },
    {
      "chrt_id": 9055602,
      "track_number": "WBCAN-TOR-0005",
      "price": 2600,
      "rid": "rid-005-b",
      "name": "Ergonomic Chair",
      "sale": 0,
      "size": "std",
      "total_price": 2600,
      "nm_id": 950061,
      "brand": "SitRight",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-5005",
  "delivery_service": "canada_post",
  "shardkey": "5",
  "sm_id": 105,
  "date_created": "2024-01-03T16:45:00Z",
  "oof_shard": "5"
}`,
	`{
  "order_uid": "ORD-0006-MB77VIC",
  "track_number": "WBAU-MEL-0006",
  "entry": "WBIL",
  "delivery": {
    "name": "Amelia Brown",
    "phone": "+61370101234",
    "zip": "3000",
    "city": "Melbourne",
    "address": "90 Collins St",
    "region": "VIC",
    "email": "amelia.brown@example.au"
  },
  "payment": {
    "transaction": "tx-006-au",
    "request_id": "req-006",
    "currency": "AUD",
    "provider": "wbpay",
    "amount": 1800,
    "payment_dt": 1719964800,
    "bank": "anz",
    "delivery_cost": 200,
    "goods_total": 1600,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1006701,
      "track_number": "WBAU-MEL-0006",
      "price": 1600,
      "rid": "rid-006-a",
      "name": "Cordless Drill",
      "sale": 0,
      "size": "std",
      "total_price": 1600,
      "nm_id": 960070,
      "brand": "BuildMate",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "sig-au-6",
  "customer_id": "cust-6006",
  "delivery_service": "auspost",
  "shardkey": "6",
  "sm_id": 106,
  "date_created": "2024-07-03T08:10:00Z",
  "oof_shard": "6"
}`,
	`{
  "order_uid": "ORD-0007-DB11IE",
  "track_number": "WBIE-DUB-0007",
  "entry": "WBIL",
  "delivery": {
    "name": "Ethan Harris",
    "phone": "+35315550123",
    "zip": "D02",
    "city": "Dublin",
    "address": "2 Dawson St",
    "region": "Dublin",
    "email": "ethan.harris@example.ie"
  },
  "payment": {
    "transaction": "tx-007-ie",
    "request_id": "",
    "currency": "EUR",
    "provider": "stripe",
    "amount": 1500,
    "payment_dt": 1694044800,
    "bank": "aib",
    "delivery_cost": 150,
    "goods_total": 1350,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1107801,
      "track_number": "WBIE-DUB-0007",
      "price": 850,
      "rid": "rid-007-a",
      "name": "Smart Speaker",
      "sale": 0,
      "size": "std",
      "total_price": 850,
      "nm_id": 970080,
      "brand": "HomeWave",
      "status": 202
    },
    {
      "chrt_id": 1107802,
      "track_number": "WBIE-DUB-0007",
      "price": 500,
      "rid": "rid-007-b",
      "name": "LED Strip Lights",
      "sale": 0,
      "size": "5m",
      "total_price": 500,
      "nm_id": 970081,
      "brand": "GlowLine",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-7007",
  "delivery_service": "anpost",
  "shardkey": "7",
  "sm_id": 107,
  "date_created": "2023-09-07T12:00:00Z",
  "oof_shard": "7"
}`,
	`{
  "order_uid": "ORD-0008-AK55NZ",
  "track_number": "WBNZ-AKL-0008",
  "entry": "WBIL",
  "delivery": {
    "name": "Noah Thompson",
    "phone": "+6495550123",
    "zip": "1010",
    "city": "Auckland",
    "address": "1 Queen St",
    "region": "Auckland",
    "email": "noah.thompson@example.nz"
  },
  "payment": {
    "transaction": "tx-008-nz",
    "request_id": "req-008",
    "currency": "NZD",
    "provider": "paypal",
    "amount": 1000,
    "payment_dt": 1727049600,
    "bank": "asb",
    "delivery_cost": 90,
    "goods_total": 910,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1208901,
      "track_number": "WBNZ-AKL-0008",
      "price": 910,
      "rid": "rid-008-a",
      "name": "Action Camera",
      "sale": 0,
      "size": "std",
      "total_price": 910,
      "nm_id": 980090,
      "brand": "GoStream",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "sig-nz-8",
  "customer_id": "cust-8008",
  "delivery_service": "nzpost",
  "shardkey": "8",
  "sm_id": 108,
  "date_created": "2024-09-23T07:30:00Z",
  "oof_shard": "8"
}`,
	`{
  "order_uid": "ORD-0009-MN33UK",
  "track_number": "WBGB-MAN-0009",
  "entry": "WBIL",
  "delivery": {
    "name": "Ava Robinson",
    "phone": "+441612345678",
    "zip": "M2 5DB",
    "city": "Manchester",
    "address": "1 King Street",
    "region": "England",
    "email": "ava.robinson@example.co.uk"
  },
  "payment": {
    "transaction": "tx-009-gb",
    "request_id": "",
    "currency": "GBP",
    "provider": "wbpay",
    "amount": 600,
    "payment_dt": 1672531200,
    "bank": "barclays",
    "delivery_cost": 60,
    "goods_total": 540,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1310001,
      "track_number": "WBGB-MAN-0009",
      "price": 540,
      "rid": "rid-009-a",
      "name": "Electric Kettle",
      "sale": 0,
      "size": "1.7L",
      "total_price": 540,
      "nm_id": 990100,
      "brand": "BoilQuick",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-9009",
  "delivery_service": "royalmail",
  "shardkey": "9",
  "sm_id": 109,
  "date_created": "2023-01-01T00:00:00Z",
  "oof_shard": "9"
}`,
	`{
  "order_uid": "ORD-0010-AU44TX",
  "track_number": "WBUS-AUS-0010",
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
    "transaction": "tx-010-us",
    "request_id": "req-010",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 3000,
    "payment_dt": 1730505600,
    "bank": "wellsfargo",
    "delivery_cost": 500,
    "goods_total": 2500,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1410201,
      "track_number": "WBUS-AUS-0010",
      "price": 1600,
      "rid": "rid-010-a",
      "name": "Yoga Mat",
      "sale": 0,
      "size": "std",
      "total_price": 1600,
      "nm_id": 991010,
      "brand": "ZenFlex",
      "status": 202
    },
    {
      "chrt_id": 1410202,
      "track_number": "WBUS-AUS-0010",
      "price": 900,
      "rid": "rid-010-b",
      "name": "Foam Block",
      "sale": 0,
      "size": "std",
      "total_price": 900,
      "nm_id": 991011,
      "brand": "ZenFlex",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "cust-1010",
  "delivery_service": "usps",
  "shardkey": "10",
  "sm_id": 110,
  "date_created": "2024-11-02T08:00:00Z",
  "oof_shard": "10"
}`,
}
