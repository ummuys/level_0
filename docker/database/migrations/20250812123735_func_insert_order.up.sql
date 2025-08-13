CREATE OR REPLACE FUNCTION orders.insert_order(data jsonb)
RETURNS text
LANGUAGE plpgsql
SECURITY INVOKER
STRICT
AS $$
    DECLARE 
        oUID text;
    BEGIN
        -- main table (info)
        INSERT INTO orders.info (order_uid, track_number, entry, locale, internal_signature,
                                 customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
        VALUES (data->>'order_uid', data->>'track_number', data->>'entry', data->>'locale', data->>'internal_signature',
                data->>'customer_id', data->>'delivery_service', data->>'shardkey',
                (data->>'sm_id')::int, (data->>'date_created')::timestamptz, data->>'oof_shard')
        RETURNING order_uid INTO oUID;


        -- payments
        INSERT INTO orders.payments (order_uid, transaction_id, request_id, currency, provider,
                                     amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES (oUID,
        data#>>'{payment,transaction}', data#>>'{payment,request_id}', data#>>'{payment,currency}',
        data#>>'{payment,provider}', (data#>>'{payment,amount}')::int,
        (data#>>'{payment,payment_dt}')::bigint, data#>>'{payment,bank}',
        (data#>>'{payment,delivery_cost}')::int, (data#>>'{payment,goods_total}')::int,
        (data#>>'{payment,custom_fee}')::int);

        -- deliveries
        INSERT INTO orders.deliveries (order_uid, name, phone, zip, city, address, region, email)
        VALUES (oUID,
          data#>>'{delivery,name}', data#>>'{delivery,phone}', data#>>'{delivery,zip}',
          data#>>'{delivery,city}', data#>>'{delivery,address}', data#>>'{delivery,region}',
          data#>>'{delivery,email}');


        -- items
        INSERT INTO orders.items (order_uid, chrt_id, track_number, price, rid, name, sale, size,
                                 total_price, nm_id, brand, status)
        SELECT oUID,
         (items->>'chrt_id')::bigint, items->>'track_number', (items->>'price')::int,
         items->>'rid', items->>'name', (items->>'sale')::int, items->>'size',
         (items->>'total_price')::int, (items->>'nm_id')::bigint, items->>'brand', (items->>'status')::smallint
        FROM jsonb_array_elements(data->'items') items;

        RETURN oUID;
    END;
$$;


        

