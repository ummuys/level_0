#!/bin/sh


c_uid="ORD-0002-LA77CA"
echo "\n Заказ вне кэше, order_uid = $c_uid"
curl -sS -o /dev/null localhost:1337/api/v1/order/$c_uid
