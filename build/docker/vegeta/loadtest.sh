#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
PUT_AMOUNT=${1:-5000}
GET_AMOUNT=${2:-2000}
ADDR=${3:-"http://application:80"}
RATE=${4:-20}
DURATION=${5-"60s"}

function randLat()
{
    RAND_LAT=$(awk -v seed=$RANDOM 'BEGIN { srand(seed);printf("%.8f\n", 1.239600 + rand()*(1.478400 - 1.239600)) }')
    echo "${RAND_LAT}"
}

function randLon()
{
    RAND_LON=$(awk -v seed=$RANDOM 'BEGIN { srand(seed);printf("%.8f\n", 103.587348 + rand()*(103.978413-103.594000)) }')
    echo "${RAND_LON}"
}

function randRadius()
{
    RADIUS_FLOAT=$(awk -v seed=$RANDOM 'BEGIN { srand(seed);printf("%.8f\n", 500 + rand()*(100000 - 500)) }')
    RADIUS=${RADIUS_FLOAT%.*}

    echo "$RADIUS"
}

function randLimit()
{
    LIMIT_FLOAT=$(awk -v seed=$RANDOM 'BEGIN { srand(seed);printf("%.8f\n", 1 + rand()*500) }')
    LIMIT=${LIMIT_FLOAT%.*}
    echo "$LIMIT"
}

pushd ${DIR}
rm -rf requests
mkdir -p requests

echo "Generating PUT requests...."
for i in `seq 1 ${PUT_AMOUNT}`;
do

cat << EOF > requests/${i}.json
{
    "latitude": $(randLat),
    "longitude": $(randLon),
    "accuracy": 1.0
}
EOF

cat << EOF > requests/${i}_2.json
{
    "latitude": $(randLat),
    "longitude": $(randLon),
    "accuracy": 1.0
}
EOF

cat << EOF >> requests/targets.txt
PUT ${ADDR}/api/v1/drivers/${i}/location
@${DIR}/requests/${i}.json

PUT ${ADDR}/api/v1/drivers/${i}/location
@${DIR}/requests/${i}_2.json
EOF

done


echo "Generating GET requests...."
for i in `seq 1 ${GET_AMOUNT}`;
do
cat << EOF >> requests/targets.txt
GET ${ADDR}/api/v1/drivers?latitude=$(randLat)&longitude=$(randLon)&radius=$(randRadius)&limit=$(randLimit)

EOF
done

echo "Start load testing...."
vegeta attack -rate=${RATE} -duration=${DURATION} -targets=requests/targets.txt | vegeta report

popd


