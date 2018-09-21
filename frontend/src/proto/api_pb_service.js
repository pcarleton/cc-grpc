// package: api
// file: api.proto

var api_pb = require("./api_pb");
var grpc = require("grpc-web-client").grpc;

var Api = (function () {
  function Api() {}
  Api.serviceName = "api.Api";
  return Api;
}());

Api.GetHealth = {
  methodName: "GetHealth",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.GetHealthRequest,
  responseType: api_pb.GetHealthResponse
};

Api.CreateReport = {
  methodName: "CreateReport",
  service: Api,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.CreateReportRequest,
  responseType: api_pb.CreateReportResponse
};

exports.Api = Api;

function ApiClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ApiClient.prototype.getHealth = function getHealth(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Api.GetHealth, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ApiClient.prototype.createReport = function createReport(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Api.CreateReport, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

exports.ApiClient = ApiClient;

