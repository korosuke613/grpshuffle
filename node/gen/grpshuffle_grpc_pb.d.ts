// package: grpshuffle
// file: grpshuffle.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as grpshuffle_pb from "./grpshuffle_pb";

interface IComputeService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    shuffle: IComputeService_IShuffle;
}

interface IComputeService_IShuffle extends grpc.MethodDefinition<grpshuffle_pb.ShuffleRequest, grpshuffle_pb.ShuffleResponse> {
    path: "/grpshuffle.Compute/Shuffle";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<grpshuffle_pb.ShuffleRequest>;
    requestDeserialize: grpc.deserialize<grpshuffle_pb.ShuffleRequest>;
    responseSerialize: grpc.serialize<grpshuffle_pb.ShuffleResponse>;
    responseDeserialize: grpc.deserialize<grpshuffle_pb.ShuffleResponse>;
}

export const ComputeService: IComputeService;

export interface IComputeServer extends grpc.UntypedServiceImplementation {
    shuffle: grpc.handleUnaryCall<grpshuffle_pb.ShuffleRequest, grpshuffle_pb.ShuffleResponse>;
}

export interface IComputeClient {
    shuffle(request: grpshuffle_pb.ShuffleRequest, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
    shuffle(request: grpshuffle_pb.ShuffleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
    shuffle(request: grpshuffle_pb.ShuffleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
}

export class ComputeClient extends grpc.Client implements IComputeClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public shuffle(request: grpshuffle_pb.ShuffleRequest, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
    public shuffle(request: grpshuffle_pb.ShuffleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
    public shuffle(request: grpshuffle_pb.ShuffleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: grpshuffle_pb.ShuffleResponse) => void): grpc.ClientUnaryCall;
}
