import { Router } from "express";
import CommodityController from "../controller/CommodityController";
import { IsAuthenticated } from "../middlewares/IsAuthenticated";

const router = Router();

router.use(IsAuthenticated);
router.get("/commodity", CommodityController.fetchCommodity);
router.get(
  "/commodity/aggregated",
  CommodityController.fetchCommodityAggregated
);

export default router;
