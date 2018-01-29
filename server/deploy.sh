#!/bin/sh -eux

cd `dirname $0`

set +x
APPLICATION=$1
VERSION=`echo $2 | sed -e s/\[/_]/-/g | awk '{print tolower($0)}'`
if [ $# -ge 3 ]; then
    echo $3 | base64 --decode > deploying_account.json
    ACCOUNT=`node -e "var a = \`cat deploying_account.json\`; console.log(a.client_email);"`
    gcloud auth activate-service-account ${ACCOUNT} \
        --key-file ./deploying_account.json \
        --project ${APPLICATION}
fi

appcfg.py update ./src \
    --oauth2_access_token $(gcloud auth print-access-token 2> /dev/null) \
    --application=$APPLICATION \
    --version=$VERSION
