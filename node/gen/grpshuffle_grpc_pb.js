// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var grpshuffle_pb = require('./grpshuffle_pb.js');

function serialize_grpshuffle_ShuffleRequest(arg) {
  if (!(arg instanceof grpshuffle_pb.ShuffleRequest)) {
    throw new Error('Expected argument of type grpshuffle.ShuffleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_grpshuffle_ShuffleRequest(buffer_arg) {
  return grpshuffle_pb.ShuffleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_grpshuffle_ShuffleResponse(arg) {
  if (!(arg instanceof grpshuffle_pb.ShuffleResponse)) {
    throw new Error('Expected argument of type grpshuffle.ShuffleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_grpshuffle_ShuffleResponse(buffer_arg) {
  return grpshuffle_pb.ShuffleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var ComputeService = exports.ComputeService = {
  shuffle: {
    path: '/grpshuffle.Compute/Shuffle',
    requestStream: false,
    responseStream: false,
    requestType: grpshuffle_pb.ShuffleRequest,
    responseType: grpshuffle_pb.ShuffleResponse,
    requestSerialize: serialize_grpshuffle_ShuffleRequest,
    requestDeserialize: deserialize_grpshuffle_ShuffleRequest,
    responseSerialize: serialize_grpshuffle_ShuffleResponse,
    responseDeserialize: deserialize_grpshuffle_ShuffleResponse,
  },
  // Shuffle.
};

exports.ComputeClient = grpc.makeGenericClientConstructor(ComputeService);
