#!/bin/sh

curl -sS -o /dev/null localhost:1337/api/v1/order/$c_uid

c_uid="ORD-0010-AU44TX"
echo "Заказ в кэше, order_uid = $c_uid"
curl -sS -o /dev/null localhost:1337/api/v1/order/$c_uid
