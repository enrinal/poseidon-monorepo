import { Response, NextFunction } from "express";
import { RequestContext } from "../interface";

export default class ClaimsController {
  public static async fetchClaims(
    req: RequestContext,
    res: Response,
    next: NextFunction
  ) {
    try {
      const loggedInUser = req.context!.auth;
      res.json(loggedInUser);
    } catch (error) {
      next(error);
    }
  }
}
