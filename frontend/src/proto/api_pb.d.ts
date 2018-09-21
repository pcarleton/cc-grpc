// package: api
// file: api.proto

import * as jspb from "google-protobuf";

export class GetHealthRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetHealthRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetHealthRequest): GetHealthRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetHealthRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetHealthRequest;
  static deserializeBinaryFromReader(message: GetHealthRequest, reader: jspb.BinaryReader): GetHealthRequest;
}

export namespace GetHealthRequest {
  export type AsObject = {
  }
}

export class HealthCheckResponse extends jspb.Message {
  getLabel(): string;
  setLabel(value: string): void;

  getStatus(): HealthStatus;
  setStatus(value: HealthStatus): void;

  getResult(): string;
  setResult(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HealthCheckResponse.AsObject;
  static toObject(includeInstance: boolean, msg: HealthCheckResponse): HealthCheckResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HealthCheckResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HealthCheckResponse;
  static deserializeBinaryFromReader(message: HealthCheckResponse, reader: jspb.BinaryReader): HealthCheckResponse;
}

export namespace HealthCheckResponse {
  export type AsObject = {
    label: string,
    status: HealthStatus,
    result: string,
  }
}

export class GetHealthResponse extends jspb.Message {
  getVersion(): string;
  setVersion(value: string): void;

  clearStatusesList(): void;
  getStatusesList(): Array<HealthCheckResponse>;
  setStatusesList(value: Array<HealthCheckResponse>): void;
  addStatuses(value?: HealthCheckResponse, index?: number): HealthCheckResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetHealthResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetHealthResponse): GetHealthResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetHealthResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetHealthResponse;
  static deserializeBinaryFromReader(message: GetHealthResponse, reader: jspb.BinaryReader): GetHealthResponse;
}

export namespace GetHealthResponse {
  export type AsObject = {
    version: string,
    statusesList: Array<HealthCheckResponse.AsObject>,
  }
}

export class CreateReportRequest extends jspb.Message {
  getMonth(): number;
  setMonth(value: number): void;

  getNamespace(): string;
  setNamespace(value: string): void;

  getAccountId(): string;
  setAccountId(value: string): void;

  getSpreadsheetId(): string;
  setSpreadsheetId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReportRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReportRequest): CreateReportRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateReportRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReportRequest;
  static deserializeBinaryFromReader(message: CreateReportRequest, reader: jspb.BinaryReader): CreateReportRequest;
}

export namespace CreateReportRequest {
  export type AsObject = {
    month: number,
    namespace: string,
    accountId: string,
    spreadsheetId: string,
  }
}

export class CreateReportResponse extends jspb.Message {
  getResult(): string;
  setResult(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateReportResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateReportResponse): CreateReportResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateReportResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateReportResponse;
  static deserializeBinaryFromReader(message: CreateReportResponse, reader: jspb.BinaryReader): CreateReportResponse;
}

export namespace CreateReportResponse {
  export type AsObject = {
    result: string,
  }
}

export enum HealthStatus {
  UNHEALTHY = 0,
  OK = 1,
}

