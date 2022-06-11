import { NextFunction, Response } from "express";
import { decodeToken } from "../helpers/tokenizer";
import { RequestContext } from "../interface";

export async function IsAuthenticated(
  req: RequestContext,
  res: Response,
  next: NextFunction
) {
  try {
    const accessToken = req.headers["authorization"] as string;
    const payload = decodeToken(accessToken);

    req.context = {
      auth: payload,
    };

    next();
  } catch (error) {
    next(error);
  }
}
