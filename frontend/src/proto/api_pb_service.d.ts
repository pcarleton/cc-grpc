// package: api
// file: api.proto

import * as api_pb from "./api_pb";
import {grpc} from "grpc-web-client";

type ApiGetHealth = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.GetHealthRequest;
  readonly responseType: typeof api_pb.GetHealthResponse;
};

type ApiCreateReport = {
  readonly methodName: string;
  readonly service: typeof Api;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.CreateReportRequest;
  readonly responseType: typeof api_pb.CreateReportResponse;
};

export class Api {
  static readonly serviceName: string;
  static readonly GetHealth: ApiGetHealth;
  static readonly CreateReport: ApiCreateReport;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor; debug?: boolean }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class ApiClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  getHealth(
    requestMessage: api_pb.GetHealthRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.GetHealthResponse|null) => void
  ): void;
  getHealth(
    requestMessage: api_pb.GetHealthRequest,
    callback: (error: ServiceError, responseMessage: api_pb.GetHealthResponse|null) => void
  ): void;
  createReport(
    requestMessage: api_pb.CreateReportRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.CreateReportResponse|null) => void
  ): void;
  createReport(
    requestMessage: api_pb.CreateReportRequest,
    callback: (error: ServiceError, responseMessage: api_pb.CreateReportResponse|null) => void
  ): void;
}

