import { Router } from "express";
import { IsAuthenticated } from "../middlewares/IsAuthenticated";
import ClaimsController from "../controller/ClaimsController";

const router = Router();

router.use(IsAuthenticated);
router.get("/claims", ClaimsController.fetchClaims);

export default router;
