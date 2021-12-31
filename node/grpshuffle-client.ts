import {
  Combination,
  ShuffleRequest,
  ShuffleResponse,
} from "./gen/grpshuffle_pb";
import { ComputeClient } from "./gen/grpshuffle_grpc_pb";
import { credentials } from "@grpc/grpc-js";

const defaultServerUrl = "localhost:13333";

export type GrpshuffleRequest = Omit<ShuffleRequest.AsObject, "targetsList"> & {
  targets: string;
};
export type GrpshuffleResponse = {
  combinations: Array<{
    targets: string[];
  }>;
};

export const callShuffle = async (
  requestParams: GrpshuffleRequest,
  serverUrl = defaultServerUrl
): Promise<GrpshuffleResponse> => {
  const request = new ShuffleRequest();
  request.setPartition(requestParams.partition);
  request.setTargetsList(requestParams.targets.split(","));
  request.setSequential(requestParams.sequential);

  const client = new ComputeClient(serverUrl, credentials.createInsecure());

  const rawResult = await new Promise<ShuffleResponse.AsObject>(
    (resolve, reject) => {
      client.shuffle(request, (error, response) => {
        if (error) {
          console.error(error);
          reject({
            code: error?.code || 500,
            message: error?.message || "something went wrong",
          });
        }

        return resolve(response.toObject());
      });
    }
  );

  const fixCombinationsList = rawResult.combinationsList.map((c) => {
    return {
      targets: c.targetsList,
    };
  });

  const result = {
    combinations: fixCombinationsList,
  };

  return result as GrpshuffleResponse;
};
