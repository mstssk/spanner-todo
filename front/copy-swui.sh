#!/bin/sh -eux

cd `dirname $0`

rm -rf ../server/src/api/swagger-ui
cp -r ./node_modules/swagger-ui-dist ../server/src/api/swagger-ui

sed -i -e "s#http://petstore.swagger.io/v2#/api#g" ../server/src/api/swagger-ui/index.html
