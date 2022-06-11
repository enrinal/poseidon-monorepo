import { Response, NextFunction } from "express";
import { RequestContext } from "../interface";
import CommodityRepository from "../repository/CommodityRepository";
import CurrencyConverterRepository from "../repository/CurrencyConverterRepository";
import * as ss from "simple-statistics";
import moment from "moment";

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

    const response = {
      message: "success",
      data: returnData,
    };

    res.json(response);
  }

  public static async fetchCommodityAggregated(
    req: RequestContext,
    res: Response,
    next: NextFunction
  ) {
    try {
      const loggedInUser = req.context!.auth;
      if (loggedInUser.role !== "admin") {
        throw new Error("Unauthorized");
      }
      var data = await CommodityRepository.fetchCommodity();
    } catch (err) {
      next(err);
    }
    const result = [];
    const objTmp = {
      [moment().format("YYYY-MM-DD")]: {} as any,
    };

    for (var i = 0; i < data.length; i++) {
      const item = data[i];
      if (!item.tgl_parsed) item.tgl_parsed = new Date();
      const tglParsed = moment(new Date(item.tgl_parsed)).format(
        "YYYY-MM-DD HH:mm:ss"
      );

      const modifiedItem = {
        ...item,
        tgl_parsed: tglParsed,
      };

      const dateOnly = moment(new Date(tglParsed)).format("YYYY-MM-DD");
      const keyObj = `${item.area_provinsi}_${dateOnly}`;
      if (!objTmp[keyObj]) objTmp[keyObj] = [];

      objTmp[keyObj].push(modifiedItem);
    }

    for (const key in objTmp) {
      const value = objTmp[key];
      const keyArr = key.split("_");

      const year = parseInt(moment(new Date(keyArr[1])).format("YYYY"));
      const month = parseInt(moment(new Date(keyArr[1])).format("MM"));
      const week = parseInt(moment(new Date(keyArr[1])).format("w"));

      var sizeArr: number[] = [];
      var priceArr: number[] = [];

      var minSize,
        maxSize,
        medianSize,
        avgSize = 0;
      var minPrice,
        maxPrice,
        medianPrice,
        avgPrice = 0;

      for (var i = 0; i < value.length; i++) {
        sizeArr.push(parseInt(value[i].size));
        priceArr.push(parseInt(value[i].price));
      }

      if (sizeArr.length > 0) {
        minSize = ss.min(sizeArr);
        maxSize = ss.max(sizeArr);
        medianSize = ss.median(sizeArr);
        avgSize = ss.mean(sizeArr);
      }

      if (priceArr.length > 0) {
        minPrice = ss.min(priceArr);
        maxPrice = ss.max(priceArr);
        medianPrice = ss.median(priceArr);
        avgPrice = ss.mean(priceArr);
      }

      result.push({
        year,
        month,
        week,
        province_area: keyArr[0],
        size_aggregate: {
          min: minSize,
          max: maxSize,
          median: medianSize,
          avg: avgSize,
        },
        price_aggregate: {
          min: minPrice,
          max: maxPrice,
          median: medianPrice,
          avg: avgPrice,
        },
      });
    }

    const response = {
      message: "success",
      data: result,
    };

    res.json(response);
  }
}
