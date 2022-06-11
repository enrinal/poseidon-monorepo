import { Request as ExpressRequest } from "express";
import { ParamsDictionary, Query } from "express-serve-static-core";

export interface TokenPayload {
  name: string;
  phone: string;
  role: string;
  exp: number;
}

export type Request<ReqBody = any> = ExpressRequest<
  ParamsDictionary,
  any,
  ReqBody,
  Query
>;

export interface Context {
  auth: TokenPayload;
}

export interface RequestContext<T = any> extends Request<T> {
  context?: Context;
}
