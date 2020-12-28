#!/bin/bash
set -ex
export TF_GIT_BRANCH=${TF_GIT_BRANCH:master}
git clone -b ${TF_GIT_BRANCH} https://github.com/tensorflow/tensorflow.git
git clone -b ${TF_GIT_BRANCH} https://github.com/tensorflow/serving.git

mkdir -p serving_client

PROTOC_OPTS='--proto_path=${HOME}/go/src -I tensorflow -I serving --go_out=serving_client'

eval "protoc $PROTOC_OPTS serving/tensorflow_serving/apis/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/config/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/util/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/sources/storage_path/*.proto"
eval "protoc $PROTOC_OPTS serving/tensorflow_serving/core/*.proto"
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/framework/*.proto"
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/example/*.proto"
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/lib/core/*.proto"
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/core/protobuf/*.proto"
eval "protoc $PROTOC_OPTS tensorflow/tensorflow/stream_executor/*.proto"
