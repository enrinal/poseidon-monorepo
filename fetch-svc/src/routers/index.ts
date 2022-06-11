import { Router } from "express";
import comodityRouter from "./commodity";
import claimsRouter from "./claims";

const router = Router();

router.use("/api/v1", comodityRouter);
router.use("/api/v1", claimsRouter);
export default router;
