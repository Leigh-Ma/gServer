#!/bin/sh

CLIENT_DIR=/Users/liupingping/malin/SecretGarden/Assets/source/ggame/Net/Protocol

./gen_handler.sh
./gen_payload.sh
./gen_msgtype.sh

cp -rf csprotocol.txt  ${CLIENT_DIR}/protocol.txt
cp -rf cspayload.proto ${CLIENT_DIR}/gNetPayLoad.proto
cp -rf cspayload.proto ${CLIENT_DIR}/gNetPayLoad.proto.txt

cd ${CLIENT_DIR}

./gen_handler.sh
./gen_payload.sh
./gen_msgtype.sh