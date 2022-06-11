import { Response, NextFunction } from "express";
import { RequestContext } from "../interface";
import CommodityRepository from "../repository/CommodityRepository";
import CurrencyConverterRepository from "../repository/CurrencyConverterRepository";

export default class CommodityController {
  public static async fetchCommodity(
    req: RequestContext,
    res: Response,
    next: NextFunction
  ) {
    try {
      const loggedInUser = req.context!.auth;
      var data = await CommodityRepository.fetchCommodity();
    } catch (err) {
      next(err);
    }

    try {
      var currency = await CurrencyConverterRepository.fetchCurrencyConverter();
    } catch (err) {
      next(err);
    }

    var returnData = [];

    for (var i = 0; i < data.length; i++) {
      returnData.push({
        ...data[i],
        price_usd: parseInt(data[i].price) / currency,
      });
    }

    res.json(returnData);
  }
}
