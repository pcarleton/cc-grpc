import {grpc} from "grpc-web-client";

import {Api} from "./proto/api_pb_service"
import {GetHealthRequest} from "./proto/api_pb"


function getHealth(idToken: string) {

  const getHealthRequest = new GetHealthRequest();
  grpc.unary(Api.GetHealth, {
    request: getHealthRequest,
    host: "http://cc.pcarleton.com:5001",
    metadata: new grpc.Metadata({"token": idToken}),
    onEnd: res => {
      const {status, statusMessage, headers, message, trailers } = res;
      if (status === grpc.Code.OK) {
        console.log("Ok!", message.toObject());
      } else {
        alert("Error: " + statusMessage);
      }
      console.log(res);
    }
  });

  return "Hello from Client!";
}

const Client = { getHealth };
export default Client;
