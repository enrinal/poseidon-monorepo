import { Router } from "express";
import CommodityController from "../controller/CommodityController";
import { IsAuthenticated } from "../middlewares/IsAuthenticated";

const router = Router();

router.use(IsAuthenticated);
router.get("/commodity", CommodityController.fetchCommodity);

export default router;
