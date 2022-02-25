import { ShuffleRequest, ShuffleResponse } from "../gen/grpshuffle_pb";
import { ComputeClient } from "../gen/grpshuffle_grpc_pb";
import { HealthCheckRequest, HealthCheckResponse } from "../gen/health_pb";
import { HealthClient } from "../gen/health_grpc_pb";
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
  request.setDivide(requestParams.divide);
  request.setTargetsList(requestParams.targets.split(","));
  request.setSequential(requestParams.sequential);

  const client = new ComputeClient(serverUrl, credentials.createInsecure());

  const rawResult = await new Promise<ShuffleResponse.AsObject>(
    (resolve, reject) => {
      client.shuffle(request, (error, response) => {
        if (error) {
          console.error(error);
          return reject({
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

export type HealthResponse = {
  status: string;
};

export const callHealth = async (
  serverUrl = defaultServerUrl
): Promise<HealthResponse> => {
  const client = new HealthClient(serverUrl, credentials.createInsecure());
  const request = new HealthCheckRequest();

  const rawResult = await new Promise<HealthCheckResponse.AsObject>(
    (resolve, reject) => {
      client.check(request, (error, response) => {
        if (error) {
          console.error(error);
          return reject({
            code: error?.code || 500,
            message: error?.message || "something went wrong",
          });
        }

        return resolve(response.toObject());
      });
    }
  );

  const result: HealthResponse = {
    status: "UNKNOWN",
  };

  switch (rawResult.status) {
    case HealthCheckResponse.ServingStatus.SERVING:
      result.status = "SERVING";
      break;
    case HealthCheckResponse.ServingStatus.NOT_SERVING:
      result.status = "NOT_SERVING";
      break;
    case HealthCheckResponse.ServingStatus.SERVICE_UNKNOWN:
      result.status = "SERVICE_UNKNOWN";
      break;
    case HealthCheckResponse.ServingStatus.UNKNOWN:
    default:
  }

  return result;
};
