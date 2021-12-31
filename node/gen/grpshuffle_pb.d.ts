// package: grpshuffle
// file: grpshuffle.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Combination extends jspb.Message { 
    clearTargetsList(): void;
    getTargetsList(): Array<string>;
    setTargetsList(value: Array<string>): Combination;
    addTargets(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Combination.AsObject;
    static toObject(includeInstance: boolean, msg: Combination): Combination.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Combination, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Combination;
    static deserializeBinaryFromReader(message: Combination, reader: jspb.BinaryReader): Combination;
}

export namespace Combination {
    export type AsObject = {
        targetsList: Array<string>,
    }
}

export class ShuffleRequest extends jspb.Message { 
    clearTargetsList(): void;
    getTargetsList(): Array<string>;
    setTargetsList(value: Array<string>): ShuffleRequest;
    addTargets(value: string, index?: number): string;
    getPartition(): number;
    setPartition(value: number): ShuffleRequest;
    getSequential(): boolean;
    setSequential(value: boolean): ShuffleRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShuffleRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ShuffleRequest): ShuffleRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShuffleRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShuffleRequest;
    static deserializeBinaryFromReader(message: ShuffleRequest, reader: jspb.BinaryReader): ShuffleRequest;
}

export namespace ShuffleRequest {
    export type AsObject = {
        targetsList: Array<string>,
        partition: number,
        sequential: boolean,
    }
}

export class ShuffleResponse extends jspb.Message { 
    clearCombinationsList(): void;
    getCombinationsList(): Array<Combination>;
    setCombinationsList(value: Array<Combination>): ShuffleResponse;
    addCombinations(value?: Combination, index?: number): Combination;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ShuffleResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ShuffleResponse): ShuffleResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ShuffleResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ShuffleResponse;
    static deserializeBinaryFromReader(message: ShuffleResponse, reader: jspb.BinaryReader): ShuffleResponse;
}

export namespace ShuffleResponse {
    export type AsObject = {
        combinationsList: Array<Combination.AsObject>,
    }
}
