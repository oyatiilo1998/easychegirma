#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/ek_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/ek_protos -I /usr/local/include --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done