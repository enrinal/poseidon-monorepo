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
      const response = {
        message: "success",
        data: loggedInUser,
      };
      res.json(response);
    } catch (error) {
      next(error);
    }
  }
}
